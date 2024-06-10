package main

import (
	"errors"
	"fmt"
)

func getCheckFn(checkName string) (func(...interface{}) chan interface{}, error) {
	switch checkName {
	case "age":
		return checkAge, nil
	case "credit":
		return checkAge, nil
	case "amounat":
		return checkAge, nil
	}

	return nil, errors.New("undefined check name")
}

func main() {
	checks := []string{"age", "credit", "amount"}
	runColenders(checks)
}

func runColenders(checks []string) {
	for _, check := range checks {
		fn, err := getCheckFn(check)
		if err != nil {
			if errors.Is(err, errors.New("undefined check name")) {
				fmt.Println("Is error!")
				continue
			}
			fmt.Printf("error in runColenders: %v\n", err)
			continue
		}
		fn()
	}
}

func checkAge(args ...interface{}) chan interface{} {
	fmt.Println("Age check")
	return make(chan interface{})
}
