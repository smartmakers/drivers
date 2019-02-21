# Stateful Golang Driver Example

This directory contains an example for a more advanced driver written in golang.
The driver manages state explicitly, which allows for keeping data around
across multiple uplinks, but also allow sending of multiple values
for the same property in a single uplink.

The kind of device this drivers is designed for is usually
more complex than what would be suitable for a tutorial like this,
so we'll work with a non-existant device type:

The hypothetical device can actively lock a door with a downlink,
but it can also report when the same door was locked or unlocked manually
by sending an uplink.

The device is built on top of the Cayenne LPP protocol and
the binary input and output types with channel 0 for
uplinks and channel 1 for downlinks.
