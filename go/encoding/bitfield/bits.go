package bitfield

import "errors"

// ExtractBits extracts a subset of bits from a bitfield.
func ExtractBits(input byte, from, to int) (byte, error) {
	if from == 0 || from > 8 {
		return 0, errors.New("index out of range")
	}

	if to == 0 || to > 8 {
		return 0, errors.New("index out of range")
	}

	bit := 1 << byte(to)
	mask := byte(bit - 1)
	shifted := (input & mask) >> byte(from-1)
	return shifted, nil
}
