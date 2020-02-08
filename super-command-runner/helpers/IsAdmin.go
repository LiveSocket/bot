package helpers

import (
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// IsAdmin WAMP call helper to check if username is admin
func IsAdmin(service *service.Service, username string) (bool, error) {
	// Call admin get endpoint
	res, err := service.SimpleCall("private.admin.get", nil, wamp.Dict{"username": username})
	if err != nil {
		return false, err
	}
	// If admin is found return true
	if len(res.Arguments) > 0 && res.Arguments[0] != nil {
		admin, err := conv.ToStringMap(res.Arguments[0])
		if err != nil {
			return false, err
		}
		if conv.ToString(admin["username"]) == username {
			return true, nil
		}
	}
	return false, nil
}
