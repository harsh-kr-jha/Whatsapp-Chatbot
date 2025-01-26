package services

import (
	"fmt"
	"log"
	"strings"

	"github.com/your-username/whatsapp-bot/internal/state"
)

// ProcessWebhookData receives the decoded webhook JSON and routes each incoming message
func ProcessWebhookData(data WhatsAppWebhook) {
	// Check minimal structure
	if len(data.Entry) == 0 || len(data.Entry[0].Changes) == 0 {
		return
	}

	messages := data.Entry[0].Changes[0].Value.Messages
	for _, msg := range messages {
		from := msg.From
		var userInput string

		// If it's a text message
		if msg.Text.Body != "" {
			userInput = msg.Text.Body
		}

		// If it's an interactive message (button or list)
		if msg.Interactive.Type == "button_reply" {
			userInput = msg.Interactive.ButtonReply.Title
		} else if msg.Interactive.Type == "list_reply" {
			userInput = msg.Interactive.ListReply.Title
		}

		// Go handle the conversation logic
		handleConversation(from, userInput)
	}
}

// handleConversation uses the state store and decides what to do next
func handleConversation(from, userInput string) {
	if strings.ToLower(userInput) == "restart" {
		state.ResetUserState(from)
	}

	currentStep := state.GetUserState(from)
	jobData := state.GetUserJobData(from)

	switch currentStep {
	case "":
		// No state means start
		log.Printf("New conversation with user %s", from)
		SendTextMessage(from, "Hello! I can help you post a job. Let's get started. What's the job title?")
		state.SetUserState(from, "waiting_job_title")

	case "waiting_job_title":
		jobData["title"] = userInput
		SendTextMessage(from, "Great. Now, please give a short job description.")
		state.SetUserState(from, "waiting_job_desc")

	case "waiting_job_desc":
		jobData["description"] = userInput
		body := "What type of job is this?"
		footer := "Choose one of the options below"
		buttons := []map[string]string{
			{"id": "full_time", "title": "Full-Time"},
			{"id": "part_time", "title": "Part-Time"},
			{"id": "contract", "title": "Contract"},
		}
		SendButtonMessage(from, body, footer, buttons)
		state.SetUserState(from, "waiting_job_type")

	case "waiting_job_type":
		jobData["type"] = userInput
		text := fmt.Sprintf("Job post summary:\nTitle: %s\nDescription: %s\nType: %s\n\nThank you!",
			jobData["title"], jobData["description"], jobData["type"])
		SendTextMessage(from, text)

		// Pretend we call microservice here:
		log.Printf("Saved job post for user %s: %+v", from, jobData)

		SendTextMessage(from, "If you want to post another job, type 'restart'.")
		state.SetUserState(from, "done")

	case "done":
		// user can either restart or do nothing
		SendTextMessage(from, "You have already posted a job. Type 'restart' to do another.")
	}
}
