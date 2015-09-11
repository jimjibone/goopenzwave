package goopenzwave

// #cgo LDFLAGS: -lopenzwave -L/usr/local/lib
// #cgo CPPFLAGS: -I/usr/local/include -I/usr/local/include/openzwave
// #include "manager.h"
// #include "notification.h"
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

type Manager struct {
	cptr          C.manager_t
	Notifications chan *Notification
}

type ManagerVersion struct {
	Major int
	Minor int
}

//
// Construction.
//

// CreateManager Creates the Manager singleton object. The Manager provides the public interface to OpenZWave, exposing all the functionality required to add Z-Wave support to an application. There can be only one Manager in an OpenZWave application. An Options object must be created and Locked first, otherwise the call to Manager::Create will fail. Once the Manager has been created, call AddWatcher to install a notification callback handler, and then call the AddDriver method for each attached PC Z-Wave controller in turn.
func CreateManager() *Manager {
	m := &Manager{}
	m.cptr = C.manager_create()
	m.Notifications = make(chan *Notification, 10)
	return m
}

// GetManager Gets a pointer to the Manager object.
func GetManager() *Manager {
	m := &Manager{}
	m.cptr = C.manager_get()
	return m
}

// DestroyManager Deletes the Manager and cleans up any associated objects.
func DestroyManager() {
	C.manager_destroy()
}

