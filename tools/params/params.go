package params

import "flag"

var (
	serverIP       = flag.String("server-ip", "127.0.0.1", "IP address for server listening")
	serverPort     = flag.String("server-port", "8080", "Port for server listening")
	demon          = flag.Bool("demon", false, "Runs server as demon")
	mongohost      = flag.String("mongohost", "127.0.0.1:27017", "Host for mongo")
	mongodb        = flag.String("mongodb", "hideseek-server", "Database to be used in mongo")
	cachetime      = flag.Float64("cachetime", 50000, "Delay for cache to be saved")
	monitoringIP   = flag.String("monitoring-ip", "127.0.0.1", "IP address for monitoring listening")
	monitoringPort = flag.String("monitoring-port", "9999", "Port for monitoring listening")
	profilecpu     = flag.Bool("profilecpu", false, "Enables profiler for CPU")
	profilemem     = flag.Bool("profilemem", false, "Enables profiler for MEM")
)

func GetServerIP() string {
	return *serverIP
}

func GetServerPort() string {
	return *serverPort
}

func IsDemon() bool {
	return *demon
}

func GetMongoHost() string {
	return *mongohost
}

func GetMongoDB() string {
	return *mongodb
}

func GetCacheTime() float64 {
	return *cachetime
}

func GetMonitoringIP() string {
	return *monitoringIP
}

func GetMonitoringPort() string {
	return *monitoringPort
}

func IsProfileCPU() bool {
	return *profilecpu
}

func IsProfileMEM() bool {
	return *profilemem
}
