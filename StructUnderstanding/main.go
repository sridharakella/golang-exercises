package StructUnderstanding

import (
	"fmt"
)

type User struct {
	Name string
}

var a User
var b User

type MyStruct struct {
	c string
}

type NewStruct struct {
	Name string
}

func main() {
	fmt.Println("hello")
	a := User{"sridhar"}
	fmt.Println(a.Name)
	ab := MyStruct{c: "sridhar"}
	fmt.Println(ab.c)
	
}
