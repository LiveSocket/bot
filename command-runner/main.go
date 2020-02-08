package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/LiveSocket/bot/command-runner/migrations"

	"github.com/LiveSocket/bot/command-runner/actions"
	"github.com/LiveSocket/bot/command-runner/commands"
	"github.com/LiveSocket/bot/command-runner/runner"
	"github.com/LiveSocket/bot/service"
)

func main() {
	s := &service.Service{}
	close := s.Init(service.Actions{
		"private.command.activate":   actions.Activate(s),
		"private.command.deactivate": actions.Deactivate(s),
		"private.command.getCustom":  actions.GetCustom(s),
		"command.add":                commands.Add(s),
		"command.commands":           commands.Commands(s),
		"command.cooldown":           commands.Cooldown(s),
		"command.del":                commands.Del(s),
		"command.disable":            commands.Disable(s),
		"command.enable":             commands.Enable(s),
	}, service.Subscriptions{
		"event.chat.message": runner.OnMessage(s),
	}, "__command_runner", migrations.CreateCustomCommandTables, migrations.AlterCustomCommands)
	defer close()

	// Wait for CTRL-c or client close while handling events.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	select {
	case <-sigChan:
	case <-s.Done():
		log.Print("Router gone, exiting")
		return
	}
}
