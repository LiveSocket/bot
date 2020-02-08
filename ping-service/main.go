package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/LiveSocket/bot/ping-service/actions"
	"github.com/LiveSocket/bot/service/socket"
)

func main() {
	s := &socket.Socket{}
	close := s.Init(socket.Actions{
		"public.ping": actions.Ping(s),
	}, nil)
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
