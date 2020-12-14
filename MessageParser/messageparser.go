package MessageParser

import (
	"strings"
	"errors"
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

func UnparseMessage(message string) (string, error) {
	reqCheck := reqCheck{request: message}
	if reqCheck.checkReq("CreateLobby") {
		return "CreateLobby", nil
	}
	if reqCheck.checkReq("AddToLobby") {
		return "AddToLobby", nil
	}
	if reqCheck.checkReq("GetMembersInLobby") {
		return "GetMembersInLobby", nil
	}
	if reqCheck.checkReq("UpdateUser") {
		return "UpdateUser", nil
	}
	if reqCheck.checkReq("GetUsersInfo"){
		return "GetUsersInfo", nil
	}
	if reqCheck.checkReq("ClosePreparingLobby") {
		return "ClosePreparingLobby", nil
	}
	if strings.Contains(reqCheck.request, "!_") {
		return "", nil
	}
	return "error", errors.New("an error happened!")
}
