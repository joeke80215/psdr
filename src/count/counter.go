package count

import (
	"fmt"
)

var (
	Total int
	Success int
	TotalCh chan int
	SuccessCh chan int
	writeCh chan bool
)

func init() {
	Total = 0
	Success = 0
	TotalCh = make(chan int)
	SuccessCh = make(chan int)
	writeCh = make(chan bool)
	go func() {TotalCh <- 0}()
	go func() {SuccessCh <- 0}()
	go func() {writeCh <- true}()
	go writeInfo()
}

func Increate (c chan int) int {
	t := <- c
	t++
	go func() {c <- t}()
	return t
}

func writeInfo () {
	for {
		<- writeCh
		Total = <- TotalCh
		Success = <- SuccessCh

		fmt.Printf("\rTotal : %d | Success : %d | Fail : %d ", Total, Success, Total - Success)
		go func() {TotalCh <- Total}()
		go func() {SuccessCh <- Success}()
	}
}

func OutputInfo () {
	writeCh <- true
}
