# SmartMakers Driver Development Kit

The SmartMakers driver SDK contains pre-built executables, documentation, and examples
for developing and deploying your own IoT device driver for the thingsHub IoT middleware.


## Contents

- [Quickstart](#quickstart)
- [Examples](#examples)
- [Contributing](#contributing)


## Quickstart Guide

Check out the [Quickstart Guide](docs/quickstart.md)
for instructions on how to get started quickly
with developing a basic driver in javascript.

After succesfully finishing the Quickstart Guide,
the next step is to learn more about the concepts
behind device drivers.

The two most important concepts are Driver Capability Levels
and Driver Versioning.
Ẃith this added knowledge,
check out the section about Driver Development below
for more in-depth information on driver development.

## Driver Metadata

Driver metadata is information about a driver which is not directly part of the driver's executable code itself.
This includes, for example, information about the driver'S name, author, as well as the devices supported by this driver.
See [here](docs/metadata.md) for more information about driver metadata.

## Driver Capability Levels

IoT devices vary greatly in their capabilities depending on the use case and the required functionality.
Some devices simply send a set of measurements, e.g. temperature and humidity, on a regular schedule.
More advanced devices bridge other protocols, e.g. ModBus or MBus, and need to split the other
protocol's payload into multiple LoRa messages.
Another commonly supported feature is configuration by downlinks,
or sending batch messages, where the  devices collects data over time and sends multiple
measurements of the same physical phenomenon at different times,
together in a single uplink.

To make full use of all of the features in any of these devices,
considerably effort needs to be spent on the server side.
Anything that the server side does not support is less accessible to the user,
or possible even entirely inaccessible.
This inherently results in a certain amount of complexity in the code which
supports the features on the server side.

However, for drivers for simpler devices, this added complexity might not even be desirable,
considering the device itself couldn't make use of any of the advanced features.

The thingsHub uses the concept of Driver Capability Levels to allow for such simple and complex
drivers to coexist in the same system and to be behave sensibly and predictably.
Driver Capability Levels are described in-depth [here](docs/capabilities.md).

Note that while Driver Capability levels provide a rough understanding of a driver's capabilities,
they should not be seen as a measurement of a driver's maturity or the completeness of a
driver's support for the device's features.
The Driver Capability Level should rather be seen as a contract
between the driver and the system that defines how data processed by the driver
should be handled by the system and vice versa.


## Driver Schemas

Driver schemas act as a contract between a driver and the system running the driver.
They provide the thingsHub with an understand on how it should interprete a device's data,
and the different features a device provides, even before receiving actual uplinks
from the device.

More information about driver schemas can be found [here](docs/schemas.md).


## Driver Versioning

Drivers are most useful when they are well-versioned.
See the [Versioning Guide](docs/versioning.md) to understand how versioning works for thingsHub drivers.


## Driver Development with Javascript and Golang

Drivers can currently be developed in [Javascript](docs/javascript.md) or [Golang](docs/golang.md).
Note that javascript drivers currently are only supported for capability levels 1 and 2,
while Golang drivers can be written for levels 1, 2, and 4.


## Examples

These sample drivers are configured with details of the [Elsys ERS](https://www.elsys.se/en/ers/)
LoRaWAN room sensor for measuring indoor environment.

| Demo | Description |
|:------|:----------|
| [`golang`](https://github.com/smartmakers/drivers/tree/master/examples/elsys-go) | Simple Elsys Go driver |
| [`javascript`](https://github.com/smartmakers/drivers/tree/master/examples/elsys-js) | Simple Elsys JavaScript driver |


## Contributing

Contributions are welcome and extremely helpful!
Please ensure that you adhere to a commit style
where logically related changes are in a single commit,
or broken up in a way that eases review if necessary.
Keep commit subject lines informative,
but short, and provide additional detail in the extended message text if needed.
If you can, mention relevant issue numbers in either the subject or the extended message.
