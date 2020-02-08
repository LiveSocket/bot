package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"

	"github.com/LiveSocket/bot/command-runner/helpers"
	"github.com/LiveSocket/bot/command-runner/models"
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

type commandsInput struct {
	Message *twitch.PrivateMessage
}

// Commands Lists available commands for the requestor
//
// !commands
//
// command.commands
// {message twitch.PrivateMessage}
//
// Returns [string] as response for chat
func Commands(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getCommandsInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Get all DB commands for channel
		dbResults, err := helpers.GetCommandsForChannel(service, input.Message.Channel)
		if err != nil {
			return socket.Error(err)
		}

		// Get all custom commands for channel
		customs, err := models.GetCustoms(service, input.Message.Channel)
		if err != nil {
			return socket.Error(err)
		}

		// Check if user is mod
		mod, err := helpers.IsMod(service, input.Message)
		if err != nil {
			return socket.Error(err)
		}

		// Create names []string to populate
		names := []string{}

		// Filter and add command names to slice
		if len(customs) != 0 {
			for _, custom := range customs {
				// Filter out "commands" as this is the command that was just run...
				// Filter out restricted commands for non-mod users
				if custom.Name != "commands" && (!custom.Restricted || mod) {
					names = append(names, "!"+custom.Name)
				}
			}
		}

		// Filter and add command names to slice
		if len(dbResults) != 0 {
			for _, command := range dbResults {
				// Filter out restricted commands for non-mod users
				if !command.Restricted || mod {
					names = append(names, fmt.Sprintf("!%s", command.Name))
				}
			}
		}

		// Return message to display in chat
		return socket.Error(fmt.Sprintf("Available commands: %s", strings.Join(names, " ")))
	}
}

func getCommandsInput(kwargs wamp.Dict) (*commandsInput, error) {
	if kwargs["message"] == nil {
		return nil, errors.New("Missing message")
	}

	message, err := conv.ToStringMap(kwargs["message"])
	if err != nil {
		return nil, err
	}
	return &commandsInput{
		Message: service.MapToPrivateMessage(message),
	}, nil
}
