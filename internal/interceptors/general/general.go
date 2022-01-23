package general

import (
	"context"

	"github.com/YarikRevich/hide-seek-server/internal/monitoring"
	"google.golang.org/grpc"
)

func Use(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	monitoring.UseMonitoring().RegisterManager().UsageHistogram().Inc()
	return handler(ctx, req)
}
