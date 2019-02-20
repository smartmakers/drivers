package main

import (
	"testing"

	. "github.com/smartmakers/drivers/go/testing"
)

func dec(payload []byte, port int) (interface{}, error) {
	return decode(payload, port)
}

func TestDecodeFeedTemperature(t *testing.T) {
	TestString(t, "00020042", 1, dec, `{"temperatures":{"feed":66}}`)
}

func TestDecodeReturnTemperature(t *testing.T) {
	TestString(t, "01020042", 1, dec, `{"temperatures":{"return":66}}`)
}
