package actions

import (
	"errors"

	"github.com/LiveSocket/bot/command-service/models"
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

type getByIDInput struct {
	Channel string
	Name    string
}

// GetByID Gets a specific command
//
// public.command.getById
// {channel string, name string}
//
// Returns [Command]
func GetByID(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *wamp.Invocation) socket.Result {
		// Get input args from call
		input, err := getGetByIDInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Find command for channel and name
		command, err := models.GetCommand(service, input.Channel, input.Name)
		if err != nil {
			return socket.Error(err)
		}

		// Return command
		return socket.Success(command)
	}
}

func getGetByIDInput(kwargs wamp.Dict) (*getByIDInput, error) {
	if kwargs["channel"] == nil {
		return nil, errors.New("Missing channel")
	}

	if kwargs["name"] == nil {
		return nil, errors.New("Missing name")
	}

	return &getByIDInput{
		Channel: conv.ToString(kwargs["channel"]),
		Name:    conv.ToString(kwargs["name"]),
	}, nil
}
