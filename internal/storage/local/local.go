package local

import (
	"github.com/YarikRevich/hide-seek-server/internal/storage/local/ammo"
	"github.com/YarikRevich/hide-seek-server/internal/storage/local/common"
	"github.com/YarikRevich/hide-seek-server/internal/storage/local/elements"
	"github.com/YarikRevich/hide-seek-server/internal/storage/local/maps"
	"github.com/YarikRevich/hide-seek-server/internal/storage/local/pcs"
	"github.com/YarikRevich/hide-seek-server/internal/storage/local/weapons"
	"github.com/YarikRevich/hide-seek-server/internal/storage/local/worlds"
)

type Local struct {
	weapons  common.Collection
	ammo     common.Collection
	pcs      common.Collection
	elements common.Collection
	maps     common.Collection
	worlds   common.Collection
}

func (l *Local) Weapons() common.Collection {
	return l.weapons
}

func (l *Local) PCs() common.Collection {
	return l.pcs
}

func (l *Local) Maps() common.Collection {
	return l.maps
}

func (l *Local) Ammo() common.Collection {
	return l.ammo
}

func (l *Local) Elements() common.Collection {
	return l.elements
}

func (l *Local) Worlds() common.Collection {
	return l.worlds
}

func New() *Local {
	return &Local{
		weapons:  weapons.New(),
		elements: elements.New(),
		maps:     maps.New(),
		pcs:      pcs.New(),
		ammo:     ammo.New(),
		worlds:   worlds.New(),
	}
}
