# Driver Capability Levels

A driver's Capability Level describes a driver's capabilities
with respect to a set of predefined levels.
The Capability Level helps the thingsHub understand how the driver works,
and how it's data should be treated.

The thingsHub uses a driver's capability level to understand if it
needs to persist a driver's state,
if it should validate the driver's data according to a schema,
if it can ask the driver to encode uplinks,
or if it can send data from a device to specific target applications.

While a higher Capability Level indicates a more capable driver,
it is frequently not required for all devices.
In particular levels above level 3 are meant to be implemented
by drivers for very powerful and complex devices.

| Level | Description |
|:------|:----------|
| 1 | Plain uplink decoding |
| 2 | Uplink decoding with a data schema |
| 3 | Downlink encoding |
| 4 | Statefulnes management |
| 5 | Declarative management |

## Level 1 - Plain Uplink Decoding

Capability Level 1 is designed
with existing javascript drivers in mind.
This level makes it possible to run off-the-web drivers with
minimal changes to the code. It frequently suffices
to a `decode` function as a well-defined entry point
to the code and the driver is ready to be used.

Because of the lack of schema, the driver's output
cannot be validated as easily.
This makes it impossible to send data for such a device
to specific target applications, if those require
predefined data schemata or device and sensor models.


## Level 2 - Uplink decoding with a data schema

Level 2 drivers are still designed for low effort,
though they require adding a data schema in addition
to the actual code of the driver.

However, this data schema unlocks additional functionality
for the driver in the thingsHub,
which is not accessible for level 1 devices.

For example, a level 2 device can encode additional types
and these types can be automatically stored in the different
target systems in the most suitable format.


## Level 3 - Downlink encoding

Level 3 are similar to level 2 devices,
but in addition they support sending downlinks to the device.

Level 3 drivers are still up-/downlink oriented.


## Level 4 - Stateful Management

Level 4 drivers can store bits of state about a drivers


## Level 5 - Declarative management
