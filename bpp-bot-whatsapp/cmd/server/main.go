package main

import (
	"bpp-bot-whatsapp/config"
	"bpp-bot-whatsapp/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load environment variables / config
	config.LoadEnv() // we'll define this in config.go

	// Setup HTTP routes
	http.HandleFunc("/webhook", handlers.WebhookHandler)
	http.HandleFunc("/notify", handlers.NotifyHandler)

	// Determine port to listen on
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("WhatsApp bot server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
