package registermanager

import (
	"github.com/YarikRevich/HideSeek-Server/internal/monitoring/register_manager/registers/common"
	"github.com/YarikRevich/HideSeek-Server/internal/monitoring/register_manager/registers/counter"
)

type RegisterManager struct {
	createdworldscounter common.Register
}

func (rm *RegisterManager) CreatedWorldsCounter() common.Register {
	return rm.createdworldscounter
}

func New() *RegisterManager {
	return &RegisterManager{
		createdworldscounter: counter.New("createdworlds"),
	}
}
