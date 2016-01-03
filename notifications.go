package goopenzwave

// #include "manager.h"
// #include "notification.h"
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

// startNotifications Calls the OpenZWave AddWatcher function. New notifications
// are received by this package and made available via the Notifications
// channel.
func startNotifications() error {
	ok := C.manager_addWatcher(cmanager, nil)
	if ok {
		return nil
	}
	return fmt.Errorf("failed to add watcher")
}

// stopNotifications Calls the OpenZWave RemoveWatcher function. This stops any
// future notifications being received.
func stopNotifications() error {
	ok := C.manager_removeWatcher(cmanager, nil)
	if ok {
		return nil
	}
	return fmt.Errorf("failed to remove watcher")
}

// goNotificationCB called by the C++ OpenZWave library when there is a new
// notification.
//export goNotificationCB
func goNotificationCB(cnotification C.notification_t, userdata unsafe.Pointer) {
	// This function is called by OpenZWave (via a C wrapper) when a
	// notification is available. All data must be extracted from the
	// notification object before we return as OpenZWave will delete the object.

	// Convert the C notification_t to Go Notification.
	notification := buildNotification(cnotification)

	// Allow the assigned handler to deal with it.
	notificationHandler(notification)
}
