package main

func main() {

	chann := make(chan int, 10)

	go func(a, b int) {
		defer close(chann)
		for i := 0; i < 10; i++ {
			chann <- a + b
		}
	}(1, 2)

	for v := range chann {
		println(v)
	}
}
