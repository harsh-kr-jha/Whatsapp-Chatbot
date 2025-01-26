package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/harsh-kr-jha/Whatsapp-Chatbot/bpp-bot-whatsapp/internal/services"
)

// NotifyHandler receives requests from other microservices
// to send push notifications to users.
func NotifyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Phone   string `json:"phone"`
		Message string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Phone == "" || req.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing phone or message"))
		return
	}

	// Send a simple text message via the services layer
	if err := services.SendTextMessage(req.Phone, req.Message); err != nil {
		log.Printf("Error sending push notification: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Notification sent"))
}
