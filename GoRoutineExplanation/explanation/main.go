package main

import (
	"fmt"
	"sync"
)

func method1(value <-chan int) {
	fmt.Println("Hello, World!")
}

func method2(value <-chan int) {
	fmt.Println(<-value)

}

func main() {
	fmt.Println("Hello, World!")

	channel := make(chan int, 3)
	channel2 := make(chan int, 3)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		channel <- 1
		method1(channel)
	}()
	go func() {
		defer wg.Done()

		channel2 <- 2
		method2(channel2)
	}()

	wg.Wait()

	printchannel := <-channel
	fmt.Println(printchannel)

}
