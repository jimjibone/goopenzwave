package goopenzwave

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/.lib/lib -L${SRCDIR}/.lib/lib -lopenzwave
// #cgo CFLAGS: -I${SRCDIR}/.lib/include -I${SRCDIR}/.lib/include/openzwave
// #cgo CPPFLAGS: -I${SRCDIR}/.lib/include -I${SRCDIR}/.lib/include/openzwave
import "C"
