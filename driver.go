package goopenzwave

// #include "gzw_manager.h"
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

// AddDriver creates a new driver for a Z-Wave controller using the path
// specified (e.g. "/dev/ttyUSB0"). It returns an error if the controller
// already exists.
//
// This method creates a Driver object for handling communications with a single
// Z-Wave controller. In the background, the driver first tries to read
// configuration data saved during a previous run. It then queries the
// controller directly for any missing information, and a refresh of the list of
// nodes that it controls. Once this information has been received, a
// DriverReady notification callback is sent, containing the Home ID of the
// controller. This Home ID is required by most of the OpenZWave Manager class
// methods.
func AddDriver(controllerPath string) error {
	cControllerPath := C.CString(controllerPath)
	defer C.free(unsafe.Pointer(cControllerPath))
	ok := bool(C.manager_addDriver(cmanager, cControllerPath))
	if ok == false {
		return fmt.Errorf("controller already exists")
	}
	return nil
}

// RemoveDriver removes the driver for a Z-Wave controller as specified, and
// closes the controller. It returns an error if the controller could not be
// found.
//
// Drivers do not need to be explicitly removed before calling Destroy - this is
// handled automatically.
func RemoveDriver(controllerPath string) error {
	cControllerPath := C.CString(controllerPath)
	defer C.free(unsafe.Pointer(cControllerPath))
	ok := bool(C.manager_removeDriver(cmanager, cControllerPath))
	if ok == false {
		return fmt.Errorf("controller not found")
	}
	return nil
}

// GetControllerNodeID returns the node ID of the Z-Wave controller.
func GetControllerNodeID(homeID uint32) uint8 {
	return uint8(C.manager_getControllerNodeId(cmanager, C.uint32_t(homeID)))
}

// GetSUCNodeID returns the node ID of the Static Update Controller.
func GetSUCNodeID(homeID uint32) uint8 {
	return uint8(C.manager_getSUCNodeId(cmanager, C.uint32_t(homeID)))
}

// IsPrimaryController returns true if the controller is a primary controller.
//
// The primary controller is the main device used to configure and control a
// Z-Wave network. There can only be one primary controller - all other
// controllers are secondary controllers.
func IsPrimaryController(homeID uint32) bool {
	return bool(C.manager_isPrimaryController(cmanager, C.uint32_t(homeID)))
}

// IsStaticUpdateController returns true if the controller is a static update
// controller.
//
// A Static Update Controller (SUC) is a controller that must never be moved in
// normal operation and which can be used by other nodes to receive information
// about network changes.
func IsStaticUpdateController(homeID uint32) bool {
	return bool(C.manager_isStaticUpdateController(cmanager, C.uint32_t(homeID)))
}

// IsBridgeController returns true if the controller is using the bridge
// controller library.
//
// A bridge controller is able to create virtual nodes that
// can be associated with other controllers to enable events to be passed on.
func IsBridgeController(homeID uint32) bool {
	return bool(C.manager_isBridgeController(cmanager, C.uint32_t(homeID)))
}

// GetLibraryVersion returns a string version of the Z-Wave API library used by
// a controller.
func GetLibraryVersion(homeID uint32) string {
	cstr := C.manager_getLibraryVersion(cmanager, C.uint32_t(homeID))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// GetLibraryTypeName returns a string containing the Z-Wave API library type
// used by a controller.
//
// The possible library types are:
// - Static Controller
// - Controller
// - Enhanced Slave
// - Slave
// - Installer
// - Routing Slave
// - Bridge Controller
// - Device Under Test
//
// The controller should never return a slave library type. For a more efficient
// test of whether a controller is a Bridge Controller, use the
// IsBridgeController method.
func GetLibraryTypeName(homeID uint32) string {
	cstr := C.manager_getLibraryTypeName(cmanager, C.uint32_t(homeID))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// GetSendQueueCount returns the count of messages in the outgoing send queue.
func GetSendQueueCount(homeID uint32) int32 {
	return int32(C.manager_getSendQueueCount(cmanager, C.uint32_t(homeID)))
}

// LogDriverStatistics will send the current driver statistics to the log file.
func LogDriverStatistics(homeID uint32) {
	C.manager_logDriverStatistics(cmanager, C.uint32_t(homeID))
}

// GetControllerInterfaceType Obtain controller interface type.
//TODO func GetControllerInterfaceType(homeID uint32) ...

// GetControllerPath returns a string of the controller interface path.
func GetControllerPath(homeID uint32) string {
	cstr := C.manager_getControllerPath(cmanager, C.uint32_t(homeID))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// GetDriverStatistics Retrieve statistics from driver.
//TODO func GetDriverStatistics(...) ...
