package collection

import (
	"time"

	"github.com/google/uuid"
)

var User = map[uuid.UUID]struct{
	Cache time.Time
	Data interface{}
}{}
