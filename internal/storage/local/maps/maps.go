package maps

import "github.com/YarikRevich/HideSeek-Server/internal/storage/local/common"

type MapsCollection struct{}

func New() common.Collection {
	return new(MapsCollection)
}
