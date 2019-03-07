package udp

import (
	"bytes"
	"fmt"
	"io"
	"net"

	"github.com/joeke80215/psdr/config"
)

type UDP struct {
	addrudp  *net.UDPAddr
	uPackage io.Reader
	ubyte    *bytes.Buffer
}

func (u UDP) Send() error {
	conn, err := net.DialUDP("udp", nil, u.addrudp)
	if err != nil {
		return err
	}

	_, err = conn.Write(u.ubyte.Bytes())
	if err != nil {
		return err
	}
	conn.Close()

	return err
}

func New(p io.Reader) *UDP {
	u := new(UDP)
	u.addrudp, _ = net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", config.Cfg.Host, config.Cfg.Port))
	u.uPackage = p
	u.ubyte = new(bytes.Buffer)
	u.ubyte.ReadFrom(u.uPackage)
	return u
}
