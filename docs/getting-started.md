# Quickstart

## 1. Installing the CLI

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

## 3. Build project

``` shell
$ drivers build
$ tree -aL 2
├── .project
├── build
│   ├── main
│   └── package.zip
└── main.go
```

## 4. Authenticate to the Driver Registry

``` shell
$ drivers login THINGS_HUB_INSTANCE_URL USERNAME -p PASSWORD
```

## 5. Upload package

To upload thedriver package to the driver registry, run:

``` shell
$ drivers push -t devel
```

For more on the push command, run `$ drivers push -h`.

