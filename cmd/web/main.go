package main

import (
	"log"
	"net/http"

	"github.com/xuoxod/ws/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Starting channel listener")

	go handlers.ListenToWsChannel()

	log.Println("Web server running on port 8080")

	_ = http.ListenAndServe(":8080", mux)

}
