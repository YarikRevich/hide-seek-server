package ammo

import "github.com/YarikRevich/HideSeek-Server/internal/storage/local/common"

type AmmoCollection struct{}

func New() common.Collection {
	return new(AmmoCollection)
}
