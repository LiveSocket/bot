package helpers

import (
	"github.com/LiveSocket/bot/mod-chat-service/models"
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// EmitMessage Emits "public.event.modchat.message.<channel>"
func EmitMessage(service *service.Service, modChat *models.ModChat) error {
	return service.Publish("public.event.modchat.message."+modChat.Channel, nil, wamp.List{modChat}, nil)
}
