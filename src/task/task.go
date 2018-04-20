package task

import (
	"bufio"
	"os"
	"../config"
	"time"
)

var (
	BreakCh chan bool
	IsFinish chan bool
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

func timer () {
	for {
		Timer <- true
		time.Sleep(time.Millisecond * time.Duration(config.Cfg.Timer))
	}
}

func breakSignal () {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "b" || scanner.Text() == "B" {
			BreakCh <- true
		}
	}
}
