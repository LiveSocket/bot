package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	routerAuth "github.com/LiveSocket/bot/nexus/auth"
	_socket "github.com/LiveSocket/bot/service/socket"

	"github.com/gammazero/nexus/v3/router"
	"github.com/gammazero/nexus/v3/router/auth"
	"github.com/gammazero/nexus/v3/wamp"
)

var ConfigInfo *Config

func main() {
	// configPtr := flag.String("config", "", "Config file to load")
	// flag.Parse()
	// if *configPtr == "" {
	// 	log.Fatal("-config=<config-file> required")
	// }

	ConfigInfo := LoadConfiguration()

	sks := &serverKeyStore{
		provider: "static",
		config:   ConfigInfo,
	}
	craAuth := auth.NewCRAuthenticator(sks, time.Second)
	socket := &_socket.Socket{}
	tAuth := routerAuth.NewTicketAuthenticator(socket)

	routerConfig := &router.Config{
		RealmConfigs: []*router.RealmConfig{
			&router.RealmConfig{
				URI:            wamp.URI(ConfigInfo.Realm),
				AnonymousAuth:  false,
				AllowDisclose:  true,
				Authenticators: []auth.Authenticator{craAuth, tAuth},
				Authorizer:     &authorizer{},
			},
		},
	}
	nxr, err := router.NewRouter(routerConfig, nil)
	if err != nil {
		panic(err)
	}
	defer nxr.Close()

	server := router.NewWebsocketServer(nxr)
	server.EnableTrackingCookie = true
	address := ConfigInfo.Host + ":" + ConfigInfo.Port
	fmt.Println("Starting server at address", address)
	closer, err := server.ListenAndServe(address)
	if err != nil {
		panic(err)
	}
	log.Printf("Websocket server listening on ws://%s/", ConfigInfo.Host+":"+ConfigInfo.Port)
	close := socket.Init(nil, nil)
	tAuth.Socket = socket
	defer close()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	closer.Close()
	println("Router shutting down...")
}
