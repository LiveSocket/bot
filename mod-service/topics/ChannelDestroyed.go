package topics

import (
	"log"

	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

func ChannelDestroyed(service *service.Service) func(*socket.Event) {
	return func(event *wamp.Event) {
		if event.ArgumentsKw["channel"] == nil {
			log.Print("ChannelDestroyed: No channel")
		}

		channel := event.ArgumentsKw["channel"]
		_, err := service.Socket.SimpleCall("private.mod.destroyAll", nil, wamp.Dict{"channel": channel})
		if err != nil {
			log.Print(err)
		}
	}
}
