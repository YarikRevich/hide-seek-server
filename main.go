package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	externalapiimp "github.com/YarikRevich/HideSeek-Server/internal/api/external-api/v1/implementation"
	externalapiproto "github.com/YarikRevich/HideSeek-Server/internal/api/external-api/v1/proto"
	"github.com/YarikRevich/HideSeek-Server/internal/cache"
	"github.com/YarikRevich/HideSeek-Server/internal/interceptors"
	"github.com/YarikRevich/HideSeek-Server/internal/monitoring"
	"github.com/YarikRevich/go-demonizer/pkg/demonizer"

	"github.com/YarikRevich/HideSeek-Server/tools/params"
	"github.com/YarikRevich/HideSeek-Server/tools/printer"
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
	cache.UseCache()

	externalapiproto.RegisterExternalServiceServer(s, externalapiimp.NewExternalService())
	s.Serve(conn)
}
