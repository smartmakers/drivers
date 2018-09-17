package driver_v2

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/smartmakers/drivers/go/device"
	"github.com/smartmakers/drivers/go/types"
)

type EncodeResponse struct {
	NewState      device.State `json:"new_state"`
	PendingUpdate device.State `json:"pending_update"`
	Payload       types.Bytes  `json:"payload"`
	Port          int          `json:"port"`
	Confirmation  bool         `json:"confirmation"`
}

// Encoder is a function which turns the current state and update into a encode response
type Encoder func(currentState, update device.State) (EncodeResponse, error)

// Decoder is a function which turns the current state and binary payload into a new state
type Decoder func(state device.State, payload []byte, fPort int) (device.State, error)

// Driver is the base for
type Driver struct {
	Decoder Decoder
	Encoder Encoder
}

// New creates and returns a new Driver
func New(decoder Decoder, encoder Encoder) *Driver {
	return &Driver{decoder, encoder}
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

func parseDecodeArgs(args []string) (state device.State, payload []byte, port int, err error) {
	if len(args) != 3 {
		err = errors.New(`decode expected 3 arguments: device state as json string, hex payload and integer port
	Usage: <driver binary> decode '{"endpoint_0_value":"true"}' 110A000F00551000 1`)
		return
	}

	if len(args[0]) > 2 {
		err = json.Unmarshal([]byte(args[0]), &state)
		if err != nil {
			return
		}
	}

	payload, err = hex.DecodeString(strings.TrimPrefix(args[1], "0x"))
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

func (d *Driver) Encode(args []string) error {
	currentState, update, err := parseEncodeArgs(args)
	if err != nil {
		return err
	}

	encoded, err := d.Encoder(currentState, update)
	if err != nil {
		return err
	}

	data, err := json.Marshal(encoded)
	if err != nil {
		return errors.New("failed to marshal in json encoded data")
	}

	fmt.Println(string(data))
	return nil
}

func parseEncodeArgs(args []string) (currentState, update device.State, err error) {
	if len(args) != 2 {
		err = errors.New(`encode expected 2 arguments: current device state and update as json strings 
	Usage: <driver binary> encode '{"endpoint_0_value":"true"}' '{"minimum_reporting_interval":60}'`)
		return
	}

	if len(args[0]) > 2 {
		err = json.Unmarshal([]byte(args[0]), &currentState)
		if err != nil {
			return
		}
	}

	if len(args[1]) <= 2 {
		err = errors.New("`update` could not be empty")
		return
	}
	err = json.Unmarshal([]byte(args[1]), &update)
	if err != nil {
		return
	}

	return
}
