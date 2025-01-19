package exception

import "fmt"

type ClientError struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Trace      interface{} `json:"trace"`
	rawError   error
}

func (clientError *ClientError) Error() string {
	return fmt.Sprintf("Error %d: %s", clientError.StatusCode, clientError.Message)
}

func NewClientError(statusCode int, message string, rawError error, traceError ...interface{}) *ClientError {
	return &ClientError{
		StatusCode: statusCode,
		rawError:   rawError,
		Message:    message,
		Trace:      traceError,
	}
}

func (clientError *ClientError) GetRawError() error {
	if clientError == nil || clientError.rawError == nil {
		return nil
	}
	return clientError.rawError
}

func ThrowClientError(clientError *ClientError) {
	panic(clientError)
}
