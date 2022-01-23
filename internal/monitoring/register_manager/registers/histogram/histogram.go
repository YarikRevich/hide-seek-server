package histogram

import (
	"github.com/YarikRevich/hide-seek-server/internal/monitoring/register_manager/registers/common"
	"github.com/prometheus/client_golang/prometheus"
)

type Histogram struct {
	register prometheus.Histogram
}

func (c *Histogram) Inc() {
	c.register.Observe(1)
}

func (c *Histogram) Dec() {
	c.register.Observe(-1)
}

func New(name string) common.Register {
	h := &Histogram{
		register: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Name: name,
			},
		),
	}
	prometheus.MustRegister(h.register)
	return h
}
