package main

import (
	"os"

	"github.com/smartmakers/drivers/go/driver"
)

const (
	feedChannel   = 0
	returnChannel = 1
)

func main() {
	drv := driver.New()
	drv.Decoder = decodePayload
	drv.Run(os.Args[1:])
}

func decodePayload(payload []byte, fPort int) (driver.DecodedPayload, error) {
	// unmarshal LPP payload first:
	uplink := Uplink{}
	err := uplink.UnmarshalBinary(payload)

	temps := Obj{}

	//  extract data from LPP payload by "contextualizing" it, e.g.
	// * channel 0 is feed temperature
	// * channel 1 is return temperature
	if v, ok := uplink[feedChannel]; ok {
		temps["feed"] = v
	}

	if v, ok := uplink[returnChannel]; ok {
		temps["return"] = v
	}

	return Obj{"temperatures": temps}, err
}

// Obj is syntactic sugar for creating untyped objects.
type Obj map[string]interface{}
