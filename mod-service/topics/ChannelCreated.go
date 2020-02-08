package topics

import (
	"log"

	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

func ChannelCreated(service *service.Service) func(event *socket.Event) {
	return func(event *wamp.Event) {
		if event.ArgumentsKw["channel"] == nil {
			log.Print("ChannelCreated: No channel")
		}

		channel := event.ArgumentsKw["channel"]
		_, err := service.Socket.SimpleCall("private.mod.create", nil, wamp.Dict{"channel": channel, "username": channel})
		if err != nil {
			log.Print(err)
		}
	}
}
