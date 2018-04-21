package sender

import (
	"log"
	"bytes"
	"../config"
	"../tcp"
	"../udp"
	"../http"
	c "../count"
)

//
//sender is connect method interface
//
type sender func () error

var (
	snd sender
)

func init() {
	//if method is "udp" and each package size over 9016 byte
	//reset package size 9016
	if config.Cfg.PackageSize > 9016 && config.Cfg.Method == "udp" {
		config.Cfg.PackageSize = 9016
		log.Println("package size > 9016 bytes in udp,reset package size = 9016 bytes")
	}

	//make a package refer package size
	mpackage := bytes.NewBuffer(make([]byte,config.Cfg.PackageSize,config.Cfg.PackageSize))

	//choose which send mothod
	switch config.Cfg.Method {
	case "tcp":
		snd = tcp.New(mpackage).Send
	case "udp":
		snd = udp.New(mpackage).Send
	case "http":
		snd = http.New(mpackage).Send
	}
}

//
//send a package to target host
//
func Handle () {
	err := snd()
	c.Counter.Mux.Lock()
	c.Counter.T++
	if err == nil {
		c.Counter.S++
	}
	go c.WriteInfo()
	c.Counter.Mux.Unlock()
}

