package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RecoveryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	// Use a named return to capture and modify the return values from deferred function

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic caught: %v\n", r)
			if clientError, ok := r.(*exception.ClientError); ok {
				grpcErrorCode := exception.HttpStatusIntoGrpcCode(clientError.StatusCode)
				switch grpcErrorCode {
				case codes.Canceled:
					err = status.Errorf(codes.Canceled, "Request dibatalkan: %s", clientError.Message)
				case codes.Unknown:
					err = status.Errorf(codes.Unknown, "Kesalahan tidak diketahui: %s", clientError.Message)
				case codes.InvalidArgument:
					errorDetails, _ := json.Marshal(clientError.Trace) // Konversi ke JSON
					err = status.Errorf(codes.InvalidArgument, "Input tidak valid: %s", string(errorDetails))
				case codes.DeadlineExceeded:
					err = status.Errorf(codes.DeadlineExceeded, "Request melebihi batas waktu: %s", clientError.Message)
				case codes.NotFound:
					err = status.Errorf(codes.NotFound, "Data tidak ditemukan: %s", clientError.Message)
				case codes.AlreadyExists:
					err = status.Errorf(codes.AlreadyExists, "Data sudah ada: %s", clientError.Message)
				case codes.PermissionDenied:
					err = status.Errorf(codes.PermissionDenied, "Akses ditolak: %s", clientError.Message)
				case codes.ResourceExhausted:
					err = status.Errorf(codes.ResourceExhausted, "Sumber daya habis: %s", clientError.Message)
				case codes.FailedPrecondition:
					err = status.Errorf(codes.FailedPrecondition, "Kondisi gagal terpenuhi: %s", clientError.Message)
				case codes.Aborted:
					err = status.Errorf(codes.Aborted, "Operasi dibatalkan: %s", clientError.Message)
				case codes.OutOfRange:
					err = status.Errorf(codes.OutOfRange, "Nilai di luar jangkauan: %s", clientError.Message)
				case codes.Unimplemented:
					err = status.Errorf(codes.Unimplemented, "Fitur belum diimplementasikan: %s", clientError.Message)
				case codes.Internal:
					err = status.Errorf(codes.Internal, "Terjadi kesalahan di server: %s", clientError.Message)
				case codes.Unavailable:
					err = status.Errorf(codes.Unavailable, "Layanan tidak tersedia: %s", clientError.Message)
				case codes.DataLoss:
					err = status.Errorf(codes.DataLoss, "Terjadi kehilangan data: %s", clientError.Message)
				case codes.Unauthenticated:
					err = status.Errorf(codes.Unauthenticated, "Unauthenticated: %s", clientError.Message)
				default:
					err = status.Errorf(codes.Internal, "Terjadi kesalahan tidak diketahui: %s", clientError.Message)
				}

			}
		}
	}()

	// Call the handler
	resp, err = handler(ctx, req)

	// Return values will be potentially modified by the deferred function
	return resp, err
}
