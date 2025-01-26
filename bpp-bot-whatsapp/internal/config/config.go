package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// We'll store global config variables here for easy access
var (
	WhatsAppAccessToken   string
	WhatsAppPhoneNumberID string
	VerifyToken           string
)

// LoadEnv loads environment variables from .env or system
func LoadEnv() {
	// Optional: load from a .env file (needs "github.com/joho/godotenv" dependency)
	_ = godotenv.Load()

	WhatsAppAccessToken = os.Getenv("WHATSAPP_ACCESS_TOKEN")
	WhatsAppPhoneNumberID = os.Getenv("WHATSAPP_PHONE_NUMBER_ID")
	VerifyToken = os.Getenv("WHATSAPP_VERIFY_TOKEN")

	// Basic check
	if WhatsAppAccessToken == "" || WhatsAppPhoneNumberID == "" || VerifyToken == "" {
		log.Println("Warning: Some env variables are missing. Check WHATSAPP_ACCESS_TOKEN, WHATSAPP_PHONE_NUMBER_ID, WHATSAPP_VERIFY_TOKEN.")
	}
}
