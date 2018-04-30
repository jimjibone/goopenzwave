# goopenzwave

Go bindings for the [OpenZWave](https://github.com/OpenZWave/open-zwave) library.

## Warning

This package is still fairly new and so the API is changing pretty rapidly, so be careful if you decide to use it. However, the API will try to mimic the C++ OpenZWave library as much as possible, if it doesn't already, so there shouldn't be many breaking changes.

Most of the C++ OpenZWave library is wrapped now, but should you find anything missing please create a new issue or fork it, implement it yourself and submit a pull request.

## Building

This package contains the open-zwave C++ library as a submodule which needs to be built before using goopenzwave.

1. `git submodule update --init`
2. `cd open-zwave`
3. `PREFIX=$(pwd)/../.lib/ make install`

## Installation

```
go get github.com/jimjibone/goopenzwave
```

Note: If you plan on distributing an executable with the goopenzwave package, make sure to copy the open-zwave config directory, found in `.lib/etc/openzwave`.

## Example: `gominozw`

This package comes with a basic example, `gominozw`, which is a replica of the original C++ OpenZWave MinOZW utility, now written in Go.

It shows how to set up the Manager with various options and listen for Notifications. Once the initial scan of devices is complete, polling for basic values is set up for the devices.

To install and use:

```
go install github.com/jimjibone/goopenzwave/gominozw
gominozw --controller /dev/ttyYourUSBDevice
```

## Notes

### Step 3 fails when trying to copy `libopenzwave.pc`

When running step 3 of the above guide you may see an error message at the end of `make`'s execution claiming that it could not copy `libopenzwave.pc`. For example: `cp: cannot create regular file '//usr/local/lib/x86_64-linux-gnu/pkgconfig/libopenzwave.pc': Permission denied`.

This can be safely ignored as this package does not require the pkg-config files.

### open-zwave build fails with `fatal error: libudev.h: No such file or directory` on Debian/Ubuntu

Try installing libudev with apt and build again.

```sh
apt-get install libudev-dev
cd open-zwave && make
```

### Crashes instantly on macOS 10.12

Do you see something like this when trying to run something with the goopenzwave package?

```
$ ./gominozw -h
zsh: killed     ./gominozw -h
```

You should try building with the `-ldflags=-s` option. E.g.: `go build -ldflags=-s`. More info at [golang/go#19734](https://github.com/golang/go/issues/19734).
