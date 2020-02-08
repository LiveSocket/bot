package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/LiveSocket/bot/command-service/actions"
	"github.com/LiveSocket/bot/command-service/migrations"
	"github.com/LiveSocket/bot/command-service/topics"
	"github.com/LiveSocket/bot/service"
)

func main() {
	s := &service.Service{}
	close := s.Init(service.Actions{
		"public.command.create":     actions.Create(s),
		"public.command.destroy":    actions.Destroy(s),
		"public.command.destroyAll": actions.DestroyAll(s),
		"public.command.get":        actions.Get(s),
		"public.command.getById":    actions.GetByID(s),
		"public.command.update":     actions.Update(s),
	}, service.Subscriptions{
		"event.channel.destroyed": topics.ChannelDestroyed(s),
	}, "__command_service", migrations.CreateCommandTable)
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
