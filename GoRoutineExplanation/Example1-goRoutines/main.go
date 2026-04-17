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

	orders := GenerateOrder(20)
	//go ProcessOrder(orders)
	go func() {

		time.Sleep(time.Second * 1)
		defer wg.Done()
		ProcessOrder(orders)

	}()

	go func() {
		time.Sleep(time.Second * 2)
		defer wg.Done()
		//UpdateOrderStatus(orders)
		for i := 0; i < 10; i++ {
			UpdateOrderStatus(orders)
		}
		//UpdateOrder(orders[0])
	}()
	go func() {
		time.Sleep(time.Second * 3)
		defer wg.Done()
		reportOrderStatus(orders[0])

	}()
	//go UpdateOrderStatus(orders)
	//go reportOrderStatus(orders[0])
	fmt.Println(orders)
	defer wg.Wait()
	currentUpdates := totalUpdates
	totalUpdates := currentUpdates
	fmt.Println(totalUpdates)
	fmt.Println("Total Updates:", currentUpdates)
	fmt.Println("All Operations Completed")
}

func GenerateOrder(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{ID: i, Status: "Pending"}
	}
	return orders

}

func ProcessOrder(orders []*Order) {
	for _, order := range orders {
		time.Sleep(time.Second)
		fmt.Println("Processing order", order.ID)
	}

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
