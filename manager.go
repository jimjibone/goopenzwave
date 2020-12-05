package goopenzwave

// #cgo pkg-config: libopenzwave
// #include "manager_wrap.h"
// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

var (
	watchers []func(notification Notification)
)

// ManagerCreate creates the Manager singleton object. The Manager provides the public interface to OpenZWave, exposing all the functionality required to add Z-Wave support to an application. There can be only one Manager in an OpenZWave application. An Options object must be created and Locked first, otherwise the call to Manager::Create will fail. Once the Manager has been created, call AddWatcher to install a notification callback handler, and then call the AddDriver method for each attached PC Z-Wave controller in turn.
// Wraps `Manager* OpenZWave::Manager::Create(...)`.
func ManagerCreate() {
	C.manager_create()
}

func ManagerAddWatcher(watcher func(notification Notification)) {
	// Manager::Get()->AddWatcher( OnNotification, NULL );
	watchers = append(watchers, watcher)
}

func ManagerAddDriver(port string) {
	// Manager::Get()->AddDriver( port );
	cport := C.CString(port)
	defer C.free(unsafe.Pointer(cport))
	_ = C.manager_add_driver(cport)
}

func ManagerRemoveDriver(port string) {
	// Manager::Get()->RemoveDriver( port );
	cport := C.CString(port)
	defer C.free(unsafe.Pointer(cport))
	_ = C.manager_remove_driver(cport)
}

func ManagerDestroy() {
	// Manager::Get()->RemoveWatcher( OnNotification, NULL );
	// Manager::Destroy();
	C.manager_destroy()
}

func ManagerVersionString() string {
	cstr := C.manager_get_version_as_string()
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

func ManagerVersionLongString() string {
	cstr := C.manager_get_version_long_as_string()
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

func ManagerVersion() (major, minor uint16) {
	version := C.manager_get_version()
	return uint16((version >> 16) & 0x00ff), uint16(version & 0x00ff)
}

// goNotificationWatcher called by the C++ OpenZWave library when there is a new notification.
//export goNotificationWatcher
func goNotificationWatcher(n_type uint8, n_vhomeid, n_vid0, n_vid1 uint32, n_byte, n_event, n_command, n_useralerttype uint8, n_comport *C.char, n_comport_size C.int) {
	// This function is called by OpenZWave (via a C wrapper) when a
	// notification is available.
	comport := C.GoStringN(n_comport, n_comport_size)
	notification := createNotification(n_type, n_vhomeid, n_vid0, n_vid1, n_byte, n_event, n_command, n_useralerttype, comport)
	for _, watcher := range watchers {
		watcher(notification)
	}
}
