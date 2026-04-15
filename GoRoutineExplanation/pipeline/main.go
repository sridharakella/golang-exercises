package main

import "fmt"

func main() {
	//input
	nums := []int{1, 2, 3, 4, 5}
	//stage 1
	datachannel := slicetoChannel(nums)
	// stage 2
	finalChannel := sq(datachannel)
	//stage 3 output the pipeline
	for n := range finalChannel {
		fmt.Println(n)
	}
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out

}

func slicetoChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
