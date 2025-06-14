package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{}
	n := 20

	for range make([]int, n) {
		arr = append(arr, rand.Intn(1000)-200)
	}

	fmt.Println(arr)
	bubbleSort(arr)
	fmt.Println(arr)
}

func bubbleSort(arr []int) {
	n := len(arr)

	for {
		pos := 0
		for i := 0; i < n-1-pos; i++ {
			if arr[i+1] < arr[i] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				pos = i + 1
			}
		}
		if pos <= 1 {
			break
		}
	}

}
