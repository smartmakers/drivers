package cayenne

import "errors"

// AnalogInput is an 16-bit integer,
// send from the device to the server.
type AnalogInput int16

// UnmarshalBinary implements the encoding.binary's
// Unmarshaler interface for an AnalogInput.
func (ai *AnalogInput) UnmarshalBinary(data []byte) error {
	if len(data) != 2 {
		return errors.New("analog input must have size 2")
	}

	*ai = AnalogInput(int16(data[1]) + 256*int16(data[0]))
	return nil
}

// An AnalogOutput is a 16-bit integer,
// send from the server to the device.
type AnalogOutput int16

// UnmarshalBinary implements the encoding.binary's
// Unmarshaler interface for an AnalogOutput.
func (ao *AnalogOutput) UnmarshalBinary(data []byte) error {
	if len(data) != 2 {
		return errors.New("analog output must have size 2")
	}

	*ao = AnalogOutput(int16(data[1]) + 256*int16(data[0]))
	return nil
}