package MessageParser

import (
	"log"
	"encoding/json"
	"github.com/liip/sheriff"
)

type Parser interface{

	Parse(Message)[]byte
	Unparse([]byte)[]*Message
	ParseSeveral([]*Message)[]byte
}

type Message struct{
	//Shows the action of request. Namely here is the table of codes
	//10: "Preparing lobby is created"
	//20: "User is added to lobby"
	//30: "Members of lobby are shown"
	//40: "Data of user is updated"
	//50: "Preparing lobby is closed and ready one is created"
	//60: "Users' info in a corrisponding preparing lobby is given"
	//70: "Users' info in a corrisponding ready lobby is given"
	//80: "Lobby is deleted"
	//90: "Health is updated"
	//200: "OK response"
	//Error responses
	//500: "Such lobby does not exist"
	//501: "User is already added"
	//502: "Lobby is not found"
	//503: "Such lobby already exists"
	Error           string	`groups:"users"`
	Type            string	`groups:"users"`
	Pos             struct{
		X int	`groups:"users"`
		Y int	`groups:"users"`
	}	`groups:"users"`
	GameInfo        struct{ 
		Health int	`groups:"users"`
		WeaponName string `groups:"users"`
		WeaponRadius int `groups:"users"`
 	}	`groups:"users"`
	PersonalInfo    struct{
		LobbyID string	`groups:"users"`
		Username string	`groups:"users"`
		HeroPicture string	`groups:"users"`
	}	`groups:"users"`
	Animation       struct{
		HeroIconUpdation int	`groups:"users"`
		HeroIconUpdationDelay int `groups:"users"`
		WeaponIconUpdation int  `groups:"users"`
		WeaponIconUpdationDelay int `groups:"users"`
		CurrentFrame int	`groups:"users"`
		CurrentFrameMatrix []float64	`groups:"users"`
	}	`groups:"users"`
	Networking      struct{	
		Index int	`groups:"users"`
	}	`groups:"users"`
	Context         struct{
		Additional      []string	`groups:"users"`
	}	`groups:"users"`
}

func (m *Message) Parse(message Message)[]byte{
	b, err := json.Marshal(message)
	if err != nil{
		log.Fatalln(err)
	}
	return b
}

func (m *Message) Unparse(message []byte)[]*Message{
	var unparsed []*Message
	err := json.Unmarshal(message, &unparsed)
	if err != nil{
		log.Fatalln(err)
	}
	return unparsed
}

func (m *Message) ParseSeveral(messages []*Message)[]byte{

	o := &sheriff.Options{
		Groups: []string{"users"},
	}
	data, err := sheriff.Marshal(o, messages)
	if err != nil{
		log.Fatalln(err)
	}
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil{
		log.Fatalln(err)
	}
	return b

}