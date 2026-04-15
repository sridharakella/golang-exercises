package main

import "time"

func main() {

	ch1 := make(chan string) // lets says we need to cache and registeries and state
	// mutex is used to gaurd the resource
	
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "hello"
	}()

	select {
	case msg := <-ch1:
		println(msg)
	default:
		println("no message")
	}
}
