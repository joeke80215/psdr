package tcp

import (
	"net"
	"io"
	"../config"
	"fmt"
	"../iountil"
)

type TCP struct {
	addrtcp *net.TCPAddr
	tPackage io.Reader
}

func (t TCP) Send () error {
	ds, err := net.DialTCP("tcp",nil, t.addrtcp)
	if err!= nil {
		return err
	}

	go iountil.Copy(ds, t.tPackage)

	return err
}

func New (p io.Reader) *TCP {
	t := new(TCP)
	t.addrtcp,_ = net.ResolveTCPAddr("tcp",fmt.Sprintf("%s:%s",config.Cfg.Host,config.Cfg.Port))
	t.tPackage = p

	return t
}

