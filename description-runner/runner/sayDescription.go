package runner

import (
	"log"

	"github.com/LiveSocket/bot/description-runner/helpers"
	"github.com/LiveSocket/bot/service"
	"github.com/gempir/go-twitch-irc/v2"
)

func sayDescription(service *service.Service, channel string, name string, message *twitch.PrivateMessage) {
	// Check if custom command
	custom, err := helpers.GetCustomCommand(service, channel, name)
	if err != nil {
		// Log error and escape
		log.Print(err)
		return
	}

	// If custom command
	if custom != nil {
		description := describeCustomCommand(service, custom, message)
		speak(service, message.Channel, description)
		return
	}

	// Check if db command
	command, err := helpers.GetCommandByID(service, channel, name)
	if err != nil {
		// Log error and escape
		log.Print(err)
		return
	}

	// If db command
	if command != nil {
		description := describeDBCommand(service, command, message)
		speak(service, message.Channel, description)
		return
	}
}

func describeCustomCommand(service *service.Service, custom *helpers.CustomCommand, message *twitch.PrivateMessage) string {
	// If command is not restricted
	if !custom.Restricted {
		return custom.Description
	}

	// Check permissions for restricted command
	isMod, err := helpers.IsMod(service, message)
	if err != nil {
		// Log error and escape
		log.Print(err)
		return ""
	}

	// If permitted
	if isMod {
		return custom.Description
	}

	// Not permitted, do nothing
	return ""
}

func describeDBCommand(service *service.Service, command *helpers.Command, message *twitch.PrivateMessage) string {
	// If command is not restricted
	if !command.Restricted {
		return command.Description
	}

	// Check permissions for restricted command
	isMod, err := helpers.IsMod(service, message)
	if err != nil {
		// Log error and escape
		log.Print(err)
		return ""
	}

	// If permitted
	if isMod {
		return command.Description
	}

	// Not permitted, do nothing
	return ""
}
