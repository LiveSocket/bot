package auth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

func AuthenticateTwitch(socket *socket.Socket, authRsp *wamp.Authenticate) (*wamp.Welcome, error) {
	request, err := http.NewRequest(http.MethodGet, "https://id.twitch.tv/oauth2/validate", nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "OAuth "+authRsp.Signature)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var j map[string]interface{}
	err = json.Unmarshal(body, &j)
	if err != nil {
		return nil, err
	}

	log.Printf("%v", j)

	mod, err := socket.SimpleCall("private.mod.get", nil, wamp.Dict{"username": j["login"]})
	if err != nil {
		return nil, err
	}

	admin, err := socket.SimpleCall("private.admin.isAdmin", nil, wamp.Dict{"username": j["login"]})
	if err != nil {
		return nil, err
	}

	if len(mod.Arguments) != 0 {
		channels := []string{}
		for _, item := range mod.Arguments {
			channels = append(channels, item.(map[string]interface{})["channel"].(string))
		}

		role := "authenticated"
		if len(admin.Arguments) > 0 {
			value, err := conv.ToBool(admin.Arguments[0])
			if err != nil {
				return nil, err
			}
			if value {
				role = "admin"
			}
		}

		// Create welcome details containing auth info.
		welcome := &wamp.Welcome{
			Details: wamp.Dict{
				"authid":     j["login"],
				"authmethod": "ticket",
				"authrole":   role,
				"authextra": map[string]interface{}{
					"access_token": authRsp.Signature,
					"channels":     channels,
				},
			},
		}
		return welcome, nil
	}
	return nil, errors.New("Invalid login")
}
