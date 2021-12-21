package params

import "flag"

var (
	ip        = flag.String("ip", "127.0.0.1", "IP address for listening")
	port      = flag.String("port", "8080", "Port for listening")
	demon     = flag.Bool("demon", false, "Runs server as demon")
	mongohost = flag.String("mongohost", "127.0.0.1:27017", "Host for mongo")
	mongodb   = flag.String("mongodb", "hideseek-server", "Database to be used in mongo")
	cachetime = flag.Float64("cachetime", 120, "Delay for cache to be saved")
)

func GetIP() string {
	return *ip
}

func GetPort() string {
	return *port
}

func GetDemon() bool {
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
