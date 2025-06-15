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
	insertionSort(arr, n)
	fmt.Println(arr)
}

// Adaptive, In-place, Stable sort algo
func bubbleSort(arr []int, n int) {
	for {
		pos := 0
		for i := 0; i < n-1-pos; i++ {
			if arr[i+1] < arr[i] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				// To Skip parts of array which is already sorted
				pos = i + 1
			}
		}
		if pos <= 1 {
			break
		}
	}
}

// Online, Adaptive, In-place, Stable sort algo
func insertionSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			arr[j+1], arr[j] = arr[j], arr[j+1]
		}
	}
}
