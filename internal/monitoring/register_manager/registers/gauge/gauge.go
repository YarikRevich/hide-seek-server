package gauge

import (
	"github.com/YarikRevich/hide-seek-server/internal/monitoring/register_manager/registers/common"
	"github.com/prometheus/client_golang/prometheus"
)

type Gauge struct {
	register prometheus.Gauge
}

func (c *Gauge) Inc() {
	c.register.Inc()
}

func (c *Gauge) Dec() {
	c.register.Dec()
}

func New(name string) common.Register {
	g := &Gauge{
		register: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: name,
			},
		),
	}
	prometheus.MustRegister(g.register)
	return g
}
