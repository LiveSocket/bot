package helpers

import (
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// GetCommandsForChannel WAMP call helper for getting all commands for a channel
func GetCommandsForChannel(service *service.Service, channel string) ([]*Command, error) {
	res, err := service.SimpleCall("public.command.get", nil, wamp.Dict{"channel": channel})
	if err != nil {
		return nil, err
	}

	if len(res.Arguments) > 0 {
		commands := []*Command{}
		for _, item := range res.Arguments {
			i, err := conv.ToStringMap(item)
			if err != nil {
				return nil, err
			}
			command, err := toCommand(i)
			if err != nil {
				return nil, err
			}
			commands = append(commands, command)
		}
		return commands, nil
	}
	return nil, nil
}
