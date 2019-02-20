# Stateful Golang Driver Example

This directory contains an example for a more complex golang driver.
The driver manages state explicitly, allow for keeping data around
across multiple uplinks as well as allow to send multiple values
for the same property in a single uplink.

The kind of device this drivers is designed for is more complex
than suitable for a tutorial like this, so we'll work with
a non-existant device type:

The device is built on top of the Cayenne LPP protocol and
implements downlinks and uplinks for a binary input.
