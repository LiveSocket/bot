package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/LiveSocket/bot/channel-service/actions"
	"github.com/LiveSocket/bot/channel-service/migrations"
	"github.com/LiveSocket/bot/channel-service/supers"
	"github.com/LiveSocket/bot/service"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	s := &service.Service{}
	close := s.Init(service.Actions{
		"private.channel.create":   actions.Create(s),
		"private.channel.delete":   actions.Delete(s),
		"private.channel.destroy":  actions.Destroy(s),
		"private.channel.get":      actions.Get(s),
		"private.channel.ignore":   actions.Ignore(s),
		"private.channel.unignore": actions.Unignore(s),
		"private.channel.update":   actions.Update(s),
		"super.channel":            supers.Channel(s),
		"super.ignore":             supers.Ignore(s),
	}, nil, "__channel_service", migrations.CreateChannelTable)
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
