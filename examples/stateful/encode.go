package main

import (
	"github.com/smartmakers/drivers/go/driver/v2"
)

// Encode encodes a downlink
func (d *Driver) Encode(req v2.EncodeRequest, resp *v2.EncodeResponse) error {
	desired := req.DesiredState.(State)
	if desired.Open != nil {
		if *desired.Open {
			resp.Payload = []byte{01, 01, 01}
			resp.Port = 1
		} else {
			resp.Payload = []byte{01, 01, 00}
			resp.Port = 1
		}
	}

	return nil
}
