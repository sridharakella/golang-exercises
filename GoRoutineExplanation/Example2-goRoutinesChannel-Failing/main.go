package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

var (
	totalUpdates int
	updateMutex  sync.Mutex
)

func main() {
	fmt.Println("hello")
	var wg sync.WaitGroup

	wg.Add(3)

	orderChan := make(chan *Order)

	//orders := GenerateOrder(20)
	//go ProcessOrder(orders)
	go func() {

		time.Sleep(time.Second * 1)
		defer wg.Done()
		for order := range GenerateOrder(20) {
			orderChan <- order
			break
		}
		close(orderChan)
	}()

	go ProcessOrder(&wg, orderChan)

	//go UpdateOrderStatus(orders)
	//go reportOrderStatus(orders[0])
	//fmt.Println(orders)
	defer wg.Wait()
	currentUpdates := totalUpdates
	totalUpdates := currentUpdates
	fmt.Println(totalUpdates)
	fmt.Println("Total Updates:", currentUpdates)
	fmt.Println("All Operations Completed")
}

func GenerateOrder(count int) chan *Order {
	orderChan := make(chan *Order, count)
	go func() {
		for i := 0; i < count; i++ {
			orderChan <- &Order{ID: i, Status: "Pending"}
		}
		close(orderChan)
	}()
	return orderChan
}

func ProcessOrder(wg *sync.WaitGroup, orders <-chan *Order) {
	for order := range orders {
		fmt.Println("Processing order:", order.ID)
		order.mu.Lock()
		order.Status = "Processing"
		order.mu.Unlock()
	}
	defer wg.Done()

}
func UpdateOrderStatus(orders []*Order) {
	updateMutex.Lock() // global lock
	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Println("Updating order status", order.ID)
	}
	status := []string{"Processing", "Shipped", "Delivered"}[rand.Intn(3)]
	for _, order := range orders {
		order.Status = status
	}
	updateMutex.Unlock() // global unlock

}

func UpdateOrder(Order *Order) *Order {
	fmt.Println("Updating order status", Order.ID)
	Order.mu.Lock()

	Order.Status = "Processing"

	Order.mu.Unlock()

	return Order
}

func reportOrderStatus(order *Order) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("Order status:", order.Status)

	}
	fmt.Println("Order status:", order.Status)
}
