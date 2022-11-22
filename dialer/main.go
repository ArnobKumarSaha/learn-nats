package main

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func getOptions() ([]nats.Option, context.Context, context.CancelFunc) {
	// Parent context cancels connecting/reconnecting altogether.
	cc, _ := context.WithTimeout(context.Background(), time.Second*5)
	ctx, cancel := context.WithCancel(cc)
	defer cancel()
	cd := &customDialer{
		ctx:             ctx,
		connectTimeout:  10 * time.Second,
		connectTimeWait: 5 * time.Second,
	}

	opts := []nats.Option{
		nats.SetCustomDialer(cd),
		nats.ReconnectWait(2 * time.Second),
		nats.ReconnectHandler(func(c *nats.Conn) {
			log.Println("Reconnected to", c.ConnectedUrl())
		}),
		nats.DisconnectHandler(func(c *nats.Conn) {
			log.Println("Disconnected from NATS")
		}),
		nats.ClosedHandler(func(c *nats.Conn) {
			log.Println("NATS connection is closed.")
		}),
		nats.NoReconnect(),
	}
	return opts, ctx, cancel
}

func main() {
	opts, ctx, cancel := getOptions()
	defer cancel()
	var err error
	var nc *nats.Conn
	go func() {
		nc, err = nats.Connect(nats.DefaultURL, opts...)
	}()

	fmt.Println("I am here")

WaitForEstablishedConnection:
	for {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Trying ...")

		// Wait for context to be canceled either by timeout
		// or because of establishing a connection...
		select {
		case <-ctx.Done():
			fmt.Println("1")
			break WaitForEstablishedConnection
		default:
			fmt.Println("2")
		}

		if nc == nil || !nc.IsConnected() {
			log.Println("Connection not ready")
			time.Sleep(200 * time.Millisecond)
			continue
		}
		break WaitForEstablishedConnection
	}
	if ctx.Err() != nil {
		log.Fatal(ctx.Err())
	}
	fmt.Println("hello world")

	for {
		if nc.IsClosed() {
			break
		}
		if err := nc.Publish("hello", []byte("world")); err != nil {
			log.Println(err)
			time.Sleep(1 * time.Second)
			continue
		}
		log.Println("Published message")
		time.Sleep(1 * time.Second)
	}

	// Disconnect and flush pending messages
	if err := nc.Drain(); err != nil {
		log.Println(err)
	}
	log.Println("Disconnected")
}
