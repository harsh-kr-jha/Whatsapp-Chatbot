package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/harsh-kr-jha/Whatsapp-Chatbot/bpp-bot-whatsapp/internal/config"
	"github.com/harsh-kr-jha/Whatsapp-Chatbot/bpp-bot-whatsapp/internal/services"
)

// WebhookHandler is the main entry for WhatsApp to POST message updates
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Verification
		verifyWebhook(w, r)
	case http.MethodPost:
		handleIncomingMessage(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// verifyWebhook handles the GET request to verify the webhook
func verifyWebhook(w http.ResponseWriter, r *http.Request) {
	mode := r.URL.Query().Get("hub.mode")
	token := r.URL.Query().Get("hub.verify_token")
	challenge := r.URL.Query().Get("hub.challenge")

	if mode == "subscribe" && token == config.VerifyToken {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(challenge))
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

// handleIncomingMessage parses inbound WhatsApp messages (POST)
func handleIncomingMessage(w http.ResponseWriter, r *http.Request) {
	var data services.WhatsAppWebhook
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Printf("Failed to decode webhook message: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Let the service layer handle the actual logic
	services.ProcessWebhookData(data)

	// Acknowledge
	w.WriteHeader(http.StatusOK)
}
