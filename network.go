package goopenzwave

// #cgo LDFLAGS: -lopenzwave -L/usr/local/lib
// #cgo CPPFLAGS: -I/usr/local/include -I/usr/local/include/openzwave
// #include "manager.h"
// #include <stdlib.h>
import "C"

// TestNetworkNode tests the network node.
//
// Sends a series of messages to a network node for testing network reliability.
func TestNetworkNode(homeID uint32, nodeID uint8, count uint32) {
	C.manager_testNetworkNode(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint32_t(count))
}

// TestNetwork tests the network.
//
// Sends a series of messages to every node on the network for testing network
// reliability.
func TestNetwork(homeID uint32, count uint32) {
	C.manager_testNetwork(cmanager, C.uint32_t(homeID), C.uint32_t(count))
}

// HealNetworkNode heals a network node by requesting that the node rediscovers
// their neighbors.
//
// Sends a ControllerCommand_RequestNodeNeighborUpdate to the node.
func HealNetworkNode(homeID uint32, nodeID uint8, doRR bool) {
	C.manager_healNetworkNode(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID), C.bool(doRR))
}

// HealNetwork heals a network by requesting node's rediscover their neighbors.
//
// Sends a ControllerCommand_RequestNodeNeighborUpdate to every node. Can take a
// while on larger networks.
func HealNetwork(homeID uint32, doRR bool) {
	C.manager_healNetwork(cmanager, C.uint32_t(homeID), C.bool(doRR))
}

// AddNode starts the Inclusion Process to add a Node to the Network. It will
// return true if the command was sent to the controller successfully.
//
// The Status of the Node Inclusion is communicated via Notifications.
// Specifically, you should monitor ControllerCommand Notifications.
func AddNode(homeID uint32, doSecurity bool) bool {
	return bool(C.manager_addNode(cmanager, C.uint32_t(homeID), C.bool(doSecurity)))
}

// RemoveNode removes a Device from the Z-Wave Network. It will return true if
// the command was sent to the controller successfully.
//
// The Status of the Node Removal is communicated via Notifications.
// Specifically, you should monitor ControllerCommand Notifications.
func RemoveNode(homeID uint32) bool {
	return bool(C.manager_removeNode(cmanager, C.uint32_t(homeID)))
}

// RemoveFailedNode removes a Failed Device from the Z-Wave Network. It will
// return true if the command was sent to the controller successfully.
//
// This Command will remove a failed node from the network. The Node should be
// on the Controllers Failed Node List, otherwise this command will fail. You
// can use the HasNodeFailed function below to test if the Controller believes
// the Node has Failed. The Status of the Node Removal is communicated via
// Notifications. Specifically, you should monitor ControllerCommand
// Notifications.
func RemoveFailedNode(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_removeFailedNode(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// HasNodeFailed checks if the Controller Believes a Node has Failed. The result
// is then communicated via a Notification. It will return true if the command
// was sent to the controller successfully.
//
// This is different from the IsNodeFailed call in that we test the Controllers
// Failed Node List, whereas the IsNodeFailed is testing our list of Failed
// Nodes, which might be different. The Results will be communicated via
// Notifications. Specifically, you should monitor the ControllerCommand
// notifications.
func HasNodeFailed(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_hasNodeFailed(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// RequestNodeNeighborUpdate will ask a Node to update its Neighbor Tables. It
// will return true if the command was sent to the controller successfully.
//
// This command will ask a Node to update its Neighbor Tables.
func RequestNodeNeighborUpdate(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_requestNodeNeighborUpdate(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// AssignReturnRoute will ask a Node to update its update its Return Route to
// the Controller. It will return true if the command was sent to the controller
// successfully.
func AssignReturnRoute(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_assignReturnRoute(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// DeleteAllReturnRoutes will ask a Node to delete all Return Routes. It will
// return true if the command was sent to the controller successfully.
//
// This command will ask a Node to delete all its return routes, and will
// rediscover when needed.
func DeleteAllReturnRoutes(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_deleteAllReturnRoutes(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// SendNodeInformation sends a NIF frame from the Controller to a Node. It will
// return true if the command was sent to the controller successfully.
//
// This command send a NIF frame from the Controller to a Node.
func SendNodeInformation(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_sendNodeInformation(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// CreateNewPrimary will create a new primary controller when old primary fails.
// Requires SUC. It will return true if the command was sent to the controller
// successfully.
//
// This command Creates a new Primary Controller when the Old Primary has
// Failed.
//
// Requires a SUC on the network to function.
func CreateNewPrimary(homeID uint32) bool {
	return bool(C.manager_createNewPrimary(cmanager, C.uint32_t(homeID)))
}

// ReceiveConfiguration will receive network configuration information from
// the primary controller. Requires secondary. This command prepares the
// controller to recieve Network Configuration from a Secondary Controller. It
// will return true if the command was sent to the controller successfully.
func ReceiveConfiguration(homeID uint32) bool {
	return bool(C.manager_receiveConfiguration(cmanager, C.uint32_t(homeID)))
}

// ReplaceFailedNode will replace a failed device with another. If the node is
// not in the controller's failed nodes list, or the node responds, this command
// will fail. You can check if a Node is in the Controllers Failed node list by
// using the HasNodeFailed method. It will return true if the command was sent
// to the controller successfully.
func ReplaceFailedNode(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_replaceFailedNode(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// TransferPrimaryRole adds a new controller to the network and make it the
// primary. The existing primary will become a secondary controller. It will
// return true if the command was sent to the controller successfully.
func TransferPrimaryRole(homeID uint32) bool {
	return bool(C.manager_transferPrimaryRole(cmanager, C.uint32_t(homeID)))
}

// RequestNetworkUpdate updates the controller with network information from the
// SUC/SIS. It will return true if the command was sent to the controller
// successfully.
func RequestNetworkUpdate(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_requestNetworkUpdate(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// ReplicationSend sends information from primary to secondary. It will return
// true if the command was sent to the controller successfully.
func ReplicationSend(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_replicationSend(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// CreateButton create a handheld button id.  It will return true if the command
// was sent to the controller successfully.
func CreateButton(homeID uint32, nodeID uint8, buttonID uint8) bool {
	return bool(C.manager_createButton(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint8_t(buttonID)))
}

// DeleteButton deletes a handheld button id. It will return true if the command
// was sent to the controller successfully.
func DeleteButton(homeID uint32, nodeID uint8, buttonID uint8) bool {
	return bool(C.manager_deleteButton(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint8_t(buttonID)))
}
