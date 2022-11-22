package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func sync(nc *nats.Conn) {
	sub, err := nc.SubscribeSync("bar")
	if err != nil {
		log.Fatal(err)
	}

	//nc.Request("bar", []byte("hello universe"), time.Second)
	//fmt.Println("request was called.")

	// Send the request, If processing is synchronous, use Request() which returns the response message.
	// PublishRequest is similar to Publish, Difference is,
	// It expects a response on the reply subject. whereas Publish waits for the response.
	nc.PublishRequest("sub", "bar", []byte("hello world")) // `no responders available for request` error will be thrown
	nc.Flush()

	// Wait for a single response
	for {
		msg, err := sub.NextMsg(1 * time.Second)
		if err != nil {
			log.Fatal(err)
		}

		response := string(msg.Data)
		fmt.Println(response)
		break
	}
	sub.Unsubscribe()
}
