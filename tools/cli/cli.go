package cli

import "flag"

var (
	ip = flag.String("ip", "127.0.0.1", "IP address for listening")
	port  = flag.String("port", "8080", "Port for listening")
)

func GetIP()string{
	return *ip
}
func GetPort()string{
	return *port
}