// Package testing is a  minimal package providing
// assertions for testing payload decoding.
package testing

import "testing"

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
