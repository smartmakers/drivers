package hex

import (
	"encoding/hex"
)

// Hex is a byte array represented as a hexadecimal string.
type Hex []byte

// MarshalText marshals the Hex object to an actual hex string
func (h Hex) MarshalText() ([]byte, error) {
	dst := make([]byte, hex.EncodedLen(len(h)))
	hex.Encode(dst, h)
	return dst, nil
}

// UnmarshalText unmarshal a hex object from a byte array
func (h *Hex) UnmarshalText(bytes []byte) error {
	dst := make([]byte, hex.DecodedLen(len(bytes)))
	_, err := hex.Decode(dst, bytes)
	*h = dst
	return err
}
