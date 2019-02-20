package testing

import (
	"encoding/hex"
	"encoding/json"
	"testing"
)

// TestString tests if decoding does not return an error but the expected string.
func TestString(t *testing.T, payload string, port int, decode func(payload []byte, port int) (interface{}, error), exp string) {
	in, err := hex.DecodeString(payload)
	ExpectNoError(t, err)

	up, err := decode(in, port)
	ExpectNoError(t, err)

	out, err := json.Marshal(up)
	ExpectNoError(t, err)

	ExpectString(t, exp, out)
}

// TestError tests if decoding does return the expected error.
func TestError(t *testing.T, payload string, port int, decode func(payload []byte, port int) (interface{}, error), exp string) {
	in, err := hex.DecodeString(payload)
	ExpectNoError(t, err)

	_, err = decode(in, port)
	ExpectError(t, err, exp)
}
