package handlers

import (
	"errors"

	"github.com/YarikRevich/HideSeek-Server/internal/collection"
	"github.com/google/uuid"
)

func RegWorld(data interface{}) (interface{}, error){
	worldData, ok := data.(map[string]interface{})["World"]
	if !ok{
		return nil, errors.New("world data is nil")
	}
	worldID, err := uuid.Parse(worldData.(map[string]interface{})["ID"].(string))
	if err != nil{
		return nil, err
	}
	userID, err := uuid.Parse(data.(map[string]interface{})["PC"].(map[string]interface{})["ID"].(string))
	if err != nil{
		return nil, err
	}
	collection.World[worldID] = struct{Data interface{}; Users []uuid.UUID}{
		Data: worldData, Users: []uuid.UUID{userID},
	}
	return nil, nil
}

