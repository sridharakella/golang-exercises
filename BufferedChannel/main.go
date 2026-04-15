package main

import (
	"fmt"
)

func main() {
	fmt.Println("buffered channel")
	message := make(chan string, 3)
	message <- "hello"
	message <- "world"
	message <- "sridhar"
	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
	// closing the channel
	fmt.Println("channel closed")
	close(message)

}
