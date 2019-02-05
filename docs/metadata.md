# Metadata

Driver metadata is information about a driver which is uploaded to the server along with
the driver itself.
Metadata can, for example,
* help the system interface with the driver correctly,
* help the user understand what the driver is about,
* and can be used to search a suitable driver for a given device.

Metada is stored in the file `.projects`, which is created when a project is initialized.
Following is an example for a metadata description:

```metadata:
  name: imaginary-sensor
  author: pupil
  supports:
  - manufacturer: none
    model: none
    firmwareversion: 0.1
  platform: amd64
  os: linux```

* `name`: Name if the driver, usually indicates the supported manufaturer or device model,
  but can be used in any way the driver author considers useful.
* `author`: The person or entity that wrote the driver. This is not necessarily the same
  as the manufacturer of the device.
* `supports`: Indicates which devices and versions thereof the driver supports. See below for more information.

## Supports Declarations

A `supports` declaration indicates which devices a driver supports. Such a declaration can have three fields:
* `manufacturer`: Name of the device's manufacturer (e.g. 'Elsys' or 'NKE')
* `model`: Name of the device's model (e.g. 'ERS Lite' or 'In'O')

Supports information is purely informational for the user and not used for any kind of validation.
This implies that the system will not use his informaton to prevent usage of an unsuitable driver.
