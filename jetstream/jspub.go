package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"

	strm "github.com/Arnobkumarsaha/learn-nats/jetstream/utils"
)

func main() {
	var nc *nats.Conn
	js := strm.JetStreamContext(nc)
	defer nc.Close()

	publish(js, "test", testMsg)
	publish(js, "test.x.y.z", testXYZMsg)

}

func publish(js nats.JetStreamContext,
	subj string,
	f func() []byte, // f is a function that returns the message to be published
) {
	ack, err := js.Publish(subj, f())
	if err != nil {
		log.Printf("publish error: %v", err)
	}
	fmt.Printf("%#v\n", ack)
}

func testMsg() []byte {
	return []byte(
		fmt.Sprintf("t - %s", time.Now().Format("15:04:05")),
	)
}

func testXYZMsg() []byte {
	return []byte(
		fmt.Sprintf("xyz - %s", time.Now().Format("15:04:05")),
	)
}
