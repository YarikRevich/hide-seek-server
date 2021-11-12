package server

import (
	"context"
	"sort"
	"time"

	"github.com/YarikRevich/HideSeek-Server/internal/api"
	"github.com/YarikRevich/HideSeek-Server/internal/collection"
	// "github.com/YarikRevich/HideSeek-Server/internal/collection"
)

type ApiServer struct {
	api.UnimplementedHideSeekServer
}

func (a *ApiServer) AddWorld(ctx context.Context, r *api.World) (*api.Status, error) {
	c := collection.UseCollection()
	c.Worlds[r.Object.Id] = r
	return &api.Status{Ok: true}, nil
}

func (a *ApiServer) AddPC(ctx context.Context, r *api.PC) (*api.Status, error) {
	c := collection.UseCollection()
	if _, ok := c.PCs[r.Object.WorldId]; !ok {
		c.PCs[r.Object.WorldId] = make(map[string]*collection.PC)
	}

	c.PCs[r.Object.WorldId][r.Object.Id] = &collection.PC{r, time.Now()}
	return &api.Status{Ok: true}, nil
}

func (a *ApiServer) AddElement(ctx context.Context, r *api.Element) (*api.Status, error) {
	return nil, nil
}

func (a *ApiServer) AddWeapon(ctx context.Context, r *api.Weapon) (*api.Status, error) {
	return nil, nil
}

func (a *ApiServer) AddAmmo(ctx context.Context, r *api.Ammo) (*api.Status, error) {
	return nil, nil
}

func (a *ApiServer) ChooseSpawns(ctx context.Context, r *api.ChooseSpawnsRequest) (*api.Status, error) {
	return nil, nil
}

func (a *ApiServer) GetWorldObjects(ctx context.Context, r *api.WorldObjectsRequest) (*api.WorldObjectsResponse, error) {
	c := collection.UseCollection()

	var weapons []*api.Weapon
	for _, v := range c.Weapons[r.WorldId] {
		weapons = append(weapons, v)
	}

	var elements []*api.Element
	for _, v := range c.Elements[r.WorldId] {
		elements = append(elements, v)
	}

	var fullPCs []*collection.PC
	for _, v := range c.PCs[r.WorldId] {
		fullPCs = append(fullPCs, v)
	}
	sort.Slice(fullPCs, func(i, j int)bool{
		return fullPCs[i].AddTime.Before(fullPCs[j].AddTime)
	})

	var pcs []*api.PC
	for _, v := range fullPCs{
		pcs = append(pcs, v.Data)
	}

	var ammo []*api.Ammo
	for _, v := range c.Ammo[r.WorldId] {
		ammo = append(ammo, v)
	}

	return &api.WorldObjectsResponse{
		Weapons:  weapons,
		Elements: elements,
		Ammos:    ammo,
		PCs:      pcs,
	}, nil
}

func (a *ApiServer) UpdatePC(ctx context.Context, r *api.PC) (*api.Status, error) {
	c := collection.UseCollection()
	c.PCs[r.Object.WorldId][r.Object.Id].Data = r
	return &api.Status{Ok: true}, nil
}

func (a *ApiServer) UpdateAmmo(ctx context.Context, r *api.Ammo) (*api.Status, error) {
	c := collection.UseCollection()
	c.Ammo[r.Object.WorldId][r.Object.Id] = r
	return &api.Status{Ok: true}, nil
}

func (a *ApiServer) RemoveWorld(ctx context.Context, r *api.RemoveWorldRequest) (*api.Status, error) {
	c := collection.UseCollection()
	delete(c.Worlds, r.WorldId)
	return &api.Status{Ok: true}, nil
}

func (a *ApiServer) RemovePC(ctx context.Context, r *api.RemovePCRequest) (*api.Status, error) {
	c := collection.UseCollection()
	delete(c.PCs[r.WorldId], r.PcId)
	return &api.Status{Ok: true}, nil
}

func (a *ApiServer) SetGameStarted(ctx context.Context, r *api.SetGameStartedRequest) (*api.Status, error) {
	c := collection.UseCollection()
	c.Games[r.WorldId] = collection.Game{Started: true}
	return &api.Status{Ok: true}, nil
}

func (a *ApiServer) IsGameStarted(ctx context.Context, r *api.IsGameStartedRequest) (*api.IsGameStartedResponse, error) {
	c := collection.UseCollection()
	if g, ok := c.Games[r.WorldId]; ok {
		return &api.IsGameStartedResponse{Started: g.Started}, nil
	}
	return &api.IsGameStartedResponse{Started: false}, nil
}

func (a *ApiServer) GetWorld(ctx context.Context, r *api.GetWorldRequest)(*api.World, error){
	c := collection.UseCollection()
	return c.Worlds[r.WorldId], nil
}

func NewApiServer() *ApiServer {
	return new(ApiServer)
}
