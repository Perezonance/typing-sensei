package main

import (
	"fmt"
	"net/http"
)

const (
	root = "/api/v1"
)

func main() {
	fmt.Println(startServer())
}

func startServer() error {
	http.HandleFunc(root + "/lexicons/public", handlers.GenLexHandler)
	http.HandleFunc(root + "/lexicons/:user", handlers.UserLexHandler)


	return nil
}