package main

import (
	todo "first-go-api/internal"
	"first-go-api/internal/transport"
	"log"
)

func main() {

	//var todos = make([]TodoItem, 0)
	svc := todo.NewService()

	server := transport.NewServer(svc)

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}

}
