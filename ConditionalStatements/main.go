package main

import (
	"fmt"
	"os"
	"runtime"
)

func init() {
	fmt.Println("init")
	if runtime.GOOS == "windows" {
		fmt.Println("running on windows")
	} else if runtime.GOOS == "linux" {
		fmt.Println("running on linux")
	} else {
		fmt.Println("running on unknown OS")
	}
}

func main() {

	args := os.Args
	fmt.Println(args)

	if len(args) > 1 {
		fmt.Println("error throw")
		os.Exit(1)
	}

	fmt.Println("hello")
	if 10 > 5 {
		println("10 is greater than 5")
	}
}
