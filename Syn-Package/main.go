package main

import (
	"fmt"
	"sync"
)

func main() {
	var data int = 0
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		data++
		defer wg.Done()
	}()

	wg.Wait()
	fmt.Println("the value in the data", data)

}
