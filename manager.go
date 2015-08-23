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

// Notification and callbacks from C:
// http://stackoverflow.com/questions/6125683/call-go-functions-from-c

// This defines the signature of our user's progress handler.
type NotificationHandler func( /*notification data, */ userdata interface{})

// This is an internal type which will pack the users callback function and
// userdata. It is an instance of this type that we will actually be sending to
// the C code.
type notificationContainer struct {
	f NotificationHandler // The user's function pointer.
	d interface{}         // The user's userdata.
}

//export goNotificationCB
func goNotificationCB( /*notification data, */ userdata unsafe.Pointer) {
	// This is the function called from the C world by the OpenZWave
	// notification system. The userdata value contains an instance of
	// *notificationContainer, We unpack it and use it's values to call the
	// actual function that our user supplied.
	watcher := (*notificationContainer)(userdata)

	// Call watcher.f with our parameters and the user's own userdata value.
	watcher.f( /*notification data, */ watcher.d)
}

func (m *Manager) AddWatcher(nh NotificationHandler, userdata interface{}) (unsafe.Pointer, bool) {
	watcher := unsafe.Pointer(&notificationContainer{nh, userdata})
	result := C.manager_addWatcher(m.manager, watcher)
	if result {
		return watcher, true
	}
	return nil, false
}

func (m *Manager) RemoveWatcher(watcher unsafe.Pointer) bool {
	result := C.manager_removeWatcher(m.manager, watcher)
	if result {
		return true
	}
	return false
}
