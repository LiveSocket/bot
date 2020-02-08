package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/LiveSocket/bot/mod-chat-service/actions"
	"github.com/LiveSocket/bot/mod-chat-service/migrations"
	"github.com/LiveSocket/bot/service"
)

func main() {
	s := &service.Service{}

	close := s.Init(service.Actions{
		"public.modchat.get":  actions.Get(s),
		"public.modchat.send": actions.Send(s),
	}, nil, "__mod_chat_service", migrations.CreateModChatTable)
	defer close()

	// Wait for CTRL-c or client close while handling events.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	select {
	case <-sigChan:
	case <-s.Done():
		log.Print("Router gone, exiting")
		return
	}
}
