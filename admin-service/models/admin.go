package models

import (
	"github.com/LiveSocket/bot/service"
)

// Admin Represents a Admin user
type Admin struct {
	Username string `db:"username" json:"username"`
	Notes    string `db:"notes" json:"notes"`
}

const getAdmin = "SELECT * FROM `admin` WHERE `username`=?"

// GetAdmin Finds an admin from the database
func GetAdmin(service *service.Service, username string) (*Admin, error) {
	admin := Admin{}
	err := service.Get(&admin, "SELECT * FROM `admin` WHERE `username`=?", username)
	return &admin, err
}
