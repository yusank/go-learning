package main

// Connect, subscribe on channel, publish into channel, read presence and history info.

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/centrifugal/centrifuge-go"
	"github.com/centrifugal/centrifugo/libcentrifugo/auth"
)



func main() {
	// Never show secret to client of your application. Keep it on your application backend only.
	secret := "secret"

	// Application user ID.
	user := "42"

	// Current timestamp as string.
	timestamp := centrifuge.Timestamp()

	// Empty info.
	info := ""

	// Generate client token so Centrifugo server can trust connection parameters received from client.
	token := auth.GenerateClientToken(secret, user, timestamp, info)

	creds := &centrifuge.Credentials{
		User:      user,
		Timestamp: timestamp,
		Info:      info,
		Token:     token,
	}

	started := time.Now()

	wsURL := "ws://10.0.0.252:8110/connection/websocket"
	conf := centrifuge.DefaultConfig
	c := centrifuge.NewCentrifuge(wsURL, creds, nil, conf)
	defer c.Close()

	err := c.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	onMessage := func(sub centrifuge.Sub, msg centrifuge.Message) error {
		l := map[string]string{}
		err := json.Unmarshal(*msg.Data, &l)

		if err != nil {
			log.Println(fmt.Sprintf("unmarshall err with : %e", err))
		}
		
		a, _ := l["userid"]
		log.Print(a)

		log.Println(fmt.Sprintf("channal %v", sub))
		log.Println(fmt.Sprintf("New message received in channel %s: %#v : %s", sub.Channel(), *msg.Info, *msg.Data))
		return nil
	}

	onJoin := func(sub centrifuge.Sub, msg centrifuge.ClientInfo) error {
		log.Println(fmt.Sprintf("User %s joined channel %s with client ID %s", msg.User, sub.Channel(), msg.Client))
		return nil
	}

	onLeave := func(sub centrifuge.Sub, msg centrifuge.ClientInfo) error {
		log.Println(fmt.Sprintf("User %s with clientID left channel %s with client ID %s", msg.User, msg.Client, sub.Channel()))
		return nil
	}

	events := &centrifuge.SubEventHandler{
		OnMessage: onMessage,
		OnJoin:    onJoin,
		OnLeave:   onLeave,
	}

	sub, err := c.Subscribe("smartestee:login", events)
	if err != nil {
		log.Fatalln(err)
	}

	data := map[string]string{
		"userid": "17",
		"server": "22",
	}
	
	dataBytes, _ := json.Marshal(data)
	err = sub.Publish(dataBytes)
	if err != nil {
		log.Fatalln(err)
	}

	history, err := sub.History()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%d messages in channel %s history", len(history), sub.Channel())

	presence, err := sub.Presence()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%d clients in channel %s", len(presence), sub.Channel())

	select{}	

	log.Printf("%s", time.Since(started))

}

