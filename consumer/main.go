package main

import (
	"github.com/getmilly/grok/logging"
	gnats "github.com/getmilly/grok/nats"
	"github.com/nats-io/go-nats"
)

func main() {
	conn, err := gnats.Connect(nats.DefaultURL, "")

	if err != nil {
		panic(err)
	}

	gnats.NewSubscriber(conn).
		SetSubject("app-subject").
		SetQueue("app-queue").
		SetHandler(func(message *gnats.Message) error {
			logging.LogWith(message).Info("incoming message")
			return nil
		}).
		Run()
}
