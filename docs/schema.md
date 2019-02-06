# Driver Schema

Schemas provide a multitude of benefits:

* Reliability in the interaction between the system and the driver
* Support for additional data types, e.g. dates, times, or
  distinction between floating point and integer numbers.

The following is an example schema:

```schema:
  properties:
    temperatures:
      type: object
      properties:
        feed:
          type: float64
          unit: °C
          description: The heater's feed temperature
        return:
          type: float64
          unit: °C
          description: The heater's return temperature
```

A device schema can contain the following types:

* `object`: An object (aka dict or map) with named key value pairs.
* `int64`: 64-bit integer number
* `float64`: 64-bit floating point number
* `string`: A sequence of Unicode characters
* To be completed

## Objects

Objects can contain sub-properties
