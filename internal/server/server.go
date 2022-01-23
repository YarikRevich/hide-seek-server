package server

import (
	"fmt"
	"net"

	"github.com/YarikRevich/hide-seek-server/internal/api/external-api/v1/implementation"
	"github.com/YarikRevich/hide-seek-server/internal/api/external-api/v1/proto"
	"github.com/YarikRevich/hide-seek-server/internal/interceptors"
	"github.com/YarikRevich/hide-seek-server/internal/monitoring"
	"github.com/YarikRevich/hide-seek-server/tools/params"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

func Run() {
	conn, err := net.Listen("tcp", fmt.Sprintf("%s:%s", params.GetServerIP(), params.GetServerPort()))
	if err != nil {
		logrus.Fatal(err)
	}

	monitoring.UseMonitoring().Run()

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			interceptors.NewInterceptorManager(),
		)}
	s := grpc.NewServer(opts...)

	grpc.UseCompressor(gzip.Name)
	// cache.UseCache()

	proto.RegisterExternalServerServiceServer(s, implementation.NewExternalServerService())
	if err := s.Serve(conn); err != nil {
		logrus.Fatal(err)
	}
}
