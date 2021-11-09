package server

import (
	"context"
	"github.com/YarikRevich/HideSeek-Server/internal/api"
)

type ApiServer struct{}

func (a *ApiServer) UploadWorldMetrics(ctx context.Context, r *api.World) *api.Status {
	return nil
}

func NewApiServer() *ApiServer {
	return new(ApiServer)
}
