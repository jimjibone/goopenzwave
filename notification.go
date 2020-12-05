package goopenzwave

import "fmt"

//go:generate stringer -trimprefix=NotificationType_ -type=NotificationType
//go:generate stringer -trimprefix=NotificationCode_ -type=NotificationCode
//go:generate stringer -trimprefix=UserAlertNotification_ -type=UserAlertNotification
//go:generate stringer -trimprefix=ControllerCommand_ -type=ControllerCommand
//go:generate stringer -trimprefix=ControllerState_ -type=ControllerState
//go:generate stringer -trimprefix=ControllerError_ -type=ControllerError

type Notification struct {
	Type                  NotificationType
	ValueID               ValueID
	Byte                  uint8
	Event                 uint8
	Command               uint8
	UserAlertNotification UserAlertNotification
	ComPort               string
}

func createNotification(n_type uint8, n_vhomeid, n_vid0, n_vid1 uint32, n_byte, n_event, n_command, n_useralerttype uint8, n_comport string) Notification {
	return Notification{
		Type:                  NotificationType(n_type),
		ValueID:               ValueID{n_vhomeid, n_vid0, n_vid1},
		Byte:                  n_byte,
		Event:                 n_event,
		Command:               n_command,
		UserAlertNotification: UserAlertNotification(n_useralerttype),
		ComPort:               n_comport,
	}
}

func (n Notification) String() string {
	out := "{" + n.Type.String() + ", ValueID:" + n.ValueID.String()
	switch n.Type {
	case NotificationType_ValueAdded: // A new node value has been added to OpenZWave's list. These notifications occur after a node has been discovered, and details of its command classes have been received. Each command class may generate one or more values depending on the complexity of the item being represented.
	case NotificationType_ValueRemoved: // A node value has been removed from OpenZWave's list. This only occurs when a node is removed.
	case NotificationType_ValueChanged: // A node value has been updated from the Z-Wave network and it is different from the previous value.
	case NotificationType_ValueRefreshed: // A node value has been updated from the Z-Wave network.
	case NotificationType_Group: // The associations for the node have changed. The application should rebuild any group information it holds about the node.
		out += fmt.Sprintf(", GroupIdx:%d", n.GetGroupIdx())
	case NotificationType_NodeNew: // A new node has been found (not already stored in zwcfg*.xml file)
	case NotificationType_NodeAdded: // A new node has been added to OpenZWave's list. This may be due to a device being added to the Z-Wave network, or because the application is initializing itself.
	case NotificationType_NodeRemoved: // A node has been removed from OpenZWave's list. This may be due to a device being removed from the Z-Wave network, or because the application is closing.
	case NotificationType_NodeProtocolInfo: // Basic node information has been received, such as whether the node is a listening device, a routing device and its baud rate and basic, generic and specific types. It is after this notification that you can call Manager::GetNodeType to obtain a label containing the device description.
	case NotificationType_NodeNaming: // One of the node names has changed (name, manufacturer, product).
		out += fmt.Sprintf(", Name:%s, Product:%s, Manufacturer:%s",
			GetNodeName(n.ValueID.HomeID(), n.ValueID.NodeID()),
			GetNodeProductName(n.ValueID.HomeID(), n.ValueID.NodeID()),
			GetNodeManufacturerName(n.ValueID.HomeID(), n.ValueID.NodeID()),
		)
	case NotificationType_NodeEvent: // A node has triggered an event. This is commonly caused when a node sends a Basic_Set command to the controller. The event value is stored in the notification.
		out += fmt.Sprintf(", NodeEvent:%d", n.GetEvent())
	case NotificationType_PollingDisabled: // Polling of a node has been successfully turned off by a call to Manager::DisablePoll
	case NotificationType_PollingEnabled: // Polling of a node has been successfully turned on by a call to Manager::EnablePoll
	case NotificationType_SceneEvent: // Scene Activation Set received (Depreciated in 1.8)
	case NotificationType_CreateButton: // Handheld controller button event created
	case NotificationType_DeleteButton: // Handheld controller button event deleted
	case NotificationType_ButtonOn: // Handheld controller button on pressed event
	case NotificationType_ButtonOff: // Handheld controller button off pressed event
	case NotificationType_DriverReady: // A driver for a PC Z-Wave controller has been added and is ready to use. The notification will contain the controller's Home ID, which is needed to call most of the Manager methods.
	case NotificationType_DriverFailed: // Driver failed to load
		out += fmt.Sprintf(", ComPort:%s", n.GetComPort())
	case NotificationType_DriverReset: // All nodes and values for this driver have been removed. This is sent instead of potentially hundreds of individual node and value notifications.
	case NotificationType_EssentialNodeQueriesComplete: // The queries on a node that are essential to its operation have been completed. The node can now handle incoming messages.
	case NotificationType_NodeQueriesComplete: // All the initialization queries on a node have been completed.
	case NotificationType_AwakeNodesQueried: // All awake nodes have been queried, so client application can expected complete data for these nodes.
	case NotificationType_AllNodesQueriedSomeDead: // All nodes have been queried but some dead nodes found.
	case NotificationType_AllNodesQueried: // All nodes have been queried, so client application can expected complete data.
	case NotificationType_Notification: // An error has occurred that we need to report.
		out += fmt.Sprintf(", Notification:%s", n.GetNotification())
	case NotificationType_DriverRemoved: // The Driver is being removed. (either due to Error or by request) Do Not Call Any Driver Related Methods after receiving this call
	case NotificationType_ControllerCommand: // When Controller Commands are executed, Notifications of Success/Failure etc are communicated via this. Notification Notification::GetEvent returns Driver::ControllerCommand and Notification::GetNotification returns Driver::ControllerState
		out += fmt.Sprintf(", ControllerCommand:%s, ControllerState:%s", ControllerCommand(n.GetEvent()), ControllerState(n.GetNotification()))
	case NotificationType_NodeReset: // The Device has been reset and thus removed from the NodeList in OZW
	case NotificationType_UserAlerts: // Warnings and Notifications Generated by the library that should be displayed to the user (eg, out of date config files)
		out += fmt.Sprintf(", UserAlerts:%s", n.GetUserAlertType())
	case NotificationType_ManufacturerSpecificDBReady: // The ManufacturerSpecific Database Is Ready
	}
	out += "}"
	return out
}

