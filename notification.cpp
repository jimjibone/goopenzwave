#include "notification.h"
#include <openzwave/Notification.h>
#include <string.h>

//
// Public member functions.
//

notification_type notification_getType(notification_t n)
{
    OpenZWave::Notification *noti = (OpenZWave::Notification*)n;
    notification_type noti_type;
    switch (noti->GetType()) {
        case OpenZWave::Notification::Type_ValueAdded:
            noti_type = notification_type_valueAdded;
            break;
        case OpenZWave::Notification::Type_ValueRemoved:
            noti_type = notification_type_valueRemoved;
            break;
        case OpenZWave::Notification::Type_ValueChanged:
            noti_type = notification_type_valueChanged;
            break;
        case OpenZWave::Notification::Type_ValueRefreshed:
            noti_type = notification_type_valueRefreshed;
            break;
        case OpenZWave::Notification::Type_Group:
            noti_type = notification_type_group;
            break;
        case OpenZWave::Notification::Type_NodeNew:
            noti_type = notification_type_nodeNew;
            break;
        case OpenZWave::Notification::Type_NodeAdded:
            noti_type = notification_type_nodeAdded;
            break;
        case OpenZWave::Notification::Type_NodeRemoved:
            noti_type = notification_type_nodeRemoved;
            break;
        case OpenZWave::Notification::Type_NodeProtocolInfo:
            noti_type = notification_type_nodeProtocolInfo;
            break;
        case OpenZWave::Notification::Type_NodeNaming:
            noti_type = notification_type_nodeNaming;
            break;
        case OpenZWave::Notification::Type_NodeEvent:
            noti_type = notification_type_nodeEvent;
            break;
        case OpenZWave::Notification::Type_PollingDisabled:
            noti_type = notification_type_pollingDisabled;
            break;
        case OpenZWave::Notification::Type_PollingEnabled:
            noti_type = notification_type_pollingEnabled;
            break;
        case OpenZWave::Notification::Type_SceneEvent:
            noti_type = notification_type_sceneEvent;
            break;
        case OpenZWave::Notification::Type_CreateButton:
            noti_type = notification_type_createButton;
            break;
        case OpenZWave::Notification::Type_DeleteButton:
            noti_type = notification_type_deleteButton;
            break;
        case OpenZWave::Notification::Type_ButtonOn:
            noti_type = notification_type_buttonOn;
            break;
        case OpenZWave::Notification::Type_ButtonOff:
            noti_type = notification_type_buttonOff;
            break;
        case OpenZWave::Notification::Type_DriverReady:
            noti_type = notification_type_driverReady;
            break;
        case OpenZWave::Notification::Type_DriverFailed:
            noti_type = notification_type_driverFailed;
            break;
        case OpenZWave::Notification::Type_DriverReset:
            noti_type = notification_type_driverReset;
            break;
        case OpenZWave::Notification::Type_EssentialNodeQueriesComplete:
            noti_type = notification_type_essentialNodeQueriesComplete;
            break;
        case OpenZWave::Notification::Type_NodeQueriesComplete:
            noti_type = notification_type_nodeQueriesComplete;
            break;
        case OpenZWave::Notification::Type_AwakeNodesQueried:
            noti_type = notification_type_awakeNodesQueried;
            break;
        case OpenZWave::Notification::Type_AllNodesQueriedSomeDead:
            noti_type = notification_type_allNodesQueriedSomeDead;
            break;
        case OpenZWave::Notification::Type_AllNodesQueried:
            noti_type = notification_type_allNodesQueried;
            break;
        case OpenZWave::Notification::Type_Notification:
            noti_type = notification_type_notification;
            break;
        case OpenZWave::Notification::Type_DriverRemoved:
            noti_type = notification_type_driverRemoved;
            break;
        case OpenZWave::Notification::Type_ControllerCommand:
            noti_type = notification_type_controllerCommand;
            break;
        case OpenZWave::Notification::Type_NodeReset:
            noti_type = notification_type_nodeReset;
            break;
    }
    return noti_type;
}

uint32_t notification_getHomeId(notification_t n)
{
    OpenZWave::Notification *noti = (OpenZWave::Notification*)n;
    return noti->GetHomeId();
}

uint8_t notification_getNodeId(notification_t n)
{
    OpenZWave::Notification *noti = (OpenZWave::Notification*)n;
    return noti->GetNodeId();
}

valueid_t notification_getValueId(notification_t n)
{
    OpenZWave::Notification *noti = (OpenZWave::Notification*)n;
    valueid_t valid = (OpenZWave::ValueID*)&(noti->GetValueID());
    return valid;
}

uint8_t notification_getGroupIdx(notification_t n)
{
    OpenZWave::Notification *noti = (OpenZWave::Notification*)n;
    return noti->GetGroupIdx();
}

uint8_t notification_getEvent(notification_t n)
{
    OpenZWave::Notification *noti = (OpenZWave::Notification*)n;
    return noti->GetEvent();
}

uint8_t notification_getButtonId(notification_t n)
{
    OpenZWave::Notification *noti = (OpenZWave::Notification*)n;
    return noti->GetButtonId();
}

uint8_t notification_getSceneId(notification_t n)
{
    OpenZWave::Notification *noti = (OpenZWave::Notification*)n;
    return noti->GetSceneId();
}

uint8_t notification_getNotification(notification_t n)
{
    OpenZWave::Notification *noti = (OpenZWave::Notification*)n;
    return noti->GetNotification();
}

uint8_t notification_getByte(notification_t n)
{
    OpenZWave::Notification *noti = (OpenZWave::Notification*)n;
    return noti->GetEvent();
}

const char* notification_getAsString(notification_t n)
{
    OpenZWave::Notification *noti = (OpenZWave::Notification*)n;
    std::string not_string = noti->GetAsString();
    char *not_cstring = new char[not_string.size()+1];
    memcpy(not_cstring, not_string.c_str(), not_string.size());
    not_cstring[not_string.size()] = 0;
    return not_cstring;
}
