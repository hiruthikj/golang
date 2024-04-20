package main

import "fmt"

func main() {
	fmt.Println(greetMe("You"))
}

func greetMe(name string) string {
	return "Hello, " + name + "!"
}
