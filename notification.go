package gozwave

// #cgo LDFLAGS: -lopenzwave -L/usr/local/lib
// #cgo CPPFLAGS: -I/usr/local/include -I/usr/local/include/openzwave
// #include "notification.h"
// #include <stdlib.h>
import "C"
import "unsafe"

// NotificationType defines a type for the notification type enum.
type NotificationType int

// NotificationCode defines a type for the notification code enum.
type NotificationCode int

const (
	NotificationTypeValueAdded                   = NotificationType(C.notification_type_valueAdded)
	NotificationTypeValueRemoved                 = NotificationType(C.notification_type_valueRemoved)
	NotificationTypeValueChanged                 = NotificationType(C.notification_type_valueChanged)
	NotificationTypeValueRefreshed               = NotificationType(C.notification_type_valueRefreshed)
	NotificationTypeGroup                        = NotificationType(C.notification_type_group)
	NotificationTypeNodeNew                      = NotificationType(C.notification_type_nodeNew)
	NotificationTypeNodeAdded                    = NotificationType(C.notification_type_nodeAdded)
	NotificationTypeNodeRemoved                  = NotificationType(C.notification_type_nodeRemoved)
	NotificationTypeNodeProtocolInfo             = NotificationType(C.notification_type_nodeProtocolInfo)
	NotificationTypeNodeNaming                   = NotificationType(C.notification_type_nodeNaming)
	NotificationTypeNodeEvent                    = NotificationType(C.notification_type_nodeEvent)
	NotificationTypePollingDisabled              = NotificationType(C.notification_type_pollingDisabled)
	NotificationTypePollingEnabled               = NotificationType(C.notification_type_pollingEnabled)
	NotificationTypeSceneEvent                   = NotificationType(C.notification_type_sceneEvent)
	NotificationTypeCreateButton                 = NotificationType(C.notification_type_createButton)
	NotificationTypeDeleteButton                 = NotificationType(C.notification_type_deleteButton)
	NotificationTypeButtonOn                     = NotificationType(C.notification_type_buttonOn)
	NotificationTypeButtonOff                    = NotificationType(C.notification_type_buttonOff)
	NotificationTypeDriverReady                  = NotificationType(C.notification_type_driverReady)
	NotificationTypeDriverFailed                 = NotificationType(C.notification_type_driverFailed)
	NotificationTypeDriverReset                  = NotificationType(C.notification_type_driverReset)
	NotificationTypeEssentialNodeQueriesComplete = NotificationType(C.notification_type_essentialNodeQueriesComplete)
	NotificationTypeNodeQueriesComplete          = NotificationType(C.notification_type_nodeQueriesComplete)
	NotificationTypeAwakeNodesQueried            = NotificationType(C.notification_type_awakeNodesQueried)
	NotificationTypeAllNodesQueriedSomeDead      = NotificationType(C.notification_type_allNodesQueriedSomeDead)
	NotificationTypeAllNodesQueried              = NotificationType(C.notification_type_allNodesQueried)
	NotificationTypeNotification                 = NotificationType(C.notification_type_notification)
	NotificationTypeDriverRemoved                = NotificationType(C.notification_type_driverRemoved)
	NotificationTypeControllerCommand            = NotificationType(C.notification_type_controllerCommand)
	NotificationTypeNodeReset                    = NotificationType(C.notification_type_nodeReset)

	NotificationCodeMsgComplete = NotificationCode(C.notification_code_msgComplete)
	NotificationCodeTimeout     = NotificationCode(C.notification_code_timeout)
	NotificationCodeNoOperation = NotificationCode(C.notification_code_noOperation)
	NotificationCodeAwake       = NotificationCode(C.notification_code_awake)
	NotificationCodeSleep       = NotificationCode(C.notification_code_sleep)
	NotificationCodeDead        = NotificationCode(C.notification_code_dead)
	NotificationCodeAlive       = NotificationCode(C.notification_code_alive)
)

