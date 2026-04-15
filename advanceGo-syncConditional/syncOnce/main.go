package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var once sync.Once

	load := func() {
		fmt.Println("Run only once initialization function")
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			once.Do(load)
			//TODO: modify so that load function gets called only once.
			//load()
			// lwe want the load function to be called only once
		}()
	}
	wg.Wait()
}
