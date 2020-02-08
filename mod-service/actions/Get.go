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

type getInput struct {
	Channel  string
	Username string
}

// Get Get a list of mods
// Behaviour is different depending on supplied arguments
//
// private.mod.get
// {channel string} - Returns all mods for the channel
// {username string} - Returns all channels for the mod
// {channel string, username} - Returns a specific mod user
//
// Returns [Mod] or [Mod...]
func Get(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getGetInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// if lookup by channel and username
		if input.Channel != "" && input.Username != "" {
			// Lookup a single mod user
			return findSpecificMod(service, input)
		}
		// if only channel is set
		if input.Channel != "" {
			// Lookup all mods in a channel
			return findModsForChannel(service, input)
		}
		// If only username is set
		// Lookup all channels for a mod
		return findChannelsForMod(service, input)
	}
}

func findSpecificMod(service *service.Service, props *getInput) socket.Result {
	// Find specific mod
	mod, err := models.GetMod(service, props.Channel, props.Username)
	if err != nil {
		return socket.Error(err)
	}

	// Return Mod
	return socket.Success(mod)
}

func findModsForChannel(service *service.Service, props *getInput) socket.Result {
	// Find all mods for channel
	mods, err := models.GetMods(service, props.Channel)
	if err != nil {
		return socket.Error(err)
	}

	// Create wamp.List of mods
	list := wamp.List{}
	for _, mod := range mods {
		list = append(list, mod)
	}

	// Return list of mods
	return client.InvokeResult{
		Args: list,
	}
}

func findChannelsForMod(service *service.Service, props *getInput) socket.Result {
	// Find all channels for mod (Mod objects)
	mods, err := models.GetChannelsForMods(service, props.Username)
	if err != nil {
		return socket.Error(err)
	}

	// Create wamp.List of mods
	list := wamp.List{}
	for _, mod := range mods {
		list = append(list, mod)
	}

	// Return list of mods
	return client.InvokeResult{
		Args: list,
	}
}

func getGetInput(kwargs wamp.Dict) (*getInput, error) {

	input := &getInput{}

	if kwargs["channel"] != nil {
		input.Channel = conv.ToString(kwargs["channel"])
	}

	if kwargs["username"] != nil {
		input.Username = conv.ToString(kwargs["username"])
	}

	if input.Channel == "" && input.Username == "" {
		return nil, errors.New("Must provide channel or username")
	}
	return input, nil
}
