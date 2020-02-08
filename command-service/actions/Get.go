package actions

import (
	"errors"

	"github.com/LiveSocket/bot/command-service/models"
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

type getInput struct {
	Channel string
}

// Get Get a list of commands for a channel
//
// public.command.get
// {channel string}
//
// Returns [Command...]
func Get(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *wamp.Invocation) socket.Result {
		// Get input args from call
		input, err := getGetInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Find all commands for channel
		commands, err := models.GetCommands(service, input.Channel)
		if err != nil {
			return socket.Error(err)
		}

		// Make genric list of commands
		list := make([]interface{}, len(commands))
		for i, command := range commands {
			list[i] = command
		}

		// Return list of commands
		return socket.Success(list...)
	}
}

func getGetInput(kwargs wamp.Dict) (*getInput, error) {
	if kwargs["channel"] == nil {
		return nil, errors.New("Missing channel")
	}

	return &getInput{Channel: conv.ToString(kwargs["channel"])}, nil
}
