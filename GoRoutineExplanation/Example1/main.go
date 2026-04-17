package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Order struct {
	ID     int
	Status string
}

func main() {
	fmt.Println("hello")
	orders := GenerateOrder(20)
	ProcessOrder(orders)
	UpdateOrderStatus(orders)
	reportOrderStatus(orders[0])
	fmt.Println(orders)
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
	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Println("Updating order status", order.ID)
	}
	status := []string{"Processing", "Shipped", "Delivered"}[rand.Intn(3)]
	for _, order := range orders {
		order.Status = status
	}
}

func reportOrderStatus(order *Order) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("Order status:", order.Status)

	}
	fmt.Println("Order status:", order.Status)
}
