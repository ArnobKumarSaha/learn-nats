package main

import (
	"fmt"
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
	respond(nc)
	sync(nc)
}

func respond(nc *nats.Conn) {
	_, err := nc.Subscribe("foo", func(msg *nats.Msg) {
		log.Println("Request received:", string(msg.Data))

		err := msg.Respond([]byte("Here you go!"))
		if err != nil {
			// `message does not have a reply` error will be shown for Publish() call
			fmt.Printf("error on Respond() %s \n", err)
			return
		}
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
