package main

import "fmt"

func channeldirection(in <-chan int, out chan<- int) (int, int) {
	return 0, 0
}

func main() {
	var data int

	go func() {
		data++
	}()
	if data == 0 {
		fmt.Println("data is 0")
	}
	// sometimes data is not incremented
}
