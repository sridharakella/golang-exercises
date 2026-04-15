package main

import (
	"awesomeProject2/internal/color"
	"fmt"
)

func main() {
	fmt.Println("colorer is working")
	redtext := color.Text(color.Red, "text")
	fmt.Println(redtext)
	yellowtext := color.Text(color.Yellow, "text")
	fmt.Println(yellowtext)
}
