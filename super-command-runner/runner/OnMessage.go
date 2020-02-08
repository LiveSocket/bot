package runner

import (
	"log"
	"regexp"

	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/LiveSocket/bot/super-command-runner/helpers"
	"github.com/gempir/go-twitch-irc/v2"
)

// OnMessage Handler for detection and execution of super commands in chat messages
func OnMessage(s *service.Service) func(*socket.Event) {
	return func(event *socket.Event) {
		// If message exists
		if len(event.Arguments) > 0 {
			// Convert json message to twitch.PrivateMessage
			message := service.MapToPrivateMessage(event.Arguments[0].(map[string]interface{}))

			// Process message
			go processMessage(s, message)
		}
	}
}

func processMessage(service *service.Service, message *twitch.PrivateMessage) {
	// Get channel record
	result, err := helpers.GetChannel(service, message.Channel)
	if err != nil {
		log.Print(err)
		return
	}
	// If no channel is found
	if result == nil {
		log.Printf("%s channel not found", message.Channel)
	}
	// Check if message is super command
	regex := regexp.MustCompile(`^@?SnareChopsBot`)
	if regex.MatchString(message.Message) {
		log.Print("Super Command Detected")
		go RunSuperCommand(service, message)
	}
}
