package Utils

import (
	"net"
	"strings"
	"fmt"
)

func GetLobbyNum(request string) string {
	var notCleaned string
	parsedText := strings.Split(request, "///")
	if len(parsedText) > 1 {
		notCleaned = strings.Split(parsedText[1], "~")[0]
	} else {
		notCleaned = parsedText[0]
	}
	availableSymbols := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	cleanedID := []string{}
	for _, word := range notCleaned {
		for _, avsym := range availableSymbols {
			if string(word) == avsym {
				cleanedID = append(cleanedID, string(word))
			}
		}
	}
	return strings.Join(cleanedID, "")
}

func SendErrorResponse(err error, id string, listener *net.UDPConn, addr *net.UDPAddr) {
	formattedResp := fmt.Sprintf("%s_%s@%s", id, "error", err)
	listener.WriteTo([]byte(formattedResp), addr)
}
