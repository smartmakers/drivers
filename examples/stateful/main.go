package main

import (
	"os"

	"github.com/smartmakers/drivers/go/driver/v2"
)

// Driver is the main driver itself.
type Driver struct{}

func main() {
	v2.Run(&Driver{}, os.Args[1:])
}
