package helpers

import (
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// CustomCommand The important parts of a custom command
type CustomCommand struct {
	Channel     string `json:"channel"`
	Name        string `json:"name"`
	Proc        string `json:"proc"`
	Enabled     bool   `json:"enabled"`
	Restricted  bool   `json:"restricted"`
	Description string `json:"description"`
}

// GetCustomCommand WAMP call helper for getting a custom command
func GetCustomCommand(service *service.Service, channel string, name string) (*CustomCommand, error) {
	// Call get command by id endpoint
	res, err := service.SimpleCall("private.command.getCustom", nil, wamp.Dict{"channel": channel, "name": name})
	if err != nil {
		return nil, err
	}
	// Check, convert, and return response
	if len(res.Arguments) > 0 && res.Arguments[0] != nil {
		command, err := conv.ToStringMap(res.Arguments[0])
		if err != nil {
			return nil, err
		}
		result, err := toCustom(command)
		return result, err
	}
	return nil, nil
}

func toCustom(command map[string]interface{}) (*CustomCommand, error) {
	enabled, err := conv.ToBool(command["enabled"])
	if err != nil {
		return nil, err
	}
	restricted, err := conv.ToBool(command["restricted"])
	if err != nil {
		return nil, err
	}
	return &CustomCommand{
		Channel:     conv.ToString(command["channel"]),
		Name:        conv.ToString(command["name"]),
		Proc:        conv.ToString(command["proc"]),
		Enabled:     enabled,
		Restricted:  restricted,
		Description: conv.ToString(command["description"]),
	}, nil
}
