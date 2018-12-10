package main

import (
	"errors"
	"os"
	"github.com/smartmakers/drivers/go/driver"
)

func main() {
	drv := driver.New()
	drv.Decoder = decodePayload
	drv.Run(os.Args[1:])
}

func decodePayload(payload []byte, fPort int) (driver.DecodedPayload, error) {
	// Implement your own decoder here
	return nil, errors.New("not yet implemented")
}
