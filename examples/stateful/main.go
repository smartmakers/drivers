package main

import (
	"os"

	"github.com/smartmakers/drivers/go/driver"
)

func main() {
	drv := driver.New()
	drv.Decoder = decodePayload
	drv.Run(os.Args[1:])
}

func decodePayload(payload []byte, fPort int) (driver.DecodedPayload, error) {
	uplink := Uplink{}
	err := uplink.UnmarshalBinary(payload)
	return uplink, err
}
