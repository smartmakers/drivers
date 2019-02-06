# Quickstart

## Installing the Driver Development Kit

In order to get started, download the contents of this repository as a zip file
or, preferrably, use `git clone` to clone the repository to your local disk.

In this repository's directory `bin/` are subdirectory for each supported platform,
which in turn contain pre-built executables for the thingsHub's drivers command line tool.

For simplicity, we assume that the right directory was added to the PATH variable.

## Running the Command Line Tool

Running the command line tool with the option `-h` gives the following output:

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

## Starting a New Driver Development Project

Run the following command to get started:

```
$ drivers init javascript
$ tree -a
├── .project
└── script.js
```

## Project Configuration

For project configuration use the following commands:

```shell
$ drivers set author <your name>
$ drivers set name <driver name>
$ drivers set supports <manufacturer> <model> <version>
```

Or alternatively edit the `.project` file accordingly:

```yaml
type: golang
metadata:
  name: "<drivern ame>"
  author: "<your name>"
  labels:
  supports:
  - manufacturer: <manufacturer>
    model: <model>
    firmware_version: <version>
  platform: "amd64"
  os: "linux"
```

## Building the Driver Package

Running the `build` subcommand builds and packages the driver.
```shell
$ drivers build
```

## Logging in to the Driver Registry

Before being able to push anything to a driver registry,
it is necessary to login to the target registry first:

```shell
$ drivers login <server> <username> -p <password>
```

Note, that this will store an access token in the `.project` file,
which will be used in subsequent calls to the registry.

## Uploading the Driver Package

To upload the driver package to the driver registry, run:

```shell
$ drivers push -t devel
```

For more on the push command, run `$ drivers push -h`.
