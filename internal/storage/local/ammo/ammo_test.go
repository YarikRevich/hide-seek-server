package ammo

import (
	"testing"

	"github.com/YarikRevich/hide-seek-server/internal/api/external-api/v1/proto"
	"github.com/franela/goblin"
	"github.com/google/uuid"
)

func TestAmmoCollection(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Test AmmoCollection", func() {
		ac := &AmmoCollection{
			ammo: make(map[string][]*proto.Ammo)}

		tests := []*proto.Ammo{
			{
				Base: &proto.Base{
					Id: uuid.NewString(),
				},
				Direction: 1,
			},
			{
				Base: &proto.Base{
					Id: uuid.NewString(),
				},
				Direction: 2,
			},
		}

		g.It("Test InsertOrUpdate", func() {
			for _, v := range tests {
				ac.InsertOrUpdate(v.Base.Id, v)
			}
		})

		g.It("Test Find", func() {
			for _, v := range tests {
				r := ac.Find(v.Base.Id)
				a, ok := r.(*proto.Ammo)
				g.Assert(ok).IsTrue()
				g.Assert(a).IsNotNil()
			}
		})

		g.It("Test Delete", func() {
			for _, v := range tests {
				ac.Delete(v.Base.Id)
				r := ac.Find(v.Base.Id)
				a, ok := r.(*proto.Ammo)
				g.Assert(ok).IsFalse()
				g.Assert(a).IsNil()
			}
		})
	})
}
