package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/LiveSocket/bot/mod-service/actions"
	"github.com/LiveSocket/bot/mod-service/migrations"
	"github.com/LiveSocket/bot/mod-service/topics"
	"github.com/LiveSocket/bot/service"
)

func main() {
	s := &service.Service{}
	close := s.Init(service.Actions{
		"private.mod.create":     actions.Create(s),
		"private.mod.destroy":    actions.Destroy(s),
		"private.mod.destroyAll": actions.DestroyAll(s),
		"private.mod.get":        actions.Get(s),
		"private.mod.isMod":      actions.IsMod(s),
	}, service.Subscriptions{
		"event.channel.created":   topics.ChannelCreated(s),
		"event.channel.destroyed": topics.ChannelDestroyed(s),
	}, "__mod_service", migrations.CreateModTable)
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
