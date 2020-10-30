package main

import (
	"github.com/shawnhankim/natsio/nats"
)

func main() {
	nats.RunServer()
	nats.RunClient()
}
