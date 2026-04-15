package main

//concurrency - refers structure that can handle multiple task, this give illusion running at same but not right, it uses context swtiching
//context switching between tasks

// parallelism- tasks can literally run in parallel at same time in multi core processor

import (
	"fmt"
	_ "math/rand"
	"sync"
)

type Order struct {
	ID          int
	OrderStatus string
	mu          sync.Mutex
}

var (
	totatOrderCount = 0
	updateMutex     sync.Mutex
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	orderChan := make(chan *Order, 20)

	go func() {
		for _, order := range generateOrder(20) {
			fmt.Println(order)
			orderChan <- order
		}
		defer wg.Done()
		fmt.Println("done with goroutine")

	}()

	go processOrder(orderChan, &wg)

	wg.Wait()
	//UpdateOrderStatus(orders)

}

func processOrder(orderChan <-chan *Order, wg *sync.WaitGroup) {

	for order := range orderChan {
		fmt.Println("Processing order:", order.ID)
		order.mu.Lock()
		order.OrderStatus = "processing"
		order.mu.Unlock()
	}

	defer wg.Done()

}
func UpdateOrderStatus(order []*Order) {

	OrderStatus := "processing"
	fmt.Println("Order status updated to:", OrderStatus)

}

func generateOrder(count int) []*Order {
	orders := make([]*Order, count)

	for i := 0; i < 10; i++ {
		orders[i] = &Order{ID: i + 1, OrderStatus: "pending"}
	}

	return orders

}
