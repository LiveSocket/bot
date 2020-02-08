package models

import (
	"time"

	"github.com/LiveSocket/bot/service"
)

// Command Represents a Command
type Command struct {
	Channel     string     `db:"channel" json:"channel"`
	Name        string     `db:"name" json:"name"`
	Response    string     `db:"response" json:"response"`
	Enabled     bool       `db:"enabled" json:"enabled"`
	Restricted  bool       `db:"restricted" json:"restricted"`
	Cooldown    uint64     `db:"cooldown" json:"cooldown"`
	Description string     `db:"description" json:"description"`
	Schedule    uint64     `db:"schedule" json:"schedule"`
	UpdatedBy   string     `db:"updated_by" json:"updated_by"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
}

const (
	createCommand      = "INSERT INTO `commands` (`channel`,`name`,`response`,`enabled`,`restricted`,`cooldown`,`schedule`,`updated_by`,`updated_at`) VALUES (:channel,:name,:response,:enabled,:restricted,:cooldown,:schedule,:updated_by,:updated_at)"
	updateCommand      = "UPDATE `commands` SET `response`=:response,`enabled`=:enabled,`restricted`=:restricted,`cooldown`=:cooldown,`description`=:description,`schedule`=:schedule,`updated_by`=:updated_by,`updated_at`=CURRENT_TIMESTAMP WHERE `channel`=:channel AND `name`=:name"
	getCommand         = "SELECT * FROM `commands` WHERE `channel`=? AND `name`=?"
	getCommands        = "SELECT * FROM `commands` WHERE `channel`=?"
	destroyAllCommands = "DELETE FROM `commands` WHERE `channel`=?"
	destroyCommand     = "DELETE FROM `commands` WHERE `channel`=? AND `name`=?"
)

func CreateCommand(service *service.Service, channel string, name string, response string, creator string) (*Command, error) {
	now := time.Now()
	command := &Command{
		Channel:     channel,
		Name:        name,
		Response:    response,
		Enabled:     true,
		Restricted:  false,
		Cooldown:    0,
		Schedule:    0,
		UpdatedBy:   creator,
		UpdatedAt:   &now,
		Description: "",
	}
	_, err := service.NamedExec(createCommand, command)
	return command, err
}

func GetCommand(service *service.Service, channel string, name string) (*Command, error) {
	command := Command{}
	err := service.Get(&command, getCommand, channel, name)
	return &command, err
}

func GetCommands(service *service.Service, channel string) ([]Command, error) {
	commands := []Command{}
	err := service.Select(&commands, getCommands, channel)
	return commands, err
}

func DestroyAllCommands(service *service.Service, channel string) error {
	_, err := service.Exec(destroyAllCommands, channel)
	return err
}

func UpdateCommand(service *service.Service, command *Command) error {
	_, err := service.NamedExec(updateCommand, command)
	return err
}

func DestroyCommand(service *service.Service, channel string, name string) error {
	_, err := service.Exec(destroyCommand, channel, name)
	return err
}
