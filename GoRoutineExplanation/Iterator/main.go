package main

import (
	"fmt"
	"iter"
)

type Iterator[V any] struct {
	seq iter.Seq[V]
}

func From[V any](seq iter.Seq[V]) Iterator[V] {
	return Iterator[V]{seq: seq}
}
func main() {
	//iterator is a function that produces one at time, with create a collection at a time.
	// go function yield involves loop body.

	//result:= From([]int {1,2,3,4,5,7,8}).Filter(func(i int) bool{ return i%2 == 0}).Map(func(i int) int{ return i*10})
	//fmt.Println(result)

	//From([]int{1,2,3,4,5,6,7,8}).Filter(func(i int) bool{ return i%2 == 0}).Map(func(i int) int{ return i*10}).ForEach(func(i int) { fmt.Println(i) })

	fmt.Println("hello")
	s := "hello"
	for _, v := range s {
		fmt.Println(v)
	}

}
