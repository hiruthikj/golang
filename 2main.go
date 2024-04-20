package main

import "fmt"

func main() {

	chanOwner := func() <-chan int {
		dataStream := make(chan int, 500)
		
		go func ()  {
			defer close(dataStream)
			for i:=0; i<=500; i++ {
				dataStream <- i
			}
		}()
		return dataStream
	}

	consumer := func(dataStream <-chan int ) {
		for data := range dataStream {
			fmt.Printf("Recived data=%v\n", data)
		}
		fmt.Println("Done")
	}

	stream := chanOwner()
	consumer(stream)

}
