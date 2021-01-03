package Utils

import (
	"fmt"
	"net"
	"strings"
	"github.com/galsondor/go-ascii"
)

func GetLobbyNum(request string) string {
	//Parses lobby id from body request for the futher actions

	var lobbyNum string
	parsedText := strings.Split(request, "///")
	if len(parsedText) > 1 {
		lobbyNum = strings.Split(parsedText[1], "~")[0]
	} else {
		lobbyNum = parsedText[0]
	}
	cleanedLobbyNum := []string{}
	for _, word := range lobbyNum {
		if ascii.IsPrint(byte(word)){
			cleanedLobbyNum = append(cleanedLobbyNum, string(word))
		}
	}
	return strings.Join(cleanedLobbyNum, "")
}

func SendErrorResponse(err error, id string, listener *net.UDPConn, addr *net.UDPAddr) {
	//If someting went wrong server uses it to send error response

	formattedResp := fmt.Sprintf("%s_%s@%s", id, "error", err)
	listener.WriteTo([]byte(formattedResp), addr)
}
