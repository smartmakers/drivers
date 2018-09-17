package runner

import (
	"errors"
	"fmt"
	"os"
)

type Driver interface {
	Decode(args []string) error
	Encode(args []string) error
}

type Runner struct {
	Driver
}

// New creates and returns a new Driver
func New(driver Driver) *Runner {
	return &Runner{driver}
}

// Run the driver with the specified arguments
func (d *Runner) Run(args []string) (success bool) {
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

func (d *Runner) run(args []string) error {
	if len(args) < 1 {
		return errors.New(`subcommand required.
	Supported commands: decode`)
	}

	cmd := args[0]
	cmdArgs := args[1:]
	switch cmd {
	case "decode":
		return d.Decode(cmdArgs)
	case "encode":
		return d.Encode(cmdArgs)

	default:
		return errors.New("invalid subcommand")
	}
}
