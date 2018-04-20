package exec

import (
	"../sender"
	"../config"
	"runtime"
	"os"
	"fmt"
	"../task"
)

func Exec () {
	exec(config.Cfg.RoutineNum,config.Cfg.PackageNum)
}

func exec (rn,pn int) {
	isExec := false
	for {
		select {
		case <- task.BreakCh :
			fmt.Print("break proccess\n")
			os.Exit(0)
		case <- task.IsFinish :
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
							if config.Cfg.Timer > 0 {
								<- task.Timer
							}
							runtime.Gosched()
							sender.Handle()
							p++
						}
					}()
					r++
				}
			}()
		}
		}
	}
}
