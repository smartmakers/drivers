package main

import (
	"os"

	"github.com/smartmakers/drivers/go/driver"
)

func main() {
	drv := driver.New()
	drv.Decoder = decode
	drv.Run(os.Args[1:])
}
