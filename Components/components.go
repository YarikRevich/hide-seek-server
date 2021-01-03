package Components

import (
	"net"
	"fmt"
	"errors"
	"strings"
)

type Data struct{
	//Main data struct for the futher data exchange

	//Contains all the preparing lobbies and each one's members' addres
	PreparingLobbies map[string][]*net.UDPAddr

	//Contains all the ready lobbies and each one's members' addres
	ReadyLobbies map[string][]*net.UDPAddr

	//Contains information about exact member of game
	ExactClient map[string]string
}

func CreateLobby(lobbyID string, id string, data *Data)error{
	//When user wants to create a lobby he uses
	//a create request and it uses this one

	for key := range data.PreparingLobbies{
		if key == lobbyID{
			return errors.New(fmt.Sprintf("%s_such lobby already exists!", id))
		}
	}
	data.PreparingLobbies[lobbyID] = []*net.UDPAddr{}
	return nil
}

func LobbyExists(lobbyID string, data *Data)bool{
	//Checks whether preparing lobby exists

	for key := range data.PreparingLobbies{
		if key == lobbyID{
			return true
		}
	}
	return false
}

func AddToLobby(lobbyID string, request string, addr *net.UDPAddr, data *Data)error{
	//When user creates or joins any lobby he should
	//be added to prepring lobbies' map. That's why
	//this one adds user to it

	value, _ := data.PreparingLobbies[lobbyID]
	if !LobbyExists(lobbyID, data){
		return errors.New("lobby does not exist")
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

func SendLobbyMembers(listener *net.UDPConn, id string, lobbyMembers string, lobbyMembersAddrs []*net.UDPAddr){
	//WARNING: it is deprecated!
	
	for _, memberAddr := range lobbyMembersAddrs{
		listener.WriteTo([]byte(fmt.Sprintf("%s_%s", lobbyMembers, id)), memberAddr)
	}
}

func ClosePreparingLobby(lobbyID string, data *Data){
	//When lobby-host wants to start the game he sends a request
	//to close preparing lobby. It deletes lobby from preparing
	//map and inserts it to the ready one

	data.ReadyLobbies[lobbyID] = data.PreparingLobbies[lobbyID]
	delete(data.PreparingLobbies, lobbyID)
}

func UpdateUser(request string, addr *net.UDPAddr, data *Data){
	//Updates internal user's data due to new comming data

	data.ExactClient[addr.String()] = strings.Split(request, "~/")[1]
}

func GetUsersInfoLobby(lobbyID string, addr *net.UDPAddr, listener *net.UDPConn, id string, data *Data)error{
	//It is used when user is at the lobbywaitroom and he wants to get all
	//the users in it and it even can show new users
	
	var usersInfo []string
	values, _ := data.PreparingLobbies[lobbyID]
	if len(values) == 0{
		return errors.New("lobby does not exist")
	}

	for _, value := range values{
		usersInfo = append(usersInfo, data.ExactClient[value.String()])
	}
	formattedResp := fmt.Sprintf("%s_GetUsersInfoLobby///%s~/%s", id, lobbyID, strings.Join(usersInfo, "/::/"))
	listener.WriteTo([]byte(formattedResp), addr)
	return nil
}

func GetUsersInfo(lobbyID string, addr *net.UDPAddr, listener *net.UDPConn, id string, data *Data){
	//When lobby-host started a game each user gets information about others
	//It sends all the newest information about all the users in lobby
	
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
	//Deletes lobby from preparing map
	
	delete(data.PreparingLobbies, lobbyID)
	listener.WriteTo([]byte(fmt.Sprintf("%s_1", id)), addr)
}

func SendOK(addr *net.UDPAddr, listener *net.UDPConn, id string){
	//Main server wants to get all the available sub-servers
	//That's why if server is here it sends ok response which
	//contains only '1' as a good sign

	listener.WriteTo([]byte(fmt.Sprintf("%s_1", id)), addr)
}