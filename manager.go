package gozwave

// #cgo LDFLAGS: -lopenzwave -L/usr/local/lib
// #cgo CPPFLAGS: -I/usr/local/include -I/usr/local/include/openzwave
// #include "manager.h"
import "C"

type Manager struct {
	manager C.manager_t
}

func CreateManager() *Manager {
	m := &Manager{}
	m.manager = C.manager_create()
	return m
}

func DestroyManager() {
	C.manager_destroy()
}

func GetManagerVersionAsString() string {
	return C.GoString(C.manager_getVersionAsString())
}

func GetManagerVersionLongAsString() string {
	return C.GoString(C.manager_getVersionLongAsString())
}
