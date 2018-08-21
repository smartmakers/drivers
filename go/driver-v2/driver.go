package driver_v2

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// State represents a device state
type State map[string]interface{}

// Decoder is a function which turns the current state and binary payload into a new state
type Decoder func(state State, payload []byte, fPort int) (State, error)

// Driver is the base for
type Driver struct {
	Decoder Decoder
}

// New creates and returns a new Driver
func New(decoder Decoder) *Driver {
	return &Driver{decoder}
}

func (d *Driver) Decode(args []string) error {
	state, payload, port, err := parseDecodeArgs(args)
	if err != nil {
		return err
	}

	decoded, err := d.Decoder(state, payload, port)
	if err != nil {
		return err
	}

	data, err := json.Marshal(decoded)
	if err != nil {
		return errors.New("failed to marshal in json decoded data")
	}

	fmt.Println(string(data))
	return nil
}

func parseDecodeArgs(args []string) (state State, payload []byte, port int, err error) {
	if len(args) != 3 {
		err = errors.New(`decode expected 3 arguments: device state as json string, hex payload and integer port
	Usage: <driver binary> decode '{"endpoint.0.value":"true"}' 110A000F00551001 1`)
		return
	}

	if len(args[0]) > 2 {
		err = json.Unmarshal([]byte(args[0]), &state)
		if err != nil {
			return
		}
	}

	payload, err = hex.DecodeString(strings.TrimLeft(args[1], "0x"))
	if err != nil {
		err = errors.New("payload argument is not in hex")
		return
	}

	val, err := strconv.ParseInt(args[2], 10, 32)
	if err != nil {
		err = errors.New("port argument is not an integer")
	}
	port = int(val)
	return
}
