package device

import "github.com/mitchellh/mapstructure"

// State represents a device state
type State map[string]interface{}

// Decode will decode the values in the state into a structure
func (s State) MapDecode(output interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: mapstructure.StringToTimeDurationHookFunc(),
		Result:     output,
	})
	if err != nil {
		return err
	}

	if err = decoder.Decode(s); err != nil {
		return err
	}

	return nil
}
