package actions

import (
	"errors"
	"log"

	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/mod-service/helpers"
	"github.com/LiveSocket/bot/mod-service/models"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

type createInput struct {
	Channel  string
	Username string
}

// Create Creates a new Mod user
//
// private.mod.create
// {channel string, username string}
//
// Returns [Mod]
func Create(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getCreateInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Create new Mod user
		mod, err := models.CreateMod(service, input.Channel, input.Username)
		if err != nil {
			return socket.Error(err)
		}

		// Emit mod created event
		err = helpers.EmitModCreated(service, mod)
		if err != nil {
			// Don't fail on error
			log.Print(err)
		}

		// Return new Mod
		return socket.Success(mod)
	}
}

func getCreateInput(kwargs wamp.Dict) (*createInput, error) {
	if kwargs["channel"] == nil {
		return nil, errors.New("Missing name")
	}

	if kwargs["username"] == nil {
		return nil, errors.New("Missing username")
	}

	return &createInput{
		Channel:  conv.ToString(kwargs["channel"]),
		Username: conv.ToString(kwargs["username"]),
	}, nil
}
