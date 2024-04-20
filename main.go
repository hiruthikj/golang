// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// var wg sync.WaitGroup
// 	words := []string{"hello", "there", "world"}

// 	// wg.Add(len(words))
//   done := make(chan (interface{}))
// 	for _, word := range words {
// 		go func(word string) {
//       <-done
// 			// defer wg.Done()
      
// 			fmt.Println(word)
// 			done <- nil
//       }(word)
//       done <- nil
//     }
//     <-done
    
// 	// wg.Wait()

// }
