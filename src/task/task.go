package task

import (
	"bufio"
	"os"
)

var (
	BreakCh chan bool
	IsFinish chan bool
	StdoutDone chan bool
)

func init() {
	BreakCh = make(chan bool)
	IsFinish = make(chan bool)
	StdoutDone = make(chan bool)
	go breakSignal()
}

func breakSignal () {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "b" || scanner.Text() == "B" {
			BreakCh <- true
		}
	}
}
