package main

import (
	"fmt"
	"math/rand"
)

// function as parameter
func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			default:
				stream <- fn()
			}
		}
	}()

	return stream
}

func main() {
	fmt.Println("started")
	done := make(chan int)
	defer close(done)
	// generator
	randomNumberFetcher := func() int { return rand.Intn(100) }
	repeatFunc(done, randomNumberFetcher)
}
