package main

import (
	"github.com/perezonance/typing-sensei/lexicon-management-service/src/pkg/handlers"

	"fmt"
	"net/http"
)

const (
	root = "/api/v1"
	port = "8080"
	addr = "0.0.0.0"

	fullAddr = addr + ":" + port
)

func main() {
	fmt.Println(startServer())
}

func startServer() error {
	//API for operations on public Lexicons
	http.HandleFunc(root + "/lexicons/public", handlers.PublicLexHandler)

	//API for operations on user Lexicons
	http.HandleFunc(root + "/lexicons/user/:user_id", handlers.UserLexHandler)

	fmt.Printf("Server listening on port:%v", port)

	//TODO: change to ListenAndServeTLS()... maybe...
	return http.ListenAndServe(fullAddr, nil)
}