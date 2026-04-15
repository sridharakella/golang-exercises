package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4) // sets logical processor

	var balance int
	var wg sync.WaitGroup

	deposit := func(amount int) {
		balance = balance + amount
	}

	withdraw := func(amount int) {
		balance -= amount

	}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			withdraw(1)
		}()
	}
	wg.Wait()
	fmt.Println(balance)
}
