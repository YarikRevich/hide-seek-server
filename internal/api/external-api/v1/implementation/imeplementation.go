package imlementation

import (
	"context"
	"sort"

	// "sort"
	// "time"

	"github.com/YarikRevich/HideSeek-Server/internal/api/external-api/v1/proto"
	"github.com/YarikRevich/HideSeek-Server/internal/storage"
	externalapistorage "github.com/YarikRevich/HideSeek-Server/internal/storage/external-api"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type ExternalService struct {
	proto.UnimplementedExternalServiceServer
}

func (h *ExternalService) UpdateWorld(ctx context.Context, world *proto.World) (*proto.Status, error) {
	storage.UseStorage().WorldStorage.Add(world.GetId(), world)
	return &proto.Status{Ok: true}, nil
}

func (h *ExternalService) GetWorld(ctx context.Context, worldID *wrappers.StringValue) (*proto.GetWorldResponse, error) {
	world, err := storage.UseStorage().WorldStorage.Get(worldID.Value)
	return &proto.GetWorldResponse{World: world}, err
}

func (a *ExternalService) AddWorld(ctx context.Context, r *proto.World) (*proto.Status, error) {
	storage.UseStorage().WorldStorage.Add(r.GetId(), r)
	return &proto.Status{Ok: true}, nil
}

func (a *ExternalService) AddMap(ctx context.Context, r *proto.Map) (*proto.Status, error) {
	storage.UseStorage().MapStorage.Add(r.Base.GetId(), r)
	return &proto.Status{Ok: true}, nil
}

func (a *ExternalService) AddPC(ctx context.Context, r *proto.PC) (*proto.Status, error) {
	s := storage.UseStorage().PCStorage

	r.LobbyNumber = int64(s.Length() + 1)
	s.Add(r.Base.GetId(), r)
	return &proto.Status{Ok: true}, nil
}

func (a *ExternalService) AddElement(ctx context.Context, r *proto.Element) (*proto.Status, error) {
	return nil, nil
}

func (a *ExternalService) AddWeapon(ctx context.Context, r *proto.Weapon) (*proto.Status, error) {
	return nil, nil
}

func (a *ExternalService) AddAmmo(ctx context.Context, r *proto.Ammo) (*proto.Status, error) {
	return nil, nil
}

func (a *ApiServer) ChooseSpawns(ctx context.Context, r *ChooseSpawnsRequest) (*Status, error) {
	return nil, nil
}

func (a *ExternalService) GetWorldProperty(ctx context.Context, worldID *wrappers.StringValue) (*proto.GetWorldPropertyResponse, error) {
	s := storage.UseStorage().ExternalApiStorage

	var pcs externalapistorage.PCs
	copy(pcs, s.PCStorage[worldID.Value])

	sort.Slice(pcs, func(i, j int) bool {
		return pcs[i].GetLobbyNumber() < pcs[j].GetLobbyNumber()
	})

	ammos, err := s.AmmoStorage.Get(worldID.Value)
	if err != nil {
		return nil, err
	}

	return &proto.GetWorldPropertyResponse{
		Weapons:  s.WeaponStorage.Get(worldID.Value),
		Elements: s.ElementStorage.Get(worldID.Value),
		Ammos:    ammos,
		PCs:      pcs,
	}, nil
}

func (a *ExternalService) UpdatePC(ctx context.Context, r *proto.PC) (*proto.Status, error) {
	v, err := storage.UseStorage().PCStorage.Get(r.Base.GetId())
	if err != nil {
		return nil, err
	}
	v.Remove(r)
	v.Add(r)
	return &proto.Status{Ok: true}, nil
}

func (a *ExternalService) UpdateAmmo(ctx context.Context, r *proto.Ammo) (*proto.Status, error) {
	v, err := storage.UseStorage().AmmoStorage.Get(r.Base.GetId())
	if err != nil {
		return nil, err
	}
	v.Remove(r)
	v.Add(r)
	return &proto.Status{Ok: true}, nil
}

func (a *ExternalService) DeleteWorld(ctx context.Context, r *wrappers.StringValue) (*proto.Status, error) {
	s := storage.UseStorage()

	worldIDString := r.GetValue()

	s.WorldStorage.Remove(worldIDString)
	s.MapStorage.Remove(worldIDString)
	s.ElementStorage.Remove(worldIDString)
	s.PCStorage.Remove(worldIDString)
	s.WeaponStorage.Remove(worldIDString)
	s.AmmoStorage.Remove(worldIDString)
	return &proto.Status{Ok: true}, nil
}

// func (a *ApiServer) RemovePC(ctx context.Context, r *RemovePCRequest) (*Status, error) {
// 	c := collection.UseCollection()
// 	delete(c.PCs[r.WorldId], r.PcId)
// 	return &Status{Ok: true}, nil
// }

func NewExternalService() *ExternalService {
	return new(ExternalService)
}
