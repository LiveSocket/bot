package runner

import (
	"log"
	"regexp"

	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
)

// OnMessage Handler for detection and handling of commands when chat messages arive
func OnMessage(s *service.Service) func(*socket.Event) {
	return func(event *socket.Event) {
		// If message exists
		if len(event.Arguments) > 0 {
			// Convert json message to twitch.PrivateMessage
			message := service.MapToPrivateMessage(event.Arguments[0].(map[string]interface{}))
			println(message.Message)

			// Check is message contains a command
			regex := regexp.MustCompile(`^!.*`)
			if regex.Match([]byte(message.Message)) {
				log.Print("Command detected")
				// Run command
				go RunCommand(s, message)
			}
		}
	}
}
