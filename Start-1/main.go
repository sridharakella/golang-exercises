package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args
	fmt.Println(args)
	fmt.Println("hello")
	fmt.Printf("hello %s\n", args[1])
	fmt.Println("hello world\nos.Args: %v\nArguments: %v\n", args, os.Args, args[1:3])

}
