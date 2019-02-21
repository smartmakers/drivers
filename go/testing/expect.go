// Package testing is a  minimal package providing
// assertions for testing payload decoding.
package testing

import (
	"encoding/json"
	"reflect"
	"testing"
)

// ExpectError asserts that no error occurred
func ExpectError(t *testing.T, err error, exp string) {
	t.Helper()

	if err == nil {
		t.Errorf("Expected error '%s' but got none", exp)
	}

	act := err.Error()
	if act != exp {
		t.Errorf("Expected error '%s' but got '%s", act, exp)
	}
}

// ExpectNoError asserts that no error occurred
func ExpectNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("Did not expect error but got  '%v'", err)
	}
}

// ExpectString asserts that the given string was returned
func ExpectString(t *testing.T, exp string, act []byte) {
	t.Helper()

	a := string(act)
	if a != exp {
		t.Errorf("Expected '%s', but got '%s'", exp, a)
	}
}

// ExpectJSON does a structural comparison of two JSON objects.
// This returns true if the JSON objects are structurally
// equal and false otherwise.
// This is preferrable to a string comparsing when
// the JSON object contains maps and it's textual representation
// is therefore not unique.
func ExpectJSON(t *testing.T, exp string, act string) {
	t.Helper()

	var expO interface{}
	var actO interface{}

	var err error
	err = json.Unmarshal([]byte(exp), &expO)
	ExpectNoError(t, err)

	err = json.Unmarshal([]byte(act), &actO)
	ExpectNoError(t, err)

	if !reflect.DeepEqual(expO, actO) {
		t.Errorf("Expected '%s', but got '%s'", exp, act)
	}
}
