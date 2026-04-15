package main

import (
	"fmt"
	"time"
)

func method(ch chan int) {
	fmt.Println(<-ch)
	time.Sleep(time.Second)
}

func main() {
	ch := make(chan int)
	go method(ch)

}
