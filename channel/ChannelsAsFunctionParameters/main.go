package main

import "fmt"

// TODO: Implement relaying of message with Channel Direction

func genMsg(ch1 chan<- string) {
	// send message on ch1
	ch1 <- "hello"
}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	// recv message on ch1
	msg := <-ch1
	fmt.Println(msg)
	ch2 <- msg
	// send it on ch2

}

func main() {
	// create ch1 and ch2
	ch1 := make(chan string)
	ch2 := make(chan string)

	// spine goroutine genMsg and relayMsg
	go genMsg(ch1)
	// recv message on ch2
	go relayMsg(ch1, ch2)

	v := <-ch2
	fmt.Println(v)
}
