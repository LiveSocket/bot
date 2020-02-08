package models

import (
	"fmt"
	"strings"

	"github.com/LiveSocket/bot/service"
)

// SuperCommand represents a super command
type SuperCommand struct {
	Proc string `db:"proc" json:"proc"`
	Name string `db:"name" json:"name"`
}

// FindSuperCommand Finds a registered super command from the database
func FindSuperCommand(service *service.Service, name string) (*SuperCommand, error) {
	super := SuperCommand{}
	err := service.Get(&super, "SELECT * FROM `super_commands` WHERE name=?", name)
	if err != nil {
		if strings.Contains(fmt.Sprint(err), "no rows") {
			return nil, nil
		}
		return nil, err
	}
	return &super, nil
}
