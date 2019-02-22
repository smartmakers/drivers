package lpp

import "errors"

// DigitalInput is a boolean, send to the device.
type DigitalInput bool

// UnmarshalBinary implements the encoding.binary Unmarshaler interface for DigitalInput
func (di *DigitalInput) UnmarshalBinary(data []byte) error {
	if len(data) != 1 {
		return errors.New("digital input must have size 1")
	}

	*di = data[0] != 0x00
	return nil
}

// MarshalBinary marshals the DigitalIO to a byte array
func (di DigitalInput) MarshalBinary() ([]byte, error) {
	if di {
		return []byte{0x01}, nil
	} else {
		return []byte{0x00}, nil
	}
}

// DigitalOutput is a boolean, send from the device.
type DigitalOutput bool

// UnmarshalBinary implements the encoding.binary Unmarshaler interface for DigitalInput
func (do *DigitalOutput) UnmarshalBinary(data []byte) error {
	if len(data) != 1 {
		return errors.New("digital input must have size 1")
	}

	*do = data[0] != 0x00
	return nil
}

// MarshalBinary marshals the DigitalIO to a byte array
func (do DigitalOutput) MarshalBinary() ([]byte, error) {
	if do {
		return []byte{0x01}, nil
	} else {
		return []byte{0x00}, nil
	}
}
