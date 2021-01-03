package main

import (
	"os"
	"net"
	"fmt"
	"strconv"
	"strings"
	"GameServer/Utils"
	C "GameServer/Components"
	"GameServer/MessageParser"
	"github.com/mbndr/figlet4go"
)

func ListenToCommands(request string, listener *net.UDPConn, addr *net.UDPAddr, data *C.Data){
	//The main part of programm. Firstly it parses
	//upcoming message and gets 'id' 'unparsedReq' and 'err'
	//Id is the index of request. unparsedReq is the body of request
	//itself. Err means when the request type is not supported.
	
	id, unparsedReq, err := MessageParser.UnparseMessage(request)
	if err != nil {
		Utils.SendErrorResponse(err, id, listener, addr)
	}
	lobbyID := Utils.GetLobbyNum(request)
	switch unparsedReq{
	case "CreateLobby":
		err := C.CreateLobby(lobbyID, id, data)
		if err != nil {
			Utils.SendErrorResponse(err, id, listener, addr)
		}else{
			C.SendOK(addr, listener, id)
		}
	case "AddToLobby":
		err := C.AddToLobby(lobbyID, request, addr, data)
		if err != nil {
			Utils.SendErrorResponse(err, id, listener, addr)
		}else{
			C.SendOK(addr, listener, id)
		}
	case "GetMembersInLobby":
		lobbyMembers, lobbyMembersAddrs, err := C.GetMembersInLobby(lobbyID, data)
		if err != nil {
			Utils.SendErrorResponse(err, id, listener, addr)
		}
		C.SendLobbyMembers(listener, id, lobbyMembers, lobbyMembersAddrs)
	case "UpdateUser":
		C.UpdateUser(request, addr, data)
	case "ClosePreparingLobby":
		C.ClosePreparingLobby(lobbyID, data)
		C.SendOK(addr, listener, id)
	case "GetUsersInfoLobby":
		err := C.GetUsersInfoLobby(lobbyID, addr, listener, id, data)
		if err != nil {
			Utils.SendErrorResponse(err, id, listener, addr)
		}
	case "GetUsersInfo":
		C.GetUsersInfo(lobbyID, addr, listener, id, data)
	case "DeleteLobby":
		C.DeleteLobby(lobbyID, addr, listener, id, data)
	case "OK":
		C.SendOK(addr, listener, id)
	}
}

func Server(){
	//Checks the env to get the adress to listen to. 

	data := new(C.Data)
	data.PreparingLobbies = make(map[string][]*net.UDPAddr)
	data.ReadyLobbies = make(map[string][]*net.UDPAddr)
	data.ExactClient = make(map[string]string)

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
		buff := make([]byte, 4096)
		var addr *net.UDPAddr
		var num int
		for{
			num, addr, err = listener.ReadFromUDP(buff)
			if !((err != nil && num == 0) || num == 0){
				break
			}
		}
		ListenToCommands(string(buff), listener, addr, data)
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
