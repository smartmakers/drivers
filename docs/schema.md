# Driver Schema

Schemas provide a multitude of benefits:

* Reliability in the interaction between the system and the driver.
* Documentation of a device's capabilities.
* Support for additional data types, e.g. dates, times, or
  distinction between floating point and integer numbers.

The following is an example of such a schema:

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

## Properties and Property Types

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

## Schemas and Data

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
