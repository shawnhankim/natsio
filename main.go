package main

import (
	"github.com/shawnhankim/natsio"
)

func main() {
	natsio.RunServer()
	natsio.RunClient()
}
