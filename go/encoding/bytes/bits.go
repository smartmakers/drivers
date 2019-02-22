package bytes

import "errors"

// Need to move this somewhere else
func Bits(input, from, to byte) (byte, error) {
	if from == 0 || from > 8 {
		return 0, errors.New("index out of range")
	}

	if to == 0 || to > 8 {
		return 0, errors.New("index out of range")
	}

	bit := 1 << (to)
	mask := byte(bit - 1)
	shifted := (input & mask) >> (from - 1)
	return shifted, nil
}
