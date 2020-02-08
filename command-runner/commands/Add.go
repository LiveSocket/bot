package commands

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/LiveSocket/bot/command-runner/helpers"
	"github.com/LiveSocket/bot/conv"

	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
)

type addInput struct {
	Channel  string
	Username string
	Name     string
	Response string
}

// Add Adds or updates a command
//
// !add !<name> <response>
//
// command.add
// [name, response...]{message twitch.PrivateMessage}
//
// Returns [string] as response for chat
func Add(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getAddInput(service, invocation)
		if err != nil {
			return socket.Error(err)
		}

		// Trim ! off command name
		input.Name = input.Name[1:len(input.Name)]

		// Check if command exists
		command, err := helpers.GetCommandByID(service, input.Channel, input.Name)
		if err != nil {
			return socket.Error(err)
		}

		// If command exists
		if command != nil {
			// Create wamp.Dict of command
			command.Response = input.Response
			command.UpdatedBy = input.Username

			// Update command
			_, err = helpers.UpdateCommand(service, command)
		} else {
			// Create new command
			err = helpers.CreateCommand(service, input.Channel, input.Name, input.Response, input.Username)
		}
		if err != nil {
			return socket.Error(err)
		}

		// Return message to display in chat
		return socket.Success(fmt.Sprintf("Command !%s added/updated", input.Name))
	}
}

func getAddInput(service *service.Service, invocation *socket.Invocation) (*addInput, error) {
	if invocation.ArgumentsKw["message"] == nil {
		return nil, errors.New("Missing message")
	}

	if len(invocation.Arguments) == 0 {
		return nil, errors.New("Missing args")
	}

	if invocation.Arguments[0] == nil {
		return nil, errors.New("Missing name")
	}

	if len(invocation.Arguments) < 2 {
		return nil, errors.New("Missing response")
	}

	// Splice response together back into a string
	parts := []string{}
	for _, arg := range invocation.Arguments[1:len(invocation.Arguments)] {
		parts = append(parts, conv.ToString(arg))
	}
	response := strings.Join(parts, " ")

	m, err := conv.ToStringMap(invocation.ArgumentsKw["message"])
	if err != nil {
		return nil, err
	}
	user, err := conv.ToStringMap(m["User"])
	if err != nil {
		return nil, err
	}
	// Create input args
	input := &addInput{
		Channel:  conv.ToString(m["Channel"]),
		Username: conv.ToString(user["Name"]),
		Name:     conv.ToString(invocation.Arguments[0]),
		Response: response,
	}

	// Validate input args
	message := validateAdd(service, input)
	if message != nil {
		return nil, message
	}

	return input, nil
}

func validateAdd(service *service.Service, data *addInput) error {
	// Check command starts with !
	if !strings.HasPrefix(data.Name, "!") {
		return errors.New("Command name must start with !")
	}

	// Check command contains at least 1 letter or number
	match, _ := regexp.MatchString("!\\w+", data.Name)
	if !match {
		return errors.New("Command name should contain at least 1 letter or number")
	}

	// Check command name does not include special characters
	match, _ = regexp.MatchString("!.*[!?.,[\\\\\\]{}%$#@^&*()_+\\-=|<>:;'\"~`]+.*", data.Name)
	if match {
		return errors.New("Command name should not include any additional special characters")
	}

	// Check command name is not already reserved
	exists, err := helpers.CustomExists(service, data.Channel, data.Name[1:len(data.Name)])
	if err != nil {
		return err
	}

	// Fail if is reserved name
	if exists {
		return errors.New("Cannot overwrite custom commands")
	}

	// Validation checks passed
	return nil
}
