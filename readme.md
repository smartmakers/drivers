# SmartMakers Driver Development Kit

The SmartMakers driver SDK contains pre-built executables, documentation, and examples
for developing and deploying your own IoT device driver for the thingsHub IoT middleware.


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
This includes, for example, information about the driver's name, author, as well as the devices supported by this driver.
See [here](docs/metadata.md) for more information about driver metadata.


## Driver Schemas

Driver schemas act as a contract between a driver and the system running the driver.
They provide the thingsHub with an understand on how it should interprete a device's data,
and the different features a device provides, even before receiving actual uplinks
from the device.

More information about driver schemas can be found [here](docs/schema.md).


## Driver Versioning

Drivers are most useful when they are well-versioned.
See the [Versioning Guide](docs/versioning.md) to understand how versioning works for thingsHub drivers.


## Driver Development with Javascript and Golang

Drivers can currently be developed in [Javascript](docs/javascript.md) or [Golang](docs/golang.md).


## Examples

| Demo | Description |
|:------|:----------|
| [`golang`](https://github.com/smartmakers/drivers/tree/master/examples/golang) | Simple golang driver |
| [`javascript`](https://github.com/smartmakers/drivers/tree/master/examples/javascript) | Simple javaScript driver |
| [`stateful`](https://github.com/smartmakers/dirvers/tree/master/examples/stateful) | Stateful example with downlinks |


## Contributing

Contributions are welcome and extremely helpful!
Please ensure that you adhere to a commit style
where logically related changes are in a single commit,
or broken up in a way that eases review if necessary.
Keep commit subject lines informative,
but short, and provide additional detail in the extended message text if needed.
If you can, mention relevant issue numbers in either the subject or the extended message.
