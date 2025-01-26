package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/harsh-kr-jha/Whatsapp-Chatbot/bpp-bot-whatsapp/internal/config"
)

// sendWhatsAppMessage calls the Cloud API to send messages
func sendWhatsAppMessage(messageBody map[string]interface{}) error {
	phoneNumberID := config.WhatsAppPhoneNumberID
	accessToken := config.WhatsAppAccessToken

	if phoneNumberID == "" || accessToken == "" {
		return fmt.Errorf("missing WhatsApp credentials")
	}

	url := "https://graph.facebook.com/v16.0/" + phoneNumberID + "/messages"

	payload, err := json.Marshal(messageBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	log.Printf("WhatsApp API response: %s", string(respBody))

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("non-2xx response: %d", resp.StatusCode)
	}

	return nil
}

// SendTextMessage for a simple text
func SendTextMessage(to, text string) error {
	message := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                to,
		"type":              "text",
		"text": map[string]string{
			"body": text,
		},
	}
	return sendWhatsAppMessage(message)
}

// SendButtonMessage sends an interactive button
func SendButtonMessage(to, body, footer string, buttons []map[string]string) error {
	message := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                to,
		"type":              "interactive",
		"interactive": map[string]interface{}{
			"type": "button",
			"body": map[string]string{
				"text": body,
			},
			"footer": map[string]string{
				"text": footer,
			},
			"action": map[string]interface{}{
				"buttons": buildButtonObjects(buttons),
			},
		},
	}

	return sendWhatsAppMessage(message)
}

func buildButtonObjects(buttons []map[string]string) []map[string]interface{} {
	var result []map[string]interface{}
	for _, btn := range buttons {
		result = append(result, map[string]interface{}{
			"type":  "reply",
			"reply": btn,
		})
	}
	return result
}

// SendListMessage example for list
func SendListMessage(to, headerText, bodyText, footerText, buttonText string, sections []map[string]interface{}) error {
	message := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                to,
		"type":              "interactive",
		"interactive": map[string]interface{}{
			"type": "list",
			"header": map[string]string{
				"type": "text",
				"text": headerText,
			},
			"body": map[string]string{
				"text": bodyText,
			},
			"footer": map[string]string{
				"text": footerText,
			},
			"action": map[string]interface{}{
				"button":   buttonText,
				"sections": sections,
			},
		},
	}
	return sendWhatsAppMessage(message)
}
