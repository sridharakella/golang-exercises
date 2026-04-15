package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")

	done := make(chan bool)
	go dowork(done)
	time.Sleep(2 * time.Second)
	close(done)
}

func dowork(done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("done channel executed")
			return
		default:
			fmt.Println("working")
		}
	}

}
