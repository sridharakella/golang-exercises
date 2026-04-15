package main

import (
	"fmt"
	"sync"
)

func main() {

	jobs := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			r, ok := <-jobs // command k pattern
			if ok {
				fmt.Println("Got this message:", r)
			} else {
				fmt.Println("Channel closed")
				return
			}
		}
		// how to puyt messages into channel
		for i := 1; i <= 3; i++ {
			jobs <- i
		}

		for i := 1; i <= 3; i++ {
			fmt.Println("Sending message:", i)
			jobs <- i

		}
	}(&wg)

	for i := 1; i <= 3; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func doublechannel() {
	jobs := make(chan int)
	done := make(chan bool)

	go func() {
		for {
			r, ok := <-jobs // command k pattern
			if ok {
				fmt.Println("Got this message:", r)
			} else {
				fmt.Println("Channel closed")
				done <- true
				return
			}
		}
		// how to puyt messages into channel
		for i := 1; i <= 3; i++ {
			jobs <- i
		}

		for i := 1; i <= 3; i++ {
			fmt.Println("Sending message:", i)
			jobs <- i

		}

	}()
	close(jobs)
	<-done

}
