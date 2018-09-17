package types

import (
	"encoding/hex"
)

type Bytes []byte

func (b Bytes) MarshalText() ([]byte, error) {
	dst := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(dst, b)
	return dst, nil
}
