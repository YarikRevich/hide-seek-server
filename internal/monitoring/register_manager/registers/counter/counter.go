package counter

import (
	"github.com/YarikRevich/hide-seek-server/internal/monitoring/register_manager/registers/common"
	"github.com/prometheus/client_golang/prometheus"
)

type Counter struct {
	register prometheus.Counter
}

func New(name string) common.Register {
	c := &Counter{
		register: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: name,
			},
		),
	}
	prometheus.MustRegister(c.register)
	return c
}
