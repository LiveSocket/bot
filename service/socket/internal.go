package socket

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gammazero/nexus/v3/client"
	"github.com/gammazero/nexus/v3/wamp"
	"github.com/gammazero/nexus/v3/wamp/crsign"
)

func open(model model) *client.Client {
	cfg := client.Config{
		Realm: model.realm,
		HelloDetails: wamp.Dict{
			"authid": model.user,
		},
		AuthHandlers: map[string]client.AuthFunc{
			"wampcra": cra(model.pword),
		},
		Logger: log.New(os.Stdout, "", 0),
	}
	println("Attempting to connect to WebSocket router...")
	c, err := client.ConnectNet(context.Background(), model.address, cfg)
	retries := 0
	for err != nil && retries < 30 {
		time.Sleep(2 * time.Second)
		println("Retrying to connect to WebSocket router...")
		c, err = client.ConnectNet(context.Background(), model.address, cfg)
		if err != nil {
			log.Print(err)
		}
		retries++
	}
	if err != nil {
		panic(err)
	}
	return c
}

func registerActions(model model) {
	if model.actions != nil {
		for proc, handler := range model.actions {
			register(model, proc, handler)
		}
	}
}

func registerSubscriptions(model model) {
	if model.subscriptions != nil {
		for topic, handler := range model.subscriptions {
			subscribe(model, topic, handler)
		}
	}
}

func register(model model, proc string, handler CallHandler) {
	fn := func(ctx context.Context, invocation *wamp.Invocation) client.InvokeResult {
		result := handler(invocation)
		return result
	}
	println("Registering:", proc)
	model.client.Register(proc, fn, wamp.Dict{wamp.OptDiscloseCaller: true, wamp.OptInvoke: wamp.InvokeRoundRobin})
}

func subscribe(model model, topic string, handler client.EventHandler) {
	println("Subscribing to:", topic)
	model.client.Subscribe(topic, handler, nil)
}

func cra(secret string) client.AuthFunc {
	return func(c *wamp.Challenge) (string, wamp.Dict) {
		return crsign.RespondChallenge(secret, c, nil), wamp.Dict{}
	}
}
