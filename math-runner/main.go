package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/LiveSocket/bot/math-runner/runner"
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

func main() {
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

func registerCommand(service *service.Service, name string, proc string, restricted bool) {
	channels, err := getChannels(service)
	if err != nil {
		panic(err)
	}

	for _, channel := range channels {
		_, err = service.SimpleCall("private.command.register", nil, wamp.Dict{
			"proc":       proc,
			"channel":    channel,
			"name":       name,
			"activated":  true,
			"restricted": restricted,
		})
		if err != nil {
			panic(err)
		}
	}
}

func getChannels(service *service.Service) ([]string, error) {
	result, err := service.SimpleCall("private.channel.get", nil, nil)
	if err != nil {
		return nil, err
	}
	results := []string{}
	for _, item := range result.Arguments {
		results = append(results, item.(map[string]interface{})["name"].(string))
	}
	return results, nil
}
