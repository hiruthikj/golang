package main

import (
	"github.com/hiruthikj/golang/learning/lucky-number/internal/random"

	"github.com/fatih/color"
)

func main() {
	green := color.New(color.FgGreen)
	green.Printf("Your lucky number is %d!\n", random.Number())
}
