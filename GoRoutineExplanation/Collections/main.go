package main

import (
	"container/heap" // heap implementation
	"container/list" // double linked list
	"container/ring"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // min-heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(map1)

	var gridy [3][3]int // rare to use
	fmt.Println(gridy)

	channel := make(chan int)
	fmt.Println(channel)

	//Double linked list
	l := list.New()
	l.PushBack("first") // append
	l.PushFront("zero") // prepend
	e := l.PushBack("second")
	l.InsertBefore("between", e)

	// iterate
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) // zero, first, between, second
	}

	l.Remove(e)
	// end of Linked list

	// heap implementation
	h := &IntHeap{5, 2, 8, 1}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Println(heap.Pop(h)) // 1 (smallest)

	// end of heap implementation

	//circular ring
	r := ring.New(3) // ring of size 3

	// fill it
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}

	// iterate — goes around the ring once
	r.Do(func(v any) {
		fmt.Println(v) // 0, 1, 2
	})

	// it wraps around
	fmt.Println(r.Next().Next().Next().Value) // back to start

	// end of circular ring

	// stack implementation ( stack is just a slice with push and pop operations)
	stack := []int{}
	stack = append(stack, 1)   // push
	top := stack[len(stack)-1] // peek
	stack = stack[:len(stack)-1]
	fmt.Println(top)

	// end of stack implementation

	// queue implementation ( queue is just a slice with push and pop operations)
	queue := []int{}
	queue = append(queue, 1) // enqueue
	front := queue[0]        // peek
	queue = queue[1:]        // dequeue (copies — O(n))

	fmt.Println(front)

	// end of queue implementation

	// set implementation, set is a map with no duplicate values,set uses a map with empty struct as value
	set := map[string]struct{}{}
	set["alice"] = struct{}{} // add
	_, exists := set["alice"] // check membership
	delete(set, "alice")      // remove
	fmt.Println(exists)
	// end of set implementation

	//Tree / BST / Trie — no built-in. You'd write your own or use a third-party library. With generics (Go 1.18+), this is much more ergonomic than before.

}
