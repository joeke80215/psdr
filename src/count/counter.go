package count

import (
	"fmt"
	"../config"
	"sync"
	"../task"
)

type counter struct {
	//already send packages
	T int
	//success
	S int
	//lock when output screen and count
	Mux sync.Mutex
}

var (
	//total packages Num
	Tc int
	//total packages Num integer to float
	tcf float32
	//count current send stage
	Counter counter
)

func init() {
	Tc = config.Cfg.RoutineNum * config.Cfg.PackageNum
	tcf = float32(Tc)
}

func WriteInfo () {
	Counter.Mux.Lock()
	fmt.Printf("\rTotal : %d | Success : %d | Fail : %d | Complete : %.2f",
		Counter.T,
		Counter.S,
		Counter.T - Counter.S,
		float32(Counter.T) / tcf)
	Counter.Mux.Unlock()

	if Counter.T == Tc {
		task.IsFinish <- true
	}
}
