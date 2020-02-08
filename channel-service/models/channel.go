package models

import (
	"time"

	"github.com/LiveSocket/bot/service"
)

type Channel struct {
	Name      string     `db:"name" json:"name"`
	BotName   string     `db:"bot_name" json:"bot_name"`
	Notes     string     `db:"notes" json:"notes"`
	Ignored   bool       `db:"ignored" json:"ignored"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

const (
	createChannel     = "INSERT INTO `channels` (`name`,`bot_name`,`notes`,`ignored`,`updated_at`) VALUES (:name,:bot_name,:notes,:ignored,CURRENT_TIMESTAMP)"
	selectChannel     = "SELECT * FROM `channels` WHERE `name`=?"
	selectChannels    = "SELECT * FROM `channels` WHERE `bot_name`=?"
	selectAllChannels = "SELECT * FROM `channels`"
	updateChannel     = "UPDATE `channels` SET `bot_name`=:bot_name,`notes`=:notes,`updated_at`=CURRENT_TIMESTAMP WHERE `name`=:name"
	deleteChannel     = "UPDATE `channels` SET `deleted_at`=CURRENT_TIMESTAMP WHERE `name`=?"
	destroyChannel    = "DELETE FROM `channels` WHERE `name`=?"
)

func CreateChannel(service *service.Service, name string, botName string) (*Channel, error) {
	now := time.Now()
	channel := &Channel{
		Name:      name,
		BotName:   botName,
		Ignored:   false,
		UpdatedAt: &now,
		Notes:     "",
	}
	_, err := service.NamedExec(createChannel, channel)
	return channel, err
}

func GetChannel(service *service.Service, name string) (*Channel, error) {
	channel := &Channel{}
	err := service.Get(&channel, selectChannel, name)
	return channel, err
}

func GetBotChannels(service *service.Service, botName string) ([]Channel, error) {
	channels := []Channel{}
	err := service.Select(&channels, selectChannels, botName)
	return channels, err
}

func GetAllChannels(service *service.Service) ([]Channel, error) {
	channels := []Channel{}
	err := service.Select(&channels, selectAllChannels)
	return channels, err
}

func UpdateChannel(service *service.Service, channel *Channel) error {
	_, err := service.NamedExec(updateChannel, channel)
	return err
}

func DeleteChannel(service *service.Service, name string) error {
	_, err := service.Exec(deleteChannel, name)
	return err
}

func DestroyChannel(service *service.Service, name string) error {
	_, err := service.DB.Exec(destroyChannel, name)
	return err
}
