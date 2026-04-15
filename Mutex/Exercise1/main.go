package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(4)
	var balance int
	var wg sync.WaitGroup

	var mu sync.Mutex

	deposit := func(amount int) {
		mu.Lock()
		balance += amount
		mu.Unlock()
		fmt.Println("balance before", balance)
	}
	fmt.Println("balance before", deposit)
	withdraw := func(amount int) {
		mu.Lock()
		defer mu.Unlock()
		balance -= amount

		fmt.Println("balance with", balance)
	}
	fmt.Println("balance before", withdraw)

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

	fmt.Println("balance before", balance)

	wg.Wait()
	fmt.Println("balance after", balance)

	fmt.Println()
}
