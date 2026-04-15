package main

import "time"

func main() {

	ch1 := make(chan string, 1)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- "hello"
	}()

	select {
	case msg := <-ch1:
		println(msg)
	case <-time.After(1 * time.Second):
		println("timeout")
	}

}
