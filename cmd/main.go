package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/YarikRevich/go-demonizer/pkg/demonizer"
	externalapiimp "github.com/YarikRevich/hide-seek-server/internal/api/external-api/v1/implementation"
	"github.com/YarikRevich/hide-seek-server/internal/api/external-api/v1/proto"
	"github.com/YarikRevich/hide-seek-server/internal/interceptors"
	"github.com/YarikRevich/hide-seek-server/internal/monitoring"

	"github.com/YarikRevich/hide-seek-server/tools/params"
	"github.com/YarikRevich/hide-seek-server/tools/printer"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

func init() {
	rand.Seed(time.Now().Unix())

	flag.Parse()

	if params.IsDemon() {
		demonizer.DemonizeThisProcess()
	}

	logrus.SetFormatter(logrus.StandardLogger().Formatter)

	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.InfoLevel)

	printer.PrintWelcomeMessage()
}

func main() {
	conn, err := net.Listen("tcp", fmt.Sprintf("%s:%s", params.GetServerIP(), params.GetServerPort()))
	if err != nil {
		logrus.Fatal(err)
	}

	monitoring.UseMonitoring().ListenAndServe()

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			interceptors.NewInterceptorManager(),
		)}
	s := grpc.NewServer(opts...)

	grpc.UseCompressor(gzip.Name)
	// cache.UseCache()

	proto.RegisterExternalServerServiceServer(s, externalapiimp.NewExternalServerService())
	if err := s.Serve(conn); err != nil {
		logrus.Fatal(err)
	}
}
