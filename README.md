# SmartMakers Drivers Integration

This guide explains how to use the SmartMakers drivers command line interface (CLI) tool for creating and managing driver projects. Some typical IoT use cases are Smart Metering, Smart Energy, Smart Farming, Fleet Tracking, etc.

## Contents

- [Quickstart](#quickstart)
- [Examples](#examples)
- [Contributing](#contributing)

## Quickstart

#### 1. Installing the CLI

In order to get started, you'll want to install SmartMakers drivers command line interface (CLI) globally. We recommend installing the drivers CLI by downloading the [pre-built binaries](https://storage.googleapis.com/sm-tools/drivers). Alternatively, you can download the CLI tool from the terminal:

``` shell
$ wget https://storage.googleapis.com/sm-tools/drivers
```

Once installed, you will be able to open create and manage driver projects in Go and JavaScript.

## How the CLI works

Each time the CLI is run, it looks for the config file, .project, in the root directory of your project. It then applies the configuration from your, and executes any commands you've requested for it to run.

``` shell
$ drivers -h
DESCRIPTION:
  CLI tool for creating and managing driver projects

COMMANDS:
    help, h  Shows a list of commands or help for one command

  Metadata:
    get  gets metadata by type
    set  sets metadata by type
    del  deletes metadata by type

  Project:
    init   command to initialize new driver project in empty working directory
    build  builds driver project

  Registry:
    login   authenticates to driver registry
    push    pushes the driver package to the driver registry
    delete  deletes the driver package with specific tag from the driver registry
    search  search for drivers in the registry

GLOBAL OPTIONS:
 --help, -h     show help
 --version, -v  print the version
```


#### 2. Start a new project

Run the following command to get started:

```
$ drivers init golang
$ tree -a
├── .project
└── main.go
```

##### The Project configuration

```yaml
type: golang
metadata:
  name: ""
  author: ""
  labels: []
  supports: []
  platform: ""
  os: ""
session:
  server: ""
  token: ""
  API: ""
```

You can also generate a JavaScript boilerplate, by specifying 'javascript' instead of golang as in this example.

#### 3. Build project

``` shell
$ drivers build
$ tree -aL 2
├── .project
├── build
│   ├── main
│   └── package.zip
└── main.go
```

#### 4. Authenticate to the Driver Registry

``` shell
$ drivers login THINGS_HUB_INSTANCE_URL USERNAME -p PASSWORD
```

#### 4. Upload package

To upload thedriver package to the driver registry, run:

``` shell
$ drivers push -t devel
```

For more on the push command, run `$ drivers push -h`.

## Tagging and versioning

When pushing a driver it can (must) be tagged, e.g. for a driver named 'bar' by author 'foo', when doing

    drivers push -t test
	
which would result in a driver identified by `foo/bar:test`.
Tags can be alphanumeric strings, e.g. git commit hashes, branch names, or semantic verions.
Multiple tags can be added at the same time:

    drivers push -t test -t temporary


### Semantic Versioning

Semantic versioning (https://semver.org/) is an informal standard on how software version numbers
change across different releases of a software product.

Version numbers follow the semantic versioning consist of a major version, a minor version,
and a patch version, separated by dots: major.minor.patch, e.g. `1.0.2`.
Changes to the major version indicate backwards-incompatible changes,
changes to the minor version number indicate backwards-compatible feature additions,
and change to the patch level indicate backwards-compatible bug fixes.

We strongly recommend use of semantic versioning for drivers.
If you decide to do so, we consider it best to adhere the following rules:

Increment the major version number when you:
* use a different schema version
* rename a field in the schema
* change the type of a field in the schema
* change the physical unit, e.g. from degree Celsius to Fahrenheit

Increment the minor version number when you:
* add new fields to the schema

Increment the patch level when you:
* fix incorrect calculations of data
* fix crashes


### Fixed and Floating Tags

For semantic versioning in driver development,
it is recommended to use the pattern of floating and fixed tags,
as commonly used in the docker community.
For this, the driver should be tagged with all version prefixes of the real version number:

    drivers push -t 1.0.0 -t 1.0 -t 1

This will create multiple tags for the same driver and a device can be configured
to use any of these:

    foo/bar:1
    foo/bar:1.0
    foo/bar:1.0.0

Right after this driver was pushed, there's no behavioral difference in using any of those tags.
However, when a new versions of the driver becomes available at a later point of time,
and the driver author applies the pattern again, the tags will be updated differently.
Imagine the driver requires patching and as a consequence, a new version 1.0.1 is released.
Still following the pattern, the driver's author, will now tag this new version like this:

    drivers push -t 1.0.1 -t 1.0 -t 1

Note that this will overwrite the tags `1` and `1.0`, but create a new tag `1.0.1`
and leave the tag `1.0.0` unchanged.
If a device was set to use version `1.0.0` it will now still use the same driver.
However, if the driver was configured to use the tags `1` or `1.0`,
it will automatically use the newer version.

If a new feature is added to the driver, a new minor version number would be released.
In this case that would be version 1.1.0, so the author would push the driver with

    drivers push -t 1.1.0 -t 1.1 -t 1
	
Notice that now only the tag `1` is overwritten, while tags `1.0` and `1.0.1` are unchanged
and tags `1.1` and `1.1.0` are created.

This means a device set to use tag `1` updates automatically now,
while a device set to use `1.0`, `1.0.0` or `1.0.1` will use the same driver as before.

Note that all of this happens as convention and is not hardcoded.
Any driver developer is free to not use semantic versioning,
so he can just as well follow a device's firmware versions,
even when those do not follow the semantic versioning standard.


## Examples

These sample drivers are configured with details of the [Elsys ERS](https://www.elsys.se/en/ers/) LoRaWAN room sensor for measuring indoor environment.

| Demo | Description |
|:------|:----------|
| [`golang`](https://github.com/smartmakers/drivers/tree/master/examples/elsys-go) | Simple Elsys Go driver |
| [`javascript`](https://github.com/smartmakers/drivers/tree/master/examples/elsys-js) | Simple Elsys JavaScript driver |

## Contributing

Contributions are welcome and extremely helpful! Please ensure that you adhere to a commit style where logically related changes are in a single commit, or broken up in a way that eases review if necessary. Keep commit subject lines informative, but short, and provide additional detail in the extended message text if needed. If you can, mention relevant issue numbers in either the subject or the extended message.
