package actions

import (
	"errors"

	"github.com/LiveSocket/bot/command-runner/models"
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

type getCustomInput struct {
	Channel string
	Name    string
}

// GetCustom Gets a list of custom commands for a channel
//
// private.command.custom.get
//
// To get list of custom commands for channel
// {channel string}
// Returns [CustomCommand...]
//
// To get specific custom command for channel
// {name string, channel string}
// Returns [CustomCommand]
func GetCustom(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getGetCustomInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Check if name is provided
		if input.Name != "" {
			// Get single custom command
			return getSingle(service, input)
		}
		// Get all custom commands for channel
		return getAll(service, input)
	}
}

func getSingle(service *service.Service, input *getCustomInput) socket.Result {
	// Find custom command
	command, err := models.GetCustom(service, input.Name, input.Channel)
	if err != nil {
		return socket.Error(err)
	}

	return socket.Success(command)
}

func getAll(service *service.Service, input *getCustomInput) socket.Result {
	// Find custom commands
	commands, err := models.GetCustoms(service, input.Channel)
	if err != nil {
		return socket.Error(err)
	}

	// Create wamp.List of results
	list := wamp.List{}
	for _, item := range commands {
		list = append(list, item)
	}

	// Return list of custom commands
	return socket.Success(list...)
}

func getGetCustomInput(kwargs wamp.Dict) (*getCustomInput, error) {
	if kwargs["channel"] == nil {
		return nil, errors.New("Missing channel")
	}

	input := &getCustomInput{Channel: conv.ToString(kwargs["channel"])}
	if kwargs["name"] != nil {
		input.Name = conv.ToString(kwargs["name"])
	}

	return input, nil
}
