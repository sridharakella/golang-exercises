package main

import "fmt"

func main() {
	//TODO: create channel owner goroutine which return channel and
	// writes data into channel and
	// closes the channel when done.
	owner := func() <-chan int {

		ch := make(chan int)

		go func() {
			defer close(ch)
			for i := 0; i < 10; i++ {
				ch <- i
			}
		}()
		return ch
	}

	consumer := func(ch <-chan int) {
		fmt.Println("Consumer channel")
		for v := range ch {
			println(v)
		}
		fmt.Println("Done")
	}

	ch := owner()
	consumer(ch)

}
