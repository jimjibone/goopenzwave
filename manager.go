package goopenzwave

// #include "manager.h"
// #include "notification.h"
// #include <stdlib.h>
import "C"
import (
	"fmt"
)

var (
	cmanager C.manager_t
)

// createManager creates the Manager singleton object. The Manager provides the
// public interface to OpenZWave, exposing all the functionality required to add
// Z-Wave support to an application. There can be only one Manager in an
// OpenZWave application. An Options object must be created and Locked first,
// otherwise the call to Manager::Create will fail. Once the Manager has been
// created, call AddWatcher to install a notification callback handler, and then
// call the AddDriver method for each attached PC Z-Wave controller in turn.
func createManager() error {
	cmanager = C.manager_create()
	if cmanager == nil {
		return fmt.Errorf("libopenzwave returned NULL pointer")
	}
	return nil
}

// getManager gets a pointer to the Manager object.
func getManager() C.manager_t {
	return cmanager
}

// destroyManager deletes the Manager and cleans up any associated objects.
func destroyManager() {
	C.manager_destroy()
}

// GetManagerVersionAsString Get the Version Number of OZW as a string.
func GetManagerVersionAsString() string {
	cString := C.manager_getVersionAsString()
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetManagerVersionLongAsString Get the Version Number including Git commit of OZW as a string.
func GetManagerVersionLongAsString() string {
	cString := C.manager_getVersionLongAsString()
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// ManagerVersion represents the OpenZWave library version as major and minor
// integers.
type ManagerVersion struct {
	Major int
	Minor int
}

// GetManagerVersion Get the Version Number as the Version Struct (Only Major/Minor returned).
func GetManagerVersion() ManagerVersion {
	var cMajor C.uint16_t
	var cMinor C.uint16_t
	C.manager_getVersion(&cMajor, &cMinor)
	return ManagerVersion{
		Major: int(cMajor),
		Minor: int(cMinor),
	}
}
