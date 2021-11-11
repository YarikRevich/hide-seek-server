package handlers

// import (
// 	"time"

// 	"github.com/YarikRevich/HideSeek-Server/internal/collection"
// 	"github.com/google/uuid"
// )

// type RegUserModelStub struct {
// 	UserID uuid.UUID `json:"ID"`
// }

// func RegUser(data interface{}) (interface{}, error) {
// 	id, err := uuid.Parse(data.(map[string]interface{})["ID"].(string))
// 	if err != nil {
// 		return nil, err
// 	}
// 	collection.User[id] = struct{Cache time.Time; Data interface{}}{
// 		Cache: time.Now().Add(time.Minute * 5), Data: data,
// 	}
// 	return nil, nil
// }
