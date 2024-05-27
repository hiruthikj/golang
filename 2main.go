package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	X, Y int
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func main() {
	v := Vertex{1, -5}
	fmt.Println(v.Abs())
}