package handlers

import (
	"github.com/YarikRevich/HideSeek-Server/internal/collection"
	"github.com/google/uuid"
)

type RegUserModelStub struct {
	UserID uuid.UUID `json:"ID"`
}

func RegUser(data interface{}) (interface{}, error) {
	id, err := uuid.Parse(data.(map[string]interface{})["ID"].(string))
	if err != nil {
		return nil, err
	}
	collection.User[id] = data
	return nil, nil
}
