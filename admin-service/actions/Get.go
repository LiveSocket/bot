package actions

import (
	"errors"

	"github.com/LiveSocket/bot/admin-service/models"
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

type getInput struct {
	Channel  string
	Username string
}

// Get Gets an Admin user
//
// private.admin.get
// {username string}
//
// Returns [Admin]
func Get(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getGetInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Find an Admin
		admin, err := models.GetAdmin(service, input.Username)
		if err != nil {
			return socket.Error(err)
		}

		// Return Admin
		return socket.Success(admin)
	}
}

func getGetInput(kwargs wamp.Dict) (*getInput, error) {
	if kwargs["username"] == nil {
		return nil, errors.New("Missing username")
	}

	return &getInput{
		Username: conv.ToString(kwargs["username"]),
	}, nil
}
