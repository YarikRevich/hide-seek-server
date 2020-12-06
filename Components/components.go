package Components

import (
	"net"
	"fmt"
	"strings"
)

type Data struct{
	PreparingLobbies map[string][]*net.UDPAddr
	ReadyLobbies map[string][]*net.UDPAddr
	ExactClient map[*net.UDPAddr]string
}

func CreateLobby(lobbyID string, data *Data)int{
	for key := range data.PreparingLobbies{
		if key == lobbyID{
			return 1
		}
	}
	data.PreparingLobbies[lobbyID] = []*net.UDPAddr{}
	return 0
}

func LobbyExists(lobbyID string, data *Data)bool{
	for key := range data.PreparingLobbies{
		if key == lobbyID{
			return true
		}
	}
	return false
}

func AddToLobby(lobbyID string, request string, addr *net.UDPAddr, data *Data)int{
	value, _ := data.PreparingLobbies[lobbyID]
	if !LobbyExists(lobbyID, data){
		return 1
	}
	value = append(value, addr)
	data.PreparingLobbies[lobbyID] = value
	data.ExactClient[addr] = strings.Split(request, "~/")[1]
	return 0
}

func GetMembersInLobby(lobbyID string, data *Data)(string, []*net.UDPAddr){
	lobbyMembers := []string{}
	lobbyMembersAddrs := []*net.UDPAddr{}
	for key := range data.PreparingLobbies{
		if strings.Contains(key, lobbyID){
			for _, value := range data.PreparingLobbies[key]{
				lobbyMembers = append(lobbyMembers, value.String())
				lobbyMembersAddrs = append(lobbyMembersAddrs, value)
			}
		}
	}
	lobbyMembersJoined := strings.Join(lobbyMembers, "//")
 	return lobbyMembersJoined, lobbyMembersAddrs
}

func SendLobbyMembers(listener *net.UDPConn, lobbyMembers string, lobbyMembersAddrs []*net.UDPAddr){
	for _, memberAddr := range lobbyMembersAddrs{
		listener.WriteTo([]byte(lobbyMembers), memberAddr)
	}
}

func ClosePreparingLobby(lobbyID string, data *Data){
	data.ReadyLobbies[lobbyID] = data.PreparingLobbies[lobbyID]
	delete(data.PreparingLobbies, lobbyID)
}

func UpdateUser(request string, addr *net.UDPAddr, data *Data){
	data.ExactClient[addr] = strings.Split(request, "~/")[1]
}

func GetUsersInfo(lobbyID string, addr *net.UDPAddr, listener *net.UDPConn, data *Data){
	var usersInfo []string
	for _, value := range data.ReadyLobbies[lobbyID]{
		usersInfo = append(usersInfo, data.ExactClient[value])
	}
	formattedResp := fmt.Sprintf("GetUsersInfo///%s~/%s", lobbyID, strings.Join(usersInfo, "/::/"))
	fmt.Println(formattedResp)
	listener.WriteTo([]byte(formattedResp), addr)
	
}