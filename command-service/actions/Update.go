package actions

import (
	"errors"

	"github.com/LiveSocket/bot/command-service/models"
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

// Update Updates a command
//
// public.command.update
// {channel string, name string, response string, enabled bool, restricted bool, cooldown uint64, scheduled uint64, updated_by string}
//
// Returns [Command]
func Update(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		command, err := getUpdateInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Update command
		err = models.UpdateCommand(service, command)
		if err != nil {
			return socket.Error(err)
		}

		// Return updated command
		return socket.Success(command)
	}
}

func getUpdateInput(kwargs wamp.Dict) (*models.Command, error) {
	if kwargs["channel"] == nil {
		return nil, errors.New("Missing channel")
	}
	if kwargs["name"] == nil {
		return nil, errors.New("Missing name")
	}

	enabled, err := conv.ToBool(kwargs["enabled"])
	if err != nil {
		return nil, err
	}
	restricted, err := conv.ToBool(kwargs["restricted"])
	if err != nil {
		return nil, err
	}
	cooldown, err := conv.ToUint64(kwargs["cooldown"])
	if err != nil {
		return nil, err
	}
	schedule, err := conv.ToUint64(kwargs["schedule"])
	if err != nil {
		return nil, err
	}
	command := &models.Command{
		Channel:     conv.ToString(kwargs["channel"]),
		Name:        conv.ToString(kwargs["name"]),
		Response:    conv.ToString(kwargs["response"]),
		Enabled:     enabled,
		Restricted:  restricted,
		Cooldown:    cooldown,
		Schedule:    schedule,
		UpdatedBy:   conv.ToString(kwargs["updated_by"]),
		Description: conv.ToString(kwargs["description"]),
	}
	return command, nil
}
