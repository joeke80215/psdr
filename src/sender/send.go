package sender

import (
	"log"
	"bytes"
	"../config"
	"../tcp"
	"../udp"
	"../http"
	"../count"
)

type sender func () error

var (
	snd sender
)

func init() {

	if config.Cfg.PackageSize > 9016 && config.Cfg.Method == "udp" {
		config.Cfg.PackageSize = 9016
		log.Println("package size > 9016 bytes in udp,reset package size = 9016 bytes")
	}

	mpackage := bytes.NewBuffer(make([]byte,config.Cfg.PackageSize,config.Cfg.PackageSize))

	switch config.Cfg.Method {
	case "tcp":
		snd = tcp.New(mpackage).Send
	case "udp":
		snd = udp.New(mpackage).Send
	case "http":
		snd = http.New(mpackage).Send
	}
}

func Handle () {
	err := snd()
	count.Increate(count.TotalCh)
	if err == nil {
		count.Increate(count.SuccessCh)
	}
	count.OutputInfo()
}

