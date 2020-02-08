package runner

import (
	"log"
	"strings"

	"github.com/LiveSocket/bot/description-runner/helpers"
	"github.com/LiveSocket/bot/service"

	"github.com/gempir/go-twitch-irc/v2"
)

func RunDescription(service *service.Service, message *twitch.PrivateMessage) {
	channel := message.Channel
	parts := strings.Split(message.Message, " ")
	name := parts[0][1:len(parts[0])]
	rest := parts[1:len(parts)]

	if len(rest) > 0 {
		editDescription(service, channel, name, message, rest)
	} else {
		sayDescription(service, channel, name, message)
	}
}

func speak(service *service.Service, channel string, message string) {
	// Only speak if there is something to say
	if channel != "" && message != "" {
		err := helpers.Speak(service, channel, message)
		if err != nil {
			// Log error
			log.Print(err)
		}
	}
}
