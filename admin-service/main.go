package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/LiveSocket/bot/admin-service/actions"
	"github.com/LiveSocket/bot/admin-service/migrations"
	"github.com/LiveSocket/bot/service"
)

func main() {
	s := &service.Service{}
	close := s.Init(service.Actions{
		"private.admin.get":     actions.Get(s),
		"private.admin.isAdmin": actions.IsAdmin(s),
	}, nil, "__admin_service", migrations.CreateAdminTable)
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
