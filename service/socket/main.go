package socket

import (
	"context"
	"fmt"
	"os"

	"github.com/gammazero/nexus/v3/client"
	"github.com/gammazero/nexus/v3/wamp"
)

// Socket a WAMP websocket connection
type Socket struct {
	*client.Client
	model model
}

// CallHandler the standard CALL handler function
type CallHandler func(invocation *Invocation) Result

type Action = CallHandler
type Subscription = client.EventHandler
type Actions = map[string]Action
type Subscriptions = map[string]Subscription
type Result = client.InvokeResult
type Invocation = wamp.Invocation
type Event = wamp.Event

type model struct {
	client        *client.Client
	address       string
	realm         string
	user          string
	pword         string
	actions       Actions
	subscriptions Subscriptions
}

func (s *Socket) Init(actions Actions, subscriptions Subscriptions) func() {
	model := model{
		address:       os.Getenv("NEXUS_ADDRESS"),
		realm:         os.Getenv("NEXUS_REALM"),
		user:          os.Getenv("NEXUS_USER"),
		pword:         os.Getenv("NEXUS_PWORD"),
		actions:       actions,
		subscriptions: subscriptions,
	}
	model.client = open(model)
	registerActions(model)
	registerSubscriptions(model)
	s.Client = model.client
	return func() {
		if err := model.client.Close(); err != nil {
			panic(err)
		}
	}
}

// SimpleCall Sends a CALL message using the provided parameters
func (socket *Socket) SimpleCall(proc string, args []interface{}, kwargs map[string]interface{}) (*wamp.Result, error) {
	return socket.Call(context.Background(), proc, wamp.Dict{wamp.OptInvoke: wamp.InvokeRoundRobin, wamp.OptTimeout: 2000}, args, kwargs, nil)
}

// Register registers a new wamp endpoint using the provided parameters
func (socket *Socket) Register(proc string, handler CallHandler) {
	register(socket.model, proc, handler)
}

// Subscribe subscribes to a topic using the provided parameters
func (socket *Socket) Subscribe(topic string, handler client.EventHandler) {
	subscribe(socket.model, topic, handler)
}

// Error Helper for returning errors from a CALL handler
func Error(err interface{}) Result {
	return Result{
		Err: wamp.URI(fmt.Sprint(err)),
	}
}

// Success Helper for returning args from a CALL handler
func Success(args ...interface{}) Result {
	list := wamp.List{}
	for _, item := range args {
		list = append(list, item)
	}
	return Result{
		Args: list,
	}
}
