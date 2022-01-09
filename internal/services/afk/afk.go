package afk

import (
	"time"

	"github.com/YarikRevich/hide-seek-server/internal/api/external-api/v1/proto"
	"github.com/YarikRevich/hide-seek-server/internal/storage"
)

//Checks if pc is afk by last activity metric
func IsAFK(worldId string, lastActivity int64) bool {
	world := storage.UseStorage().Local().Worlds().Find(worldId).(*proto.World)
	return int64(time.Since(time.Unix(lastActivity, 0)).Seconds()) >= world.GameSettings.AFKTimeout
}
