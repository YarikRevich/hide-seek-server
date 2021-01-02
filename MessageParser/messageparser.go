package MessageParser

import (
	"errors"
	"strings"
)

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

func (r reqCheck) getId()string{
	return strings.Split(r.request, "_")[0]
}

func UnparseMessage(message string) (string, string, error) {
	reqCheck := reqCheck{request: message}
	if reqCheck.checkReq("CreateLobby") {
		return reqCheck.getId(), "CreateLobby", nil
	}
	if reqCheck.checkReq("AddToLobby") {
		return reqCheck.getId(), "AddToLobby", nil
	}
	if reqCheck.checkReq("GetMembersInLobby") {
		return reqCheck.getId(), "GetMembersInLobby", nil
	}
	if reqCheck.checkReq("UpdateUser") {
		return reqCheck.getId(), "UpdateUser", nil
	}
	if reqCheck.checkReq("GetUsersInfo"){
		return reqCheck.getId(), "GetUsersInfo", nil
	}
	if reqCheck.checkReq("ClosePreparingLobby") {
		return reqCheck.getId(), "ClosePreparingLobby", nil
	}
	if reqCheck.checkReq("OK"){
		return reqCheck.getId() ,"OK", nil
	}
	if reqCheck.checkReq("GetUsersInfoLobby"){
		return reqCheck.getId(), "GetUsersInfoLobby", nil
	}
	if reqCheck.checkReq("DeleteLobby"){
		return reqCheck.getId(), "DeleteLobby", nil
	}
	if strings.Contains(reqCheck.request, "!_") {
		return "", "", nil
	}
	return "error", "error", errors.New("an error happened")
}
