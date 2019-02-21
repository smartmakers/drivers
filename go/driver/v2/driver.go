package v2

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/smartmakers/drivers/go/types"
)

// State is a generic type which used in requests and responses to interface with the driver.
//
// The runner guarantees that any function in the Driver interface
// which is called with an argument of this type,
// will receive the same type as the Driver returns in the function NewState().
type State interface{}

// DecodeRequest is are request for the stateful decoding of a binary payload on a specified port.
//
// A DecodeRequest contains the current state (which can be used to store a devices configuration),
// the binary payload to be decoded and the port on which the uplink was sent.
type DecodeRequest struct {
	// ReportedState is the current known state of the device.
	// This is guaranteed to be of the same type as NewState's return value.
	ReportedState State `json:"current_state"`

	// Payload is the uplink's binary payload
	Payload []byte `json:"payload"`

	// Port is the uplink's LoRaWAN port
	Port int `json:"port"`
}

// DecodeResponse is the response from a succesful decoding of a binary payload.
//
// A DecodeResponse contains an updated state and a slice of updates
// which are meant to be sent to the consumer of the device's data.
type DecodeResponse struct {
	// NewState is the state after decoding the uplink.
	// This is guaranteed to be of the same type as NewState's return value.
	NewState State    `json:"new_state"`
	Updates  []Update `json:"updates"`
}

// Update is a single message send out to any receiver who is interested
// in live updates.
type Update struct {
	Timestamp time.Time `json:"timestamp"`
	Values    State     `json:"values"`
}

// EncodeRequest represents the request to encode a downlink from a desired state.
type EncodeRequest struct {
	ReportedState State `json:"current_state"`
	DesiredState  State `json:"desired_state"`
}

// EncodeResponse represents the response for a request to encode a downlink.
type EncodeResponse struct {
	NewState     State       `json:"new_state"`
	Payload      types.Bytes `json:"payload"`
	Port         int         `json:"port"`
	Confirmation bool        `json:"confirmation"`
}

// Driver is the core interface for implementing drivers.
type Driver interface {
	UnmarshalState(bytes []byte, state *State) error
	Decode(DecodeRequest, *DecodeResponse) error
	Encode(EncodeRequest, *EncodeResponse) error
}

// runner is a thin convenience wrapper around the Driver
//
// A runner unmarshals requests from the command line,
// executes the drivers decode or encode function,
// marshals the response and prints it out on stdout.
type runner struct {
	Driver
}

func Run(driver Driver, args []string) (success bool) {
	success = false

	defer func() {
		p := recover()
		if p != nil {
			if err, ok := p.(error); ok {
				fmt.Fprintln(os.Stderr, "Crash:", err.Error())
			} else {
				fmt.Fprintln(os.Stderr, "Crash:", p)
			}
		}
	}()

	// Run the actual driver
	r := runner{driver}
	err := r.runUnsafe(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err.Error())
		return false
	}

	return true
}

func (r *runner) runUnsafe(args []string) error {
	if len(args) < 1 {
		return errors.New(`subcommand required.
	Supported commands: decode`)
	}

	cmd := args[0]
	cmdArgs := args[1:]
	switch cmd {
	case "decode":
		return r.runDecode(cmdArgs)
	case "encode":
		return r.runEncode(cmdArgs)

	default:
		return errors.New("invalid subcommand")
	}
}

func (r *runner) runDecode(args []string) error {
	req, err := r.parseDecodeRequest(args)
	if err != nil {
		return err
	}

	resp := DecodeResponse{Updates: []Update{}}
	err = r.Decode(*req, &resp)
	if err != nil {
		return err
	}

	data, err := json.Marshal(resp)
	if err != nil {
		return errors.New("failed to marshal in json decoded data")
	}

	fmt.Println(string(data))
	return nil
}

func (r *runner) parseDecodeRequest(args []string) (*DecodeRequest, error) {
	if len(args) != 3 {
		return nil, errors.New(`decode expected 3 arguments: device state as json string, hex payload and integer port
	Usage: <driver binary> decode '{"endpoint_0_value":"true"}' 110A000F00551000 1`)
	}

	var currentState State
	err := r.UnmarshalState([]byte(args[0]), &currentState)
	if err != nil {
		return nil, err
	}

	payload, err := hex.DecodeString(strings.TrimPrefix(args[1], "0x"))
	if err != nil {
		return nil, err
	}

	val, err := strconv.ParseInt(args[2], 10, 32)
	if err != nil {
		err = errors.New("port argument is not an integer")
	}

	return &DecodeRequest{currentState, payload, int(val)}, nil
}

func (r *runner) runEncode(args []string) error {
	req, err := r.parseEncodeRequest(args)
	if err != nil {
		return err
	}

	resp := EncodeResponse{}
	err = r.Encode(*req, &resp)
	if err != nil {
		return err
	}

	data, err := json.Marshal(resp)
	if err != nil {
		return errors.New("failed to marshal encode response to json")
	}

	fmt.Println(string(data))
	return nil
}

func (r *runner) parseEncodeRequest(args []string) (*EncodeRequest, error) {
	if len(args) != 2 {
		return nil, errors.New(`encode expected 2 arguments: current device state and update as json strings
	Usage: <driver binary> encode '{"endpoint_0_value":"true"}' '{"endpoint_0":{"reporting_interval":{"minimum":"30s","maximum":"12h","reported_change":0}}}'`)
	}

	if len(args[1]) == 0 || args[1] == "{}" {
		return nil, errors.New("`update` must be non-empty")
	}

	var reportedState State
	err := r.UnmarshalState([]byte(args[0]), &reportedState)
	if err != nil {
		return nil, err
	}

	var desiredState State
	err = r.UnmarshalState([]byte(args[1]), &desiredState)
	if err != nil {
		return nil, err
	}

	return &EncodeRequest{reportedState, desiredState}, nil
}
