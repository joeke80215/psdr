package config

import (
	"flag"
)

type config struct {
	Host string
	Port string
	RoutineNum int
	PackageNum int
	Sleep int
	PackageSize int
	Method string
}

var (
	Cfg *config
	host = flag.String("H","127.0.0.1","Target host")
	port = flag.String("P","8080","Target host port")
	rn = flag.Int("rn",2,"Execute how many Routines")
	pn = flag.Int("pn",5,"Send how many packages")
	ps = flag.Int("ps",1024,"Package size (byte)")
	mt = flag.String("mt","tcp","Request method(tcp/udp/http)")
)

func init() {
	flag.Parse()
	Cfg = new(config)
	Cfg.RoutineNum = *rn
	Cfg.PackageNum = *pn
	Cfg.Host = *host
	Cfg.Port = *port
	Cfg.PackageSize = *ps
	Cfg.Method = *mt
}
