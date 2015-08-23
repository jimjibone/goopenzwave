package gozwave

// #cgo LDFLAGS: -lopenzwave -L/usr/local/lib
// #cgo CPPFLAGS: -I/usr/local/include -I/usr/local/include/openzwave
// #include "manager.h"
// #include <stdlib.h>
import "C"
import "unsafe"

type Manager struct {
	manager C.manager_t
}

func CreateManager() *Manager {
	m := &Manager{}
	m.manager = C.manager_create()
	return m
}

func Get() *Manager {
	m := &Manager{}
	m.manager = C.manager_get()
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

func (m *Manager) AddDriver(controllerPath string) bool {
	cControllerPath := C.CString(controllerPath)
	result := C.manager_addDriver(m.manager, cControllerPath)
	C.free(unsafe.Pointer(cControllerPath))
	if result {
		return true
	}
	return false
}

func (m *Manager) RemoveDriver(controllerPath string) bool {
	cControllerPath := C.CString(controllerPath)
	result := C.manager_removeDriver(m.manager, cControllerPath)
	C.free(unsafe.Pointer(cControllerPath))
	if result {
		return true
	}
	return false
}

func (m *Manager) AddWatcher() bool {
	result := C.manager_addWatcher(m.manager)
	if result {
		return true
	}
	return false
}

func (m *Manager) RemoveWatcher() bool {
	result := C.manager_removeWatcher(m.manager)
	if result {
		return true
	}
	return false
}
