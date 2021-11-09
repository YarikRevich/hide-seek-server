package collection

import (
	"time"

	"github.com/google/uuid"
)

var World = map[uuid.UUID]struct {
	Data interface{}
	Cache time.Time
	PCs []interface{}
	Elements interface{}
}{}


