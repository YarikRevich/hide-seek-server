package weapons

import "github.com/YarikRevich/HideSeek-Server/internal/storage/local/common"

//	storage map[uuid.UUID]interface{}
type WeaponsCollection struct{}

func New() common.Collection {
	return new(WeaponsCollection)
}
