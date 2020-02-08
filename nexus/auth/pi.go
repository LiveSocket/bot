package auth

import (
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

func AuthenticatePi(socket *socket.Socket, username string, authRsp *wamp.Authenticate) (*wamp.Welcome, error) {
	_, err := socket.SimpleCall("private.pi.authenticate", nil, wamp.Dict{"username": username, "password": authRsp.Signature})
	if err != nil {
		return nil, err
	}

	// Create welcome details containing auth info.
	welcome := &wamp.Welcome{
		Details: wamp.Dict{
			"authid":     username,
			"authmethod": "ticket",
			"authrole":   "pi",
			"authextra":  map[string]interface{}{},
		},
	}
	return welcome, nil
}
