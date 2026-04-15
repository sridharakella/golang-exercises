package main

import "time"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Millisecond)
		ch1 <- "hello"
	}()

	go func() {
		time.Sleep(2 * time.Millisecond)
		ch2 <- "world"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			println(msg)
		case msg := <-ch2:
			println(msg)
		}
	}

}
