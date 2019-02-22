package lpp

import (
	"errors"
	"fmt"
)

// Uplink represents a single Low Power Payload encoded uplink.
//
// A single LPP uplink can contain multiple
// data values, one per channel.
// However, this decoder assumes that the same channel
// always uses the same data type.
// This also means a channel is either for uplinks or
// for downlinks, but never for both at the same time.
// It does not seem to be entirely clear if this asusmption
// is guaranteed by the LPP specification.
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
	ty, length, err := dataType(code)
	if err != nil {
		return 0, err
	}

	if length > len(payload) {
		return 0, errors.New("payload size does not match data types")
	}

	err = ty.UnmarshalBinary(payload[0:length])
	if err != nil {
		return 0, err
	}

	(*u)[ch] = ty
	return length, nil
}

// dataType returns the DataType and it's length for a type code.
func dataType(code byte) (Data, int, error) {
	switch code {
	case 0x00:
		return new(DigitalInput), 1, nil
	case 0x01:
		return new(DigitalOutput), 1, nil
	case 0x02:
		return new(AnalogInput), 2, nil
	case 0x03:
		return new(AnalogOutput), 2, nil
	default:
		return nil, 0, fmt.Errorf("unsupported data type '%v'", code)
	}
}

// Data is the primary interface representing payload data.
//
// This is an empty interface: we currently only use it for
// unmarshaling it from binary and for marshaling JSON from it.
type Data interface {
	UnmarshalBinary(payload []byte) error
}
