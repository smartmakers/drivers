# Cayenne Low Power Payload (LPP)

This is a golang implementation of the Cayenne Low Power Payload specification.

Note that this can be used as a generic driver for LPP devices,
but also as a package for use in another device-specific, LPP-based driver written in golang.

This implementation currently supports decoding for the following data types:

* Digital input
* Digital output
* Analog input
* Analog output
