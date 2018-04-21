package task

import (
	"bufio"
	"os"
	"../config"
	"time"
)

var (
	//break signal
	BreakCh chan bool
	//finish signal
	IsFinish chan bool
	//each package sleep control signal
	Timer chan bool
)

func init() {
	BreakCh = make(chan bool)
	IsFinish = make(chan bool)
	if config.Cfg.Timer > 0 {
		Timer = make(chan bool)
		go timer()
	}
	go breakSignal()
}

//
//each package sleep timer
//
func timer () {
	for {
		Timer <- true
		time.Sleep(time.Millisecond * time.Duration(config.Cfg.Timer))
	}
}

//
//input "b" or "B" and press return
//
//
func breakSignal () {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "b" || scanner.Text() == "B" {
			BreakCh <- true
		}
	}
}
