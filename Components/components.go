package Components

import (
	"GameServer/MessageParser"
	"fmt"
	"log"
	"net"
	"strconv"
)

type Data struct{
	//Main data struct for the futher data exchange

	//Contains all the preparing lobbies and each one's members' addres
	PreparingLobbies map[string][]*net.UDPAddr

	//Contains all the ready lobbies and each one's members' addres
	ReadyLobbies map[string][]*net.UDPAddr

	//Contains information about exact member of game
	//Each cell consists of for parts namely: the first one
	//is user's position, the next one is the info of a game
	//process, then goes some personal information and finally
	//some metadata for animation work
	ExactClient map[string]*MessageParser.Message
}

func CreateLobby(result *MessageParser.Message, listener net.Conn, addr net.Addr, data *Data)string{
	//When user wants to create a lobby he uses
	//a create request and it uses this one

	for key := range data.PreparingLobbies{
		if key == result.PersonalInfo.LobbyID{
			return "502"
		}
	}
	data.PreparingLobbies[result.PersonalInfo.LobbyID] = []*net.UDPAddr{}

	return "10"
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

func AddToLobby(result *MessageParser.Message, listener *net.UDPConn, addr *net.UDPAddr, data *Data)string{
	//When user creates or joins any lobby he should
	//be added to prepring lobbies' map. That's why
	//this one adds user to it

	value, _ := data.PreparingLobbies[result.PersonalInfo.LobbyID]
	if !LobbyExists(result.PersonalInfo.LobbyID, data){
		return "500"
	}
	for _, user := range value{
		if addr.String() == user.String(){
			return "501"
		}
	}
	value = append(value, addr)
	data.PreparingLobbies[result.PersonalInfo.LobbyID] = value
	data.ExactClient[addr.String()] = result

	return "20"
}

func SendLobbyMembers(listener *net.UDPConn, id int, lobbyMembers string, lobbyMembersAddrs []*net.UDPAddr){
	//WARNING: it is deprecated!
	
	for _, memberAddr := range lobbyMembersAddrs{
		listener.WriteTo([]byte(fmt.Sprintf("%d_%s", lobbyMembers, id)), memberAddr)
	}
}

func ClosePreparingLobby(lobbyID string, data *Data)string{
	//When lobby-host wants to start the game he sends a request
	//to close preparing lobby. It deletes lobby from preparing
	//map and inserts it to the ready one

	data.ReadyLobbies[lobbyID] = data.PreparingLobbies[lobbyID]
	delete(data.PreparingLobbies, lobbyID)
	return "50"
}

func UpdateUser(request *MessageParser.Message, addr *net.UDPAddr, data *Data)string{
	//Updates internal user's data due to new comming data

	data.ExactClient[addr.String()] = request
	return "40"
}

func GetUsersInfoPrepLobby(result *MessageParser.Message, addr *net.UDPAddr, listener *net.UDPConn, data *Data)([]*MessageParser.Message, string){
	//It is used when user is at the lobbywaitroom and he wants to get all
	//the users in it and it even can show new users
	
	var usersInfo []*MessageParser.Message
	prep_values, _ := data.PreparingLobbies[result.PersonalInfo.LobbyID]

	if len(prep_values) == 0{
		usersInfo = append(usersInfo, result)
		return usersInfo, "502"
	}

	for _, value := range prep_values{
		usersInfo = append(usersInfo, data.ExactClient[value.String()])
	}

	return usersInfo, "60"
}

func GetUsersInfoReadyLobby(result *MessageParser.Message, addr *net.UDPAddr, listener *net.UDPConn, data *Data)([]*MessageParser.Message, string){
	//When lobby-host started a game each user gets information about others
	//It sends all the newest information about all the users in lobby
	
	var usersInfo []*MessageParser.Message

	ready_values, _ := data.ReadyLobbies[result.PersonalInfo.LobbyID]
	
	if len(ready_values) == 0{
		usersInfo = append(usersInfo, result)
		return usersInfo, "502"
	}

	for _, value := range ready_values{
		usersInfo = append(usersInfo, data.ExactClient[value.String()])
	}
	return usersInfo, "70"
}

func UpdateUsersHealth(result *MessageParser.Message, data *Data)([]*MessageParser.Message, string){
	//Updates user's info about his health points

	var usersInfo []*MessageParser.Message
	ready_values, _ := data.ReadyLobbies[result.PersonalInfo.LobbyID]
	for _, value := range ready_values{
		if data.ExactClient[value.String()].PersonalInfo.Username == result.Context.Additional[0]{

			hp, err := strconv.Atoi(result.Context.Additional[1])
			if err != nil{
				log.Fatalln(err)
			}
			data.ExactClient[value.String()].GameInfo.Health -= hp
			usersInfo = append(usersInfo, data.ExactClient[value.String()])
		}
	}
	return usersInfo, "90"

}

func DeleteLobby(lobbyID string, addr *net.UDPAddr, listener *net.UDPConn, id int, data *Data)string{
	//Deletes lobby from preparing map
	
	delete(data.PreparingLobbies, lobbyID)
	delete(data.ReadyLobbies, lobbyID)
	return "80"
}

func SendAnswerS(results []*MessageParser.Message, code string, listener *net.UDPConn, addr *net.UDPAddr){
	
	if results != nil{
		for _, value := range results{
			value.Error = code
		}
		parser := MessageParser.Parser(new(MessageParser.Message))
		b := parser.ParseSeveral(results)
		fmt.Println(string(b), addr.String())
		listener.WriteTo(b, addr)
	}
}

func SendAnswer(result *MessageParser.Message, code string, listener *net.UDPConn, addr *net.UDPAddr){
	//Main server wants to get all the available sub-servers
	//That's why if server is here it sends ok response which
	//contains only '200' as a good sign

	result.Error = code
	parser := MessageParser.Parser(new(MessageParser.Message))
	b := parser.Parse(*result)
	listener.WriteTo(b, addr)
}

func UpdateExactUser(result *MessageParser.Message, addr *net.UDPAddr, data *Data){
	if result.Type != "OK"{
		user, ok := data.ExactClient[addr.String()]
		if ok{
			user.Error = result.Error
			user.Type = result.Type
			user.Pos = result.Pos
			user.PersonalInfo = result.PersonalInfo
			user.Animation = result.Animation
			user.Networking = result.Networking
		}
	}
}