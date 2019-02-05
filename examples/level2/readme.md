# Example Driver Capability Level 2

This directory contains an example for a capability level 2 driver.
A level can be written in javascript, but also in golang.
This particular example illustrates how to implement a simple
driver in golang.

Note, that this driver also defines a driver schema in the file `.projects`.
This schema acts as a documentation of the device's (and the drivers)
capabilities, but it is also used as a contract between the driver
and the thingsHub.
