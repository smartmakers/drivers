package bcd

import (
	"fmt"
	"math"
)

// NonPackedBigEndianBCD is a non-packed, little-endian, binary coded decimal value.
type NonPackedBigEndianBCD uint64

// UnmarshalBinary implements the binary.Unmarshaler interface.
func (bcd *NonPackedBigEndianBCD) UnmarshalBinary(bytes []byte) error {
	var res uint64

	for i := 0; i < len(bytes); i++ {
		digit := uint64(bytes[i])
		if digit > 9 {
			*bcd = 0
			return fmt.Errorf("Bad digit value for BCD: %d", digit)
		}

		multiplier := uint64(math.Pow(10, float64(len(bytes)-i-1)))
		res += multiplier * digit
	}

	*bcd = NonPackedBigEndianBCD(res)
	return nil
}

// NonPackedLittleEndianBCD is a non-packed, little-endian, binary coded decimal value.
type NonPackedLittleEndianBCD uint64

// UnmarshalBinary implements the binary.Unmarshaler interface.
func (bcd *NonPackedLittleEndianBCD) UnmarshalBinary(bytes []byte) error {
	var res uint64

	for i := 0; i < len(bytes); i++ {
		digit := uint64(bytes[i])
		if digit > 9 {
			*bcd = 0
			return fmt.Errorf("Bad digit value for BCD: %d", digit)
		}

		multiplier := uint64(math.Pow(10, float64(i)))
		res += multiplier * digit
	}

	*bcd = NonPackedLittleEndianBCD(res)
	return nil
}