// GetManagerVersionAsString Get the Version Number of OZW as a string.
func GetManagerVersionAsString() string {
	cString := C.manager_getVersionAsString()
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetManagerVersionLongAsString Get the Version Number including Git commit of OZW as a string.
func GetManagerVersionLongAsString() string {
	cString := C.manager_getVersionLongAsString()
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetManagerVersion Get the Version Number as the Version Struct (Only Major/Minor returned).
func GetManagerVersion() ManagerVersion {
	var cMajor C.uint16_t
	var cMinor C.uint16_t
	C.manager_getVersion(&cMajor, &cMinor)
	return ManagerVersion{
		Major: int(cMajor),
		Minor: int(cMinor),
	}
}

//
// Configuration.
//

// WriteConfig For saving the Z-Wave network configuration so that the entire network does not need to be polled every time the application starts.
func (m *Manager) WriteConfig(homeID uint32) {
	C.manager_writeConfig(m.cptr, C.uint32_t(homeID))
}

// GetOptions Gets a pointer to the locked Options object.
func (m *Manager) GetOptions() *Options {
	return &Options{
		options: C.manager_getOptions(m.cptr),
	}
}

//
// Drivers.
//

// AddDriver Creates a new driver for a Z-Wave controller. This method creates a Driver object for handling communications with a single Z-Wave controller. In the background, the driver first tries to read configuration data saved during a previous run. It then queries the controller directly for any missing information, and a refresh of the list of nodes that it controls. Once this information has been received, a DriverReady notification callback is sent, containing the Home ID of the controller. This Home ID is required by most of the OpenZWave Manager class methods.
func (m *Manager) AddDriver(controllerPath string) error {
	cControllerPath := C.CString(controllerPath)
	result := bool(C.manager_addDriver(m.cptr, cControllerPath))
	C.free(unsafe.Pointer(cControllerPath))
	return result
}

// RemoveDriver Removes the driver for a Z-Wave controller, and closes the controller. Drivers do not need to be explicitly removed before calling Destroy - this is handled automatically.
func (m *Manager) RemoveDriver(controllerPath string) error {
	cControllerPath := C.CString(controllerPath)
	result := bool(C.manager_removeDriver(m.cptr, cControllerPath))
	C.free(unsafe.Pointer(cControllerPath))
	return result
}

// GetControllerNodeID Get the node ID of the Z-Wave controller.
func (m *Manager) GetControllerNodeID(homeID uint32) uint8 {
	return uint8(C.manager_getControllerNodeId(m.cptr, C.uint32_t(homeID)))
}

// GetSUCNodeID Get the node ID of the Static Update Controller.
func (m *Manager) GetSUCNodeID(homeID uint32) uint8 {
	return uint8(C.manager_getSUCNodeId(m.cptr, C.uint32_t(homeID)))
}

// IsPrimaryController Query if the controller is a primary controller. The primary controller is the main device used to configure and control a Z-Wave network. There can only be one primary controller - all other controllers are secondary controllers.
func (m *Manager) IsPrimaryController(homeID uint32) bool {
	return bool(C.manager_isPrimaryController(m.cptr, C.uint32_t(homeID)))
}

// IsStaticUpdateController Query if the controller is a static update controller. A Static Update Controller (SUC) is a controller that must never be moved in normal operation and which can be used by other nodes to receive information about network changes.
func (m *Manager) IsStaticUpdateController(homeID uint32) bool {
	return bool(C.manager_isStaticUpdateController(m.cptr, C.uint32_t(homeID)))
}

// IsBridgeController Query if the controller is using the bridge controller library. A bridge controller is able to create virtual nodes that can be associated with other controllers to enable events to be passed on.
func (m *Manager) IsBridgeController(homeID uint32) bool {
	return bool(C.manager_isBridgeController(m.cptr, C.uint32_t(homeID)))
}

// GetLibraryVersion Get the version of the Z-Wave API library used by a controller.
func (m *Manager) GetLibraryVersion(homeID uint32) string {
	cString := C.manager_getLibraryVersion(m.cptr, C.uint32_t(homeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetLibraryTypeName Get a string containing the Z-Wave API library type used
// by a controller. The possible library types are:
//
// - Static Controller
// - Controller
// - Enhanced Slave
// - Slave
// - Installer
// - Routing Slave
// - Bridge Controller
// - Device Under Test
//
// The controller should never return a slave library type. For a more efficient
// test of whether a controller is a Bridge Controller, use the
// IsBridgeController method.
func (m *Manager) GetLibraryTypeName(homeID uint32) string {
	cString := C.manager_getLibraryTypeName(m.cptr, C.uint32_t(homeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetSendQueueCount Get count of messages in the outgoing send queue.
func (m *Manager) GetSendQueueCount(homeID uint32) int32 {
	return int32(C.manager_getSendQueueCount(m.cptr, C.uint32_t(homeID)))
}

// LogDriverStatistics Send current driver statistics to the log file.
func (m *Manager) LogDriverStatistics(homeID uint32) {
	C.manager_logDriverStatistics(m.cptr, C.uint32_t(homeID))
}

// GetControllerInterfaceType Obtain controller interface type.
//TODO func (m *Manager) GetControllerInterfaceType(homeID uint32) ...

// GetControllerPath Obtain controller interface path.
func (m *Manager) GetControllerPath(homeID uint32) string {
	cString := C.manager_getControllerPath(m.cptr, C.uint32_t(homeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

//
// Polling Z-Wave devices.
//

// GetPollInterval Get the time period between polls of a node's state.
func (m *Manager) GetPollInterval() int32 {
	return int32(C.manager_getPollInterval(m.cptr))
}

// SetPollInterval Set the time period between polls of a node's state. Due to patent concerns, some devices do not report state changes automatically to the controller. These devices need to have their state polled at regular intervals. The length of the interval is the same for all devices. To even out the Z-Wave network traffic generated by polling, OpenZWave divides the polling interval by the number of devices that have polling enabled, and polls each in turn. It is recommended that if possible, the interval should not be set shorter than the number of polled devices in seconds (so that the network does not have to cope with more than one poll per second).
func (m *Manager) SetPollInterval(milliseconds int32, intervalBetweenPolls bool) {
	C.manager_setPollInterval(m.cptr, C.int32_t(milliseconds), C.bool(intervalBetweenPolls))
}

// EnablePoll Enable the polling of a device's state.
func (m *Manager) EnablePoll(valueid *ValueID, intensity uint8) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_enablePoll(m.cptr, cValueid, C.uint8_t(intensity)))
}

// DisablePoll Disable the polling of a device's state.
func (m *Manager) DisablePoll(valueid *ValueID) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_disablePoll(m.cptr, cValueid))
}

// IsPolled Determine the polling of a device's state.
func (m *Manager) IsPolled(valueid *ValueID) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_isPolled(m.cptr, cValueid))
}

// SetPollIntensity Set the frequency of polling (0=none, 1=every time through the list, 2-every other time, etc).
func (m *Manager) SetPollIntensity(valueid *ValueID, intensity uint8) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	C.manager_setPollIntensity(m.cptr, cValueid, C.uint8_t(intensity))
}

// GetPollIntensity Get the polling intensity of a device's state.
func (m *Manager) GetPollIntensity(valueid *ValueID) uint8 {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return uint8(C.manager_getPollIntensity(m.cptr, cValueid))
}

//
// Node information.
//

// RefreshNodeInfo Trigger the fetching of fixed data about a node. Causes the node's data to be obtained from the Z-Wave network in the same way as if it had just been added. This method would normally be called automatically by OpenZWave, but if you know that a node has been changed, calling this method will force a refresh of all of the data held by the library. This can be especially useful for devices that were asleep when the application was first run. This is the same as the query state starting from the beginning.
func (m *Manager) RefreshNodeInfo(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_refreshNodeInfo(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// RequestNodeState Trigger the fetching of dynamic value data for a node. Causes the node's values to be requested from the Z-Wave network. This is the same as the query state starting from the associations state.
func (m *Manager) RequestNodeState(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_requestNodeState(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// RequestNodeDynamic Trigger the fetching of just the dynamic value data for a node. Causes the node's values to be requested from the Z-Wave network. This is the same as the query state starting from the dynamic state.
func (m *Manager) RequestNodeDynamic(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_requestNodeDynamic(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// IsNodeListeningDevice Get whether the node is a listening device that does not go to sleep.
func (m *Manager) IsNodeListeningDevice(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_isNodeListeningDevice(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// IsNodeFrequentListeningDevice Get whether the node is a frequent listening device that goes to sleep but can be woken up by a beam. Useful to determine node and controller consistency.
func (m *Manager) IsNodeFrequentListeningDevice(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_isNodeFrequentListeningDevice(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// IsNodeBeamingDevice Get whether the node is a beam capable device.
func (m *Manager) IsNodeBeamingDevice(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_isNodeBeamingDevice(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// IsNodeRoutingDevice Get whether the node is a routing device that passes messages to other nodes.
func (m *Manager) IsNodeRoutingDevice(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_isNodeRoutingDevice(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// IsNodeSecurityDevice Get the security attribute for a node. True if node supports security features.
func (m *Manager) IsNodeSecurityDevice(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_isNodeSecurityDevice(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeMaxBaudRate Get the maximum baud rate of a node's communications.
func (m *Manager) GetNodeMaxBaudRate(homeID uint32, nodeID uint8) uint32 {
	return uint32(C.manager_getNodeMaxBaudRate(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeVersion Get the version number of a node.
func (m *Manager) GetNodeVersion(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_getNodeVersion(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeSecurity Get the security byte of a node.
func (m *Manager) GetNodeSecurity(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_getNodeSecurity(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// IsNodeZWavePlus Is this a ZWave+ Supported Node?
func (m *Manager) IsNodeZWavePlus(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_isNodeZWavePlus(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeBasicType Get the basic type of a node.
func (m *Manager) GetNodeBasicType(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_getNodeBasic(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeGenericType Get the generic type of a node.
func (m *Manager) GetNodeGenericType(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_getNodeGeneric(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeSpecificType Get the specific type of a node.
func (m *Manager) GetNodeSpecificType(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_getNodeSpecific(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeType Get a human-readable label describing the node The label is taken from the Z-Wave specific, generic or basic type, depending on which of those values are specified by the node.
func (m *Manager) GetNodeType(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeType(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetNodeNeighbours Get the bitmap of this node's neighbors.
//TODO func (m *Manager) GetNodeNeighbours(homeID uint32, nodeID uint8) (uint32, uint8 nodeNeighbors)

// GetNodeManufacturerName Get the manufacturer name of a device The manufacturer name would normally be handled by the Manufacturer Specific commmand class, taking the manufacturer ID reported by the device and using it to look up the name from the manufacturer_specific.xml file in the OpenZWave config folder. However, there are some devices that do not support the command class, so to enable the user to manually set the name, it is stored with the node data and accessed via this method rather than being reported via a command class Value object.
func (m *Manager) GetNodeManufacturerName(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeManufacturerName(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetNodeProductName Get the product name of a device The product name would normally be handled by the Manufacturer Specific commmand class, taking the product Type and ID reported by the device and using it to look up the name from the manufacturer_specific.xml file in the OpenZWave config folder. However, there are some devices that do not support the command class, so to enable the user to manually set the name, it is stored with the node data and accessed via this method rather than being reported via a command class Value object.
func (m *Manager) GetNodeProductName(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeProductName(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetNodeName Get the name of a node The node name is a user-editable label for the node that would normally be handled by the Node Naming commmand class, but many devices do not support it. So that a node can always be named, OpenZWave stores it with the node data, and provides access through this method and SetNodeName, rather than reporting it via a command class Value object. The maximum length of a node name is 16 characters.
func (m *Manager) GetNodeName(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeName(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetNodeLocation Get the location of a node The node location is a user-editable string that would normally be handled by the Node Naming commmand class, but many devices do not support it. So that a node can always report its location, OpenZWave stores it with the node data, and provides access through this method and SetNodeLocation, rather than reporting it via a command class Value object.
func (m *Manager) GetNodeLocation(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeLocation(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetNodeManufacturerID Get the manufacturer ID of a device The manufacturer ID is a four digit hex code and would normally be handled by the Manufacturer Specific commmand class, but not all devices support it. Although the value reported by this method will be an empty string if the command class is not supported and cannot be set by the user, the manufacturer ID is still stored with the node data (rather than being reported via a command class Value object) to retain a consistent approach with the other manufacturer specific data.
func (m *Manager) GetNodeManufacturerID(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeManufacturerId(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetNodeProductType Get the product type of a device The product type is a four digit hex code and would normally be handled by the Manufacturer Specific commmand class, but not all devices support it. Although the value reported by this method will be an empty string if the command class is not supported and cannot be set by the user, the product type is still stored with the node data (rather than being reported via a command class Value object) to retain a consistent approach with the other manufacturer specific data.
func (m *Manager) GetNodeProductType(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeProductType(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetNodeProductID Get the product ID of a device The product ID is a four digit hex code and would normally be handled by the Manufacturer Specific commmand class, but not all devices support it. Although the value reported by this method will be an empty string if the command class is not supported and cannot be set by the user, the product ID is still stored with the node data (rather than being reported via a command class Value object) to retain a consistent approach with the other manufacturer specific data.
func (m *Manager) GetNodeProductID(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeProductId(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// SetNodeManufacturerName Set the manufacturer name of a device The manufacturer name would normally be handled by the Manufacturer Specific commmand class, taking the manufacturer ID reported by the device and using it to look up the name from the manufacturer_specific.xml file in the OpenZWave config folder. However, there are some devices that do not support the command class, so to enable the user to manually set the name, it is stored with the node data and accessed via this method rather than being reported via a command class Value object.
func (m *Manager) SetNodeManufacturerName(homeID uint32, nodeID uint8, manufacturerName string) {
	cString := C.CString(manufacturerName)
	C.manager_setNodeManufacturerName(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID), cString)
	C.free(unsafe.Pointer(cString))
}

// SetNodeProductName Set the product name of a device The product name would normally be handled by the Manufacturer Specific commmand class, taking the product Type and ID reported by the device and using it to look up the name from the manufacturer_specific.xml file in the OpenZWave config folder. However, there are some devices that do not support the command class, so to enable the user to manually set the name, it is stored with the node data and accessed via this method rather than being reported via a command class Value object.
func (m *Manager) SetNodeProductName(homeID uint32, nodeID uint8, productName string) {
	cString := C.CString(productName)
	C.manager_setNodeProductName(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID), cString)
	C.free(unsafe.Pointer(cString))
}

// SetNodeName Set the name of a node The node name is a user-editable label for the node that would normally be handled by the Node Naming commmand class, but many devices do not support it. So that a node can always be named, OpenZWave stores it with the node data, and provides access through this method and GetNodeName, rather than reporting it via a command class Value object. If the device does support the Node Naming command class, the new name will be sent to the node. The maximum length of a node name is 16 characters.
func (m *Manager) SetNodeName(homeID uint32, nodeID uint8, nodeName string) {
	cString := C.CString(nodeName)
	C.manager_setNodeName(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID), cString)
	C.free(unsafe.Pointer(cString))
}

// SetNodeLocation Set the location of a node The node location is a user-editable string that would normally be handled by the Node Naming commmand class, but many devices do not support it. So that a node can always report its location, OpenZWave stores it with the node data, and provides access through this method and GetNodeLocation, rather than reporting it via a command class Value object. If the device does support the Node Naming command class, the new location will be sent to the node.
func (m *Manager) SetNodeLocation(homeID uint32, nodeID uint8, location string) {
	cString := C.CString(location)
	C.manager_setNodeLocation(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID), cString)
	C.free(unsafe.Pointer(cString))
}

// SetNodeOn Turns a node on This is a helper method to simplify basic control of a node. It is the equivalent of changing the level reported by the node's Basic command class to 255, and will generate a ValueChanged notification from that class. This command will turn on the device at its last known level, if supported by the device, otherwise it will turn it on at 100%.
func (m *Manager) SetNodeOn(homeID uint32, nodeID uint8) {
	C.manager_setNodeOn(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID))
}

// SetNodeOff Turns a node off This is a helper method to simplify basic control of a node. It is the equivalent of changing the level reported by the node's Basic command class to zero, and will generate a ValueChanged notification from that class.
func (m *Manager) SetNodeOff(homeID uint32, nodeID uint8) {
	C.manager_setNodeOff(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID))
}

// SetNodeLevel Sets the basic level of a node This is a helper method to simplify basic control of a node. It is the equivalent of changing the value reported by the node's Basic command class and will generate a ValueChanged notification from that class.
func (m *Manager) SetNodeLevel(homeID uint32, nodeID uint8, level uint8) {
	C.manager_setNodeLevel(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint8_t(level))
}

// IsNodeInfoReceived Get whether the node information has been received.
func (m *Manager) IsNodeInfoReceived(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_isNodeInfoReceived(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeClassInformation Get whether the node has the defined class available or not.
func (m *Manager) GetNodeClassInformation(homeID uint32, nodeID uint8, commandClassID uint8) (bool, string, uint8) {
	cClassName := C.string_emptyString()
	var cClassVersion C.uint8_t
	result := bool(C.manager_getNodeClassInformation(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint8_t(commandClassID), cClassName, &cClassVersion))
	goClassName := C.GoString(cClassName.data)
	goClassVersion := uint8(cClassVersion)
	C.string_freeString(cClassName)
	return result, goClassName, goClassVersion
}

// IsNodeAwake Get whether the node is awake or asleep.
func (m *Manager) IsNodeAwake(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_isNodeAwake(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// IsNodeFailed Get whether the node is working or has failed.
func (m *Manager) IsNodeFailed(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_isNodeFailed(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeQueryStage Get whether the node's query stage as a string.
func (m *Manager) GetNodeQueryStage(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeQueryStage(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetNodeDeviceType Get the node device type as reported in the Z-Wave+ Info report.
func (m *Manager) GetNodeDeviceType(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_getNodeDeviceType(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeDeviceTypeString Get the node device type as reported in the Z-Wave+ Info report.
func (m *Manager) GetNodeDeviceTypeString(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeDeviceTypeString(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetNodeRole Get the node role as reported in the Z-Wave+ Info report.
func (m *Manager) GetNodeRole(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_getNodeRole(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeRoleString Get the node role as reported in the Z-Wave+ Info report.
func (m *Manager) GetNodeRoleString(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeRoleString(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetNodePlusType Get the node PlusType as reported in the Z-Wave+ Info report.
func (m *Manager) GetNodePlusType(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_getNodePlusType(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodePlusTypeString Get the node PlusType as reported in the Z-Wave+ Info report.
func (m *Manager) GetNodePlusTypeString(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodePlusTypeString(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

//
// Values.
//

// GetValueLabel Gets the user-friendly label for the value.
func (m *Manager) GetValueLabel(valueid *ValueID) string {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.manager_getValueLabel(m.cptr, cValueid)
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// SetValueLabel Sets the user-friendly label for the value.
func (m *Manager) SetValueLabel(valueid *ValueID, value string) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.CString(value)
	C.manager_setValueLabel(m.cptr, cValueid, cString)
	C.free(unsafe.Pointer(cString))
}

// GetValueUnits Gets the units that the value is measured in.
func (m *Manager) GetValueUnits(valueid *ValueID) string {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.manager_getValueUnits(m.cptr, cValueid)
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// SetValueUnits Sets the units that the value is measured in.
func (m *Manager) SetValueUnits(valueid *ValueID, value string) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.CString(value)
	C.manager_setValueUnits(m.cptr, cValueid, cString)
	C.free(unsafe.Pointer(cString))
}

// GetValueHelp Gets a help string describing the value's purpose and usage.
func (m *Manager) GetValueHelp(valueid *ValueID) string {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.manager_getValueHelp(m.cptr, cValueid)
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// SetValueHelp Sets a help string describing the value's purpose and usage.
func (m *Manager) SetValueHelp(valueid *ValueID, value string) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.CString(value)
	C.manager_setValueHelp(m.cptr, cValueid, cString)
	C.free(unsafe.Pointer(cString))
}

// GetValueMin Gets the minimum that this value may contain.
func (m *Manager) GetValueMin(valueid *ValueID) int32 {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return int32(C.manager_getValueMin(m.cptr, cValueid))
}

// GetValueMax Gets the maximum that this value may contain.
func (m *Manager) GetValueMax(valueid *ValueID) int32 {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return int32(C.manager_getValueMax(m.cptr, cValueid))
}

// IsValueReadOnly Test whether the value is read-only.
func (m *Manager) IsValueReadOnly(valueid *ValueID) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_isValueReadOnly(m.cptr, cValueid))
}

// IsValueWriteOnly Test whether the value is write-only.
func (m *Manager) IsValueWriteOnly(valueid *ValueID) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_isValueWriteOnly(m.cptr, cValueid))
}

// IsValueSet Test whether the value has been set.
func (m *Manager) IsValueSet(valueid *ValueID) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_isValueSet(m.cptr, cValueid))
}

// IsValuePolled Test whether the value is currently being polled.
func (m *Manager) IsValuePolled(valueid *ValueID) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_isValuePolled(m.cptr, cValueid))
}

// GetValueAsBool Gets a value as a bool.
func (m *Manager) GetValueAsBool(valueid *ValueID) (bool, bool) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cBool C.bool
	result := bool(C.manager_getValueAsBool(m.cptr, cValueid, &cBool))
	return result, bool(cBool)
}

// GetValueAsByte Gets a value as an 8-bit unsigned integer.
func (m *Manager) GetValueAsByte(valueid *ValueID) (bool, byte) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cByte C.uint8_t
	result := bool(C.manager_getValueAsByte(m.cptr, cValueid, &cByte))
	return result, byte(cByte)
}

// GetValueAsFloat Gets a value as a float.
func (m *Manager) GetValueAsFloat(valueid *ValueID) (bool, float32) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cFloat C.float
	result := bool(C.manager_getValueAsFloat(m.cptr, cValueid, &cFloat))
	return result, float32(cFloat)
}

// GetValueAsInt Gets a value as a 32-bit signed integer.
func (m *Manager) GetValueAsInt(valueid *ValueID) (bool, int32) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cInt C.int32_t
	result := bool(C.manager_getValueAsInt(m.cptr, cValueid, &cInt))
	return result, int32(cInt)
}

// GetValueAsShort Gets a value as a 16-bit signed integer.
func (m *Manager) GetValueAsShort(valueid *ValueID) (bool, int16) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cShort C.int16_t
	result := bool(C.manager_getValueAsShort(m.cptr, cValueid, &cShort))
	return result, int16(cShort)
}

// GetValueAsString Gets a value as a string. Creates a string representation of a value, regardless of type.
func (m *Manager) GetValueAsString(valueid *ValueID) (bool, string) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.string_emptyString()
	result := bool(C.manager_getValueAsString(m.cptr, cValueid, cString))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return result, goString
}

// GetValueAsRaw Gets a value as a collection of bytes.
func (m *Manager) GetValueAsRaw(valueid *ValueID) (bool, []byte) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cBytes := C.string_emptyBytes()
	result := bool(C.manager_getValueAsRaw(m.cptr, cValueid, cBytes))
	goBytes := make([]byte, int(cBytes.length))
	for i := 0; i < int(cBytes.length); i++ {
		goBytes[i] = byte(C.string_byteAt(cBytes, C.size_t(i)))
	}
	return result, goBytes
}

// GetValueListSelectionAsString Gets the selected item from a list (as a string).
func (m *Manager) GetValueListSelectionAsString(valueid *ValueID) (bool, string) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.string_emptyString()
	result := bool(C.manager_getValueListSelectionAsString(m.cptr, cValueid, cString))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return result, goString
}

// GetValueListSelectionAsInt32 Gets the selected item from a list (as an integer).
func (m *Manager) GetValueListSelectionAsInt32(valueid *ValueID) (bool, int32) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cInt C.int32_t
	result := bool(C.manager_getValueListSelectionAsInt32(m.cptr, cValueid, &cInt))
	return result, int32(cInt)
}

// GetValueListItems Gets the list of items from a list value.
func (m *Manager) GetValueListItems(valueid *ValueID) (bool, []string) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cStringList := C.string_emptyStringList()
	result := bool(C.manager_getValueListItems(m.cptr, cValueid, cStringList))
	goStringList := make([]string, int(cStringList.length))
	for i := 0; i < int(cStringList.length); i++ {
		cString := C.string_stringAt(cStringList, C.size_t(i))
		goStringList[i] = C.GoString(cString.data)
	}
	C.string_freeStringList(cStringList)
	return result, goStringList
}

// GetValueFloatPrecision Gets a float value's precision.
func (m *Manager) GetValueFloatPrecision(valueid *ValueID) (bool, uint8) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cPrecision C.uint8_t
	result := bool(C.manager_getValueFloatPrecision(m.cptr, cValueid, &cPrecision))
	return result, uint8(cPrecision)
}

// SetValueBool Sets the state of a bool. Due to the possibility of a device being asleep, the command is assumed to suceed, and the value held by the node is updated directly. This will be reverted by a future status message from the device if the Z-Wave message actually failed to get through. Notification callbacks will be sent in both cases.
func (m *Manager) SetValueBool(valueid *ValueID, value bool) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setValueBool(m.cptr, cValueid, C.bool(value)))
}

// SetValueUint8 Sets the value of a byte. Due to the possibility of a device being asleep, the command is assumed to suceed, and the value held by the node is updated directly. This will be reverted by a future status message from the device if the Z-Wave message actually failed to get through. Notification callbacks will be sent in both cases.
func (m *Manager) SetValueUint8(valueid *ValueID, value uint8) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setValueUint8(m.cptr, cValueid, C.uint8_t(value)))
}

// SetValueFloat Sets the value of a decimal. It is usually better to handle decimal values using strings rather than floats, to avoid floating point accuracy issues. Due to the possibility of a device being asleep, the command is assumed to succeed, and the value held by the node is updated directly. This will be reverted by a future status message from the device if the Z-Wave message actually failed to get through. Notification callbacks will be sent in both cases.
func (m *Manager) SetValueFloat(valueid *ValueID, value float32) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setValueFloat(m.cptr, cValueid, C.float(value)))
}

// SetValueInt32 Sets the value of a 32-bit signed integer. Due to the possibility of a device being asleep, the command is assumed to suceed, and the value held by the node is updated directly. This will be reverted by a future status message from the device if the Z-Wave message actually failed to get through. Notification callbacks will be sent in both cases.
func (m *Manager) SetValueInt32(valueid *ValueID, value int32) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setValueInt32(m.cptr, cValueid, C.int32_t(value)))
}

// SetValueInt16 Sets the value of a 16-bit signed integer. Due to the possibility of a device being asleep, the command is assumed to suceed, and the value held by the node is updated directly. This will be reverted by a future status message from the device if the Z-Wave message actually failed to get through. Notification callbacks will be sent in both cases.
func (m *Manager) SetValueInt16(valueid *ValueID, value int16) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setValueInt16(m.cptr, cValueid, C.int16_t(value)))
}

// SetValueBytes Sets the value of a collection of bytes. Due to the possibility of a device being asleep, the command is assumed to suceed, and the value held by the node is updated directly. This will be reverted by a future status message from the device if the Z-Wave message actually failed to get through. Notification callbacks will be sent in both cases.
func (m *Manager) SetValueBytes(valueid *ValueID, value []byte) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cBytes := C.string_emptyBytes()
	C.string_initBytes(cBytes, C.size_t(len(value)))
	for i := range value {
		C.string_setByteAt(cBytes, C.uint8_t(value[i]), C.size_t(i))
	}
	result := bool(C.manager_setValueBytes(m.cptr, cValueid, cBytes))
	C.string_freeBytes(cBytes)
	return result
}

// SetValueString Sets the value from a string, regardless of type. Due to the possibility of a device being asleep, the command is assumed to suceed, and the value held by the node is updated directly. This will be reverted by a future status message from the device if the Z-Wave message actually failed to get through. Notification callbacks will be sent in both cases.
func (m *Manager) SetValueString(valueid *ValueID, value string) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.CString(value)
	result := bool(C.manager_setValueString(m.cptr, cValueid, cString))
	C.free(unsafe.Pointer(cString))
	return result
}

// SetValueListSelection Sets the selected item in a list. Due to the possibility of a device being asleep, the command is assumed to suceed, and the value held by the node is updated directly. This will be reverted by a future status message from the device if the Z-Wave message actually failed to get through. Notification callbacks will be sent in both cases.
func (m *Manager) SetValueListSelection(valueid *ValueID, selectedItem string) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.CString(selectedItem)
	result := bool(C.manager_setValueListSelection(m.cptr, cValueid, cString))
	C.free(unsafe.Pointer(cString))
	return result
}

// RefreshValue Refreshes the specified value from the Z-Wave network. A call to this function causes the library to send a message to the network to retrieve the current value of the specified ValueID (just like a poll, except only one-time, not recurring).
func (m *Manager) RefreshValue(valueid *ValueID) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_refreshValue(m.cptr, cValueid))
}

// SetChangeVerified Sets a flag indicating whether value changes noted upon a refresh should be verified. If so, the library will immediately refresh the value a second time whenever a change is observed. This helps to filter out spurious data reported occasionally by some devices.
func (m *Manager) SetChangeVerified(valueid *ValueID, verify bool) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	C.manager_setChangeVerified(m.cptr, cValueid, C.bool(verify))
}

// GetChangeVerified Determine if value changes upon a refresh should be verified. If so, the library will immediately refresh the value a second time whenever a change is observed. This helps to filter out spurious data reported occasionally by some devices.
func (m *Manager) GetChangeVerified(valueid *ValueID) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_getChangeVerified(m.cptr, cValueid))
}

// PressButton Starts an activity in a device. Since buttons are write-only values that do not report a state, no notification callbacks are sent.
func (m *Manager) PressButton(valueid *ValueID) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_pressButton(m.cptr, cValueid))
}

// ReleaseButton Stops an activity in a device. Since buttons are write-only values that do not report a state, no notification callbacks are sent.
func (m *Manager) ReleaseButton(valueid *ValueID) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_releaseButton(m.cptr, cValueid))
}

//
// Climate control schedules.
//

// GetNumSwitchPoints Get the number of switch points defined in a schedule.
func (m *Manager) GetNumSwitchPoints(valueid *ValueID) uint8 {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return uint8(C.manager_getNumSwitchPoints(m.cptr, cValueid))
}

// SetSwitchPoint Set a switch point in the schedule. Inserts a new switch point into the schedule, unless a switch point already exists at the specified time in which case that switch point is updated with the new setback value instead. A maximum of nine switch points can be set in the schedule.
func (m *Manager) SetSwitchPoint(valueid *ValueID, hours, minutes uint8, setback int8) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setSwitchPoint(m.cptr, cValueid, C.uint8_t(hours), C.uint8_t(minutes), C.int8_t(setback)))
}

// RemoveSwitchPoint Remove a switch point from the schedule. Removes the switch point at the specified time from the schedule.
func (m *Manager) RemoveSwitchPoint(valueid *ValueID, hours, minutes uint8) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_removeSwitchPoint(m.cptr, cValueid, C.uint8_t(hours), C.uint8_t(minutes)))
}

// ClearSwitchPoints Clears all switch points from the schedule.
func (m *Manager) ClearSwitchPoints(valueid *ValueID) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	C.manager_clearSwitchPoints(m.cptr, cValueid)
}

// GetSwitchPoint Gets switch point data from the schedule. Retrieves the time and setback values from a switch point in the schedule.
func (m *Manager) GetSwitchPoint(valueid *ValueID, idx uint8) (bool, uint8, uint8, int8) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cHours C.uint8_t
	var cMinutes C.uint8_t
	var cSetback C.int8_t
	result := bool(C.manager_getSwitchPoint(m.cptr, cValueid, C.uint8_t(idx), &cHours, &cMinutes, &cSetback))
	return result, uint8(cHours), uint8(cMinutes), int8(cSetback)
}

//
// Switch all.
//

// SwitchAllOn Switch all devices on. All devices that support the SwitchAll command class will be turned on.
func (m *Manager) SwitchAllOn(homeID uint32) {
	C.manager_switchAllOn(m.cptr, C.uint32_t(homeID))
}

// SwitchAllOff Switch all devices off. All devices that support the SwitchAll command class will be turned off.
func (m *Manager) SwitchAllOff(homeID uint32) {
	C.manager_switchAllOff(m.cptr, C.uint32_t(homeID))
}

//
// Configuration parameters.
//

// SetConfigParam Set the value of a configurable parameter in a device. Some devices have various parameters that can be configured to control the device behaviour. These are not reported by the device over the Z-Wave network, but can usually be found in the device's user manual. This method returns immediately, without waiting for confirmation from the device that the change has been made.
func (m *Manager) SetConfigParam(homeID uint32, nodeID uint8, param uint8, value int32, size uint8) bool {
	return bool(C.manager_setConfigParam(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint8_t(param), C.int32_t(value), C.uint8_t(size)))
}

// RequestConfigParam Request the value of a configurable parameter from a device. Some devices have various parameters that can be configured to control the device behaviour. These are not reported by the device over the Z-Wave network, but can usually be found in the device's user manual. This method requests the value of a parameter from the device, and then returns immediately, without waiting for a response. If the parameter index is valid for this device, and the device is awake, the value will eventually be reported via a ValueChanged notification callback. The ValueID reported in the callback will have an index set the same as _param and a command class set to the same value as returned by a call to Configuration::StaticGetCommandClassId.
func (m *Manager) RequestConfigParam(homeID uint32, nodeID uint8, param uint8) {
	C.manager_requestConfigParam(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint8_t(param))
}

// RequestAllConfigParam Request the values of all known configurable parameters from a device.
func (m *Manager) RequestAllConfigParam(homeID uint32, nodeID uint8) {
	C.manager_requestAllConfigParams(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID))
}

//
// Groups.
//

// GetNumGroups Gets the number of association groups reported by this node In Z-Wave, groups are numbered starting from one. For example, if a call to GetNumGroups returns 4, the _groupIdx value to use in calls to GetAssociations, AddAssociation and RemoveAssociation will be a number between 1 and 4.
func (m *Manager) GetNumGroups(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_getNumGroups(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetAssociations Gets the associations for a group. Makes a copy of the list of associated nodes in the group, and returns it in an array of uint8's. The caller is responsible for freeing the array memory with a call to delete [].
//TODO func (m *Manager) GetAssociations(homeID uint32, nodeID uint8, groupIDx uint8, uint8 **o_associations) ...
// GetAssociations Gets the associations for a group. Makes a copy of the list of associated nodes in the group, and returns it in an array of InstanceAssociation's. The caller is responsible for freeing the array memory with a call to delete [].
//TODO func (m *Manager) GetAssociations(homeID uint32, nodeID uint8, groupIDx uint8, InstanceAssociation **o_associations) ...

// GetMaxAssociations Gets the maximum number of associations for a group.
func (m *Manager) GetMaxAssociations(homeID uint32, nodeID uint8, groupIDx uint8) uint8 {
	return uint8(C.manager_getMaxAssociations(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint8_t(groupIDx)))
}

// GetGroupLabel Returns a label for the particular group of a node. This label is populated by the device specific configuration files.
func (m *Manager) GetGroupLabel(homeID uint32, nodeID uint8, groupIDx uint8) string {
	cString := C.manager_getGroupLabel(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint8_t(groupIDx))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// AddAssociation Adds a node to an association group. Due to the possibility of a device being asleep, the command is assumed to suceed, and the association data held in this class is updated directly. This will be reverted by a future Association message from the device if the Z-Wave message actually failed to get through. Notification callbacks will be sent in both cases.
func (m *Manager) AddAssociation(homeID uint32, nodeID uint8, groupIDx uint8, targetNodeID uint8, instance uint8) {
	C.manager_addAssociation(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint8_t(groupIDx), C.uint8_t(targetNodeID), C.uint8_t(instance))
}

// RemoveAssociation Removes a node from an association group. Due to the possibility of a device being asleep, the command is assumed to suceed, and the association data held in this class is updated directly. This will be reverted by a future Association message from the device if the Z-Wave message actually failed to get through. Notification callbacks will be sent in both cases.
func (m *Manager) RemoveAssociation(homeID uint32, nodeID uint8, groupIDx uint8, targetNodeID uint8, instance uint8) {
	C.manager_removeAssociation(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint8_t(groupIDx), C.uint8_t(targetNodeID), C.uint8_t(instance))
}

//
// Notifications.
//

// goNotificationCB called by the C++ OpenZWave library when there is a new
// notification.
//export goNotificationCB
func goNotificationCB(notification C.notification_t, userdata unsafe.Pointer) {
	// This function is called by OpenZWave (via a C wrapper) when a
	// notification is available. All data must be extracted from the
	// notification object before we return as OpenZWave will delete the object.
	m := (*Manager)(userdata)

	// Convert the C notification_t to Go Notification.
	noti := buildNotification(notification)

	// Send the Notification on the channel.
	m.Notifications <- noti
}

// StartNotifications Calls the OpenZWave AddWatcher function. New notifications
// are received by this package and made available via the Notifications
// channel.
func (m *Manager) StartNotifications() error {
	themanager := unsafe.Pointer(m)
	result := C.manager_addWatcher(m.cptr, themanager)
	if result {
		return nil
	}
	return fmt.Errorf("failed to add watcher")
}

// StopNotifications Calls the OpenZWave RemoveWatcher function. This stops any
// future notifications being received.
func (m *Manager) StopNotifications() error {
	themanager := unsafe.Pointer(m)
	result := C.manager_removeWatcher(m.cptr, themanager)
	if result {
		return nil
	}
	return fmt.Errorf("failed to remove watcher")
}

//
// Controller commands.
//

// ResetController Hard Reset a PC Z-Wave Controller. Resets a controller and erases its network configuration settings. The controller becomes a primary controller ready to add devices to a new network.
func (m *Manager) ResetController(homeID uint32) {
	C.manager_resetController(m.cptr, C.uint32_t(homeID))
}

// SoftReset Soft Reset a PC Z-Wave Controller. Resets a controller without erasing its network configuration settings.
func (m *Manager) SoftReset(homeID uint32) {
	C.manager_softReset(m.cptr, C.uint32_t(homeID))
}

// CancelControllerCommand Cancels any in-progress command running on a controller.
func (m *Manager) CancelControllerCommand(homeID uint32) {
	C.manager_cancelControllerCommand(m.cptr, C.uint32_t(homeID))
}

//
// Network commands.
//

// TestNetworkNode Test network node. Sends a series of messages to a network node for testing network reliability.
func (m *Manager) TestNetworkNode(homeID uint32, nodeID uint8, count uint32) {
	C.manager_testNetworkNode(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint32_t(count))
}

// TestNetwork Test network. Sends a series of messages to every node on the network for testing network reliability.
func (m *Manager) TestNetwork(homeID uint32, count uint32) {
	C.manager_testNetwork(m.cptr, C.uint32_t(homeID), C.uint32_t(count))
}

// HealNetworkNode Heal network node by requesting the node rediscover their neighbors. Sends a ControllerCommand_RequestNodeNeighborUpdate to the node.
func (m *Manager) HealNetworkNode(homeID uint32, nodeID uint8, doRR bool) {
	C.manager_healNetworkNode(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID), C.bool(doRR))
}

// HealNetwork Heal network by requesting node's rediscover their neighbors. Sends a ControllerCommand_RequestNodeNeighborUpdate to every node. Can take a while on larger networks.
func (m *Manager) HealNetwork(homeID uint32, doRR bool) {
	C.manager_healNetwork(m.cptr, C.uint32_t(homeID), C.bool(doRR))
}

// AddNode Start the Inclusion Process to add a Node to the Network. The Status of the Node Inclusion is communicated via Notifications. Specifically, you should monitor ControllerCommand Notifications.
func (m *Manager) AddNode(homeID uint32, doSecurity bool) bool {
	return bool(C.manager_addNode(m.cptr, C.uint32_t(homeID), C.bool(doSecurity)))
}

// RemoveNode Remove a Device from the Z-Wave Network The Status of the Node Removal is communicated via Notifications. Specifically, you should monitor ControllerCommand Notifications.
func (m *Manager) RemoveNode(homeID uint32) bool {
	return bool(C.manager_removeNode(m.cptr, C.uint32_t(homeID)))
}

// RemoveFailedNode Remove a Failed Device from the Z-Wave Network This Command will remove a failed node from the network. The Node should be on the Controllers Failed Node List, otherwise this command will fail. You can use the HasNodeFailed function below to test if the Controller believes the Node has Failed. The Status of the Node Removal is communicated via Notifications. Specifically, you should monitor ControllerCommand Notifications.
func (m *Manager) RemoveFailedNode(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_removeFailedNode(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// HasNodeFailed Check if the Controller Believes a Node has Failed. This is different from the IsNodeFailed call in that we test the Controllers Failed Node List, whereas the IsNodeFailed is testing our list of Failed Nodes, which might be different. The Results will be communicated via Notifications. Specifically, you should monitor the ControllerCommand notifications.
func (m *Manager) HasNodeFailed(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_hasNodeFailed(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// RequestNodeNeighborUpdate Ask a Node to update its Neighbor Tables This command will ask a Node to update its Neighbor Tables.
func (m *Manager) RequestNodeNeighborUpdate(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_requestNodeNeighborUpdate(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// AssignReturnRoute Ask a Node to update its update its Return Route to the Controller This command will ask a Node to update its Return Route to the Controller.
func (m *Manager) AssignReturnRoute(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_assignReturnRoute(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// DeleteAllReturnRoutes Ask a Node to delete all Return Route. This command will ask a Node to delete all its return routes, and will rediscover when needed.
func (m *Manager) DeleteAllReturnRoutes(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_deleteAllReturnRoutes(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// SendNodeInformation Send a NIF frame from the Controller to a Node. This command send a NIF frame from the Controller to a Node.
func (m *Manager) SendNodeInformation(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_sendNodeInformation(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// CreateNewPrimary Create a new primary controller when old primary fails. Requires SUC. This command Creates a new Primary Controller when the Old Primary has Failed. Requires a SUC on the network to function.
func (m *Manager) CreateNewPrimary(homeID uint32) bool {
	return bool(C.manager_createNewPrimary(m.cptr, C.uint32_t(homeID)))
}

// ReceiveConfiguration Receive network configuration information from primary controller. Requires secondary. This command prepares the controller to recieve Network Configuration from a Secondary Controller.
func (m *Manager) ReceiveConfiguration(homeID uint32) bool {
	return bool(C.manager_receiveConfiguration(m.cptr, C.uint32_t(homeID)))
}

// ReplaceFailedNode Replace a failed device with another. If the node is not in the controller's failed nodes list, or the node responds, this command will fail. You can check if a Node is in the Controllers Failed node list by using the HasNodeFailed method.
func (m *Manager) ReplaceFailedNode(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_replaceFailedNode(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// TransferPrimaryRole Add a new controller to the network and make it the primary. The existing primary will become a secondary controller.
func (m *Manager) TransferPrimaryRole(homeID uint32) bool {
	return bool(C.manager_transferPrimaryRole(m.cptr, C.uint32_t(homeID)))
}

// RequestNetworkUpdate Update the controller with network information from the SUC/SIS.
func (m *Manager) RequestNetworkUpdate(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_requestNetworkUpdate(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// ReplicationSend Send information from primary to secondary.
func (m *Manager) ReplicationSend(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_replicationSend(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// CreateButton Create a handheld button id.
func (m *Manager) CreateButton(homeID uint32, nodeID uint8, buttonID uint8) bool {
	return bool(C.manager_createButton(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint8_t(buttonID)))
}

// DeleteButton Delete a handheld button id.
func (m *Manager) DeleteButton(homeID uint32, nodeID uint8, buttonID uint8) bool {
	return bool(C.manager_deleteButton(m.cptr, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint8_t(buttonID)))
}

//
// Scene commands.
//

// GetNumScenes Gets the number of scenes that have been defined.
func (m *Manager) GetNumScenes() uint8 {
	return uint8(C.manager_getNumScenes(m.cptr))
}

// GetAllScenes Gets a list of all the SceneIds.
//TODO(jimjibone) func (m *Manager) GetAllScenes(...) ...

// RemoveAllScenes Remove all the SceneIds.
func (m *Manager) RemoveAllScenes(homeID uint32) {
	C.manager_removeAllScenes(m.cptr, C.uint32_t(homeID))
}

// CreateScene Create a new Scene passing in Scene ID.
func (m *Manager) CreateScene() uint8 {
	return uint8(C.manager_createScene(m.cptr))
}

// RemoveScene Remove an existing Scene.
func (m *Manager) RemoveScene(sceneID uint8) bool {
	return bool(C.manager_removeScene(m.cptr, C.uint8_t(sceneID)))
}

// AddSceneValueBool Add a bool Value ID to an existing scene.
func (m *Manager) AddSceneValueBool(sceneID uint8, valueid *ValueID, value bool) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_addSceneValueBool(m.cptr, C.uint8_t(sceneID), cValueid, C.bool(value)))
}

// AddSceneValueUint8 Add a bool Value ID to an existing scene.
func (m *Manager) AddSceneValueUint8(sceneID uint8, valueid *ValueID, value uint8) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_addSceneValueUint8(m.cptr, C.uint8_t(sceneID), cValueid, C.uint8_t(value)))
}

// AddSceneValueFloat Add a decimal Value ID to an existing scene.
func (m *Manager) AddSceneValueFloat(sceneID uint8, valueid *ValueID, value float32) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_addSceneValueFloat(m.cptr, C.uint8_t(sceneID), cValueid, C.float(value)))
}

// AddSceneValueInt32 Add a 32-bit signed integer Value ID to an existing scene.
func (m *Manager) AddSceneValueInt32(sceneID uint8, valueid *ValueID, value int32) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_addSceneValueInt32(m.cptr, C.uint8_t(sceneID), cValueid, C.int32_t(value)))
}

// AddSceneValueInt16 Add a 16-bit signed integer Value ID to an existing scene.
func (m *Manager) AddSceneValueInt16(sceneID uint8, valueid *ValueID, value int16) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_addSceneValueInt16(m.cptr, C.uint8_t(sceneID), cValueid, C.int16_t(value)))
}

// AddSceneValueString Add a string Value ID to an existing scene.
func (m *Manager) AddSceneValueString(sceneID uint8, valueid *ValueID, value string) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.CString(value)
	result := bool(C.manager_addSceneValueString(m.cptr, C.uint8_t(sceneID), cValueid, cString))
	C.free(unsafe.Pointer(cString))
	return result
}

// AddSceneValueListSelectionString Add the selected item list Value ID to an existing scene (as a string).
func (m *Manager) AddSceneValueListSelectionString(sceneID uint8, valueid *ValueID, value string) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.CString(value)
	result := bool(C.manager_addSceneValueListSelectionString(m.cptr, C.uint8_t(sceneID), cValueid, cString))
	C.free(unsafe.Pointer(cString))
	return result
}

// AddSceneValueListSelectionInt32 Add the selected item list Value ID to an existing scene (as a integer).
func (m *Manager) AddSceneValueListSelectionInt32(sceneID uint8, valueid *ValueID, value int32) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_addSceneValueListSelectionInt32(m.cptr, C.uint8_t(sceneID), cValueid, C.int32_t(value)))
}

// RemoveSceneValue Remove the Value ID from an existing scene.
// TODO: bool RemoveSceneValue (uint8 const _sceneId, ValueID const &_valueId) ...

// SceneGetValues Retrieves the scene's list of values.
// TODO: int SceneGetValues (uint8 const _sceneId, vector< ValueID > *o_value) ...

// GetSceneValueAsBool Retrieves a scene's value as a bool.
func (m *Manager) GetSceneValueAsBool(sceneID uint8, valueid *ValueID) (bool, bool) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cBool C.bool
	result := bool(C.manager_sceneGetValueAsBool(m.cptr, C.uint8_t(sceneID), cValueid, &cBool))
	return result, bool(cBool)
}

// GetSceneValueAsByte Retrieves a scene's value as an 8-bit unsigned integer.
func (m *Manager) GetSceneValueAsByte(sceneID uint8, valueid *ValueID) (bool, byte) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cByte C.uint8_t
	result := bool(C.manager_sceneGetValueAsByte(m.cptr, C.uint8_t(sceneID), cValueid, &cByte))
	return result, byte(cByte)
}

// GetSceneValueAsFloat Retrieves a scene's value as a float.
func (m *Manager) GetSceneValueAsFloat(sceneID uint8, valueid *ValueID) (bool, float32) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cFloat C.float
	result := bool(C.manager_sceneGetValueAsFloat(m.cptr, C.uint8_t(sceneID), cValueid, &cFloat))
	return result, float32(cFloat)
}

// GetSceneValueAsInt Retrieves a scene's value as a 32-bit signed integer.
func (m *Manager) GetSceneValueAsInt(sceneID uint8, valueid *ValueID) (bool, int32) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cInt C.int32_t
	result := bool(C.manager_sceneGetValueAsInt(m.cptr, C.uint8_t(sceneID), cValueid, &cInt))
	return result, int32(cInt)
}

// GetSceneValueAsShort Retrieves a scene's value as a 16-bit signed integer.
func (m *Manager) GetSceneValueAsShort(sceneID uint8, valueid *ValueID) (bool, int16) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cShort C.int16_t
	result := bool(C.manager_sceneGetValueAsShort(m.cptr, C.uint8_t(sceneID), cValueid, &cShort))
	return result, int16(cShort)
}

// GetSceneValueAsString Retrieves a scene's value as a string.
func (m *Manager) GetSceneValueAsString(sceneID uint8, valueid *ValueID) (bool, string) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.string_emptyString()
	result := bool(C.manager_sceneGetValueAsString(m.cptr, C.uint8_t(sceneID), cValueid, cString))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return result, goString
}

// GetSceneValueListSelectionString Retrieves a scene's value as a list (as a string).
func (m *Manager) GetSceneValueListSelectionString(sceneID uint8, valueid *ValueID) (bool, string) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.string_emptyString()
	result := bool(C.manager_sceneGetValueListSelectionString(m.cptr, C.uint8_t(sceneID), cValueid, cString))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return result, goString
}

// GetSceneValueListSelectionInt32 Retrieves a scene's value as a list (as a integer).
func (m *Manager) GetSceneValueListSelectionInt32(sceneID uint8, valueid *ValueID) (bool, int32) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cInt C.int32_t
	result := bool(C.manager_sceneGetValueListSelectionInt32(m.cptr, C.uint8_t(sceneID), cValueid, &cInt))
	return result, int32(cInt)
}

// SetSceneValueBool Set a bool Value ID to an existing scene's ValueID.
func (m *Manager) SetSceneValueBool(sceneID uint8, valueid *ValueID, value bool) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setSceneValueBool(m.cptr, C.uint8_t(sceneID), cValueid, C.bool(value)))
}

// SetSceneValueUint8 Set a byte Value ID to an existing scene's ValueID.
func (m *Manager) SetSceneValueUint8(sceneID uint8, valueid *ValueID, value uint8) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setSceneValueUint8(m.cptr, C.uint8_t(sceneID), cValueid, C.uint8_t(value)))
}

// SetSceneValueFloat Set a decimal Value ID to an existing scene's ValueID.
func (m *Manager) SetSceneValueFloat(sceneID uint8, valueid *ValueID, value float32) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setSceneValueFloat(m.cptr, C.uint8_t(sceneID), cValueid, C.float(value)))
}

// SetSceneValueInt32 Set a 32-bit signed integer Value ID to an existing scene's ValueID.
func (m *Manager) SetSceneValueInt32(sceneID uint8, valueid *ValueID, value int32) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setSceneValueInt32(m.cptr, C.uint8_t(sceneID), cValueid, C.int32_t(value)))
}

// SetSceneValueInt16 Set a 16-bit integer Value ID to an existing scene's ValueID.
func (m *Manager) SetSceneValueInt16(sceneID uint8, valueid *ValueID, value int16) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setSceneValueInt16(m.cptr, C.uint8_t(sceneID), cValueid, C.int16_t(value)))
}

// SetSceneValueString Set a string Value ID to an existing scene's ValueID.
func (m *Manager) SetSceneValueString(sceneID uint8, valueid *ValueID, value string) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.CString(value)
	result := bool(C.manager_setSceneValueString(m.cptr, C.uint8_t(sceneID), cValueid, cString))
	C.free(unsafe.Pointer(cString))
	return result
}

// SetSceneValueListSelectionString Set the list selected item Value ID to an existing scene's ValueID (as a string).
func (m *Manager) SetSceneValueListSelectionString(sceneID uint8, valueid *ValueID, value string) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.CString(value)
	result := bool(C.manager_setSceneValueListSelectionString(m.cptr, C.uint8_t(sceneID), cValueid, cString))
	C.free(unsafe.Pointer(cString))
	return result
}

// SetSceneValueListSelectionInt32 Set the list selected item Value ID to an existing scene's ValueID (as a integer).
func (m *Manager) SetSceneValueListSelectionInt32(sceneID uint8, valueid *ValueID, value int32) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setSceneValueListSelectionInt32(m.cptr, C.uint8_t(sceneID), cValueid, C.int32_t(value)))
}

// GetSceneLabel Returns a label for the particular scene.
func (m *Manager) GetSceneLabel(sceneID uint8) string {
	cString := C.manager_getSceneLabel(m.cptr, C.uint8_t(sceneID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// SetSceneLabel Sets a label for the particular scene.
func (m *Manager) SetSceneLabel(sceneID uint8, value string) {
	cString := C.CString(value)
	C.manager_setSceneLabel(m.cptr, C.uint8_t(sceneID), cString)
	C.free(unsafe.Pointer(cString))
}

// SceneExists Check if a Scene ID is defined.
func (m *Manager) SceneExists(sceneID uint8) bool {
	return bool(C.manager_sceneExists(m.cptr, C.uint8_t(sceneID)))
}

// ActivateScene Activate given scene to perform all its actions.
func (m *Manager) ActivateScene(sceneID uint8) bool {
	return bool(C.manager_activateScene(m.cptr, C.uint8_t(sceneID)))
}

//
// Statistics retreival interface.
//

// GetDriverStatistics Retrieve statistics from driver.
//TODO func (m *Manager) GetDriverStatistics(...) ...
// GetNodeStatistics Retrieve statistics per node.
//TODO func (m *Manager) GetNodeStatistics(...) ...
