package main

import (
	"fmt"
	"net/http"
)

const (
	root = "/api/v1/test"
)

func main() {
	fmt.Println("starting up server")
	fmt.Println(startServer())
}

func startServer() error {
	http.HandleFunc(root + "/random", h.)
}