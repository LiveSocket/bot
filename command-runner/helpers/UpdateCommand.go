package helpers

import (
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// UpdateCommand WAMP call helper for updating a command
func UpdateCommand(service *service.Service, command *Command) (*Command, error) {
	dict := wamp.Dict{
		"channel":     command.Channel,
		"name":        command.Name,
		"response":    command.Response,
		"enabled":     command.Enabled,
		"restricted":  command.Restricted,
		"cooldown":    command.Cooldown,
		"description": command.Description,
		"schedule":    command.Schedule,
		"updated_by":  command.UpdatedBy,
	}
	result, err := service.SimpleCall("public.command.update", nil, dict)
	if err != nil {
		return nil, err
	}
	cMap, err := conv.ToStringMap(result.Arguments[0])
	if err != nil {
		return nil, err
	}
	c, err := toCommand(cMap)
	return c, err
}
