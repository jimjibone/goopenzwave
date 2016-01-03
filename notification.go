package goopenzwave

// #include "notification.h"
// #include <stdlib.h>
import "C"
import (
	"fmt"
)

// NotificationType defines a type for the notification type enum.
type NotificationType int

const (
	NotificationTypeValueAdded NotificationType = iota
	NotificationTypeValueRemoved
	NotificationTypeValueChanged
	NotificationTypeValueRefreshed
	NotificationTypeGroup
	NotificationTypeNodeNew
	NotificationTypeNodeAdded
	NotificationTypeNodeRemoved
	NotificationTypeNodeProtocolInfo
	NotificationTypeNodeNaming
	NotificationTypeNodeEvent
	NotificationTypePollingDisabled
	NotificationTypePollingEnabled
	NotificationTypeSceneEvent
	NotificationTypeCreateButton
	NotificationTypeDeleteButton
	NotificationTypeButtonOn
	NotificationTypeButtonOff
	NotificationTypeDriverReady
	NotificationTypeDriverFailed
	NotificationTypeDriverReset
	NotificationTypeEssentialNodeQueriesComplete
	NotificationTypeNodeQueriesComplete
	NotificationTypeAwakeNodesQueried
	NotificationTypeAllNodesQueriedSomeDead
	NotificationTypeAllNodesQueried
	NotificationTypeNotification
	NotificationTypeDriverRemoved
	NotificationTypeControllerCommand
	NotificationTypeNodeReset
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

// NotificationCode defines a type for the notification code enum.
type NotificationCode int

const (
	NotificationCodeMsgComplete NotificationCode = iota // C.notification_code_msgComplete
	NotificationCodeTimeout                             // C.notification_code_timeout
	NotificationCodeNoOperation                         // C.notification_code_noOperation
	NotificationCodeAwake                               // C.notification_code_awake
	NotificationCodeSleep                               // C.notification_code_sleep
	NotificationCodeDead                                // C.notification_code_dead
	NotificationCodeAlive                               // C.notification_code_alive
)

func (nc NotificationCode) String() string {
	switch nc {
	case NotificationCodeMsgComplete:
		return "MsgComplete"
	case NotificationCodeTimeout:
		return "Timeout"
	case NotificationCodeNoOperation:
		return "NoOperation"
	case NotificationCodeAwake:
		return "Awake"
	case NotificationCodeSleep:
		return "Sleep"
	case NotificationCodeDead:
		return "Dead"
	case NotificationCodeAlive:
		return "Alive"
	}
	return "UNKNOWN"
}

// Notification is a container for the C++ OpenZWave library Notification class.
type Notification struct {
	Type         NotificationType
	HomeID       uint32
	NodeID       uint8
	ValueID      *ValueID
	GroupIDX     *uint8
	Event        *uint8
	ButtonID     *uint8
	SceneID      *uint8
	Notification *NotificationCode
}

// buildNotification builds a new Notification filled with the relevant
// information from the OpenZWave::Notification as received from the OpenZWave
// library.
func buildNotification(n C.notification_t) *Notification {
	notification := &Notification{
		HomeID:  uint32(C.notification_getHomeId(n)),
		NodeID:  uint8(C.notification_getNodeId(n)),
		ValueID: buildValueID(C.notification_getValueId(n)),
	}

	switch C.notification_getType(n) {
	case C.notification_type_valueAdded:
		notification.Type = NotificationTypeValueAdded
	case C.notification_type_valueRemoved:
		notification.Type = NotificationTypeValueRemoved
	case C.notification_type_valueChanged:
		notification.Type = NotificationTypeValueChanged
	case C.notification_type_valueRefreshed:
		notification.Type = NotificationTypeValueRefreshed
	case C.notification_type_group:
		notification.Type = NotificationTypeGroup
	case C.notification_type_nodeNew:
		notification.Type = NotificationTypeNodeNew
	case C.notification_type_nodeAdded:
		notification.Type = NotificationTypeNodeAdded
	case C.notification_type_nodeRemoved:
		notification.Type = NotificationTypeNodeRemoved
	case C.notification_type_nodeProtocolInfo:
		notification.Type = NotificationTypeNodeProtocolInfo
	case C.notification_type_nodeNaming:
		notification.Type = NotificationTypeNodeNaming
	case C.notification_type_nodeEvent:
		notification.Type = NotificationTypeNodeEvent
	case C.notification_type_pollingDisabled:
		notification.Type = NotificationTypePollingDisabled
	case C.notification_type_pollingEnabled:
		notification.Type = NotificationTypePollingEnabled
	case C.notification_type_sceneEvent:
		notification.Type = NotificationTypeSceneEvent
	case C.notification_type_createButton:
		notification.Type = NotificationTypeCreateButton
	case C.notification_type_deleteButton:
		notification.Type = NotificationTypeDeleteButton
	case C.notification_type_buttonOn:
		notification.Type = NotificationTypeButtonOn
	case C.notification_type_buttonOff:
		notification.Type = NotificationTypeButtonOff
	case C.notification_type_driverReady:
		notification.Type = NotificationTypeDriverReady
	case C.notification_type_driverFailed:
		notification.Type = NotificationTypeDriverFailed
	case C.notification_type_driverReset:
		notification.Type = NotificationTypeDriverReset
	case C.notification_type_essentialNodeQueriesComplete:
		notification.Type = NotificationTypeEssentialNodeQueriesComplete
	case C.notification_type_nodeQueriesComplete:
		notification.Type = NotificationTypeNodeQueriesComplete
	case C.notification_type_awakeNodesQueried:
		notification.Type = NotificationTypeAwakeNodesQueried
	case C.notification_type_allNodesQueriedSomeDead:
		notification.Type = NotificationTypeAllNodesQueriedSomeDead
	case C.notification_type_allNodesQueried:
		notification.Type = NotificationTypeAllNodesQueried
	case C.notification_type_notification:
		notification.Type = NotificationTypeNotification
	case C.notification_type_driverRemoved:
		notification.Type = NotificationTypeDriverRemoved
	case C.notification_type_controllerCommand:
		notification.Type = NotificationTypeControllerCommand
	case C.notification_type_nodeReset:
		notification.Type = NotificationTypeNodeReset
	}

	switch notification.Type {
	case NotificationTypeCreateButton, NotificationTypeDeleteButton, NotificationTypeButtonOn, NotificationTypeButtonOff:
		if notification.ButtonID == nil {
			notification.ButtonID = new(uint8)
		}
		*(notification.ButtonID) = uint8(C.notification_getButtonId(n))

	case NotificationTypeNodeEvent:
		if notification.Event == nil {
			notification.Event = new(uint8)
		}
		*(notification.Event) = uint8(C.notification_getEvent(n))

	case NotificationTypeGroup:
		if notification.GroupIDX == nil {
			notification.GroupIDX = new(uint8)
		}
		*(notification.GroupIDX) = uint8(C.notification_getGroupIdx(n))

	case NotificationTypeNotification:
		if notification.Notification == nil {
			notification.Notification = new(NotificationCode)
		}
		switch C.notification_getNotification(n) {
		case C.notification_code_msgComplete:
			*notification.Notification = NotificationCodeMsgComplete
		case C.notification_code_timeout:
			*notification.Notification = NotificationCodeTimeout
		case C.notification_code_noOperation:
			*notification.Notification = NotificationCodeNoOperation
		case C.notification_code_awake:
			*notification.Notification = NotificationCodeAwake
		case C.notification_code_sleep:
			*notification.Notification = NotificationCodeSleep
		case C.notification_code_dead:
			*notification.Notification = NotificationCodeDead
		case C.notification_code_alive:
			*notification.Notification = NotificationCodeAlive
		}

	case NotificationTypeControllerCommand:
		if notification.Event == nil {
			notification.Event = new(uint8)
		}
		*(notification.Event) = uint8(C.notification_getEvent(n))
		if notification.Notification == nil {
			notification.Notification = new(NotificationCode)
		}
		*(notification.Notification) = NotificationCode(C.notification_getNotification(n))

	case NotificationTypeSceneEvent:
		if notification.SceneID == nil {
			notification.SceneID = new(uint8)
		}
		*(notification.SceneID) = uint8(C.notification_getSceneId(n))
	}

	return notification
}

func (n *Notification) String() string {
	var pointed []string
	if n.GroupIDX != nil {
		pointed = append(pointed, fmt.Sprintf("GroupIDX: %d", *(n.GroupIDX)))
	}
	if n.Event != nil {
		pointed = append(pointed, fmt.Sprintf("Event: %d", *(n.Event)))
	}
	if n.ButtonID != nil {
		pointed = append(pointed, fmt.Sprintf("ButtonID: %d", *(n.ButtonID)))
	}
	if n.SceneID != nil {
		pointed = append(pointed, fmt.Sprintf("SceneID: %d", *(n.SceneID)))
	}
	if n.Notification != nil {
		pointed = append(pointed, fmt.Sprintf("Notification: %d", *(n.Notification)))
	}
	output := fmt.Sprintf("Notification{Type: %s, HomeID: %d, NodeID: %d, ValueID: %s", n.Type, n.HomeID, n.NodeID, n.ValueID)
	for i := range pointed {
		output = fmt.Sprintf("%s, %s", output, pointed[i])
	}
	return output
}
