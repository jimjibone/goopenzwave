package goopenzwave

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/.lib -L${SRCDIR}/.lib -lopenzwave
// #cgo CFLAGS: -I${SRCDIR}/.lib/include -I${SRCDIR}/.lib/include/openzwave
// #cgo CPPFLAGS: -I${SRCDIR}/.lib/include -I${SRCDIR}/.lib/include/openzwave
import "C"
