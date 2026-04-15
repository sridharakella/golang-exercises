package main

import "time"

func main() {

	chans := make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second)
		chans <- "hello"

	}()

	// implement non blocking wiht select and default statement.
	for i := 0; i < 2; i++ {
		select {
		case msg := <-chans:
			println(msg)
		default:
			println("no message")
		}
	}

}
