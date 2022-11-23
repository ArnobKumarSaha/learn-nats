package backend

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"k8s.io/apimachinery/pkg/util/wait"
	"time"
)

type Options struct {
	AckWait time.Duration
	Stream  string
	Name    string
	// others
}

func DefaultOptions() Options {
	// return Options with name=os.Hostname(), & stream = "scanner"
	return Options{}
}

type Manager struct {
	nc     *nats.Conn
	sub    *nats.Subscription
	stream string
	// others
}

func New(nc *nats.Conn, opts Options) *Manager {
	return &Manager{
		nc:     nc,
		stream: opts.Stream,
		// set the other options from opts
	}
}

func (mgr *Manager) Start(ctx context.Context, jsmOpts ...nats.JSOpt) error {
	mgr.nc.QueueSubscribe(fmt.Sprintf("%s.report", mgr.stream), "scanner-backend", func(msg *nats.Msg) {
		img := string(msg.Data)
		data, err := DownloadReport(nil, img)
		if err != nil {
			// convert the error to metav1.Status
			// If it is docker produced `tooManyRequests` error:
			//     sleep for 1 hour
			// mgr.submitScanRequest()
		}
		msg.Respond(data)
	})

	// exactly same for Summary
	// ensure the stream with name=scanner & subject=scanner.queue.*
	// create nats consumer with name `workers`
	// Create a subscription for durable pull consumer

	// Launch two workers to process Foo resources
	for i := 0; i < 2; i++ {
		go wait.Until(mgr.runWorker, 5*time.Second, ctx.Done())
	}

	return nil
}

func (mgr *Manager) submitScanRequest(img string) {
	SubmitScanRequest(mgr.nc, fmt.Sprintf("%s.queue.scan", mgr.stream), img)
}

func SubmitScanRequest(nc *nats.Conn, subj, img string) {
	nc.Request(subj, []byte(img), time.Second)
}

func (mgr *Manager) runWorker() {
	for {
		err := mgr.processNextMsg()
		if err != nil {
			break
		}
	}
}

func (mgr *Manager) processNextMsg() (err error) {
	// fetch the messages one by one from mgr.sub
	// acknowledge the fetch
	// If the report doesn't exist yet: Calculate & upload it
	return nil
}
