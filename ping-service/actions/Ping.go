package actions

import (
	"github.com/LiveSocket/bot/service/socket"
)

// Ping Accepts pings
//
// public.ping
//
// Returns nothing
func Ping(s *socket.Socket) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		return socket.Success()
	}
}
