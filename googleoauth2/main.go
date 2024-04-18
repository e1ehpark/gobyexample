package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	server := &http.Server{
		Addr: fmt.Sprintf(":8000"),
		Handler: handlers.New(),
	}

	log.Printf("Starting Http Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClose {
		log.Printf(%v", err)
	} else {
		log.Println("Server closed!")
	}
}