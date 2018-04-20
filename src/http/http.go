package http

import (
	"net/http"
	"io"
	"../config"
	"fmt"
)

type HTTP struct {
	addrhttp string
	hPackage io.Reader
}

func (h HTTP) Send () error {
	req, err := http.NewRequest("POST", h.addrhttp, h.hPackage)

	req.Header.Set("Content-Type", "application/byte")

	client := &http.Client{}
	_, err = client.Do(req)

	//defer resp.Body.Close()
	if err != nil {
		return err
	}

	return nil
}

func New (p io.Reader) *HTTP {
	h := new(HTTP)
	h.addrhttp = fmt.Sprintf("http://%s:%s",config.Cfg.Host,config.Cfg.Port)
	h.hPackage = p

	return h
}
