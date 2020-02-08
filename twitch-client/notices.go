package main

import (
	"log"

	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
	"github.com/gempir/go-twitch-irc/v2"
)

func UserNoticeHandler(service *service.Service, client *Client) func(twitch.UserNoticeMessage) {
	return func(message twitch.UserNoticeMessage) {
		switch message.MsgID {
		case "resub":
		case "sub":
		case "subgift":
		case "anonsubgift":
		case "submysterygift":
		case "anonsubmysterygift":
		case "primepaidupgrade":
			broadcast("subscription", service, client, message)
			return
		case "giftpaidupgrade":
		case "anongiftpaidupgrade":

		case "raid":
			broadcast("raid", service, client, message)
			return
		}
	}
}

func broadcast(name string, service *service.Service, client *Client, message twitch.UserNoticeMessage) {
	if err := service.Publish("event."+name, nil, wamp.List{message}, nil); err != nil {
		log.Print(err)
	}
}
