package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		fmt.Println("usage : /http-get")
		os.Exit(1)
	}

	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Println("URL is inavlid format:", err)
		os.Exit(1)
	}
	resp, err := http.Get(args[1])
	fmt.Println(resp.Status)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
