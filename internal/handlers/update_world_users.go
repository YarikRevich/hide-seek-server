package handlers

import (
	"github.com/YarikRevich/HideSeek-Server/internal/collection"
	"github.com/google/uuid"
)

func UpdateWorldUsersHandler(data interface{}) (interface{}, error) {
	worldID, err := uuid.Parse(data.(map[string]interface{})["ID"].(string))
	if err != nil{
		return nil, err
	}
	users := collection.World[worldID].Users
	var r []interface{}
	for _, user := range users {
		v, ok := collection.User[user]
		if !ok {
			continue
		}
		r = append(r, v)
	}
	return r, nil
}
