package lpp

import (
	"testing"

	. "github.com/smartmakers/drivers/go/testing"
)

// adapter for Decode function and testing package.
func dec(payload []byte, port int) (interface{}, error) {
	return Decode(payload, port)
}

func TestBinaryInputFalse(t *testing.T) {
	TestString(t, "000000", 1, dec, `{"0":false}`)
}

func TestBinaryInputTrue(t *testing.T) {
	TestString(t, "000001", 1, dec, `{"0":true}`)
}

func TestBinaryOutputFalse(t *testing.T) {
	TestString(t, "020100", 1, dec, `{"2":false}`)
}

func TestBinaryOutputTrue(t *testing.T) {
	TestString(t, "020101", 1, dec, `{"2":true}`)
}

func TestAnalogInput(t *testing.T) {
	TestString(t, "05020042", 1, dec, `{"5":66}`)
}

func TestAnalogOutput(t *testing.T) {
	TestString(t, "06030000", 1, dec, `{"6":0}`)
}

func TestAnalogSize(t *testing.T) {
	TestError(t, "070300", 1, dec, "payload size does not match data types")
}
