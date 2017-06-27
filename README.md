# goopenzwave

Go bindings for the [OpenZWave](https://github.com/OpenZWave/open-zwave) library.

## Warning

This package is still fairly new and so the API is changing pretty rapidly, so be careful if you decide to use it. However, the API will try to mimic the C++ OpenZWave library as much as possible, if it doesn't already, so there shouldn't be many breaking changes.

Most of the C++ OpenZWave library is wrapped now, but should you find anything missing please create a new issue or fork it, implement it yourself and submit a pull request.

## Prerequisites

This package does not include the OpenZWave library, so you will need to install the library yourself. Below are some instructions for various platforms.

### Mac OS X and Linux (Ubuntu)

1. Clone the OpenZWave git repository: `git clone git@github.com:OpenZWave/open-zwave.git`
2. Go into the directory and run Make: `cd open-zwave && make` (get :coffee: at this point)
3. Install the library: `make install` (may need to run as root)

## Installation

```
go get github.com/jimjibone/goopenzwave
```

_Notice how there was no need to run `make`_ :wink:

## Example: `gominozw`

This package comes with a basic example, `gominozw`, which is a replica of the original C++ OpenZWave MinOZW utility, now written in Go.

It shows how to set up the Manager with various options and listen for Notifications. Once the initial scan of devices is complete, polling for basic values is set up for the devices.

To install and use:

```
go install github.com/jimjibone/goopenzwave/gominozw
gominozw --controller /dev/ttyYourUSBDevice
```
