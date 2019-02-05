package main

import (
	"encoding/json"
	"testing"
)

func TestBinaryInputFalse(t *testing.T) {
	in := []byte{0x00, 0x00, 0x00}
	up, err := decodePayload(in, 1)
	ExpectNoError(t, err)

	out, err := json.Marshal(up)
	ExpectNoError(t, err)

	ExpectString(t, `{"0":false}`, out)
}

func TestBinaryInputTrue(t *testing.T) {
	in := []byte{0x00, 0x00, 0x01}
	res, err := decodePayload(in, 1)
	ExpectNoError(t, err)

	out, err := json.Marshal(res)
	ExpectNoError(t, err)

	ExpectString(t, `{"0":true}`, out)
}

func TestBinaryOutputFalse(t *testing.T) {
	in := []byte{0x02, 0x01, 0x00}
	up, err := decodePayload(in, 1)
	ExpectNoError(t, err)

	out, err := json.Marshal(up)
	ExpectNoError(t, err)

	ExpectString(t, `{"2":false}`, out)
}

func TestBinaryOutputTrue(t *testing.T) {
	in := []byte{0x02, 0x01, 0x01}
	res, err := decodePayload(in, 1)
	ExpectNoError(t, err)

	out, err := json.Marshal(res)
	ExpectNoError(t, err)

	ExpectString(t, `{"2":true}`, out)
}

func ExpectNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func ExpectString(t *testing.T, exp string, act []byte) {
	t.Helper()
	a := string(act)
	if a != exp {
		t.Errorf("Expected '%s', but got '%s'", exp, a)
	}
}
