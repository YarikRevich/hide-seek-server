package monitoring

import (
	"context"
	"net/http"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var instance *Monitoring

type Monitoring struct {
	metrics *grpc_prometheus.ServerMetrics
}

func (m *Monitoring) NewPrometheusInterceptor() func(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error) {
	return m.metrics.UnaryServerInterceptor()
}

func (m *Monitoring) Init() {
	reg := prometheus.NewRegistry()
	m.metrics = grpc_prometheus.NewServerMetrics()
	reg.MustRegister(m.metrics)
	httpServer := &http.Server{
		Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
		Addr:    "127.0.0.1:9099",
	}
	go logrus.Fatal(httpServer.ListenAndServe())
}

func UseMonitoring() *Monitoring {
	if instance == nil {
		instance = new(Monitoring)
	}
	return instance
}