// GetGroupIdx for association group that has been changed. Only valid in NotificationType_Group notifications.
func (n Notification) GetGroupIdx() uint8 {
	if NotificationType_Group == n.Type {
		return n.Byte
	}
	return 0
}

// GetEvent returns the event value of a notification. Only valid in NotificationType_NodeEvent and
// NotificationType_ControllerCommand notifications.
func (n Notification) GetEvent() uint8 {
	if NotificationType_NodeEvent == n.Type || NotificationType_ControllerCommand == n.Type {
		return n.Event
	}
	return 0
}

// GetButtonID of a notification. Only valid in NotificationType_CreateButton, NotificationType_DeleteButton,
// NotificationType_ButtonOn and NotificationType_ButtonOff notifications.
func (n Notification) GetButtonID() uint8 {
	if NotificationType_CreateButton == n.Type || NotificationType_DeleteButton == n.Type || NotificationType_ButtonOn == n.Type || NotificationType_ButtonOff == n.Type {
		return n.Byte
	}
	return 0
}

/**
* GetNotification code from a notification. Only valid for NotificationType_Notification or NotificationType_ControllerCommand notifications.
 */
func (n Notification) GetNotification() NotificationCode {
	if NotificationType_Notification == n.Type || NotificationType_ControllerCommand == n.Type {
		return NotificationCode(n.Byte)
	}
	return 0
}

// GetCommand returns the (controller) command from a notification. Only valid for NotificationType_ControllerCommand notifications.
func (n Notification) GetCommand() uint8 {
	if NotificationType_ControllerCommand == n.Type {
		return n.Command
	}
	return 0
}

// GetRetry is a helper function to return the timeout to wait for. Only valid for NotificationType_UserAlerts - UserAlertNotification_ApplicationStatus_Retry.
func (n Notification) GetRetry() uint8 {
	if NotificationType_UserAlerts == n.Type && UserAlertNotification_ApplicationStatus_Retry == n.UserAlertNotification {
		return n.Byte
	}
	return 0
}

// GetUserAlertType returns the User Alert Type Enum to determine what this message is about.
func (n Notification) GetUserAlertType() UserAlertNotification {
	return n.UserAlertNotification
}

// GetComPort returns the Comport associated with the DriverFailed Message.
func (n Notification) GetComPort() string {
	return n.ComPort
}

// NotificationType for various Z-Wave events sent to the watchers registered with the ManagerAddWatcher method.
type NotificationType uint8

