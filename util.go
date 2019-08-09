package comms

import (
	"fmt"
	nats "github.com/nats-io/nats.go"
)

func NewNATS(address string) *nats.EncodedConn {
	nc, err := nats.Connect(address)
	if(err != nil) {
		fmt.Printf("Cannot connect to NATS: %s\n", err)
		panic(err)
	}
	fmt.Printf("Connected.\nEncoding...\n")
	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if(err != nil) {
		fmt.Printf("Cannot encode the NATS connection with JSON for some reason: %s\n", err)
		panic(err)
	}
	return c
}