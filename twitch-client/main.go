package main

import (
	"log"
	"os"
	"strings"

	"github.com/gammazero/nexus/v3/wamp"

	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gempir/go-twitch-irc/v2"
)

func main() {
	client := &Client{
		Twitch:      twitch.NewClient(os.Getenv("BOT_UNAME"), os.Getenv("BOT_PWORD")),
		BotUsername: os.Getenv("BOT_UNAME"),
		BotPassword: os.Getenv("BOT_PWORD"),
	}

	s := &service.Service{}
	close := s.Init(service.Actions{
		"private.twitch.chat.say": speaker(s, client),
	}, nil, "")
	defer close()

	// Register the speak endpoint

	client.Twitch.OnConnect(func() {
		client.Twitch.OnPrivateMessage(MessageHandler(s, client))
		// client.Twitch.OnUserJoinMessage(UserJoinHandler(s, client))
		// client.Twitch.OnUserPartMessage(UserPartHandler(s, client))
		// client.Twitch.OnUserNoticeMessage(UserNoticeHandlers(s, client))
	})

	if err := client.Start(s); err != nil {
		panic(err)
	}
}

func MessageHandler(service *service.Service, client *Client) func(twitch.PrivateMessage) {
	return func(message twitch.PrivateMessage) {
		if message.User.Name != strings.ToLower(client.BotUsername) {
			if err := service.Publish("event.chat.message", nil, wamp.List{message}, nil); err != nil {
				log.Print(err)
			}
		}
	}
}

func speaker(service *service.Service, client *Client) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		if invocation.Arguments == nil {
			log.Print("Cannot speak without instructions")
			return socket.Error("Cannot speak without instructions")
		}
		if len(invocation.Arguments) < 2 {
			log.Print("Cannot speak without channel and message")
			return socket.Error("Cannot speak without channel and message")
		}
		log.Printf("%v", invocation.Arguments[1])
		channel := conv.ToString(invocation.Arguments[0])
		message := conv.ToString(invocation.Arguments[1])
		client.Twitch.Say(channel, "/me "+message)
		return socket.Success()
	}
}
