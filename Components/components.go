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

func CreateLobby(lobbyID string, id string, data *Data)error{
	for key := range data.PreparingLobbies{
		if key == lobbyID{
			return errors.New(fmt.Sprintf("%s_such lobby already exists!", id))
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

func AddToLobby(lobbyID string, request string, addr *net.UDPAddr, id string, data *Data)error{
	value, _ := data.PreparingLobbies[lobbyID]
	if !LobbyExists(lobbyID, data){
		return errors.New(fmt.Sprintf("%s_lobby does not exist", id))
	}
	for _, user := range value{
		if addr.String() == user.String(){
			return nil
		}
	}
	value = append(value, addr)
	data.PreparingLobbies[lobbyID] = value
	data.ExactClient[addr.String()] = strings.Split(request, "~/")[1]
	return nil
}

func GetMembersInLobby(lobbyID string, id string, data *Data)(string, []*net.UDPAddr, error){
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
		return "", []*net.UDPAddr{}, errors.New(fmt.Sprintf("%s_lobby does not exist", id))
	}
	lobbyMembersJoined := strings.Join(lobbyMembers, "//")
 	return lobbyMembersJoined, lobbyMembersAddrs, nil
}

func SendLobbyMembers(listener *net.UDPConn, id string, lobbyMembers string, lobbyMembersAddrs []*net.UDPAddr){
	for _, memberAddr := range lobbyMembersAddrs{
		listener.WriteTo([]byte(fmt.Sprintf("%s_%s", lobbyMembers, id)), memberAddr)
	}
}

func ClosePreparingLobby(lobbyID string, data *Data){
	data.ReadyLobbies[lobbyID] = data.PreparingLobbies[lobbyID]
	delete(data.PreparingLobbies, lobbyID)
}

func UpdateUser(request string, addr *net.UDPAddr, data *Data){
	data.ExactClient[addr.String()] = strings.Split(request, "~/")[1]
}

func GetUsersInfoLobby(lobbyID string, addr *net.UDPAddr, listener *net.UDPConn, id string, data *Data)error{
	var usersInfo []string
	values, _ := data.PreparingLobbies[lobbyID]
	if len(values) == 0{
		return errors.New(fmt.Sprintf("%s_lobby does not exist", id))
	}

	for _, value := range values{
		usersInfo = append(usersInfo, data.ExactClient[value.String()])
	}
	formattedResp := fmt.Sprintf("%s_GetUsersInfoLobby///%s~/%s", id, lobbyID, strings.Join(usersInfo, "/::/"))
	listener.WriteTo([]byte(formattedResp), addr)
	return nil
}

func GetUsersInfo(lobbyID string, addr *net.UDPAddr, listener *net.UDPConn, id string, data *Data){
	var usersInfo []string
	for _, value := range data.ReadyLobbies[lobbyID]{
		if value.String() != addr.String(){
			usersInfo = append(usersInfo, data.ExactClient[value.String()])
		}
	}
	formattedResp := fmt.Sprintf("%s_GetUsersInfo///%s~/%s", id, lobbyID, strings.Join(usersInfo, "/::/"))
	listener.WriteTo([]byte(formattedResp), addr)
}

func DeleteLobby(lobbyID string, addr *net.UDPAddr, listener *net.UDPConn, id string, data *Data){
	delete(data.PreparingLobbies, lobbyID)
	listener.WriteTo([]byte(fmt.Sprintf("%s_1", id)), addr)
}

func SendOK(addr *net.UDPAddr, listener *net.UDPConn, id string){
	listener.WriteTo([]byte(fmt.Sprintf("%s_1", id)), addr)
}