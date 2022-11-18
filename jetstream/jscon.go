package main

import (
	strm "github.com/Arnobkumarsaha/learn-nats/jetstream/utils"
	"github.com/nats-io/nats.go"
)

func main() {
	var nc *nats.Conn
	js := strm.JetStreamContext(nc)
	defer nc.Close()

	const strName = "tst"
	strm.AddConsumer(js, strName, "test-dot", "test.>")
	strm.AddConsumer(js, strName, "test", "test")
}
