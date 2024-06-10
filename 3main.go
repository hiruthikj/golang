package main

import "fmt"

func main() {
    m1 := map[string]int{"one": 1, "two": 2}
    m2 := map[string]int{"one": 1, "two": 2}

    // This will cause a compilation error
    // invalid operation: m1 == m2 (map can only be compared to nil)
    fmt.Println(m1 == m2)
    panic(1)
}

