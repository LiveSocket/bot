package models

import (
	"database/sql"
	"time"

	"github.com/LiveSocket/bot/service"
)

// ModChat Represents a ModChat message
type ModChat struct {
	ID        uint      `db:"id" json:"id"`
	Channel   string    `db:"channel" json:"channel"`
	Message   string    `db:"message" json:"message"`
	Name      string    `db:"name" json:"name"`
	Timestamp time.Time `db:"timestamp" json:"timestamp"`
}

// NewModChat Creates a new ModChat message
func NewModChat(channel string, message string, name string) *ModChat {
	now := time.Now()
	return &ModChat{
		Channel:   channel,
		Message:   message,
		Name:      name,
		Timestamp: now,
	}
}

// FindLatestMessagesForChannel Finds the latest x number of mod chat messages for a channel
func FindLatestMessagesForChannel(service *service.Service, channel string, offset uint64, limit uint64) ([]ModChat, error) {
	messages := []ModChat{}
	err := service.Select(&messages, "SELECT * FROM `mod_chat` WHERE `channel`=? ORDER BY `timestamp` DESC LIMIT ?,?", channel, offset, limit)
	return messages, err
}

// CreateModChat Creates a ModChat record in the database
func CreateModChat(service *service.Service, modChat *ModChat) (sql.Result, error) {
	return service.NamedExec("INSERT INTO `mod_chat` (`channel`, `message`,`name`,`timestamp`) VALUES (:channel,:message,:name,:timestamp)", modChat)
}
