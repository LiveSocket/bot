package helpers

import (
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/super-command-runner/models"
	"github.com/gammazero/nexus/v3/wamp"
)

// FindSuperCommand WAMP call to find a super command by name
func FindSuperCommand(service *service.Service, name string) (*models.SuperCommand, error) {
	// Call super command get
	res, err := service.SimpleCall("private.super.command.get", nil, wamp.Dict{"name": name})
	if err != nil {
		return nil, err
	}
	// Check and return response
	if len(res.Arguments) > 0 && res.Arguments[0] != nil {
		super, err := conv.ToStringMap(res.Arguments[0])
		if err != nil {
			return nil, err
		}
		return &models.SuperCommand{
			Proc: conv.ToString(super["proc"]),
			Name: conv.ToString(super["name"]),
		}, nil
	}
	return nil, nil
}
