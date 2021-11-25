package main

import (
	"net"
	"os"

	externalapiimp "github.com/YarikRevich/HideSeek-Server/internal/api/external-api/v1/implementation"
	externalapiproto "github.com/YarikRevich/HideSeek-Server/internal/api/external-api/v1/proto"
	"github.com/YarikRevich/HideSeek-Server/internal/cache"
	"github.com/YarikRevich/HideSeek-Server/internal/interceptors"
	"github.com/YarikRevich/HideSeek-Server/tools/printer"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

// 	switch result[0].Type{
// 	case "CreateLobby":
// 		code := C.CreateLobby(result[0], listener, addr, data)
// 		C.SendAnswerS(result, code, listener, addr)
// 	case "AddToLobby":
// 		code := C.AddToLobby(result[0], listener, addr, data)
// 		C.SendAnswerS(result, code, listener, addr)
// 	case "ClosePreparingLobby":
// 		code := C.ClosePreparingLobby(result[0].PersonalInfo.LobbyID, data)
// 		C.SendAnswerS(result, code, listener, addr)
// 	case "GetUsersInfoPrepLobby":
// 		usersInfo, code := C.GetUsersInfoPrepLobby(result[0], addr, listener, data)
// 		C.SendAnswerS(usersInfo, code, listener, addr)
// 	case "GetUsersInfoReadyLobby":
// 		usersInfo, code := C.GetUsersInfoReadyLobby(result[0], addr, listener, data)
// 		C.SendAnswerS(usersInfo, code, listener, addr)
// 	case "UpdateUsersHealth":
// 		usersInfo, code := C.UpdateUsersHealth(result[0], data)
// 		C.SendAnswerS(usersInfo, code, listener, addr)
// 	case "DeleteLobby":
// 		code := C.DeleteLobby(result[0].PersonalInfo.LobbyID, addr, listener, result[0].Networking.Index, data)
// 		C.SendAnswerS(result, code, listener, addr)
// 	case "OK":
// 		C.SendAnswerS(result, "200", listener, addr)
// 	}

func init() {
	logrus.SetFormatter(logrus.StandardLogger().Formatter)

	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.WarnLevel)

	printer.PrintWelcomeMessage()
}

func main() {
	conn, err := net.Listen("tcp", ":8090")
	if err != nil {
		logrus.Fatal(err)
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(interceptors.NewInterceptorManager())}
	s := grpc.NewServer(opts...)

	grpc.UseCompressor(gzip.Name)
	cache.UseCache()

	externalapiproto.RegisterExternalServiceServer(s, externalapiimp.NewExternalService())
	s.Serve(conn)
}
