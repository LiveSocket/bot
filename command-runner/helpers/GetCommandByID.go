package helpers

import (
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// Command The important parts of a command
type Command struct {
	Channel     string `json:"channel"`
	Name        string `json:"name"`
	Response    string `json:"response"`
	Enabled     bool   `json:"enabled"`
	Restricted  bool   `json:"restricted"`
	Cooldown    uint64 `json:"cooldown"`
	Description string `json:"description"`
	Schedule    uint64 `json:"schedule"`
	UpdatedBy   string `json:"updated_by"`
}

// GetCommandByID WAMP call helper for getting a command by ID
func GetCommandByID(service *service.Service, channel string, name string) (*Command, error) {
	// Call get command by id endpoint
	res, err := service.SimpleCall("public.command.getById", nil, wamp.Dict{"channel": channel, "name": name})
	if err != nil {
		return nil, err
	}
	// Check, convert, and return response
	if len(res.Arguments) > 0 && res.Arguments[0] != nil {
		command, err := conv.ToStringMap(res.Arguments[0])
		if err != nil {
			return nil, err
		}
		result, err := toCommand(command)
		return result, err
	}
	return nil, nil
}

func toCommand(command map[string]interface{}) (*Command, error) {
	enabled, err := conv.ToBool(command["enabled"])
	if err != nil {
		return nil, err
	}
	restricted, err := conv.ToBool(command["restricted"])
	if err != nil {
		return nil, err
	}
	cooldown, err := conv.ToUint64(command["cooldown"])
	if err != nil {
		return nil, err
	}
	schedule, err := conv.ToUint64(command["schedule"])
	if err != nil {
		return nil, err
	}
	return &Command{
		Channel:     conv.ToString(command["channel"]),
		Name:        conv.ToString(command["name"]),
		Response:    conv.ToString(command["response"]),
		Enabled:     enabled,
		Restricted:  restricted,
		Cooldown:    cooldown,
		Description: conv.ToString(command["description"]),
		Schedule:    schedule,
		UpdatedBy:   conv.ToString(command["updated_by"]),
	}, nil
}
