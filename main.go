package main

import (
	"os"
	"runtime"

	"github.com/joeke80215/psdr/exec"
)

func init() {
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
}

func main() {
	exec.Exec()
}
