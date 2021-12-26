package interceptors

import (
	"context"

	externalapiinterceptor "github.com/YarikRevich/hide-seek-server/internal/interceptors/external-api"
	"google.golang.org/grpc"
)

// type Interceptors struct{}

// //Postpones cache expiration for choosen methods
// func (i *Interceptors) Cache(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
// 	var id string
// 	switch pr := req.(type){
// 	case *api.World:
// 		id = pr.Base.Id
// 	case *api.PC:
// 		id = pr.Object.Id
// 	case *api.Weapon:
// 		id = pr.Object.Id
// 	case *api.Ammo:
// 		id = pr.Object.Id
// 	case *api.Element:
// 		id = pr.Object.Id
// 	}
// 	cache.UseCache().Postpone(id)
// 	return handler(ctx, req)
// }

// func (i *Interceptors) Get() []grpc.UnaryServerInterceptor {
// 	return []grpc.UnaryServerInterceptor{
// 		i.Cache,
// 	}
// }

// func NewInterceptors() *Interceptors{
// 	return new(Interceptors)
// }

const (
	External = "external"
	Internal = "internal"
)

func GetReqType() string {
	return External
}

func NewInterceptorManager() func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		switch GetReqType() {
		case External:
			return externalapiinterceptor.NewExternalApiInterceptor().Use(
				ctx, req, info, handler)
		case Internal:

		}

		return handler(ctx, req)
	}
}
