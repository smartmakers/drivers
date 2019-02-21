package main

import (
	"encoding/json"

	"github.com/smartmakers/drivers/go/driver/v2"
)

type State struct {
	Open *bool `json:"open,omitempty"`
}

// UnmarshalState unmarshals a JSON String into a v2.State.
func (d *Driver) UnmarshalState(bytes []byte, state *v2.State) error {
	if len(bytes) == 0 || string(bytes) == "{}" {
		*state = State{}
		return nil
	}

	var st State
	err := json.Unmarshal(bytes, &st)
	if err != nil {
		return err
	}

	*state = st
	return nil
}
