package runner

import (
	"fmt"
	"log"
	"regexp"

	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

// OnMessage Handler for detection and handling of commands when chat messages arive
func OnMessage(s *service.Service) func(event *socket.Event) {
	return func(event *wamp.Event) {
		// If message exists
		if len(event.Arguments) > 0 {
			// Convert json message to twitch.PrivateMessage
			mMap, err := conv.ToStringMap(event.Arguments[0])
			if err != nil {
				fmt.Printf("%v", err)
				return
			}
			message := service.MapToPrivateMessage(mMap)
			println(message.Message)

			// Check is message contains a description
			regex := regexp.MustCompile(`^\?.+`)
			if regex.Match([]byte(message.Message)) {
				log.Print("Description detected")
				// Run command
				go RunDescription(s, message)
			}
		}
	}
}
