package monitoring

import (
	"fmt"
	"net/http"

	registermanager "github.com/YarikRevich/hide-seek-server/internal/monitoring/register_manager"
	"github.com/YarikRevich/hide-seek-server/tools/params"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var instance *Monitoring

type Monitoring struct {
	registerManager *registermanager.RegisterManager
}

func (m *Monitoring) ListenAndServe() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		logrus.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", params.GetMonitoringIP(), params.GetMonitoringPort()), nil))
	}()
}

func UseMonitoring() *Monitoring {
	if instance == nil {
		instance = &Monitoring{
			registerManager: registermanager.New(),
		}
	}
	return instance
}
