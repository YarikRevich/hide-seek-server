package Components

import (
	"net"
	"fmt"
	"strings"
	"errors"
)

type Data struct{
	PreparingLobbies map[string][]*net.UDPAddr
	ReadyLobbies map[string][]*net.UDPAddr
	ExactClient map[string]string
}

func CreateLobby(lobbyID string, data *Data)error{
	for key := range data.PreparingLobbies{
		if key == lobbyID{
			return errors.New("such lobby already exists!")
		}
	}
	data.PreparingLobbies[lobbyID] = []*net.UDPAddr{}
	return nil
}

func LobbyExists(lobbyID string, data *Data)bool{
	for key := range data.PreparingLobbies{
		if key == lobbyID{
			return true
		}
	}
	return false
}

func AddToLobby(lobbyID string, request string, addr *net.UDPAddr, data *Data)error{
	value, _ := data.PreparingLobbies[lobbyID]
	if !LobbyExists(lobbyID, data){
		return errors.New("lobby does not exist")
	}
	value = append(value, addr)
	data.PreparingLobbies[lobbyID] = value
	data.ExactClient[addr.String()] = strings.Split(request, "~/")[1]
	return nil
}

func GetMembersInLobby(lobbyID string, data *Data)(string, []*net.UDPAddr, error){
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
	if len(lobbyMembers) == 0{
		return "", []*net.UDPAddr{}, errors.New("lobby does not exist")
	}
	lobbyMembersJoined := strings.Join(lobbyMembers, "//")
 	return lobbyMembersJoined, lobbyMembersAddrs, nil
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
	data.ExactClient[addr.String()] = strings.Split(request, "~/")[1]
}

func GetUsersInfo(lobbyID string, addr *net.UDPAddr, listener *net.UDPConn, data *Data){
	var usersInfo []string
	for _, value := range data.ReadyLobbies[lobbyID]{
		if value.String() != addr.String(){
			usersInfo = append(usersInfo, data.ExactClient[value.String()])
		}
	}
	formattedResp := fmt.Sprintf("GetUsersInfo///%s~/%s", lobbyID, strings.Join(usersInfo, "/::/"))
	listener.WriteTo([]byte(formattedResp), addr)
}