package collection

import "github.com/google/uuid"

var World = map[uuid.UUID]struct {
	Data interface{}
	Users []uuid.UUID
}{}
