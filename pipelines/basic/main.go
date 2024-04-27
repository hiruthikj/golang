package main

import "fmt"

func main() {
	done := make(chan interface{})
	c := gen(done, 1,2,3)
	op := sq(done, c)

	fmt.Println(<-op)
	fmt.Println(<-op)
	fmt.Println(<-op)
	fmt.Println(<-op)
	a, b:= <-op
	fmt.Println(a, b)
}