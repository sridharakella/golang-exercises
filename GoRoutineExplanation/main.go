package main

import (
	"fmt"
)

func call() {
	charchannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}
	for _, char := range chars {
		charchannel <- char
	}
	charchannel <- "hello"
	close(charchannel)

	for result := range charchannel {
		fmt.Println(result)
	}
}

func main() {
	// concurrency primitives

	go func() {
		for index := 0; index < 3; index++ {
			call()

		}
	}()

	chanel := make(chan string)
	anotherchanel := make(chan string)

	go func() {
		fmt.Println("world")
		chanel <- "message"
	}()

	go func() {
		anotherchanel <- "another message"
	}()

	//channelresponse := <-chanel // this channel will be blocking fork and join ( join is called) jointpoint.
	//fmt.Println(channelresponse)
	// select statemetn is block the unless the case

	// for select loop
	// done channel
	// pipelines

	for {
		select {
		case channelresponse := <-chanel:
			fmt.Println(channelresponse)
			break
		case channelresponse2 := <-anotherchanel:
			fmt.Println(channelresponse2)
			break
		default:
			fmt.Println("no message")
			break
		}
	}

	select {
	case channelresponse := <-chanel:
		fmt.Println(channelresponse)
	case channelresponse2 := <-anotherchanel:
		fmt.Println(channelresponse2)
	default:
		fmt.Println("no message")
	}
}
