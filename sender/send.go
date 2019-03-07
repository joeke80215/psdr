package sender

import (
	"log"
	"bytes"
	"github.com/joeke80215/psdr/config"
	"github.com/joeke80215/psdr/tcp"
	"github.com/joeke80215/psdr/udp"
	"github.com/joeke80215/psdr/http"
	c "github.com/joeke80215/psdr/count"
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
	c.Counter.Mux.Lock()
	c.Counter.T++
	if err == nil {
		c.Counter.S++
	}
	go c.WriteInfo()
	c.Counter.Mux.Unlock()
}

