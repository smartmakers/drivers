# Encoding

The encoding package provides encoding for
frequently occuring data types in low-level payload specifications.

For example, the binary coded decimal encoding (BCD) is
frequently used to encode numerical values in binary payloads.

## Package `bcd`

The `bcd` package provides a type for encoding BCD values.
This conveniently marshals and unmarshal BCD values.

## Package `hex`

The `hex` package provides convenient encoding and decoding
for hexadecimal values.

## Package `bitfield`

Provides helper functions for extracting bits from a slice of bytes.
