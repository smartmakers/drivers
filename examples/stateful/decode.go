package main

import (
	"errors"
	"time"

	"github.com/smartmakers/drivers/go/cayenne"

	"github.com/smartmakers/drivers/go/driver/v2"
)

const (
	InputChannel  = 0
	OutputChannel = 1

	Open = "open"
)

// Decode decodes an uplink
func (d *Driver) Decode(req v2.DecodeRequest, resp *v2.DecodeResponse) error {
	newState := req.ReportedState.(State)

	uplink, err := cayenne.Decode(req.Payload, req.Port)
	if err != nil {
		return err
	}

	for channel, data := range *uplink {
		switch channel {
		case OutputChannel:
			// check if the output channel actually is a boolean
			if b, ok := data.(*cayenne.DigitalInput); ok {
				b := bool(*b)
				newState.Open = &b
				resp.Updates = append(resp.Updates, v2.Update{
					Timestamp: time.Now(),
					Values:    State{Open: &b}})
				continue
			}

			return errors.New("ouput channel must be boolean")
		case InputChannel:
			// This dynamically publishes the existence of the
			// InputChannel to the Application (and it's type).
			// We do not need this, as the user already knows about
			// the device's capabilites from the schema.
			continue
		default:
			return errors.New("unsupported channel")
		}
	}

	resp.NewState = newState
	return nil
}
