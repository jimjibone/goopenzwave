# goopenzwave

Go bindings for the [OpenZWave](https://github.com/OpenZWave/open-zwave) library (version 1.6+).


## Warning

This package is still fairly new and so the API is changing pretty rapidly, so be careful if you decide to use it. However, the API will try to mimic the C++ OpenZWave library as much as possible, if it doesn't already, so there shouldn't be many breaking changes.

Most of the C++ OpenZWave library is wrapped now, but should you find anything missing please create a new issue or fork it, implement it yourself and submit a pull request.


## Installing OpenZWave

This package requires a system installation of [OpenZWave](https://github.com/OpenZWave/open-zwave). pkg-config is then used during the build of this package to get the OpenZWave library and headers.

Note that package managers may install an old version of the library so a manual build/install from source is preferred.

### macOS

1. Install [homebrew](https://brew.sh)
2. Install the OpenZWave library and dependencies: `brew install open-zwave pkg-config` (v1.6.962 http://old.openzwave.com/downloads/)

### Debian (and other Linuxes)

1. Install pkg-config: `sudo apt install pkg-config`
2. Get the source: `wget http://old.openzwave.com/downloads/openzwave-1.6.962.tar.gz && tar xzf openzwave-1.6.962.tar.gz`
3. Build and install it: `cd openzwave-1.6.962 && sudo make install`
4. Run ldconfig to update library links and cache: `sudo ldconfig`
5. See the [open-zwave/INSTALL](https://github.com/OpenZWave/open-zwave/blob/master/INSTALL) file for more information


## Example: `gominozw`

This package comes with a basic example, `gominozw`, which is based on the original C++ OpenZWave MinOZW utility, now written in Go.

It shows how to set up the Manager with various options and listen for Notifications. Once the initial scan of devices is complete the device state information is printed to the console.

To install and use:

1. Build application: `go build -o gominozw ./gominozw`
2. Create config and user directories: `mkdir -p config user`
3. Copy openzwave config files:
    - macOS: `cp -r /usr/local/Cellar/open-zwave/1.6.962/etc/openzwave/ ./config`
    - Debian: `cp -r /usr/local/etc/openzwave/ ./config`
6. Run it: `./gominozw run -c ./config/openzwave -u ./user -p /dev/tty...`
