package implementation

import (
	"context"
	"sort"

	"github.com/YarikRevich/hide-seek-server/internal/api/external-api/v1/proto"
	"github.com/YarikRevich/hide-seek-server/internal/storage"

	// externalapistorage "github.com/YarikRevich/hide-seek-server/internal/storage/external-api"
	"github.com/YarikRevich/hide-seek-server/tools/utils"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type ExternalService struct {
	proto.UnimplementedExternalServiceServer
}

func (h *ExternalService) UpdateWorld(ctx context.Context, world *proto.World) (*wrappers.BoolValue, error) {
	storage.UseStorage().Local().Worlds().InsertOrUpdate(world.GetId(), world)
	return &wrappers.BoolValue{Value: true}, nil
}

func (a *ExternalService) UpdateMap(ctx context.Context, worldMap *proto.Map) (*wrappers.BoolValue, error) {
	storage.UseStorage().Local().Worlds().InsertOrUpdate(worldMap.Base.Parent.GetId(), worldMap)
	return &wrappers.BoolValue{Value: true}, nil
}

//Assigns random spawns to the pcs set by this world
func (a *ExternalService) AssignRandomSpawnsToPCs(ctx context.Context, r *proto.AssignRandomSpawnsToPCsRequest) (*wrappers.BoolValue, error) {
	s := storage.UseStorage()
	worldMap := s.Local().Maps().Find(r.GetWorldId()).(*proto.Map)
	pcs := s.Local().PCs().Find(r.GetWorldId()).([]*proto.PC)
	randomSpawnsIndicies := utils.GetSequenceOfRandomInt(len(pcs), 0, len(pcs)-1)
	for i := 0; i < len(pcs)-1; i++ {
		pcs[i].Base.Spawn = worldMap.GetSpawns()[randomSpawnsIndicies[i]]
	}
	return nil, nil
}

func (a *ExternalService) GetWorld(ctx context.Context, worldID *wrappers.StringValue) (*proto.GetWorldResponse, error) {
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

func (a *ExternalService) UpdatePC(ctx context.Context, r *proto.PC) (*wrappers.BoolValue, error) {
	storage.UseStorage().Local().PCs().InsertOrUpdate(r.Base.GetId(), r)
	return &wrappers.BoolValue{Value: true}, nil
}

func (a *ExternalService) UpdateAmmo(ctx context.Context, r *proto.Ammo) (*wrappers.BoolValue, error) {
	storage.UseStorage().Local().Ammo().InsertOrUpdate(r.Base.GetId(), r)
	return &wrappers.BoolValue{Value: true}, nil
}

func (a *ExternalService) DeleteWorld(ctx context.Context, r *wrappers.StringValue) (*wrappers.BoolValue, error) {
	s := storage.UseStorage().Local()

	worldIDString := r.GetValue()

	// s.WorldStorage.Remove(worldIDString)
	s.Maps().Delete(worldIDString)
	s.Elements().Delete(worldIDString)
	s.PCs().Delete(worldIDString)
	s.Weapons().Delete(worldIDString)
	s.Ammo().Delete(worldIDString)
	return &wrappers.BoolValue{Value: true}, nil
}

// func (a *ApiServer) RemovePC(ctx context.Context, r *RemovePCRequest) (*Status, error) {
// 	c := collection.UseCollection()
// 	delete(c.PCs[r.WorldId], r.PcId)
// 	return &Status{Ok: true}, nil
// }

func NewExternalService() *ExternalService {
	return new(ExternalService)
}
