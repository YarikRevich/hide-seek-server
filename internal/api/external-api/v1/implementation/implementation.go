package implementation

import (
	"context"
	"sort"

	"github.com/YarikRevich/hide-seek-server/internal/services/afk"
	"github.com/YarikRevich/hide-seek-server/internal/storage"

	"github.com/YarikRevich/hide-seek-server/internal/api/external-api/v1/proto"
	"github.com/YarikRevich/hide-seek-server/tools/utils"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type ExternalServerService struct {
	proto.UnimplementedExternalServerServiceServer
}

func (h *ExternalServerService) InsertOrUpdateWorld(ctx context.Context, world *proto.World) (*empty.Empty, error) {
	storage.UseStorage().Local().Worlds().InsertOrUpdate(world.GetId(), world)
	return &empty.Empty{}, nil
}

func (a *ExternalServerService) InsertOrUpdateMap(ctx context.Context, worldMap *proto.Map) (*empty.Empty, error) {
	storage.UseStorage().Local().Maps().InsertOrUpdate(worldMap.Base.GetId(), worldMap)
	return &empty.Empty{}, nil
}

func (a *ExternalServerService) InsertOrUpdatePC(ctx context.Context, r *proto.PC) (*empty.Empty, error) {
	l := storage.UseStorage().Local()
	l.PCs().InsertOrUpdate(r.Base.Parent.Parent.GetId(), r)
	if afk.IsAFK(r.Base.Parent.Parent.GetId(), r.LastActivity) {
		l.PCs().Delete(r.Base.GetId())
		l.Cooldown().InsertOrUpdate(r.Base.Parent.Parent.GetId(), r.Base.GetId())
	}
	return &empty.Empty{}, nil
}

func (a *ExternalServerService) InsertOrUpdateAmmo(ctx context.Context, r *proto.Ammo) (*empty.Empty, error) {
	storage.UseStorage().Local().Ammo().InsertOrUpdate(r.Base.GetId(), r)
	return &empty.Empty{}, nil
}

func (a *ExternalServerService) InsertOrUpdateCooldown(ctx context.Context, r *proto.InsertOrUpdateCooldownRequest) (*empty.Empty, error) {
	storage.UseStorage().Local().Cooldown().InsertOrUpdate(r.GetWorldId(), r.GetPcId())
	return &empty.Empty{}, nil
}

//Assigns random spawns to the pcs set by this world
func (a *ExternalServerService) AssignRandomSpawnsToPCs(ctx context.Context, r *proto.AssignRandomSpawnsToPCsRequest) (*empty.Empty, error) {
	l := storage.UseStorage().Local()
	worldMap := l.Maps().Find(r.GetWorldId()).(*proto.Map)
	pcs := l.PCs().Find(r.GetWorldId()).([]*proto.PC)
	randomSpawnsIndicies := utils.GetSequenceOfRandomInt(len(pcs), 0, len(pcs)-1)
	for i := 0; i < len(pcs)-1; i++ {
		pcs[i].Base.Spawn = worldMap.GetSpawns()[randomSpawnsIndicies[i]]
	}
	return nil, nil
}

//Deletes world object and its related objects
func (a *ExternalServerService) DeleteWorld(ctx context.Context, r *wrappers.StringValue) (*empty.Empty, error) {
	s := storage.UseStorage().Local()
	worldID := r.GetValue()

	s.Worlds().Delete(worldID)
	s.Maps().Delete(worldID)
	s.Elements().Delete(worldID)
	s.PCs().Delete(worldID)
	s.Weapons().Delete(worldID)
	s.Ammo().Delete(worldID)
	return &empty.Empty{}, nil
}

func (a *ExternalServerService) DeletePC(ctx context.Context, r *wrappers.StringValue) (*empty.Empty, error) {
	storage.UseStorage().Local().PCs().Delete(r.Value)
	return &empty.Empty{}, nil
}

func (a *ExternalServerService) DeleteCooldown(ctx context.Context, r *proto.DeleteCooldownRequest) (*empty.Empty, error) {
	storage.UseStorage().Local().Cooldown().Delete([2]string{r.GetWorldId(), r.GetPcId()})
	return &empty.Empty{}, nil
}

func (a *ExternalServerService) FindWorldObjects(ctx context.Context, worldID *wrappers.StringValue) (*proto.GetWorldResponse, error) {
	s := storage.UseStorage().Local()

	pcs := s.PCs().Find(worldID.Value).([]*proto.PC)
	copyPCs := make([]*proto.PC, len(pcs), cap(pcs))
	copy(copyPCs, pcs)

	sort.Slice(copyPCs, func(i, j int) bool {
		return copyPCs[i].GetLobbyNumber() < copyPCs[j].GetLobbyNumber()
	})

	return &proto.GetWorldResponse{
		World:    s.Worlds().Find(worldID.GetValue()).(*proto.World),
		Weapons:  s.Weapons().Find(worldID.GetValue()).([]*proto.Weapon),
		Elements: s.Elements().Find(worldID.GetValue()).([]*proto.Element),
		Ammos:    s.Ammo().Find(worldID.GetValue()).([]*proto.Ammo),
		PCs:      copyPCs,
	}, nil
}

func (a *ExternalServerService) FindCooldown(ctx context.Context, r *proto.FindCooldownRequest) (*wrappers.BoolValue, error) {
	storage.UseStorage().Local().Cooldown().Find([2]string{r.GetWorldId(), r.GetPcId()})
	return &wrappers.BoolValue{}, nil
}

func NewExternalServerService() *ExternalServerService {
	return new(ExternalServerService)
}
