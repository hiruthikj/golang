package main

import "sync"

func orDone[T any](done <-chan interface{}, mainChan <-chan T) <-chan T {
	combinedChan := make(chan T)
	go func() {
		defer close(combinedChan)
		for {
			select {
			case <-done:
				return
			case item, ok := <-mainChan:
				if !ok {
					return
				}
				select {
				case <-done:
					return
				case combinedChan <- item:

				}
			}
		}
	}()

	return combinedChan
}

func gen(done <-chan interface{}, nums ...int) <-chan int {
	stream := make(chan int)

	go func() {
		defer close(stream)
		for _, num := range nums {
			select {
			case <-done:
				return
			case stream <- num:
			}
		}
	}()

	return stream
}

func fanIn[T any](done <-chan interface{}, inboundChans ...<-chan T) <-chan T {
	outboundChan := make(chan T)

	go func() {
		defer close(outboundChan)
		var wg sync.WaitGroup
		wg.Add(len(inboundChans))

		for _, inboundChan := range(inboundChans) {
			go func(inboundChan <-chan T){
				defer wg.Done()
				for {
					select {
					case <-done:
						return
					case v, ok := <-inboundChan:
						if !ok {
							return
						}
						select {
						case <-done:
							return
						case outboundChan <- v:
						}
					}
				}
			}(inboundChan)
		}
		wg.Wait()
	}()


	return outboundChan
}