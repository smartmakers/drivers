# Driver Schema

Schemas provide a multitude of benefits:

* Reliability in the interaction between the system and the driver.
* Documentation of a device's capabilities.
* Support for additional data types, e.g. dates, times, or
  distinction between floating point and integer numbers.

The following is an example of such a schema:

```
schema:
  version: "2"
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

A `schema` has a version and the default version is **1**, which doesn't support much functionality apart from allowing you to define the schema however you want.
From version **2** however we support schema validation and we also enforce the driver to decode the data according to the schema. In order for this functionality to apply the schema version must be specified otherwise it will fallback to default version(*1*) and the functionality is ignored.


## Version 2 Schema

### Properties and Property Types

A schema is a tree of properties.
Each property has a `type` and optionally a `description`.
Leaf properties can also have an optional `unit`.

Schemas are commonly written as YAML or JSON code.
In this guide, we will stick to YAML for readability,
e.g. the following is a simple floating-point property
which represents a temperature:

```yaml
temperature:
  type: float64
  description: The temperature in degree celsius. Can be between -51,2°C and +51,1°C.
  unit: °C
```

Valid `type`s are:

* `object`: An object (aka dict or map) with named key value pairs.
* `boolean`: Either true or false.
* `int64`: 64-bit integer number
* `float64`: 64-bit floating point number
* `string`: A sequence of Unicode characters

The `description` should provide information about what the property represents.
The `unit` contains the physical unit, if applicable, of the measured phenomenon.

The types `boolean`, `int64`, `float64`, `string` work as expected.

An `object` is different in that it is used for grouping sub-properties:

```yaml
room:
  properties:
    temperature:
      type: float64
    humidity:
      type: int64
```

### Schemas and Data

Schemas are used to validate the data returned by a driver.
A driver may send the data listed in the schema,
though all fields are considered optional.

Let's assume the following schema:

```yaml
room:
  properties:
    temperature:
      type: float64
    humidity:
      type: int64
```

This will validate the following structure:

```json
{
  "room": {
    "temperature": 23.5,
    "humidity": 70
  }
}
```

Also the following is perfectly valid, even though it might seem incomplete:

```json
{
  "room": {
    "temperature": 23.5
  }
}
```

Finally, an empty object is always valid:

```json
{}
```

But any property that is not listed in the schema is not valid:

```json
{
  "engine": 20
}
```

And also an incorrect type is not valid with respect to the schema:

```json
{
  "room": {
    "temperature": "I'm a string!"
  }
}
```

## Version 2.1 Schema

Backwards compatible with 2. Added functionality for `configurable` properties.
Only properties marked as `configurable` are allow to be set in the `desired state` of a device.

Eg.:

```
schema:
  version: "2.1"
  properties:
    humidity:
      type: int64
      unit: '%'
      description: 0-100%
      properties: {}
      configurable: false
      default: null
    light:
      type: int64
      unit: Lux
      description: 0-65535 Lux
      properties: {}
      configurable: false
      default: null
    settings:
      type: object
      unit: ""
      description: Device settings.
      properties:
        lightper:
          type: int64
          unit: ""
          description: |-
            Interval in seconds for the sensor(Light) to wake up and sample data.
            Value * Timebase(SplPer)= Light sample time.
          properties: {}
          configurable: true
          default: 120

```

  **Note:** Objects cannot be set as `configurable` only simple properties. An `object` however can contain `configurable` properties.
