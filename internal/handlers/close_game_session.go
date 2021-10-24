package handlers

import (
	"github.com/YarikRevich/HideSeek-Server/internal/collection"
	"github.com/google/uuid"
)

func CloseGameSession(data interface{}) (interface{}, error){
	worldID, err := uuid.Parse(data.(map[string]interface{})["ID"].(string))
	if err != nil{
		return nil, err
	}
	delete(collection.World, worldID)
	return nil, nil
}
