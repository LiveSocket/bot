package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/super-command-runner/runner"
)

func main() {
	// Create service
	s := &service.Service{}
	close := s.Init(nil, service.Subscriptions{
		"event.chat.message": runner.OnMessage(s),
	}, "")
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
