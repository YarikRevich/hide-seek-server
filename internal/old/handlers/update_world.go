package handlers

import (
	"fmt"

	// "github.com/YarikRevich/HideSeek-Server/internal/collection"
	"github.com/google/uuid"
)

// "time"

// "github.com/YarikRevich/HideSeek-Server/internal/collection"
// "github.com/google/uuid"

type UpdateWorldModel struct {
}

func UpdateWorldHandler(data interface{}) (interface{}, error) {
	worldID, err := uuid.Parse(data.(map[string]interface{})["World"].(map[string]interface{})["ID"].(string))
	if err != nil {
		return nil, err
	}
	fmt.Println(worldID)
	// pcID, err := uuid.Parse(parsedData["PCID"].(string))
	// if err != nil {
	// 	return nil, err
	// }

	// collection.World[worldID].

	// users := collection.World[worldID].Users
	// var r []interface{}
	// for _, user := range users {
	// 	v, ok := collection.User[user]
	// 	if !ok {
	// 		continue
	// 	}
	// 	r = append(r, v.Data)
	// }
	// collection.World[worldID].Cache.Add(time.Minute * 5)
	// return collection, nil
	// return collection.World[worldID].Data, nil
	return nil, nil
}
