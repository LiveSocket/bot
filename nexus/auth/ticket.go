package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/router/auth"
	"github.com/gammazero/nexus/v3/wamp"
)

type TicketAuthenticator struct {
	auth.TicketAuthenticator
	Socket *socket.Socket
}

func NewTicketAuthenticator(socket *socket.Socket) *TicketAuthenticator {
	return &TicketAuthenticator{Socket: socket}
}

func (t *TicketAuthenticator) Authenticate(sid wamp.ID, details wamp.Dict, client wamp.Peer) (*wamp.Welcome, error) {
	// The HELLO.Details.authid|string is the authentication ID (e.g. username)
	// the client wishes to authenticate as. For Ticket-based authentication,
	// this MUST be provided.
	authID, _ := wamp.AsString(details["authid"])
	if authID == "" {
		return nil, errors.New("missing authid")
	}

	// Challenge Extra map is empty since the ticket challenge only asks for a
	// ticket (using authmethod) and provides no additional challenge info.
	err := client.Send(&wamp.Challenge{
		AuthMethod: t.AuthMethod(),
		Extra:      wamp.Dict{},
	})
	if err != nil {
		return nil, err
	}

	// Read AUTHENTICATE response from client.
	msg, err := wamp.RecvTimeout(client, 5*time.Second)
	if err != nil {
		return nil, err
	}
	authRsp, ok := msg.(*wamp.Authenticate)
	if !ok {
		return nil, fmt.Errorf("unexpected %v message received from client %v",
			msg.MessageType(), client)
	}

	if authID == "access_token" {
		return AuthenticateTwitch(t.Socket, authRsp)
	}
	return AuthenticatePi(t.Socket, authID, authRsp)
}
