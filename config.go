package goopenzwave

// #cgo LDFLAGS: -lopenzwave -L/usr/local/lib
// #cgo CPPFLAGS: -I/usr/local/include -I/usr/local/include/openzwave
// #include "manager.h"
import "C"

// WriteConfig saves the Z-Wave network configuration. This is so that the
// entire network does not need to be polled every time the application starts.
func WriteConfig(homeID uint32) {
	C.manager_writeConfig(cmanager, C.uint32_t(homeID))
}