func (nt NotificationType) String() string {
	switch nt {
	case NotificationTypeValueAdded:
		return "ValueAdded"
	case NotificationTypeValueRemoved:
		return "ValueRemoved"
	case NotificationTypeValueChanged:
		return "ValueChanged"
	case NotificationTypeValueRefreshed:
		return "ValueRefreshed"
	case NotificationTypeGroup:
		return "Group"
	case NotificationTypeNodeNew:
		return "NodeNew"
	case NotificationTypeNodeAdded:
		return "NodeAdded"
	case NotificationTypeNodeRemoved:
		return "NodeRemoved"
	case NotificationTypeNodeProtocolInfo:
		return "NodeProtocolInfo"
	case NotificationTypeNodeNaming:
		return "NodeNaming"
	case NotificationTypeNodeEvent:
		return "NodeEvent"
	case NotificationTypePollingDisabled:
		return "PollingDisabled"
	case NotificationTypePollingEnabled:
		return "PollingEnabled"
	case NotificationTypeSceneEvent:
		return "SceneEvent"
	case NotificationTypeCreateButton:
		return "CreateButton"
	case NotificationTypeDeleteButton:
		return "DeleteButton"
	case NotificationTypeButtonOn:
		return "ButtonOn"
	case NotificationTypeButtonOff:
		return "ButtonOff"
	case NotificationTypeDriverReady:
		return "DriverReady"
	case NotificationTypeDriverFailed:
		return "DriverFailed"
	case NotificationTypeDriverReset:
		return "DriverReset"
	case NotificationTypeEssentialNodeQueriesComplete:
		return "EssentialNodeQueriesComplete"
	case NotificationTypeNodeQueriesComplete:
		return "NodeQueriesComplete"
	case NotificationTypeAwakeNodesQueried:
		return "AwakeNodesQueried"
	case NotificationTypeAllNodesQueriedSomeDead:
		return "AllNodesQueriedSomeDead"
	case NotificationTypeAllNodesQueried:
		return "AllNodesQueried"
	case NotificationTypeNotification:
		return "Notification"
	case NotificationTypeDriverRemoved:
		return "DriverRemoved"
	case NotificationTypeControllerCommand:
		return "ControllerCommand"
	case NotificationTypeNodeReset:
		return "NodeReset"
	}
	return "UNKNOWN"
}

// Notification is a container for the C++ OpenZWave library Notification class.
type Notification struct {
	Type         NotificationType
	HomeID       uint32
	NodeID       uint8
	ValueID      ValueID
	GroupIDX     uint8
	Event        uint8
	ButtonID     uint8
	SceneID      uint8
	Notification uint8
	Name         string
}

func buildNotification(n C.notification_t) Notification {
	notification := Notification{
		Type:    NotificationType(C.notification_getType(n)),
		HomeID:  uint32(C.notification_getHomeId(n)),
		NodeID:  uint8(C.notification_getNodeId(n)),
		ValueID: ValueID{valueid: C.notification_getValueId(n)},
	}

	cstr := C.notification_getAsString(n)
	notification.Name = C.GoString(cstr)
	C.free(unsafe.Pointer(cstr))

	// switch notification.Type {
	// case NotificationTypeCreateButton:
	// case NotificationTypeDeleteButton:
	// case NotificationTypeButtonOn:
	// case NotificationTypeButtonOff:
	// 	notification.ButtonID = uint8(C.notification_getButtonId(n))
	// 	break

	// case NotificationTypeNodeEvent:
	// 	notification.Event = uint8(C.notification_getEvent(n))
	// 	break

	// case NotificationTypeGroup:
	// 	notification.GroupIDX = uint8(C.notification_getGroupIdx(n))
	// 	break

	// case NotificationTypeNotification:
	// 	notification.Notification = uint8(C.notification_getNotification(n))
	// 	break

	// case NotificationTypeControllerCommand:
	// 	notification.Event = uint8(C.notification_getEvent(n))
	// 	notification.Notification = uint8(C.notification_getNotification(n))
	// 	break

	// case NotificationTypeSceneEvent:
	// 	notification.SceneID = uint8(C.notification_getSceneId(n))
	// 	break
	// }

	return notification
}

// func (n *Notification) GetType() NotificationType {
// 	return NotificationType(C.notification_getType(n.notification))
// }

// func (n *Notification) GetHomeId() uint32 {
// 	return uint32(C.notification_getHomeId(n.notification))
// }

// func (n *Notification) GetNodeId() uint8 {
// 	return uint8(C.notification_getNodeId(n.notification))
// }

// func (n *Notification) GetValueId() *ValueId {
// 	val := &ValueId{}
// 	val.valueid = C.notification_getValueId(n.notification)
// 	return val
// }

// func (n *Notification) GetGroupIdx() uint8 {
// 	return uint8(C.notification_getGroupIdx(n.notification))
// }

// func (n *Notification) GetEvent() uint8 {
// 	return uint8(C.notification_getEvent(n.notification))
// }

// func (n *Notification) GetButtonId() uint8 {
// 	return uint8(C.notification_getButtonId(n.notification))
// }

// func (n *Notification) GetSceneId() uint8 {
// 	return uint8(C.notification_getSceneId(n.notification))
// }

// func (n *Notification) GetNotification() uint8 {
// 	return uint8(C.notification_getNotification(n.notification))
// }

// func (n *Notification) GetByte() uint8 {
// 	return uint8(C.notification_getByte(n.notification))
// }

// func (n *Notification) GetAsString() string {
// 	cstr := C.notification_getAsString(n.notification)
// 	str := C.GoString(cstr)
// 	C.free(unsafe.Pointer(cstr))
// 	return str
// }
