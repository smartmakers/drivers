package driver

import (
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strconv"
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
func New() *Driver {
	return &Driver{}
}

// Run the driver with the specified arguments
func (d *Driver) Run(args []string) (success bool) {
	// Return false on panics
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

	// Run it
	err := d.run(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err.Error())
		return false
	}

	return true
}

func (d *Driver) run(args []string) error {
	if len(args) < 1 {
		return errors.New("subcommand required")
	}

	cmd := args[0]
	switch cmd {
	case "decode":
		// TODO: decode the payload
		bytes, err := hex.DecodeString(args[1])
		if err != nil {
			return errors.New("payload argument is not in hex")
		}

		port, err := strconv.ParseInt(args[2], 10, 32)
		if err != nil {
			return errors.New("port argument is not a string")
		}

		_, err = d.Decoder(bytes, int(port))
		if err != nil {
			return err
		}

		// TODO: print out the decoded payload to stdout

		return nil
	default:
		return errors.New("invalid subcommand")
	}
}
