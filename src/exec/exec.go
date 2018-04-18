package exec

import (
	"../sender"
	"sync"
	"../config"
	"runtime"
	"bufio"
	"os"
	"fmt"
)

var (
	breakCh chan bool
	isFinish chan bool
)

func init() {
	breakCh = make(chan bool)
	isFinish = make(chan bool)
	go breakSignal()
}

func breakSignal () {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "b" || scanner.Text() == "B" {
			breakCh <- true
		}
	}
}

func Exec () {
	exec(config.Cfg.RoutineNum,config.Cfg.PackageNum)
}

func exec (rn,pn int) {
	isExec := false
	var wg sync.WaitGroup
	wg.Add(rn * pn)
	for {
		select {
		case <- breakCh :
			fmt.Print("break proccess\n")
			os.Exit(0)
		case <- isFinish :
			fmt.Print("\nfinish proccess\n")
			os.Exit(0)
		default :
		if !isExec {
			isExec = true
			go func() {
				r := 0
				for r < rn {
					go func() {
						p := 0
						for p < pn {
							runtime.Gosched()
							sender.Handle()
							p++
							wg.Done()
						}
					}()
					r++
				}
				wg.Wait()
				isFinish <- true
			}()
		}
		}
	}
}
