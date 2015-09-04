# gozwave

Go bindings for the [OpenZWave](https://github.com/OpenZWave/open-zwave) library.

## Warning

This is still a fairly incomplete wrapper of the C++ OpenZWave code. It almost completely wraps the Manager class as well as the ValueID and Notification classes to allow for basic use of the library.

Enough is available to implement the MinOZW example application in Go.

More is planned.

## Prerequisites

This package does not include the OpenZWave library, so you will need to install the library yourself. Below are some instructions for various platforms.

### Mac OS X and Linux (Ubuntu)

1. Clone the OpenZWave git repository: `git clone git@github.com:OpenZWave/open-zwave.git`
2. Go into the directory and run Make: `cd open-zwave && make` (get :coffee: at this point)
3. Install the library: `make install` (may need to run as root)

## Installation

```
go get gitlab.com/jimjibone/gozwave
```

_Notice how there was no need to run `make` :wink:_

## Examples/Tools

This package comes with an example, `gominozw`, and a tool, `gozw`.

### `gominozw`

This is a replica of the original MinOZW utility, from the original OpenZWave repository, but now written in Go.

It shows how to set up the Manager with various options and listen for Notifications. Once the initial scan of devices is complete, polling for basic values is set up for the devices.

To install:

```
go install gitlab.com/jimjibone/gozwave/tools/gominozw
```

### `gozw`

This has yet to be written, but it will serve a web app from which you can view devices and their values, as well as modify the state of them as appropriate.
