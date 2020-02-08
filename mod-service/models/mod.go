package models

import (
	"time"

	"github.com/LiveSocket/bot/service"
)

// Mod Represents a Mod user
type Mod struct {
	Channel   string     `db:"channel" json:"channel"`
	Username  string     `db:"username" json:"username"`
	CreatedAt *time.Time `db:"created_at" json:"createdAt"`
}

const (
	createMod          = "INSERT INTO `mods` (`channel`,`username`,`created_at`) VALUES (?,?,CURRENT_TIMESTAMP)"
	getMod             = "SELECT * FROM `mods` WHERE `channel`=? AND `username`=?"
	getMods            = "SELECT * FROM `mods` WHERE `channel`=?"
	getChannelsForMods = "SELECT * FROM `mods` WHERE `username`=?"
	destroyAllMods     = "DELETE FROM `mods` WHERE `channel`=?"
	destroyMod         = "DELETE FROM `mods` WHERE `channel`=? AND `username`=?"
)

// CreateMod Creates a new Mod
func CreateMod(service *service.Service, channel string, username string) (*Mod, error) {
	_, err := service.Exec(createMod, channel, username)
	if err != nil {
		return nil, err
	}
	return GetMod(service, channel, username)
}

func GetMod(service *service.Service, channel string, username string) (*Mod, error) {
	mod := Mod{}
	err := service.Get(&mod, getMod, channel, username)
	return &mod, err
}

func GetMods(service *service.Service, channel string) ([]Mod, error) {
	mods := []Mod{}
	err := service.Select(&mods, getMods, channel)
	return mods, err
}

func GetChannelsForMods(service *service.Service, username string) ([]Mod, error) {
	mods := []Mod{}
	err := service.Select(&mods, getChannelsForMods, username)
	return mods, err
}

func DestroyAllMods(service *service.Service, channel string) error {
	_, err := service.Exec(destroyAllMods, channel)
	return err
}

func DestroyMod(service *service.Service, channel, username string) error {
	_, err := service.Exec(destroyMod, channel, username)
	return err
}
