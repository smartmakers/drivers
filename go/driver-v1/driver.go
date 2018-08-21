package driver_v1

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// DecodedPayload represents a decoded payload
type DecodedPayload interface{}

// Decoder is a function which turns a binary payload into a decoded payload
type Decoder func(payload []byte, fPort int) (DecodedPayload, error)

// Driver is the base for
type Driver struct {
	Decoder Decoder
}

// New creates and returns a new Driver
func New(decoder Decoder) *Driver {
	return &Driver{decoder}
}

func (d *Driver) Decode(args []string) error {
	payload, port, err := parseDecodeArgs(args)
	if err != nil {
		return err
	}

	decoded, err := d.Decoder(payload, port)
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

func parseDecodeArgs(args []string) (payload []byte, port int, err error) {
	if len(args) != 2 {
		err = errors.New(`decode expected 2 arguments: hex payload and integer port
	Usage: <driver binary> decode 110A000F00551001 1`)
		return
	}

	payload, err = hex.DecodeString(strings.TrimLeft(args[0], "0x"))
	if err != nil {
		err = errors.New("payload argument is not in hex")
		return
	}

	val, err := strconv.ParseInt(args[1], 10, 32)
	if err != nil {
		err = errors.New("port argument is not an integer")
	}
	port = int(val)
	return
}
