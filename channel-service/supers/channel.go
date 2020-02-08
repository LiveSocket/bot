package supers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/LiveSocket/bot/channel-service/helpers"
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
)

type channelInput struct {
	SubCommand string
	Channel    string
	BotName    string
}

// Channel Adds or destroys a channel
//
// Add a channel
// ["add <channel> <botName>"]{message twitch.PrivateMessage}
//
// Destroy a channel
// ["destroy <channel>"]{message twitch.PrivateMessage}
//
// Returns message for chat
func Channel(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getChannelInput(invocation)
		if err != nil {
			return socket.Error(err)
		}

		switch input.SubCommand {
		case "add":
			return channelAdd(service, input)
		case "destroy":
			return channelDestroy(service, input)
		default:
			return socket.Error("Invalid subcommand")
		}
	}
}

func channelAdd(service *service.Service, input *channelInput) socket.Result {
	// Validate input
	if input.Channel == "" {
		return socket.Error("Missing channel name")
	}
	if input.BotName == "" {
		return socket.Error("Missing bot name")
	}

	// Create channel
	_, err := helpers.CreateChannel(service, input.Channel, input.BotName)
	if err != nil {
		return socket.Error(err)
	}

	// Return message for chat
	return socket.Success(fmt.Sprintf("%s channel added", input.Channel))
}

func channelDestroy(service *service.Service, input *channelInput) socket.Result {
	// Validate input
	if input.Channel == "" {
		return socket.Error("Missing channel name")
	}

	// Destroy channel
	err := helpers.DestroyChannel(service, input.Channel)
	if err != nil {
		return socket.Error(err)
	}

	// Return message for chat
	return socket.Success(fmt.Sprintf("%s channel added", input.Channel))
}

func getChannelInput(invocation *socket.Invocation) (*channelInput, error) {
	input := &channelInput{}
	if len(invocation.Arguments) == 0 {
		return nil, errors.New("Missing args")
	}

	if len(invocation.Arguments) > 0 {
		input.SubCommand = strings.ToLower(conv.ToString(invocation.Arguments[0]))
	}

	if len(invocation.Arguments) > 1 {
		input.Channel = strings.ToLower(conv.ToString(invocation.Arguments[1]))
	}

	if len(invocation.Arguments) > 2 {
		input.BotName = conv.ToString(invocation.Arguments[2])
	}

	return input, nil
}
