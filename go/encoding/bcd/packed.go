package bcd

import (
	"fmt"
	"math"

	"github.com/smartmakers/drivers/go/encoding/bitfield"
)

// PackedBigEndianBCD is a packed, big endian, binary coded decimal value.
type PackedBigEndianBCD uint64

// UnmarshalBinary implements the binary.Unmarshaler interface.
func (bcd *PackedBigEndianBCD) UnmarshalBinary(array []byte) error {
	var res uint64

	for i := 0; i < len(array); i++ {
		// can ignore errors here, because they depend only on 'from' and 'to'
		// which are hard coded to valid values
		upperDigit, _ := bitfield.ExtractBits(array[i], 5, 8)
		lowerDigit, _ := bitfield.ExtractBits(array[i], 1, 4)
		if lowerDigit > 9 || upperDigit > 9 {
			*bcd = 0
			return fmt.Errorf("Bad digit value for BCD: %x", array[i])
		}

		bytePos := len(array) - i - 1
		multiplier := uint64(math.Pow(100, float64(bytePos)))
		res += (multiplier * uint64(lowerDigit)) + (10 * multiplier * uint64(upperDigit))
	}

	*bcd = PackedBigEndianBCD(res)
	return nil
}

// PackedLittleEndianBCD is a packed, little endian, binary coded decimal value.
type PackedLittleEndianBCD uint64

// UnmarshalBinary implements the binary.Unmarshaler interface.
func (bcd *PackedLittleEndianBCD) UnmarshalBinary(array []byte) error {
	var res uint64

	for i := 0; i < len(array); i++ {
		// can ignore errors here, because they depend only on 'from' and 'to'
		// which are hard coded to valid values
		upperDigit, _ := bitfield.ExtractBits(array[i], 1, 4)
		lowerDigit, _ := bitfield.ExtractBits(array[i], 5, 8)
		if lowerDigit > 9 || upperDigit > 9 {
			return fmt.Errorf("Bad digit value for BCD: %x", array[i])
		}

		bytePos := i
		multiplier := uint64(math.Pow(100, float64(bytePos)))
		res += (multiplier * uint64(lowerDigit)) + (10 * multiplier * uint64(upperDigit))
	}

	*bcd = PackedLittleEndianBCD(res)
	return nil
}
