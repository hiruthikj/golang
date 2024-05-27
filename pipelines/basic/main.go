package main

import (
	"fmt"
	"runtime"
)

func main() {
	totalInputs := 10

	ls := make([]int, totalInputs)
	done := make(chan interface{})
	defer close(done)
	
	for i := 0; i < totalInputs; i++ {
		ls[i] = i+1
	}
	c := gen(done, ls...)

	cpuCount := runtime.NumCPU()

	outboundChans := make([]<-chan int, cpuCount)
	for i:= 0; i < cpuCount; i++ {
		outboundChans[i] = sq(done, c)
	}

	for o := range fanIn(done, outboundChans...) {
		fmt.Println(o)
	}
}