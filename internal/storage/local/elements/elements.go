package elements

import "github.com/YarikRevich/HideSeek-Server/internal/storage/local/common"

type ElementsCollection struct{}

func New() common.Collection {
	return new(ElementsCollection)
}
