package main

import (
	"fmt"
)

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("hello")
	a := MyStruct{"sridhar", 10}
	a.method()
	b := &MyStruct{}
	b.method2()

}

type MyStruct struct {
	c string
	b int
}

func (m *MyStruct) method() {
	fmt.Println("method with pointer")
	m.c = "sridhar"
	fmt.Println(m.c)
	m.b = 50
	fmt.Println(m.b)
}
func (m MyStruct) method2() {
	fmt.Println("without pointer")
	m.b = 10
	fmt.Println(m.b)
	m.c = "ravi"
	fmt.Println(m.c)
}
