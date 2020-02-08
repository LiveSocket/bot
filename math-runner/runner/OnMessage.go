package runner

import (
	"log"
	"regexp"

	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
)

// OnMessage handles detection and calculation of math in messages
func OnMessage(s *service.Service) func(*socket.Event) {
	return func(event *socket.Event) {
		// If message exists
		if len(event.Arguments) > 0 {
			// Convert json message to twitch.PrivateMessage
			message := service.MapToPrivateMessage(event.Arguments[0].(map[string]interface{}))

			// Check if message is math
			regex := regexp.MustCompile(`^\s*\(?(\-?\+?[0-9.]+)\s*([+*-/%x^])\s*(\-?\+?[0-9.]+)\)?`)
			if match := regex.FindStringSubmatch(message.Message); match != nil {
				log.Print("Math detected")
				// Run math
				go RunMath(s, message, match[1], match[3], match[2])
			}
		}
	}
}
