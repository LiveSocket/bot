package runner

import (
	"fmt"
	"log"
	"strings"

	"github.com/LiveSocket/bot/command-runner/helpers"

	"github.com/LiveSocket/bot/command-runner/models"
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
	"github.com/gempir/go-twitch-irc/v2"
)

func RunCommand(service *service.Service, message *twitch.PrivateMessage) {
	channel := message.Channel
	parts := strings.Split(message.Message, " ")
	name := parts[0][1:len(parts[0])]
	rest := parts[1:len(parts)]
	args := wamp.List{}
	for _, item := range rest {
		args = append(args, item)
	}

	kwargs := wamp.Dict{"message": message}

	// Check if custom command
	custom, err := models.GetEnabledCustom(service, name, channel)
	if err != nil {
		// Log error and escape
		log.Println("error fetching custom commands", err)
		return
	}

	// If custom command
	if custom != nil {
		response := runCustomCommand(service, custom, message, args, kwargs)
		speak(service, message.Channel, response)
		return
	}

	// Check if db command
	command, err := helpers.GetCommandByID(service, channel, name)
	if err != nil {
		// Log error and escape
		log.Println("error fetching db command", err)
		return
	}

	// If db command
	if command != nil {
		response := runDBCommand(service, command, message, args, kwargs)
		speak(service, message.Channel, response)
		return
	}
}

func runCustomCommand(service *service.Service, custom *models.Custom, message *twitch.PrivateMessage, args wamp.List, kwargs wamp.Dict) string {
	// Add twitch message to kwargs
	kwargs["twitch"] = message
	// If command is not restricted
	if !custom.Restricted {
		return getCustomCommandResult(service, custom, args, kwargs)
	}

	// Check permissions for restricted command
	isMod, err := helpers.IsMod(service, message)
	if err != nil {
		// Log error and escape
		fmt.Println("error checking if mod", err)
		return ""
	}

	// If permitted
	if isMod {
		return getCustomCommandResult(service, custom, args, kwargs)
	}

	// Not permitted, do nothing
	return ""
}

func runDBCommand(service *service.Service, command *helpers.Command, message *twitch.PrivateMessage, args wamp.List, kwargs wamp.Dict) string {
	// If command is not restricted
	if !command.Restricted {
		return command.Response
	}

	// Check permissions for restricted command
	isMod, err := helpers.IsMod(service, message)
	if err != nil {
		// Log error and escape
		fmt.Println("error checking if mod", err)
		return ""
	}

	// If permitted
	if isMod {
		return command.Response
	}

	// Not permitted, do nothing
	return ""
}

func getCustomCommandResult(service *service.Service, custom *models.Custom, args wamp.List, kwargs wamp.Dict) string {
	// Call custom command and get result
	result, err := service.SimpleCall(custom.Proc, args, kwargs)
	if err != nil {
		fmt.Println("error calling custom command", err)
		return fmt.Sprint(err)
	} else if len(result.Arguments) > 0 && result.Arguments[0] != nil {
		return result.Arguments[0].(string)
	}
	return ""
}

func speak(service *service.Service, channel string, message string) {
	// Only speak if there is something to say
	if channel != "" && message != "" {
		err := helpers.Speak(service, channel, message)
		if err != nil {
			// Log error
			log.Println("error speaking", err)
		}
	}
}
