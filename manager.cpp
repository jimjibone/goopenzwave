#include "manager.h"
#include <openzwave/Manager.h>
#include <openzwave/Notification.h>

#include <iostream>

//
// Construction.
//

manager_t manager_create()
{
	OpenZWave::Manager *man = OpenZWave::Manager::Create();
	return (manager_t)man;
}

manager_t manager_get()
{
	OpenZWave::Manager *man = OpenZWave::Manager::Get();
	return (manager_t)man;
}

void manager_destroy()
{
	OpenZWave::Manager::Destroy();
}

const char* manager_getVersionAsString()
{
	return OpenZWave::Manager::getVersionAsString().c_str();
}

const char* manager_getVersionLongAsString()
{
	return OpenZWave::Manager::getVersionLongAsString().c_str();
}

// static ozwversion getVersion();

//
// Configuration.
//

//
// Drivers.
//
bool manager_addDriver(manager_t m, const char *controllerPath)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	if (strcasecmp(controllerPath, "usb") == 0) {
		return man->AddDriver("HID Controller", OpenZWave::Driver::ControllerInterface_Hid);
	} else {
		return man->AddDriver(controllerPath);
	}
}

bool manager_removeDriver(manager_t m, const char *controllerPath)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	if (strcasecmp(controllerPath, "usb") == 0) {
		return man->RemoveDriver("HID Controller");
	} else {
		return man->RemoveDriver(controllerPath);
	}
}

//
// Polling Z-Wave devices.
//

//
// Node information.
//

//
// Values.
//

//
// Climate control schedules.
//

//
// Switch all.
//

//
// Configuration parameters.
//

//
// Groups.
//

//
// Notifications.
//

static void manager_notificationHandler(OpenZWave::Notification const* notification, void* userdata)
{
	std::cout << "manager_notificationHandler was called" << std::endl;
	// Now we need to convert the OpenZWave::Notification data into a Go
	// friendly type.
	notification_t noti = (notification_t)notification;
	goNotificationCB(noti, userdata);
}

bool manager_addWatcher(manager_t m, void *userdata)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->AddWatcher(manager_notificationHandler, userdata);
}

bool manager_removeWatcher(manager_t m, void *userdata)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->RemoveWatcher(manager_notificationHandler, userdata);
}

//
// Controller commands.
//

//
// Network commands.
//

//
// Scene commands.
//

//
// Statistics retreival interface.
//
