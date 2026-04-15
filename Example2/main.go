package main

import (
	"fmt"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}

func main() {

	// main is also goroutine.
	//var wg sync.WaitGroup
	//wg.Add(3)

	// channel are used to communicate between goroutine

	//cncurrency is structred

	mychannel := make(chan string)

	go func() {
		mychannel <- "data"
	}()

	go someFunc("1") // that gives the run this concurrently
	go someFunc("2")
	go someFunc("3") // goroutine needs to some in 2, we need channel to communicate
	// send the message
	go someFunc("4")

	time.Sleep(time.Second * 2)

	fmt.Println("hello")

	msg := <-mychannel
	fmt.Println(msg)

}
