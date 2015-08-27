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
	Byte         uint8
	Name         string
}

func buildNotification(n C.notification_t) Notification {
	notification := Notification{
		Type:         NotificationType(C.notification_getType(n)),
		HomeID:       uint32(C.notification_getHomeId(n)),
		NodeID:       uint8(C.notification_getNodeId(n)),
		ValueID:      ValueID{valueid: C.notification_getValueId(n)},
		GroupIDX:     uint8(C.notification_getGroupIdx(n)),
		Event:        uint8(C.notification_getEvent(n)),
		ButtonID:     uint8(C.notification_getButtonId(n)),
		SceneID:      uint8(C.notification_getSceneId(n)),
		Notification: uint8(C.notification_getNotification(n)),
		Byte:         uint8(C.notification_getByte(n)),
	}

	cstr := C.notification_getAsString(n)
	notification.Name = C.GoString(cstr)
	C.free(unsafe.Pointer(cstr))

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
