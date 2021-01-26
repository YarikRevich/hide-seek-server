package main

import (
	"os"
	"net"
	"fmt"
	"strconv"
	"strings"
	C "GameServer/Components"
	"GameServer/MessageParser"
	"github.com/mbndr/figlet4go"
)

func ListenToCommands(request []byte, listener *net.UDPConn, addr *net.UDPAddr, data *C.Data){
	//The main part of programm. Firstly it parses
	//upcoming message and gets 'id' 'unparsedReq' and 'err'
	//Id is the index of request. unparsedReq is the body of request
	//itself. Err means when the request type is not supported.

	parser := MessageParser.Parser(&MessageParser.Message{})
	result := parser.Unparse(request)

	C.UpdateExactUser(result[0], addr, data)

	switch result[0].Type{
	case "CreateLobby":
		code := C.CreateLobby(result[0], listener, addr, data)
		C.SendAnswerS(result, code, listener, addr)
	case "AddToLobby":
		code := C.AddToLobby(result[0], listener, addr, data)
		C.SendAnswerS(result, code, listener, addr)
	case "ClosePreparingLobby":
		code := C.ClosePreparingLobby(result[0].PersonalInfo.LobbyID, data)
		C.SendAnswerS(result, code, listener, addr)
	case "GetUsersInfoPrepLobby":
		usersInfo, code := C.GetUsersInfoPrepLobby(result[0], addr, listener, data)
		C.SendAnswerS(usersInfo, code, listener, addr)
	case "GetUsersInfoReadyLobby":
		usersInfo, code := C.GetUsersInfoReadyLobby(result[0], addr, listener, data)
		C.SendAnswerS(usersInfo, code, listener, addr)
	case "UpdateUsersHealth":
		usersInfo, code := C.UpdateUsersHealth(result[0], data)
		C.SendAnswerS(usersInfo, code, listener, addr)
	case "DeleteLobby":
		code := C.DeleteLobby(result[0].PersonalInfo.LobbyID, addr, listener, result[0].Networking.Index, data)
		C.SendAnswerS(result, code, listener, addr)
	case "OK":
		C.SendAnswerS(result, "200", listener, addr)
	}
}

func Server(){
	//Checks the env to get the adress to listen to. 

	data := new(C.Data)
	data.PreparingLobbies = make(map[string][]*net.UDPAddr)
	data.ReadyLobbies = make(map[string][]*net.UDPAddr)
	data.ExactClient = make(map[string]*MessageParser.Message)

	envaddr := os.Getenv("GAMESERVER_ADDR")
	if len(envaddr) == 0{
		fmt.Println("GAMESERVER is not written!")
		os.Exit(0)
	}
	splittedAddr := strings.Split(envaddr, ":")
	ip, portStr := splittedAddr[0], splittedAddr[1]
	port, err := strconv.Atoi(portStr)
	if err != nil{
		panic(err)
	}

	addr := net.UDPAddr{
		Port: port,
		IP: net.ParseIP(ip),
	}
	listener, _ := net.ListenUDP("udp", &addr)

	DrawWelcomeMessage()

	for{
		notCleanedBuff := make([]byte, 4096)
		_, addr, err := listener.ReadFromUDP(notCleanedBuff)
		if err != nil{
			continue
		}
		var cleanedBuff []byte
		for _, value := range notCleanedBuff{
			if value != 0{
				cleanedBuff = append(cleanedBuff, value)
			}
		}
		ListenToCommands(cleanedBuff, listener, addr, data)
	}
}

func DrawWelcomeMessage(){
	//Writes hello message on the screen when the programm is started

	renderer := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorGreen,
	}
	text, err := renderer.RenderOpts("Hide&Seek server!", options)
	if err != nil{
		panic(err)
	}
	fmt.Print(text)
	fmt.Println("This server just process players...")
}

func main(){
	Server()
}
