# Simple Golang Driver With Schema

This directory contains an example for a golang driver that provides a schema.

The drivers is designed for a hypothetical device which builts it's use-case-specific
payload specification on top of the generic Cayenne LPP specification.
The device is a water heater and the use-case-specific payload format
uses channel 0 (analog output) for the feed temperature and
channel 1 (also analog output) for the return temperature.
