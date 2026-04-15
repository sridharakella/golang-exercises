package main

import "fmt"

func main() {

	channal := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, char := range chars {
		channal <- char
	}
	//close(channal)

	for result := range channal {
		fmt.Println(result)
	}

}
