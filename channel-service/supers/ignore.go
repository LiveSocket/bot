package supers

import (
	"errors"
	"strings"

	"github.com/LiveSocket/bot/channel-service/helpers"
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
)

type ignoreInput struct {
	SubCommand string
	Channel    string
}

// Ignore Manages ignore list for channels
//
// Add a channel to the ignore list
// ["ignore add <channel>"]{message twitch.PrivateMessage}
//
// Remove a channel from the ignore list
// ["ignore remove <channel>"]{message twitch.PrivateMessage}
//
// Returns respose for chat
func Ignore(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getIgnoreInput(invocation)
		if err != nil {
			return socket.Error(err)
		}

		// Switch on sub command
		switch input.SubCommand {
		case "add":
			return add(service, input)
		case "remove":
			return remove(service, input)
		default:
			return socket.Error("Invalid sub command")
		}
	}
}

func add(service *service.Service, input *ignoreInput) socket.Result {
	// Validate input
	if input.Channel == "" {
		return socket.Error("Missing channel")
	}

	// ignore channel
	err := helpers.IgnoreChannel(service, input.Channel)
	if err != nil {
		return socket.Error(err)
	}

	// Return response for chat
	return socket.Success("%s added to channel ignore list", input.Channel)
}

func remove(service *service.Service, input *ignoreInput) socket.Result {
	// Validate input
	if input.Channel == "" {
		return socket.Error("Missing channel")
	}

	// unignore channel
	err := helpers.UnignoreChannel(service, input.Channel)
	if err != nil {
		return socket.Error(err)
	}

	// Return response for chat
	return socket.Success("%s removed from channel ignore list", input.Channel)
}

func getIgnoreInput(invocation *socket.Invocation) (*ignoreInput, error) {
	input := &ignoreInput{}
	if len(invocation.Arguments) == 0 {
		return nil, errors.New("Missing args")
	}

	if len(invocation.Arguments) > 0 {
		input.SubCommand = strings.ToLower(conv.ToString(invocation.Arguments[0]))
	}

	if len(invocation.Arguments) > 1 {
		input.Channel = strings.ToLower(conv.ToString(invocation.Arguments[1]))
	}

	return input, nil
}
