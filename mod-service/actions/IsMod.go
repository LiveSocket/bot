package actions

import (
	"errors"

	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/mod-service/models"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/client"
	"github.com/gammazero/nexus/v3/wamp"
)

type isModInput struct {
	Channel  string
	Username string
}

// IsMod Checks if mod record exists for username + channel
//
// private.mod.isMod
// {channel string, username string}
//
// Returns [bool]
func IsMod(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *wamp.Invocation) client.InvokeResult {
		// Get input args from call
		input, err := getIsModInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Find mod record
		mod, err := models.GetMod(service, input.Channel, input.Username)
		if err != nil {
			return socket.Error(err)
		}

		// Return true if mod exists
		return socket.Success(mod != nil)
	}
}

func getIsModInput(kwargs wamp.Dict) (*isModInput, error) {
	if kwargs["channel"] == nil {
		return nil, errors.New("Missing channel")
	}

	if kwargs["username"] == nil {
		return nil, errors.New("Missing username")
	}

	return &isModInput{
		Channel:  conv.ToString(kwargs["channel"]),
		Username: conv.ToString(kwargs["username"]),
	}, nil
}
