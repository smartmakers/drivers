package main

import (
	"errors"
	"fmt"
)

// Uplink represents a single Cayenne LPP uplink.
//
// A single Cayenne LPP uplink can contain multiple
// data values, one per channel.
// However, this decoder assumes that the same channel
// always uses the same data type.
// This also means a channel is either for uplinks or
// for downlinks, but never for both at the same time.
type Uplink map[byte]Data

// UnmarshalBinary unmarshal an Uplink from a binary payload.
func (u *Uplink) UnmarshalBinary(payload []byte) error {
	for len(payload) > 2 {
		ch := payload[0]
		ty := payload[1]
		consumed, err := u.unmarshalData(ch, ty, payload[2:])
		if err != nil {
			return err
		}

		payload = payload[2+consumed:]
	}

	if len(payload) != 0 {
		return errors.New("Did not consume the full payload")
	}

	return nil
}

// unmarshalData unmarshals a single data field
func (u *Uplink) unmarshalData(ch, code byte, payload []byte) (int, error) {
	ty, length, err := Type(code)
	if err != nil {
		return 0, err
	}

	err = ty.UnmarshalBinary(payload[0:length])
	if err != nil {
		return 0, err
	}

	(*u)[ch] = ty
	return length, nil
}

// Data is the primary interface representing payload data.
//
// This is an empty interface: we currently only use it for
// unmarshaling it from binary and for marshaling JSON from it.
type Data interface {
	UnmarshalBinary(payload []byte) error
}

// Type returns the DataType and it's length for a type code.
func Type(code byte) (Data, int, error) {
	switch code {
	case 0x00:
		return new(DigitalInput), 1, nil
	case 0x01:
		return new(DigitalOutput), 1, nil
	default:
		return nil, 0, fmt.Errorf("unsupported data type '%v'", code)
	}
}
