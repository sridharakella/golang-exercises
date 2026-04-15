package main

import (
	"fmt"
	"time"
)

func sayHello(message string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println("Hello, sridhar message!", delay)
}

func main() {
	fmt.Println("Hello() from goroutine")

	go sayHello("Hello", time.Second*2)
	go sayHello("Hi", time.Second*1)
	go sayHello("Hey", time.Second*0)
	time.Sleep(time.Second * 3)
	fmt.Println("last message")
}
