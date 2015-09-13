# goopenzwave

Go bindings for the [OpenZWave](https://github.com/OpenZWave/open-zwave) library.

## Warning

This package is still fairly new and so the API is changing pretty rapidly, so be casreful if you decide to use it. I expect the API to become stable soon, however, as it is largely based on the OpenZWave library. Things may change to make usage more like idiomatic Go.

Most of the OpenZWave library is wrapped now, but should you find anything missing please create a new issue or fork it, implement it yourself and submit a pull request.

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

_Notice how there was no need to run `make` :wink:_

## Examples/Tools

This package comes with an example, `gominozw`, and a tool, `gozwd`.

### `gominozw`

This is a replica of the original MinOZW utility, from the original OpenZWave repository, but now written in Go.

It shows how to set up the Manager with various options and listen for Notifications. Once the initial scan of devices is complete, polling for basic values is set up for the devices.

To install and use:

```
go install github.com/jimjibone/goopenzwave/tools/gominozw
gominozw --controller /path/to/your/controller
```

### `gozwd`

__This is still in progress.__

It serves a web app from which you can view devices and their values, as well as
modify the state of them as appropriate.

To build, the best thing to do is to get the project, build it and then build the web assets (NodeJS and Gulp are required).

```
go get github.com/jimjibone/goopenzwave
cd $GOPATH/github.com/jimjibone/goopenzwave/tools/gozwd
npm install
gulp # press ctrl-c once it completes
./gozwd --controller /path/to/your/controller
```
