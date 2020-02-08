package actions

import (
	"github.com/LiveSocket/bot/channel-service/models"
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

type getInput struct {
	Name    string
	BotName string
}

// Get Gets a list of channels
//
// private.channel.get
//
// If botName is supplied returns list of channels matching botName
// {botName string}
//
// If name is supplied returns channel with matching name
// {name string}
//
// If no arguments supplied returns list of all channels
// {}
//
// Returns [Channel...]
func Get(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *wamp.Invocation) socket.Result {
		// Get input args from call
		input, err := getGetInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// If query by channel name
		if input.Name != "" {
			return getByName(service, input)
		}

		// If query by bot name
		if input.BotName != "" {
			return getByBotName(service, input)
		}

		// If no query params
		return getAll(service, input)
	}
}

func getByName(service *service.Service, input *getInput) socket.Result {
	// Find channel with matching name
	channel, err := models.GetChannel(service, input.Name)
	if err != nil {
		return socket.Error(err)
	}
	// Return single channel
	return socket.Success(channel)
}

func getByBotName(service *service.Service, input *getInput) socket.Result {
	// Find channels by bot name
	channels, err := models.GetBotChannels(service, input.BotName)
	if err != nil {
		return socket.Error(err)
	}

	// Convert channels to wamp.List
	list := wamp.List{}
	for _, channel := range channels {
		list = append(list, channel)
	}

	// Return list of channels
	return socket.Result{
		Args: list,
	}
}

func getAll(service *service.Service, input *getInput) socket.Result {
	// Find all channels
	channels, err := models.GetAllChannels(service)
	if err != nil {
		return socket.Error(err)
	}

	// Convert to wamp.List
	list := wamp.List{}
	for _, channel := range channels {
		list = append(list, channel)
	}

	// Return list of channels
	return socket.Result{
		Args: list,
	}

}

func getGetInput(kwargs wamp.Dict) (*getInput, error) {
	if kwargs["name"] != nil {
		return &getInput{Name: conv.ToString(kwargs["name"])}, nil
	}
	if kwargs["botName"] != nil {
		return &getInput{BotName: conv.ToString(kwargs["botName"])}, nil
	}
	return nil, nil
}
