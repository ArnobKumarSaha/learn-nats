package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatalln(err)
	}

	defer nc.Close()

	nc.Subscribe("foo", func(msg *nats.Msg) {
		log.Println("Request received:", string(msg.Data))

		msg.Respond([]byte("Here you go!"))
	})

	// this block is added by me
	if err := nc.Publish("foo", []byte("Message")); err != nil {
		log.Fatalln(err)
	}

	reply, err := nc.Request("foo", []byte("Give me data"), 4*time.Second)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Got Reply:", string(reply.Data))
}
