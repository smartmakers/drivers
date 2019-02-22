package main

import (
	"os"

	"github.com/smartmakers/drivers/go/driver/v1"
)

func main() {
	drv := v1.New()
	drv.Decoder = decode
	drv.Run(os.Args[1:])
}
