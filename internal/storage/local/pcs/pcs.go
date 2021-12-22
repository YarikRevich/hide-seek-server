package pcs

import "github.com/YarikRevich/HideSeek-Server/internal/storage/local/common"

type PCsCollection struct{}

func New() common.Collection {
	return new(PCsCollection)
}
