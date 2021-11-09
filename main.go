package main

import (
	"log"
	"os"

	"github.com/YarikRevich/HideSeek-Server/internal/cache"
	"github.com/YarikRevich/HideSeek-Server/internal/collection"
	"github.com/YarikRevich/HideSeek-Server/internal/handlers"
	"github.com/YarikRevich/HideSeek-Server/tools/printer"
	"github.com/YarikRevich/game-networking/pkg/config"
	"github.com/YarikRevich/game-networking/pkg/server"
	"github.com/sirupsen/logrus"
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

func init(){
	logrus.SetFormatter(logrus.StandardLogger().Formatter)
	
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.WarnLevel)

	printer.PrintWelcomeMessage()
}

func main(){
	conn := server.Listen(config.Config{
		IP: "127.0.0.1",
		Port: "8090"})

	// conn.AddHandler("reg_user", handlers.RegUser)
	// conn.AddHandler("reg_world", handlers.RegWorld)
	conn.AddHandler("update_world", handlers.UpdateWorldHandler)
	conn.AddHandler("close_game_session", handlers.CloseGameSession)
	conn.AddHandler("init_world_user_spawns", handlers.InitWorldUserSpawns)
	
	cache.UseCache().Start()

	log.Fatalln(conn.WaitForInterrupt())
}
