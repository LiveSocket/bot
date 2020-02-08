package runner

import (
	"log"
	"strings"

	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/super-command-runner/helpers"
	"github.com/gammazero/nexus/v3/wamp"
	"github.com/gempir/go-twitch-irc/v2"
)

// RunSuperCommand Run the super command and reply if necessary
func RunSuperCommand(service *service.Service, message *twitch.PrivateMessage) {
	// Process arguments
	slice := strings.Split(message.Message, " ")
	parts := slice[1:len(slice)]
	command := parts[0]

	// Check if user is admin
	isAdmin, err := helpers.IsAdmin(service, message.User.Name)
	if err != nil {
		log.Print(err)
		return
	}

	// If user is not an admin, stop processing
	if !isAdmin {
		return
	}

	// Find super command to run
	super, err := helpers.FindSuperCommand(service, command)
	if err != nil {
		log.Print(err)
		return
	}
	if super == nil {
		return
	}

	// Convert message into args
	rest := parts[1:len(parts)]
	args := wamp.List{}
	for _, item := range rest {
		args = append(args, item)
	}

	// Create kwargs for command
	kwargs := wamp.Dict{"message": message}

	// Run Super Command
	result, err := service.SimpleCall(super.Proc, args, kwargs)
	if err != nil {
		log.Print(err)
		return
	}
	// If something to say, speak message
	if result != nil && len(result.Arguments) > 0 && conv.ToString(result.Arguments[0]) != "" {
		err := helpers.Speak(service, message.Channel, conv.ToString(result.Arguments[0]))
		if err != nil {
			log.Print(err)
			return
		}
	}
}
