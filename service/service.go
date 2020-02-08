package service

import (
	"github.com/LiveSocket/bot/service/db"
	"github.com/LiveSocket/bot/service/healthcheck"
	"github.com/LiveSocket/bot/service/socket"
)

const (
	ALL         = 0x7
	SOCKET      = 0x1
	DB          = 0x2
	HEALTHCHECK = 0x4
)

type Service struct {
	*db.DB
	*socket.Socket
}

type model struct {
	db     *db.DB
	socket *socket.Socket
}

type Actions = map[string]socket.Action
type Subscriptions = map[string]socket.Subscription

// Init Creates a new service using standard livesocket project settings
func (service *Service) Init(actions Actions, subscriptions Subscriptions, migrationTable string, migrations ...db.Migration) func() {

	d, dbClose := db.Init(migrationTable, migrations...)
	s := &socket.Socket{}
	socketClose := s.Init(actions, subscriptions)

	service.DB = d
	service.Socket = s
	healthcheck.Init()

	return func() {
		dbClose()
		socketClose()
	}
}
