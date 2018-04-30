package goopenzwave

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/.lib/lib,-rpath,${SRCDIR}/.lib/lib64 -L${SRCDIR}/.lib/lib -L${SRCDIR}/.lib/lib64 -lopenzwave
// #cgo CFLAGS: -I${SRCDIR}/.lib/include -I${SRCDIR}/.lib/include/openzwave
// #cgo CPPFLAGS: -I${SRCDIR}/.lib/include -I${SRCDIR}/.lib/include/openzwave
import "C"
