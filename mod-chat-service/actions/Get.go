package actions

import (
	"errors"

	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/mod-chat-service/models"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

type getInput struct {
	Channel string
	Offset  uint64
	Limit   uint64
}

// Get Get a list of modchat messages
//
// public.modchat.get
// {channel string, offset uint, limit uint}
//
// Returns [ModChat...]
func Get(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getGetInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Find requested amout of messages for channel
		messages, err := models.FindLatestMessagesForChannel(service, input.Channel, input.Offset, input.Limit)
		if err != nil {
			return socket.Error(err)
		}

		// Create wamp.List of messages
		list := wamp.List{}
		for _, message := range messages {
			list = append(list, message)
		}

		// Return list of messages
		return socket.Success(list)
	}
}

func getGetInput(kwargs wamp.Dict) (*getInput, error) {
	if kwargs["channel"] == nil {
		return nil, errors.New("Missing channel")
	}

	if kwargs["offset"] == nil {
		return nil, errors.New("Missing offset")
	}

	if kwargs["limit"] == nil {
		return nil, errors.New("Missing limit")
	}

	offset, err := conv.ToUint64(kwargs["offset"])
	if err != nil {
		return nil, err
	}
	limit, err := conv.ToUint64(kwargs["limit"])
	if err != nil {
		return nil, err
	}

	return &getInput{
		Channel: conv.ToString(kwargs["channel"]),
		Offset:  offset,
		Limit:   limit,
	}, nil
}
