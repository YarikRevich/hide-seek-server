package externalapiinterceptor

import (
	"context"

	"google.golang.org/grpc"
)

type ExternalApiInterceptor struct{}

func (eai *ExternalApiInterceptor) Use(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// logrus.Info(info.FullMethod, info.Server)
	return handler(ctx, req)
}

func NewExternalApiInterceptor() *ExternalApiInterceptor {
	return new(ExternalApiInterceptor)
}
