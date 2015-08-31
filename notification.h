#ifndef GOZWAVE_NOTIFICATION
#define GOZWAVE_NOTIFICATION

#include <stdint.h>
#include <stdbool.h>
#include <stddef.h>
#include "string_helpers.h"
#include "valueid.h"

#ifdef __cplusplus
extern "C" {
#endif

    // Types.
    typedef void* notification_t;

    // enum notification_type
    typedef enum {
        notification_type_valueAdded = 0,
        notification_type_valueRemoved,
        notification_type_valueChanged,
        notification_type_valueRefreshed,
        notification_type_group,
        notification_type_nodeNew,
        notification_type_nodeAdded,
        notification_type_nodeRemoved,
        notification_type_nodeProtocolInfo,
        notification_type_nodeNaming,
        notification_type_nodeEvent,
        notification_type_pollingDisabled,
        notification_type_pollingEnabled,
        notification_type_sceneEvent,
        notification_type_createButton,
        notification_type_deleteButton,
        notification_type_buttonOn,
        notification_type_buttonOff,
        notification_type_driverReady,
        notification_type_driverFailed,
        notification_type_driverReset,
        notification_type_essentialNodeQueriesComplete,
        notification_type_nodeQueriesComplete,
        notification_type_awakeNodesQueried,
        notification_type_allNodesQueriedSomeDead,
        notification_type_allNodesQueried,
        notification_type_notification,
        notification_type_driverRemoved,
        notification_type_controllerCommand,
        notification_type_nodeReset
    } notification_type;

    // enum notification_code
    typedef enum {
        notification_code_msgComplete = 0,
        notification_code_timeout,
        notification_code_noOperation,
        notification_code_awake,
        notification_code_sleep,
        notification_code_dead,
        notification_code_alive
    } notification_code;

    // Public member functions.
    notification_type notification_getType(notification_t n);
    uint32_t notification_getHomeId(notification_t n);
    uint8_t notification_getNodeId(notification_t n);
    valueid_t notification_getValueId(notification_t n);
    uint8_t notification_getGroupIdx(notification_t n);
    uint8_t notification_getEvent(notification_t n);
    uint8_t notification_getButtonId(notification_t n);
    uint8_t notification_getSceneId(notification_t n);
    uint8_t notification_getNotification(notification_t n);
    uint8_t notification_getByte(notification_t n);
    string_t* notification_getAsString(notification_t n);

#ifdef __cplusplus
}
#endif

#endif // define GOZWAVE_NOTIFICATION
