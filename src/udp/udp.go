package udp

import (
	"net"
	"io"
	"../config"
	"fmt"
	"bytes"
)

type UDP struct {
	addrudp *net.UDPAddr
	uPackage io.Reader
	ubyte *bytes.Buffer
}

func (u UDP) Send () error {
	conn, err := net.DialUDP("udp", nil, u.addrudp)
	if err!= nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(u.ubyte.Bytes())
	if err != nil {
		return err
	}

	return err
}

func New (p io.Reader) *UDP {
	u := new(UDP)
	u.addrudp,_ = net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s",config.Cfg.Host,config.Cfg.Port))
	u.uPackage = p
	u.ubyte = new(bytes.Buffer)
	u.ubyte.ReadFrom(u.uPackage)
	return u
}