const (
	NotificationType_ValueAdded                   NotificationType = iota // A new node value has been added to OpenZWave's list. These notifications occur after a node has been discovered, and details of its command classes have been received. Each command class may generate one or more values depending on the complexity of the item being represented.
	NotificationType_ValueRemoved                                         // A node value has been removed from OpenZWave's list. This only occurs when a node is removed.
	NotificationType_ValueChanged                                         // A node value has been updated from the Z-Wave network and it is different from the previous value.
	NotificationType_ValueRefreshed                                       // A node value has been updated from the Z-Wave network.
	NotificationType_Group                                                // The associations for the node have changed. The application should rebuild any group information it holds about the node.
	NotificationType_NodeNew                                              // A new node has been found (not already stored in zwcfg*.xml file)
	NotificationType_NodeAdded                                            // A new node has been added to OpenZWave's list. This may be due to a device being added to the Z-Wave network, or because the application is initializing itself.
	NotificationType_NodeRemoved                                          // A node has been removed from OpenZWave's list. This may be due to a device being removed from the Z-Wave network, or because the application is closing.
	NotificationType_NodeProtocolInfo                                     // Basic node information has been received, such as whether the node is a listening device, a routing device and its baud rate and basic, generic and specific types. It is after this notification that you can call Manager::GetNodeType to obtain a label containing the device description.
	NotificationType_NodeNaming                                           // One of the node names has changed (name, manufacturer, product).
	NotificationType_NodeEvent                                            // A node has triggered an event. This is commonly caused when a node sends a Basic_Set command to the controller. The event value is stored in the notification.
	NotificationType_PollingDisabled                                      // Polling of a node has been successfully turned off by a call to Manager::DisablePoll
	NotificationType_PollingEnabled                                       // Polling of a node has been successfully turned on by a call to Manager::EnablePoll
	NotificationType_SceneEvent                                           // Scene Activation Set received (Depreciated in 1.8)
	NotificationType_CreateButton                                         // Handheld controller button event created
	NotificationType_DeleteButton                                         // Handheld controller button event deleted
	NotificationType_ButtonOn                                             // Handheld controller button on pressed event
	NotificationType_ButtonOff                                            // Handheld controller button off pressed event
	NotificationType_DriverReady                                          // A driver for a PC Z-Wave controller has been added and is ready to use. The notification will contain the controller's Home ID, which is needed to call most of the Manager methods.
	NotificationType_DriverFailed                                         // Driver failed to load
	NotificationType_DriverReset                                          // All nodes and values for this driver have been removed. This is sent instead of potentially hundreds of individual node and value notifications.
	NotificationType_EssentialNodeQueriesComplete                         // The queries on a node that are essential to its operation have been completed. The node can now handle incoming messages.
	NotificationType_NodeQueriesComplete                                  // All the initialization queries on a node have been completed.
	NotificationType_AwakeNodesQueried                                    // All awake nodes have been queried, so client application can expected complete data for these nodes.
	NotificationType_AllNodesQueriedSomeDead                              // All nodes have been queried but some dead nodes found.
	NotificationType_AllNodesQueried                                      // All nodes have been queried, so client application can expected complete data.
	NotificationType_Notification                                         // An error has occurred that we need to report.
	NotificationType_DriverRemoved                                        // The Driver is being removed. (either due to Error or by request) Do Not Call Any Driver Related Methods after receiving this call
	NotificationType_ControllerCommand                                    // When Controller Commands are executed, Notifications of Success/Failure etc are communicated via this Notification Notification::GetEvent returns Driver::ControllerCommand and Notification::GetNotification returns Driver::ControllerState
	NotificationType_NodeReset                                            // The Device has been reset and thus removed from the NodeList in OZW
	NotificationType_UserAlerts                                           // Warnings and Notifications Generated by the library that should be displayed to the user (eg, out of date config files)
	NotificationType_ManufacturerSpecificDBReady                          // The ManufacturerSpecific Database Is Ready
)

// NotificationCodes for NotificationType_Notification convey some extra information defined here.
type NotificationCode uint8

const (
	NotificationCode_MsgComplete NotificationCode = iota // Completed messages
	NotificationCode_Timeout                             // Messages that timeout will send a Notification with this code.
	NotificationCode_NoOperation                         // Report on NoOperation message sent completion
	NotificationCode_Awake                               // Report when a sleeping node wakes up
	NotificationCode_Sleep                               // Report when a node goes to sleep
	NotificationCode_Dead                                // Report when a node is presumed dead
	NotificationCode_Alive                               // Report when a node is revived
)

// UserAlertNotification for messages that should be displayed to users to inform them of potential issues such as Out of Date configuration files etc.
type UserAlertNotification uint8

