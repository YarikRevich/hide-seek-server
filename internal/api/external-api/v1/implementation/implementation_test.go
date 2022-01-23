package implementation_test

import (
	"context"
	"testing"
	"time"

	external "github.com/YarikRevich/hide-seek-server/internal/api/external-api/v1/proto"
	"github.com/YarikRevich/hide-seek-server/internal/server"
	"github.com/YarikRevich/hide-seek-server/tools/utils/testclient"
	"github.com/franela/goblin"
	"github.com/google/uuid"
)

func TestExternalServerService_InsertOrUpdateWorld(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Testing InsertOrUpdateWorld", func() {
		g.Before(func() {
			go server.Run()
		})
		time.Sleep(3 * time.Second)
		ce, _ := testclient.Run()
		g.It("Testing if correct", func() {
			r, err := ce.InsertOrUpdateWorld(context.Background(), &external.World{
				Id: uuid.New().String(),
				GameSettings: &external.GameSettings{
					IsGameStarted: false,
					IsWorldExist:  false,
					AFKTimeout:    10,
				},
			})
			g.Assert(err).IsNil()
			g.Assert(r.String()).IsZero()
		})
	})
}

func BenchmarkExternalServerService_InsertOrUpdateWorld(t *testing.B) {
	go server.Run()
	time.Sleep(3 * time.Second)
	ce, _ := testclient.Run()

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		t.StartTimer()
		_, err := ce.InsertOrUpdateWorld(context.Background(), &external.World{
			Id: uuid.New().String(),
			GameSettings: &external.GameSettings{
				IsGameStarted: false,
				IsWorldExist:  false,
				AFKTimeout:    10,
			},
		})
		if err != nil {
			t.Error(err)
		}
		t.StopTimer()
	}
}
