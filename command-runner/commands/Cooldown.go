package commands

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/LiveSocket/bot/command-runner/helpers"
	"github.com/LiveSocket/bot/conv"

	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

type cooldownInput struct {
	Channel string
	Name    string
	Amount  *uint64
}

// Cooldown Sets or displays cooldown for a command
//
// `!cooldown !<name> <seconds>` Sets cooldown time
// `!cooldown !<name>` Displays current cooldown time
//
// command.cooldown
// [name string, amount string]{message twitch.PrivateMessage}
// `amount` should be uint64 parsable
//
// Returns [string] as response for chat
func Cooldown(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getCooldownInput(service, invocation)
		if err != nil {
			return socket.Error(err)
		}

		// If amount is present, then set cooldown
		if input.Amount != nil {
			return setCooldown(service, input)
		}
		// else get cooldown time
		return getCooldown(service, input)
	}
}

func getCooldown(service *service.Service, input *cooldownInput) socket.Result {
	// Find command to show cooldown for
	command, err := helpers.GetCommandByID(service, input.Channel, input.Name)
	if err != nil {
		return socket.Error(err)
	}

	// If no command found
	if command == nil {
		return socket.Success(fmt.Sprintf("Command !%s doesn't exist", input.Name))
	}

	// Return message to display in chat
	return socket.Success(fmt.Sprintf("!%s cooldown is %d", input.Name, command.Cooldown))
}

func setCooldown(service *service.Service, input *cooldownInput) socket.Result {
	// Find command to set
	command, err := helpers.GetCommandByID(service, input.Channel, input.Name)
	if err != nil {
		return socket.Error(err)
	}

	// Set cooldown time
	command.Cooldown = *input.Amount

	// Update command
	result, err := helpers.UpdateCommand(service, command)
	if err != nil {
		return socket.Error(err)
	}

	// Return message to display in chat
	return socket.Success(fmt.Sprintf("!%s cooldown set to %d", result.Name, result.Cooldown))
}

func getCooldownInput(service *service.Service, invocation *socket.Invocation) (*cooldownInput, error) {
	if invocation.ArgumentsKw["message"] == nil {
		return nil, errors.New("Missing message")
	}

	if len(invocation.Arguments) == 0 {
		return nil, errors.New("Missing args")
	}

	if invocation.Arguments[0] == nil {
		return nil, errors.New("Missing command name")
	}

	name := conv.ToString(invocation.Arguments[0])
	message, err := conv.ToStringMap(invocation.ArgumentsKw["message"])
	if err != nil {
		return nil, err
	}
	input := &cooldownInput{
		Channel: conv.ToString(message["Channel"]),
		Name:    name[1:len(name)],
	}

	// Check if amount is provided
	if len(invocation.Arguments) > 2 {
		// Attempt to parse to uint64
		amount, err := strconv.ParseUint(conv.ToString(invocation.Arguments[1]), 10, 64)
		if err != nil {
			return nil, errors.New("Invalid cooldown time")
		}
		input.Amount = &amount
	}

	// Validate
	err = validateCooldown(service, input)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func validateCooldown(service *service.Service, input *cooldownInput) error {
	// Find command to get or set cooldown
	result, err := service.SimpleCall("public.command.getById", nil, wamp.Dict{"channel": input.Channel, "name": input.Name})
	if err != nil {
		return err
	}

	// Check if command exists
	if len(result.Arguments) == 0 {
		return errors.New("Command does not exist")
	}

	// Validation passed
	return nil
}
