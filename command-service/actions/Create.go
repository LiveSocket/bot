package actions

import (
	"errors"
	"log"

	"github.com/LiveSocket/bot/command-service/models"
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

type createInput struct {
	Channel  string
	Name     string
	Response string
	Username string
}

// Create Creates a new command
//
// public.command.create
// {name string, botName string, response string, username string}
//
// Returns [Command]
func Create(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getCreateInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Create new command
		command, err := models.CreateCommand(service, input.Channel, input.Name, input.Response, input.Username)
		if err != nil {
			return socket.Error(err)
		}

		// Emit created event
		if err := service.Publish("event.command.created", nil, wamp.List{command}, nil); err != nil {
			log.Print(err)
		}

		// Return new command
		return socket.Success(command)
	}
}

func getCreateInput(kwargs wamp.Dict) (*createInput, error) {
	if kwargs["channel"] == nil {
		return nil, errors.New("Missing channel")
	}

	if kwargs["name"] == nil {
		return nil, errors.New("Missing name")
	}

	if kwargs["response"] == nil {
		return nil, errors.New("Missing response")
	}

	if kwargs["username"] == nil {
		return nil, errors.New("Missing username")
	}

	return &createInput{
		Channel:  conv.ToString(kwargs["channel"]),
		Name:     conv.ToString(kwargs["name"]),
		Response: conv.ToString(kwargs["response"]),
		Username: conv.ToString(kwargs["username"]),
	}, nil
}