const (
	UserAlertNotification_None                       UserAlertNotification = iota // No Alert Currently Present
	UserAlertNotification_ConfigOutOfDate                                         // One of the Config Files is out of date. Use GetNodeId to determine which node is effected.
	UserAlertNotification_MFSOutOfDate                                            // the manufacturer_specific.xml file is out of date.
	UserAlertNotification_ConfigFileDownloadFailed                                // A Config File failed to download
	UserAlertNotification_DNSError                                                // A error occurred performing a DNS Lookup
	UserAlertNotification_NodeReloadRequired                                      // A new Config file has been discovered for this node, and its pending a Reload to Take affect
	UserAlertNotification_UnsupportedController                                   // The Controller is not running a Firmware Library we support
	UserAlertNotification_ApplicationStatus_Retry                                 // Application Status CC returned a Retry Later Message
	UserAlertNotification_ApplicationStatus_Queued                                // Command Has been Queued for later execution
	UserAlertNotification_ApplicationStatus_Rejected                              // Command has been rejected
)

// ControllerCommand to be used with the BeginControllerCommand method.
type ControllerCommand uint8

const (
	ControllerCommand_None                      ControllerCommand = iota // No command.
	ControllerCommand_AddDevice                                          // Add a new device or controller to the Z-Wave network.
	ControllerCommand_CreateNewPrimary                                   // Add a new controller to the Z-Wave network. Used when old primary fails. Requires SUC.
	ControllerCommand_ReceiveConfiguration                               // Receive Z-Wave network configuration information from another controller.
	ControllerCommand_RemoveDevice                                       // Remove a device or controller from the Z-Wave network.
	ControllerCommand_RemoveFailedNode                                   // Move a node to the controller's failed nodes list. This command will only work if the node cannot respond.
	ControllerCommand_HasNodeFailed                                      // Check whether a node is in the controller's failed nodes list.
	ControllerCommand_ReplaceFailedNode                                  // Replace a non-responding node with another. The node must be in the controller's list of failed nodes for this command to succeed.
	ControllerCommand_TransferPrimaryRole                                // Make a different controller the primary.
	ControllerCommand_RequestNetworkUpdate                               // Request network information from the SUC/SIS.
	ControllerCommand_RequestNodeNeighborUpdate                          // Get a node to rebuild its neighbour list.  This method also does RequestNodeNeighbors
	ControllerCommand_AssignReturnRoute                                  // Assign a network return routes to a device.
	ControllerCommand_DeleteAllReturnRoutes                              // Delete all return routes from a device.
	ControllerCommand_SendNodeInformation                                // Send a node information frame
	ControllerCommand_ReplicationSend                                    // Send information from primary to secondary
	ControllerCommand_CreateButton                                       // Create an id that tracks handheld button presses
	ControllerCommand_DeleteButton                                       // Delete id that tracks handheld button presses
)

// ControllerState is reported via the callback handler passed into the BeginControllerCommand method.
type ControllerState uint8

const (
	ControllerState_Normal     ControllerState = iota // No command in progress.
	ControllerState_Starting                          // The command is starting.
	ControllerState_Cancel                            // The command was canceled.
	ControllerState_Error                             // Command invocation had error(s) and was aborted
	ControllerState_Waiting                           // Controller is waiting for a user action.
	ControllerState_Sleeping                          // Controller command is on a sleep queue wait for device.
	ControllerState_InProgress                        // The controller is communicating with the other device to carry out the command.
	ControllerState_Completed                         // The command has completed successfully.
	ControllerState_Failed                            // The command has failed.
	ControllerState_NodeOK                            // Used only with ControllerCommand_HasNodeFailed to indicate that the controller thinks the node is OK.
	ControllerState_NodeFailed                        // Used only with ControllerCommand_HasNodeFailed to indicate that the controller thinks the node has failed.
)

// ControllerError provides some more information about controller failures.
type ControllerError uint8

const (
	ControllerError_None           ControllerError = iota
	ControllerError_ButtonNotFound                 // Button
	ControllerError_NodeNotFound                   // Button
	ControllerError_NotBridge                      // Button
	ControllerError_NotSUC                         // CreateNewPrimary
	ControllerError_NotSecondary                   // CreateNewPrimary
	ControllerError_NotPrimary                     // RemoveFailedNode, AddNodeToNetwork
	ControllerError_IsPrimary                      // ReceiveConfiguration
	ControllerError_NotFound                       // RemoveFailedNode
	ControllerError_Busy                           // RemoveFailedNode, RequestNetworkUpdate
	ControllerError_Failed                         // RemoveFailedNode, RequestNetworkUpdate
	ControllerError_Disabled                       // RequestNetworkUpdate error
	ControllerError_Overflow                       // RequestNetworkUpdate error
)
