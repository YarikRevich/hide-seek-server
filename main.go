package main

import (
	"net"
	"fmt"
	"github.com/mbndr/figlet4go"
	"GameServer/MessageParser"
	"GameServer/Utils"
	C "GameServer/Components"
)

func ListenToCommands(request string, listener *net.UDPConn, addr *net.UDPAddr, data *C.Data){
	unparsedReq, err := MessageParser.UnparseMessage(request)
	if err != nil {
		Utils.SendErrorResponse(err, listener, addr)
	}
	lobbyID := Utils.GetLobbyNum(request)
	switch unparsedReq{
	case "CreateLobby":
		err := C.CreateLobby(lobbyID, data)
		if err != nil {
			Utils.SendErrorResponse(err, listener, addr)
		}
	case "AddToLobby":
		err := C.AddToLobby(lobbyID, request, addr, data)
		if err != nil {
			Utils.SendErrorResponse(err, listener, addr)
		}
	case "GetMembersInLobby":
		lobbyMembers, lobbyMembersAddrs, err := C.GetMembersInLobby(lobbyID, data)
		if err != nil {
			Utils.SendErrorResponse(err, listener, addr)
		}
		C.SendLobbyMembers(listener, lobbyMembers, lobbyMembersAddrs)
	case "UpdateUser":
		C.UpdateUser(request, addr, data)
	case "ClosePreparingLobby":
		C.ClosePreparingLobby(lobbyID, data)
	case "GetUsersInfo":
		C.GetUsersInfo(lobbyID, addr, listener, data)
	}
}

//~~~~~~~~Main server part~~~~~~~~


func Server1(stopChan chan bool){
	data := new(C.Data)
	data.PreparingLobbies = make(map[string][]*net.UDPAddr)
	data.ReadyLobbies = make(map[string][]*net.UDPAddr)
	data.ExactClient = make(map[string]string)
	addr := net.UDPAddr{
		Port: 9001,
		IP: net.ParseIP("127.0.0.1"),
	}
	listener, _ := net.ListenUDP("udp", &addr)
	for{
		buff := make([]byte, 2048)
		_, addr, err := listener.ReadFromUDP(buff)
		if err != nil{
			continue
		}
		ListenToCommands(string(buff), listener, addr, data)
	}
	stopChan <- true
}

func Server2(stopChan chan bool){
	addr := net.UDPAddr{
		Port: 9002,
		IP: net.ParseIP("127.0.0.1"),
	}
	listener, _ := net.ListenUDP("udp", &addr)
	for{
		buff := make([]byte, 32)
		_, _, err := listener.ReadFromUDP(buff)
		if err != nil{
			continue
		}
		println(string(buff))
	}
	stopChan <- true
}

func DrawWelcomeMessage(){
	renderer := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorGreen,
	}
	text, _ := renderer.RenderOpts("Hide&Seek server!", options)
	fmt.Print(text)
	fmt.Println("This server just process players...")
}

func main(){
	ServerStopChan1 := make(chan bool)
	ServerStopChan2 := make(chan bool)
	go Server1(ServerStopChan1)
	go Server2(ServerStopChan2)
	DrawWelcomeMessage()
	select {
	case <- ServerStopChan1:
		return
	case <- ServerStopChan2:
		return
	}
}
