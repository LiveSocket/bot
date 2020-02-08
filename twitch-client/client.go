package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
	"github.com/gempir/go-twitch-irc/v2"
)

type Client struct {
	Twitch      *twitch.Client
	BotUsername string
	BotPassword string
	Joined      []string
}

type OnMessage = func(channel string, user twitch.User, message twitch.Message)

func (client *Client) Start(service *service.Service) error {

	err := errors.New("")
	for err != nil {
		time.Sleep(2 * time.Second)
		println("Attempting to connect to channels...")
		err = client.connectToChannels(service)
		fmt.Println(err)
	}
	return err
}

func (client *Client) connectToChannels(service *service.Service) error {

	channels, err := client.getChannels(service)
	if err != nil {
		return err
	}

	for _, channel := range channels {
		fmt.Println("Joining ", channel, "...")
		client.Joined = append(client.Joined, channel)
		client.Twitch.Join(channel)
	}

	fmt.Println("Connecting...")
	err = client.Twitch.Connect()
	return err
}

func (client *Client) getChannels(service *service.Service) ([]string, error) {
	response, err := service.SimpleCall("private.channel.get", nil, wamp.Dict{"botName": client.BotUsername})
	if err != nil {
		return nil, err
	}
	if response.Arguments == nil {
		return nil, nil
	}

	result := make([]string, len(response.Arguments))
	for i, item := range response.Arguments {
		base, err := conv.ToStringMap(item)
		if err != nil {
			return nil, err
		}
		result[i] = conv.ToString(base["name"])
	}
	return result, nil
}
