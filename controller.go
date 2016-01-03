package goopenzwave

// #include "manager.h"
import "C"

// ResetController performs a hard reset on a PC Z-Wave Controller.
//
// Resets a controller and erases its network configuration settings. The
// controller becomes a primary controller ready to add devices to a new
// network.
func ResetController(homeID uint32) {
	C.manager_resetController(cmanager, C.uint32_t(homeID))
}

// SoftReset performs a soft reset on a PC Z-Wave Controller.
//
// Resets a controller without erasing its network configuration settings.
func SoftReset(homeID uint32) {
	C.manager_softReset(cmanager, C.uint32_t(homeID))
}

// CancelControllerCommand cancels any in-progress command running on a
// controller.
func CancelControllerCommand(homeID uint32) {
	C.manager_cancelControllerCommand(cmanager, C.uint32_t(homeID))
}
