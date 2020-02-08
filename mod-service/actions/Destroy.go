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

type destroyInput struct {
	Channel  string
	Username string
}

// Destroy Destroys a mod user
//
// private.mod.destroy
// {channel string, username string}
//
// Returns nothing
func Destroy(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getDestroyInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Find the mod to destroy
		err = models.DestroyMod(service, input.Channel, input.Username)
		if err != nil {
			return socket.Error(err)
		}

		// Emit mod destroyed event
		err = helpers.EmitModDestroyed(service, input.Channel, input.Username)
		if err != nil {
			// Don't fail on error
			log.Print(err)
		}

		// Return success but nothing
		return socket.Success()
	}
}

func getDestroyInput(kwargs wamp.Dict) (*destroyInput, error) {
	if kwargs["channel"] == nil {
		return nil, errors.New("Missing channel")
	}

	if kwargs["username"] == nil {
		return nil, errors.New("Missing username")
	}

	return &destroyInput{
		Channel:  conv.ToString(kwargs["channel"]),
		Username: conv.ToString(kwargs["username"]),
	}, nil
}
