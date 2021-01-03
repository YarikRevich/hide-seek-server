package MessageParser

import (
	"errors"
	"strings"
)

//Main checkers of request types

type reqCheck struct {
	//Struct for messageparsering

	//Contains comming request
	request string
}

func (r reqCheck) checkReq(reqtype string) bool {
	//Parsers comming request and returns its type gotten from body

	parser := func(splittedreq string)bool{
		parsedText := strings.Split(splittedreq, "///")
		if parsedText[0] == reqtype {
			return true
		}	
		return false
	}
	splittedReq := strings.Split(r.request, "_")
	if len(splittedReq) == 1{
		return parser(splittedReq[0])
	}
	return parser(splittedReq[1])
}

func (r reqCheck) getId()string{
	//Parsers id from request and return it

	splitted := strings.Split(r.request, "_")
	if len(splitted) == 1{
		return "1"
	}
	return splitted[0]
}

func UnparseMessage(message string) (string, string, error) {
	//Unparsers comming request and checks whether it is good
	//if it is good it returns request's id, its type and error

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
		return "" ,"OK", nil
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
