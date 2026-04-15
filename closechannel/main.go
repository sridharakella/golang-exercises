package main

import "fmt"

func main() {
	jobs := make(chan int, 5)

	done := make(chan bool)

	go func() {
		for {
			r, ok := <-jobs
			if ok {
				fmt.Println("Got this message:", r)

			} else {
				fmt.Println("Channel closed")
				done <- true
				return
			}

			fmt.Println(r)
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("Sending message:", i)

	}

	close(jobs)

	<-done
}
