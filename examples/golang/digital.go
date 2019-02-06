package main

import "errors"

// DigitalInput is a boolean, send from the device to the server.
type DigitalInput bool

// UnmarshalBinary implements the encoding.binary Unmarshaler interface for DigitalInput
func (di *DigitalInput) UnmarshalBinary(data []byte) error {
	if len(data) != 1 {
		return errors.New("digital input must have size 1")
	}

	*di = data[0] != 0x00
	return nil
}

type DigitalOutput bool

func (do *DigitalOutput) UnmarshalBinary(data []byte) error {
	if len(data) != 1 {
		return errors.New("digital input must have size 1")
	}

	*do = data[0] != 0x00
	return nil
}
