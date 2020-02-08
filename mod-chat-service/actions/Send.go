package actions

import (
	"errors"

	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/mod-chat-service/helpers"
	"github.com/LiveSocket/bot/mod-chat-service/models"
	"github.com/LiveSocket/bot/service"

	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

// public.modchat.send

type sendInput struct {
	Channel string
	Message string
}

// Send Sends a modchat message
//
// public.modchat.send
// {channel string, message string}
//
// Returns nothing
func Send(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getSendInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Create new mod chat message
		message := models.NewModChat(input.Channel, input.Message, invocation.Details["caller_authid"].(string))

		// Save to db
		result, err := models.CreateModChat(service, message)
		if err != nil {
			return socket.Error(err)
		}

		// Get the id for the new ModChat message
		id, err := result.LastInsertId()
		if err != nil {
			return socket.Error(err)
		}

		// Set the ID on the message
		message.ID = uint(id)

		// Emit the new message to all listeners
		helpers.EmitMessage(service, message)

		// Return success
		return socket.Success()
	}
}

func getSendInput(kwargs wamp.Dict) (*sendInput, error) {
	if kwargs["channel"] == nil {
		return nil, errors.New("Missing channel")
	}

	if kwargs["message"] == nil {
		return nil, errors.New("Missing message")
	}

	return &sendInput{
		Channel: conv.ToString(kwargs["channel"]),
		Message: conv.ToString(kwargs["message"]),
	}, nil
}
