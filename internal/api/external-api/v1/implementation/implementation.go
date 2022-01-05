package implementation

import (
	"context"
	"sort"

	"github.com/YarikRevich/hide-seek-server/internal/storage"

	"github.com/YarikRevich/hide-seek-server/internal/api/external-api/v1/proto"
	"github.com/YarikRevich/hide-seek-server/tools/utils"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type ExternalServerService struct {
	proto.UnimplementedExternalServerServiceServer
}

func (h *ExternalServerService) UpdateWorld(ctx context.Context, world *proto.World) (*empty.Empty, error) {
	storage.UseStorage().Local().Worlds().InsertOrUpdate(world.GetId(), world)
	return &empty.Empty{}, nil
}

func (a *ExternalServerService) UpdateMap(ctx context.Context, worldMap *proto.Map) (*empty.Empty, error) {
	storage.UseStorage().Local().Maps().InsertOrUpdate(worldMap.Base.Parent.GetId(), worldMap)
	return &empty.Empty{}, nil
}

//Assigns random spawns to the pcs set by this world
func (a *ExternalServerService) AssignRandomSpawnsToPCs(ctx context.Context, r *proto.AssignRandomSpawnsToPCsRequest) (*empty.Empty, error) {
	s := storage.UseStorage()
	worldMap := s.Local().Maps().Find(r.GetWorldId()).(*proto.Map)
	pcs := s.Local().PCs().Find(r.GetWorldId()).([]*proto.PC)
	randomSpawnsIndicies := utils.GetSequenceOfRandomInt(len(pcs), 0, len(pcs)-1)
	for i := 0; i < len(pcs)-1; i++ {
		pcs[i].Base.Spawn = worldMap.GetSpawns()[randomSpawnsIndicies[i]]
	}
	return nil, nil
}

func (a *ExternalServerService) GetWorld(ctx context.Context, worldID *wrappers.StringValue) (*proto.GetWorldResponse, error) {
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

func (a *ExternalServerService) UpdatePC(ctx context.Context, r *proto.PC) (*empty.Empty, error) {
	storage.UseStorage().Local().PCs().InsertOrUpdate(r.Base.GetId(), r)
	return &empty.Empty{}, nil
}

func (a *ExternalServerService) UpdateAmmo(ctx context.Context, r *proto.Ammo) (*empty.Empty, error) {
	storage.UseStorage().Local().Ammo().InsertOrUpdate(r.Base.GetId(), r)
	return &empty.Empty{}, nil
}

func (a *ExternalServerService) DeleteWorld(ctx context.Context, r *wrappers.StringValue) (*empty.Empty, error) {
	s := storage.UseStorage().Local()

	worldIDString := r.GetValue()

	// s.WorldStorage.Remove(worldIDString)
	s.Maps().Delete(worldIDString)
	s.Elements().Delete(worldIDString)
	s.PCs().Delete(worldIDString)
	s.Weapons().Delete(worldIDString)
	s.Ammo().Delete(worldIDString)
	return &empty.Empty{}, nil
}

func (a *ExternalServerService) DeletePC(ctx context.Context, r *wrappers.StringValue) (*empty.Empty, error) {
	storage.UseStorage().Local().PCs().Delete(r.Value)
	return &empty.Empty{}, nil
}

func NewExternalServerService() *ExternalServerService {
	return new(ExternalServerService)
}
