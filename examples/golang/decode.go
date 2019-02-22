package main

import (
	"fmt"

	"github.com/smartmakers/drivers/go/driver/v1"
	"github.com/smartmakers/drivers/go/lpp"
)

const (
	feedChannel   = 0
	returnChannel = 1
)

func decode(payload []byte, port int) (v1.DecodedPayload, error) {
	// unmarshal generic Cayenne LPP payload first
	uplink, err := lpp.Decode(payload, port)
	if err != nil {
		return nil, err
	}

	// convert generic payload to device-specific payload:
	temps := Obj{}
	for channel, data := range *uplink {
		// Need to also check for the right data types here,
		// i.e. is this really as described in the data schema?
		switch channel {
		case feedChannel:
			temps["feed"] = data
		case returnChannel:
			temps["return"] = data
		default:
			return nil, fmt.Errorf("unsupported channel %v", channel)
		}
	}

	return Obj{"temperatures": temps}, err
}

// Obj is syntactic sugar for creating untyped objects.
type Obj map[string]interface{}
