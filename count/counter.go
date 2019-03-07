package count

import (
	"fmt"
	"sync"

	"github.com/joeke80215/psdr/config"
	"github.com/joeke80215/psdr/task"
)

type counter struct {
	T   int
	S   int
	Mux sync.Mutex
}

var (
	Tc      int
	tcf     float32
	Counter counter
)

func init() {
	Tc = config.Cfg.RoutineNum * config.Cfg.PackageNum
	tcf = float32(Tc)
}

func WriteInfo() {
	Counter.Mux.Lock()
	fmt.Printf("\rTotal : %d | Success : %d | Fail : %d | Complete : %.2f",
		Counter.T,
		Counter.S,
		Counter.T-Counter.S,
		float32(Counter.T)/tcf)
	Counter.Mux.Unlock()

	if Counter.T == Tc {
		task.IsFinish <- true
	}
}
