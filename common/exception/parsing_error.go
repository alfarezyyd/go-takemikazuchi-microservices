package exception

import (
	"encoding/json"
	"errors"
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func ParseGormError(err error) *ClientError {
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return &ClientError{
				Message:    "Record not found",
				StatusCode: http.StatusNotFound,
			}

		case errors.Is(err, gorm.ErrDuplicatedKey):
			return &ClientError{
				Message:    "Data already exists",
				StatusCode: http.StatusConflict,
			}

		// Handle MySQL/Postgres specific errors
		case errors.Is(err, gorm.ErrForeignKeyViolated):
			return &ClientError{
				Message:    "Related record not found",
				StatusCode: http.StatusBadRequest,
			}

		case errors.Is(err, gorm.ErrDuplicatedKey):
			return &ClientError{
				Message:    "Duplicate entry",
				StatusCode: http.StatusConflict,
			}
		case errors.Is(err, gorm.ErrInvalidData):
			return &ClientError{
				Message:    "Invalid data",
				StatusCode: http.StatusBadRequest,
			}
		default:
			var clientError *ClientError
			isClientError := errors.As(err, &clientError)
			if isClientError {
				return NewClientError(clientError.StatusCode, clientError.Message, clientError.rawError)
			}
			return NewClientError(http.StatusInternalServerError, "Database error occurred", errors.New("database error occurred"))
		}
	}
	return nil
}

func ParseValidationError(validationError error, engTranslator ut.Translator) {
	if validationError != nil {
		parsedMap := make(map[string]interface{})
		for _, fieldError := range validationError.(validator.ValidationErrors) {
			parsedMap[fieldError.Field()] = fieldError.Translate(engTranslator)
		}
		panic(NewClientError(http.StatusBadRequest, ErrBadRequest, validationError, parsedMap))
	}
}

func ParseGrpcError(err error) {
	if err != nil {
		// Cek apakah error berasal dari gRPC
		statusRequest, ok := status.FromError(err)
		var errorDetail interface{}
		var statusMessage string = statusRequest.Message()
		if ok {
			// Mapping kode gRPC ke kode HTTP
			var httpStatus int
			switch statusRequest.Code() {
			case codes.NotFound:
				httpStatus = http.StatusNotFound // 404
			case codes.InvalidArgument:
				statusRequest.Message()
				splitMessage := strings.Split(statusRequest.Message(), ": ")
				statusMessage = splitMessage[0]
				unparsedErrorDetail := splitMessage[1]
				fmt.Println(splitMessage)
				// Buat variable untuk menampung hasil unmarshal
				var parsedData []map[string]interface{}

				// Lakukan unmarshal dengan menangkap error
				err := json.Unmarshal([]byte(unparsedErrorDetail), &parsedData)
				if err != nil {
					fmt.Println("Error Unmarshal:", err)
					return
				}
				errorDetail = parsedData
				httpStatus = http.StatusBadRequest // 400
			case codes.DeadlineExceeded:
				httpStatus = http.StatusGatewayTimeout // 504
			case codes.Unavailable:
				httpStatus = http.StatusServiceUnavailable // 503
			case codes.Internal:
				httpStatus = http.StatusInternalServerError // 500
			case codes.Unauthenticated:
				httpStatus = http.StatusUnauthorized // 401
			case codes.PermissionDenied:
				httpStatus = http.StatusForbidden // 403
			case codes.AlreadyExists:
				httpStatus = http.StatusConflict // 409
			case codes.FailedPrecondition:
				httpStatus = http.StatusPreconditionFailed // 412
			case codes.Aborted:
				httpStatus = http.StatusConflict // 409
			case codes.OutOfRange:
				httpStatus = http.StatusBadRequest // 400
			case codes.Unimplemented:
				httpStatus = http.StatusNotImplemented // 501
			default:
				httpStatus = http.StatusInternalServerError // 500
			}
			panic(NewClientError(httpStatus, statusMessage, nil, errorDetail))
		} else {
			panic(NewClientError(http.StatusInternalServerError, err.Error(), nil))
		}
	}
}

// Map HTTP status codes to gRPC error codes
func HttpStatusIntoGrpcCode(httpStatus int) codes.Code {
	switch httpStatus {
	case 400:
		return codes.InvalidArgument
	case 401:
		return codes.Unauthenticated
	case 403:
		return codes.PermissionDenied
	case 404:
		return codes.NotFound
	case 409:
		return codes.AlreadyExists
	case 412:
		return codes.FailedPrecondition
	case 500:
		return codes.Internal
	case 503:
		return codes.Unavailable
	case 504:
		return codes.DeadlineExceeded
	default:
		return codes.Unknown
	}
}

// Convert ClientError to gRPC error
func ParseIntoGrpcError(clientError *ClientError) error {
	if clientError == nil {
		return nil
	}

	grpcCode := HttpStatusIntoGrpcCode(clientError.StatusCode)
	return status.Error(grpcCode, clientError.Message)
}
