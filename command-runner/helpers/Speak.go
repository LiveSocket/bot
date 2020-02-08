package helpers

import (
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// Speak WAMP call helper to speak a message in chat
func Speak(service *service.Service, channel string, message string) error {
	// Call twitch chat say endpoint
	_, err := service.SimpleCall("private.twitch.chat.say", wamp.List{channel, message}, nil)
	return err
}
