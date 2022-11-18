package utils

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

var (
	NatsURL = "nats://127.0.0.1:4222"
)

// JetStreamContext takes a nats connections and returns a JetStreamContext.
func JetStreamContext(nc *nats.Conn) nats.JetStreamContext {
	if nc == nil {
		var err error
		nc, err = nats.Connect(NatsURL)
		if err != nil {
			log.Fatalf("could not connect to NATS: %v", err)
		}
	}

	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("could not create JetStream context: %v", err)
	}

	return js
}

func Create(js nats.JetStreamContext, name string) *nats.StreamInfo {
	fmt.Printf("Creating stream: %q\n", name)
	strInfo, err := js.AddStream(&nats.StreamConfig{
		Name:     name,
		Subjects: []string{"test.>", "test"},
		MaxAge:   0, // 0 means keep forever
		Storage:  nats.FileStorage,
	})
	if err != nil {
		log.Panicf("could not create stream: %v", err)
	}

	prettyPrint(strInfo)
	return strInfo
}

func prettyPrint(x interface{}) {
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		log.Fatalf("could not prettyPrint: %v", err)
	}
	fmt.Println(string(b))
}

func Delete(js nats.JetStreamContext, name string) {
	fmt.Printf("Deleting stream: %q\n", name)
	if err := js.DeleteStream(name); err != nil {
		log.Printf("error deleting stream: %v", err)
	}
}

func AddConsumer(js nats.JetStreamContext, strmName, consName, consFilter string) {
	info, err := js.AddConsumer(strmName, &nats.ConsumerConfig{
		Durable:   consName,
		AckPolicy: nats.AckExplicitPolicy,
		// MaxAckPending: 1,      // default value is 20,000
		FilterSubject: consFilter,
	})
	if err != nil {
		log.Panicf("could not add consumer: %v", err)
	}
	prettyPrint(info)
}
