package implementation

import (
	"context"
	"sort"

	"github.com/YarikRevich/HideSeek-Server/internal/api/external-api/v1/proto"
	"github.com/YarikRevich/HideSeek-Server/internal/storage"
	externalapistorage "github.com/YarikRevich/HideSeek-Server/internal/storage/external-api"
	"github.com/YarikRevich/HideSeek-Server/tools/utils"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type ExternalService struct {
	proto.UnimplementedExternalServiceServer
}

func (h *ExternalService) UpdateWorld(ctx context.Context, world *proto.World) (*wrappers.BoolValue, error) {
	storage.UseStorage().WorldStorage.Add(world.GetId(), world)
	return &wrappers.BoolValue{Value: true}, nil
}

func (a *ExternalService) UpdateMap(ctx context.Context, r *proto.Map) (*wrappers.BoolValue, error) {
	storage.UseStorage().MapStorage.Add(r.Base.Parent.GetId(), r)
	return &wrappers.BoolValue{Value: true}, nil
}

// func (a *ExternalService) AddElement(ctx context.Context, r *proto.Element) (*wrappers.BoolValue, error) {
// 	return nil, nil
// }

// func (a *ExternalService) AddWeapon(ctx context.Context, r *proto.Weapon) (*wrappers.BoolValue, error) {
// 	return nil, nil
// }

// func (a *ExternalService) AddAmmo(ctx context.Context, r *proto.Ammo) (*wrappers.BoolValue, error) {
// 	return nil, nil
// }

//Assigns random spawns to the pcs set by this world
func (a *ExternalService) AssignRandomSpawnsToPCs(ctx context.Context, r *proto.AssignRandomSpawnsToPCsRequest) (*wrappers.BoolValue, error) {
	s := storage.UseStorage()
	mps := s.MapStorage.Get(r.GetWorldId())
	pcs := s.PCStorage.Get(r.GetWorldId())
	randomSpawnsIndicies := utils.GetSequenceOfRandomInt(len(pcs), 0, len(pcs)-1)
	for i := 0; i < len(pcs)-1; i++ {
		pcs[i].Base.Spawn = mps.GetSpawns()[randomSpawnsIndicies[i]]
	}
	return nil, nil
}

func (a *ExternalService) GetWorld(ctx context.Context, worldID *wrappers.StringValue) (*proto.GetWorldResponse, error) {
	s := storage.UseStorage().ExternalApiStorage

	pcs := s.PCStorage.Get(worldID.Value)
	copyPCs := make(externalapistorage.PCs, len(pcs), cap(pcs))
	copy(copyPCs, pcs)

	sort.Slice(copyPCs, func(i, j int) bool {
		return copyPCs[i].GetLobbyNumber() < copyPCs[j].GetLobbyNumber()
	})

	return &proto.GetWorldResponse{
		World:    s.WorldStorage.Get(worldID.Value),
		Weapons:  s.WeaponStorage.Get(worldID.Value),
		Elements: s.ElementStorage.Get(worldID.Value),
		Ammos:    s.AmmoStorage.Get(worldID.Value),
		PCs:      copyPCs,
	}, nil
}

func (a *ExternalService) UpdatePC(ctx context.Context, r *proto.PC) (*wrappers.BoolValue, error) {
	v := storage.UseStorage().PCStorage.Get(r.Base.GetId())
	v.Remove(r)
	v.Add(r)

	s := storage.UseStorage().PCStorage

	r.LobbyNumber = int64(s.Length() + 1)
	s.Add(r.Base.Parent.Parent.GetId(), r)

	return &wrappers.BoolValue{Value: true}, nil

	// return &wrappers.BoolValue{true}, nil
}

func (a *ExternalService) UpdateAmmo(ctx context.Context, r *proto.Ammo) (*wrappers.BoolValue, error) {
	v := storage.UseStorage().AmmoStorage.Get(r.Base.GetId())
	v.Remove(r)
	v.Add(r)
	return &wrappers.BoolValue{Value: true}, nil
}

func (a *ExternalService) DeleteWorld(ctx context.Context, r *wrappers.StringValue) (*wrappers.BoolValue, error) {
	s := storage.UseStorage()

	worldIDString := r.GetValue()

	// s.WorldStorage.Remove(worldIDString)
	s.MapStorage.Remove(worldIDString)
	s.ElementStorage.Remove(worldIDString)
	s.PCStorage.Remove(worldIDString)
	s.WeaponStorage.Remove(worldIDString)
	s.AmmoStorage.Remove(worldIDString)
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
