#include "manager_wrap.h"
#include "openzwave/Manager.h"
#include "openzwave/Notification.h"
#include "openzwave/value_classes/ValueID.h"
#include "_cgo_export.h" // defines goNotificationWatcher
#ifndef GO_COMPLIER
    // void goNotificationWatcher(uint8_t, uint32_t, uint32_t, uint32_t, uint8_t, uint8_t, uint8_t, uint8_t, char*, long);
    // goNotificationWatcher(unsigned char, unsigned int, unsigned int, unsigned int, unsigned char, unsigned char, unsigned char, unsigned char, char*, unsigned long)
#endif

static void manager_notificationWatcher(OpenZWave::Notification const* notification, void* userdata);

void manager_create()
{
    OpenZWave::Manager::Create();
    OpenZWave::Manager::Get()->AddWatcher(manager_notificationWatcher, nullptr);
}

bool manager_add_driver(const char* port)
{
    return OpenZWave::Manager::Get()->AddDriver(port);
}

bool manager_remove_driver(const char* port)
{
    return OpenZWave::Manager::Get()->RemoveDriver(port);
}

void manager_destroy()
{
    OpenZWave::Manager::Get()->RemoveWatcher(manager_notificationWatcher, nullptr);
    OpenZWave::Manager::Destroy();
}

// char* manager_strdup(const char *src)
// {
//     char *dst = static_cast<char*>(malloc(strlen(src) + 1)); // Space for length plus nul
//     if (dst == nullptr) return nullptr;    // No memory
//     strcpy(dst, src);                      // Copy the characters
//     return dst;                            // Return the new string
// }

char* manager_get_version_as_string()
{
    std::string str = OpenZWave::Manager::Get()->getVersionAsString();
    return strdup(str.c_str());
}

char* manager_get_version_long_as_string()
{
    std::string str = OpenZWave::Manager::Get()->getVersionLongAsString();
    return strdup(str.c_str());
}

uint32_t manager_get_version()
{
    return OpenZWave::Manager::Get()->getVersion()._v;
}

static void manager_notificationWatcher(OpenZWave::Notification const* n, void*)
{
    // uint8 groupidx = n->GetGroupIdx(); // Get the index of the association group that has been changed. Only valid in Notification::Type_Group notifications.
    // uint8 event = n->GetEvent(); // Get the event value of a notification. Only valid in Notification::Type_NodeEvent and Notification::Type_ControllerCommand notifications.
    // uint8 buttonid = n->GetButtonId(); // Get the button id of a notification. Only valid in Notification::Type_CreateButton, Notification::Type_DeleteButton, Notification::Type_ButtonOn and Notification::Type_ButtonOff notifications.
    // uint8 notification = n->GetNotification(); // Get the notification code from a notification. Only valid for Notification::Type_Notification or Notification::Type_ControllerCommand notifications.
    // uint8 command = n->GetCommand(); // Get the (controller) command from a notification. Only valid for Notification::Type_ControllerCommand notifications.
    // uint8 byteval = n->GetByte(); // Helper function to simplify wrapping the notification class. Should not normally need to be called.
    // uint8 retry = n->GetRetry(); // Helper function to return the Timeout to wait for. Only valid for Notification::Type_UserAlerts - Notification::Alert_ApplicationStatus_Retry.
    // OpenZWave::Notification::UserAlertNotification uan = n->GetUserAlertType(); // Retrieve the User Alert Type Enum to determine what this message is about.
    // std::string comport = n->GetComPort(); // Return the Comport associated with the DriverFailed Message.

    uint8_t n_type = static_cast<uint8_t>(n->GetType());
    uint32_t n_vhomeid = n->GetValueID().GetHomeId();
    uint32_t n_vid0 = static_cast<uint32_t>(n->GetValueID().GetId() & 0xFFFFFFFF);
    uint32_t n_vid1 = static_cast<uint32_t>((n->GetValueID().GetId() >> 32) & 0xFFFFFFFF);
    uint8 n_byte = n->GetByte();
    uint8 n_event = 0;
    if ((OpenZWave::Notification::Type_NodeEvent == n_type) || (OpenZWave::Notification::Type_ControllerCommand == n_type)) {
        n_event = n->GetEvent();
    }
    uint8 n_command = 0;
    if (OpenZWave::Notification::Type_ControllerCommand == n_type) {
        n_command = n->GetCommand();
    }
    uint8_t n_useralerttype = static_cast<uint8_t>(n->GetUserAlertType());
    std::string n_comport = n->GetComPort();

    goNotificationWatcher(n_type, n_vhomeid, n_vid0, n_vid1, n_byte, n_event, n_command, n_useralerttype, const_cast<char*>(n_comport.c_str()), static_cast<int>(n_comport.size()));
}
