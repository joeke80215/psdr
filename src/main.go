package main

import (
	"./exec"
	"runtime"
	"os"
)

func init () {
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
}

func main() {
	exec.Exec()
}
