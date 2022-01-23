package registermanager

import (
	"github.com/YarikRevich/hide-seek-server/internal/monitoring/register_manager/registers/common"
	"github.com/YarikRevich/hide-seek-server/internal/monitoring/register_manager/registers/gauge"
	"github.com/YarikRevich/hide-seek-server/internal/monitoring/register_manager/registers/histogram"
)

type RegisterManager struct {
	worldscounter  common.Register
	usagehistogram common.Register
	playerscounter common.Register
}

func (rm *RegisterManager) WorldsGauge() common.Register {
	return rm.worldscounter
}

func (rm *RegisterManager) PlayersGauge() common.Register {
	return rm.playerscounter
}

func (rm *RegisterManager) UsageHistogram() common.Register {
	return rm.usagehistogram
}

func New() *RegisterManager {
	return &RegisterManager{
		worldscounter:  gauge.New("worlds"),
		usagehistogram: histogram.New("usage"),
		playerscounter: gauge.New("players"),
	}
}
