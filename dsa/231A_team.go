package main

import "fmt"

func main() {
	var problems int
	fmt.Scanf("%d\n", &problems)

	var a, b, c int
	result := 0
	for i := 0; i < problems; i++ {
		fmt.Scanf("%d %d %d\n", &a, &b, &c)

		result += a&b | b&c | c&a
	}
	fmt.Println(result)
}
