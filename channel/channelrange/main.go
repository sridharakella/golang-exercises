package main

import "fmt"

func main() {
	c := make(chan int)

	go func(a, b int) {
		for i := 0; i < 10; i++ {
			c <- a + b
		}
		close(c)
	}(1, 2)

	for value := range c {
		fmt.Println(value)
		//defer close(c)
	}

}
