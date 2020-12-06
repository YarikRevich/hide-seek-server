package MessageParser

import (
	"strings"
)

// func checkClientExists(conn string)bool{
// 	for _, value := range addrs{
// 		if value.String() == conn{
// 			return true
// 		}
// 	}
// 	return false
// }

//Main checkers of request types

type reqCheck struct {
	request string
}

func (r reqCheck) checkReq(reqtype string) bool {
	parsedText := strings.Split(r.request, "///")
	if parsedText[0] == reqtype {
		return true
	}
	return false
}

func UnparseMessage(message string) (string, int) {
	reqCheck := reqCheck{request: message}
	if reqCheck.checkReq("CreateLobby") {
		return "CreateLobby", 0
	}
	if reqCheck.checkReq("AddToLobby") {
		return "AddToLobby", 0
	}
	if reqCheck.checkReq("GetMembersInLobby") {
		return "GetMembersInLobby", 0
	}
	if reqCheck.checkReq("UpdateUser") {
		return "UpdateUser", 0
	}
	if reqCheck.checkReq("GetUsersInfo"){
		return "GetUsersInfo", 0
	}
	if reqCheck.checkReq("ClosePreparingLobby") {
		return "ClosePreparingLobby", 0
	}
	if strings.Contains(reqCheck.request, "!_") {
		return "", 0
	}
	return "error", 2
}
