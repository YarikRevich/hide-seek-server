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
		err := C.AddToLobby(lobbyID, request, addr, id, data)
		if err != nil {
			Utils.SendErrorResponse(err, id, listener, addr)
		}else{
			C.SendOK(addr, listener, id)
		}
	case "GetMembersInLobby":
		lobbyMembers, lobbyMembersAddrs, err := C.GetMembersInLobby(lobbyID, id, data)
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

//~~~~~~~~Main server part~~~~~~~~


func Server(){
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
		buff := make([]byte, 2048)
		_, addr, err := listener.ReadFromUDP(buff)
		if err != nil{
			continue
		}
		ListenToCommands(string(buff), listener, addr, data)
	}
}

func DrawWelcomeMessage(){
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
