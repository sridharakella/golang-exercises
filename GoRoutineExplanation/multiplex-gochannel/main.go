package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")
	chanel1 := make(chan string)
	chanel2 := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		chanel1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		chanel2 <- "two"
	}()

	select {
	case msg := <-chanel1:
		fmt.Println(msg)
	case msg2 := <-chanel2:
		fmt.Println(msg2)
	default:
		fmt.Println("no message")
	}
}
