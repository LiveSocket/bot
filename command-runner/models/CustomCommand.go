package models

import (
	"github.com/LiveSocket/bot/service"
)

// CustomCommand Represents a custom command
type Custom struct {
	Name        string `db:"name" json:"name"`
	Proc        string `db:"proc" json:"proc"`
	Channel     string `db:"channel" json:"channel"`
	Enabled     bool   `db:"enabled" json:"enabled"`
	Restricted  bool   `db:"restricted" json:"restricted"`
	Description string `db:"description" json:"description"`
}

const (
	getCustom    = "SELECT * FROM `custom_commands` WHERE `name`=? AND `channel`=?"
	getCustoms   = "SELECT * FROM `custom_commands` WHERE `channel`=?"
	createCustom = "INSERT INTO `custom_commands` (`channel`,`name`,`proc`,`description`,`enabled`,`restricted`) VALUES (?,?,?,?,?,?)"
	updateCustom = "UPDATE `custom_commands` SET `enabled`=:enabled,`restricted`=:restricted WHERE name=:name AND channel=:channel"
)

// GetCustom Finds a custom command for the channel
func GetCustom(service *service.Service, name string, channel string) (*Custom, error) {
	command := Custom{}
	err := service.Get(&command, getCustom, name, channel)
	return &command, err
}

// GetEnabledCustom Finds an enabled custom command for the channel
func GetEnabledCustom(service *service.Service, name string, channel string) (*Custom, error) {
	command, err := GetCustom(service, name, channel)
	if err != nil {
		return nil, err
	}
	if command != nil && !command.Enabled {
		return nil, nil
	}
	return command, nil
}

func GetCustoms(service *service.Service, channel string) ([]Custom, error) {
	commands := []Custom{}
	err := service.Select(&commands, getCustoms, channel)
	return commands, err
}

func CreateCustom(service *service.Service, channel, name, proc, description string, enabled, restricted bool) error {
	_, err := service.Exec(createCustom, channel, name, proc, description, enabled, restricted)
	return err
}

func UpdateCustom(service *service.Service, command *Custom) error {
	_, err := service.NamedExec(updateCustom, command)
	return err
}
