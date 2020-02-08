package actions

import (
	"errors"

	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/mod-service/models"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

type destroyAllInput struct {
	Channel string
}

// DestroyAll Destroys all mod users for a channel
//
// private.mod.destroyAll
// {channel string}
//
// Returns nothing
func DestroyAll(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getDestroyAllInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Destroy all mods for channel
		err = models.DestroyAllMods(service, input.Channel)
		if err != nil {
			return socket.Error(err)
		}

		// Return success but nothing
		return socket.Success()
	}
}

func getDestroyAllInput(kwargs wamp.Dict) (*destroyAllInput, error) {
	if kwargs["channel"] == nil {
		return nil, errors.New("Missing channel")
	}

	return &destroyAllInput{Channel: conv.ToString(kwargs["channel"])}, nil
}
