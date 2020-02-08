package runner

import (
	"fmt"
	"log"
	"strings"

	"github.com/LiveSocket/bot/description-runner/helpers"
	"github.com/LiveSocket/bot/service"
	"github.com/gempir/go-twitch-irc/v2"
)

func editDescription(service *service.Service, channel string, name string, message *twitch.PrivateMessage, rest []string) {
	// Check if custom command
	custom, err := helpers.GetCustomCommand(service, channel, name)
	if err != nil {
		// Log error and escape
		log.Print(err)
		return
	}

	if custom != nil {
		// Editing descriptions of custom commands is not allowed therefore speak an error message
		speak(service, channel, "Cannot edit custom command descriptions")
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
		command.Description = strings.Join(rest, " ")

		_, err := helpers.UpdateCommand(service, command)
		if err != nil {
			// Log error and escape
			log.Print(err)
			return
		}
		speak(service, message.Channel, fmt.Sprintf("!%s description updated", command.Name))
		return
	}
}
