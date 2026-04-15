package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Atomic")
	runtime.GOMAXPROCS(4)

	var counter uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {

			defer wg.Done()
			for j := 0; j < 1000; j++ {
				
				atomic.AddUint64(&counter, 1)

			}
		}()

	}
}
