package main

import (
	"fmt"
	"time"
)

type User struct {
	Name string
}

func main() {

	messages := make(chan string) // Unbuffered channel
	users := make(chan User)

	go func() {
		fmt.Println("sending message")
		users <- User{"sridhar"}

	}()

	go func() {
		fmt.Println("sending message")
		messages <- "hello"
		fmt.Println(<-messages)
	}()
	go func() {
		fmt.Println("sending message")
		messages <- "hello"
		fmt.Println(<-messages)
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("last message")
	msg := <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)

	usr := <-users
	fmt.Println(usr)
}
