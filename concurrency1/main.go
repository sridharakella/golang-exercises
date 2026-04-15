package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(message string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(delay)
	fmt.Println("Hello, sridhar message!", delay)
}

func main() {
	var wg sync.WaitGroup

	/*
			// add outside of the go routine
			// you must dcrease the counter wg.Done() inside the go routine
			// Don't foget to call wg.wait()
		// always pass the reference of wg group variable, not a reference.

	*/
	totalJobs := 5

	for i := 0; i < totalJobs; i++ {
		wg.Add(1)
		go sayHello("Hello", time.Second*2, &wg)
	}
	wg.Add(3)
	fmt.Println("Hello() from goroutine")

	go sayHello("Hello", time.Second*2, &wg)
	go sayHello("Hi", time.Second*1, &wg)
	go sayHello("Hey", time.Second*0, &wg)
	go sayHello("hello", time.Second*5, &wg)
	//time.Sleep(time.Second * 3)
	fmt.Println("last message")
	wg.Wait()
}
