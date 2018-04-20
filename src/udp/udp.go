package udp

import (
	"net"
	"io"
	"../config"
	"fmt"
	"../iountil"
)

type UDP struct {
	addrudp *net.UDPAddr
	uPackage io.Reader
}

func (u UDP) Send () error {
	conn, err := net.DialUDP("udp", nil, u.addrudp)
	if err!= nil {
		return err
	}

	go iountil.Copy(conn,u.uPackage)

	return nil
}

func New (p io.Reader) *UDP {
	u := new(UDP)
	u.addrudp,_ = net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s",config.Cfg.Host,config.Cfg.Port))
	u.uPackage = p
	return u
}
