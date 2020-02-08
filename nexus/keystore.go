package main

import (
	"crypto/sha256"
	"errors"

	"golang.org/x/crypto/pbkdf2"
)

type serverKeyStore struct {
	provider string // static
	config   Config
}

func (ks *serverKeyStore) AuthKey(authid, authmethod string) ([]byte, error) {
	switch authmethod {
	case "wampcra":
		if authid != ks.config.User {
			return nil, errors.New("Invalid User")
		}

		return pbkdf2.Key([]byte(ks.config.Password), []byte(ks.config.Salt), ks.config.Iterations, ks.config.KeyLen, sha256.New), nil
	default:
		return nil, errors.New("unsupported authmethod")
	}
}

func (ks *serverKeyStore) PasswordInfo(authid string) (string, int, int) {
	return ks.config.Salt, ks.config.KeyLen, ks.config.Iterations
}

func (ks *serverKeyStore) Provider() string { return ks.provider }

func (ks *serverKeyStore) AuthRole(authid string) (string, error) {
	if authid != ks.config.User {
		return "", errors.New("no such user: " + authid)
	}
	return ks.config.Role, nil
}

// func (ks *serverKeyStore) OnWelcome(authid string, welcome *wamp.Welcome, details wamp.Dict) error {

// 	return nil
// }
