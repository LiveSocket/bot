package actions

import (
	"errors"

	"github.com/LiveSocket/bot/admin-service/models"
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

type isAdminInput struct {
	Username string
}

// IsAdmin Checks if a user is an Admin
//
// private.admin.isAdmin
// {username string}
//
// Returns [bool]
func IsAdmin(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getIsAdminInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Find an Admin
		admin, err := models.GetAdmin(service, input.Username)
		if err != nil {
			return socket.Error(err)
		}

		// Return
		return socket.Success(admin != nil)
	}
}

func getIsAdminInput(kwargs wamp.Dict) (*isAdminInput, error) {
	if kwargs["username"] == nil {
		return nil, errors.New("Missing username")
	}

	return &isAdminInput{
		Username: conv.ToString(kwargs["username"]),
	}, nil
}
