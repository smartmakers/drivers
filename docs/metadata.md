# Metadata

Driver metadata is information about a driver which is uploaded to the server along with
the driver itself.
Metadata, for example,

* helps the system interface with the driver correctly,
* helps the user understand what the driver is about,
* and can be used to identify a suitable driver for a given device in the registry.

Metada is stored in the file `.projects`, which is created when a project is initialized.
Following is an example for a metadata description:

```metadata:
  name: house-sensor
  author: john
  supports:
  - manufacturer: jack
    model: house
    firmwareversion: 0.1
  platform: amd64
  os: linux
```

* `name`: Name if the driver, usually indicates the supported manufaturer or device model,
  but can be used in any way the driver author considers useful.
* `author`: The person or entity that wrote the driver. This is not necessarily the same
  as the manufacturer of the device.
* `platform`: The target platform for which a driver was built.
  Currently, the only option here is `amd64`.
* `os`: The target operating system for which a driver was built.
  Currently, the only two options here are `linux` and `amd64`.
* `supports`: Indicates which devices and versions thereof the driver supports. See below for more information.

## Supports Declarations

A driver's author of a driver is not nececssarily the same entity as the manufacturer of a device.
For this reason, a driver's metadata allows explicitly listing the supported devices and their manufacturers.

A `supports` declaration indicates which devices a driver supports. Such a declaration can have three fields:
* `manufacturer`: Name of the device's manufacturer (e.g. 'Elsys' or 'NKE')
* `model`: Name of the device's model (e.g. 'ERS Lite' or 'In'O')
* `firmware_version`: The supported firmware version. As formats for firmware version vary widely,
  this can be any string.

Supports information is purely informational for the user and not used for any kind of validation.
This implies that the system will not use his informaton to prevent usage of an unsuitable driver.
