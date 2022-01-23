package testclient

import (
	external "github.com/YarikRevich/hide-seek-server/internal/api/external-api/v1/proto"
	internal "github.com/YarikRevich/hide-seek-server/internal/api/internal-api/v1/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Run() (external.ExternalServerServiceClient, internal.InternalServiceClient) {
	opts := []grpc.DialOption{grpc.WithInsecure()}

	d, err := grpc.Dial("127.0.0.1:8080", opts...)
	if err != nil {
		logrus.Fatalln(err)
	}
	return external.NewExternalServerServiceClient(d), internal.NewInternalServiceClient(d)
}
