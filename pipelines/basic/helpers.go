package main

import "math"

func sq(done <-chan interface{}, inputStream <-chan int) <-chan int {
	outputStream := make(chan int)

	go func() {
		defer close(outputStream)
		for num := range orDone(done, inputStream) {
			outputStream <- int(math.Pow(float64(num), 2))
		}
	}()

	return outputStream
}
