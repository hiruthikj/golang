package main

import (
	"log"
	"net/http"
)

func checkStatus(done <-chan interface{}, urls ...string) <-chan *http.Response {
	responses := make(chan *http.Response)
	go func() {
		defer close(responses)

		for _, url := range urls {
			resp, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
				continue
			}
			select {
			case <-done:
				return
			case responses <- resp:
			}
		}
	}()

	return responses

}

func main() {
	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://google.com", "https://badhost"}

	responses := checkStatus(done, urls...)
	for response := range responses {
		log.Println("Response:", response.Status)
	}

}
