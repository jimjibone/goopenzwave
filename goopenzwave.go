package goopenzwave

// #cgo pkg-config: libopenzwave
// #cgo CFLAGS: -I/usr/local/include
// #cgo CPPFLAGS: -I/usr/local/include
// #include "manager.h"
// #include "notification.h"
// #include <stdlib.h>
import "C"
import (
	"fmt"
)

var (
	started             bool
	notificationHandler NotificationHandler
)

// NotificationHandler defines the format for a function that will handle new
// Notification's as they arrive.
type NotificationHandler func(notification *Notification)

// Start will create a new OpenZWave Manager, starting the library execution.
// The OpenZWave Options must be created and locked before calling this. See the
// Options struct. Also pass a NotificationHandler function to this as
// notifications will also be started.
func Start(handler NotificationHandler) error {
	// Only start once.
	if started {
		return fmt.Errorf("already started")
	}

	// Create the manager.
	err := createManager()
	if err != nil {
		return fmt.Errorf("failed to create manager: %s", err)
	}

	// Start notifications.
	err = startNotifications()
	if err != nil {
		return err
	}

	// Set the notification handler.
	notificationHandler = handler

	return nil
}

// Stop will stop notifications and destroy the manager. Do this just before you
// quit your app. Don't forget to destroy the Options object after calling this.
func Stop() error {
	// Check we have started.
	if started {
		return fmt.Errorf("already started")
	}

	// Stop notifications.
	err := stopNotifications()
	if err != nil {
		return err
	}

	// Destroy the manager.
	destroyManager()

	return nil
}
