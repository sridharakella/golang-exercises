package main

import "fmt"

const (
	monday = iota + 1
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

// exercise from 01-exercise/o2-channel/05-channel
func main() {
	//TODO: create channel owner goroutine which return channel and
	// writes data into channel and
	// closes the channel when done.

	owner := func() <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; i < 10; i++ {
				ch <- i
			}
			defer close(ch)
		}()
		return ch
	}

	consumer := func(ch <-chan int) {
		// read values from channel
		for v := range ch {
			fmt.Printf("Received: %d\n", v)
		}
		fmt.Println("Done receiving!")
	}

	ch := owner()
	consumer(ch)
}
