package goopenzwave

// #cgo LDFLAGS: -lopenzwave -L/usr/local/lib
// #cgo CPPFLAGS: -I/usr/local/include -I/usr/local/include/openzwave
// #include "manager.h"
// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

// RefreshNodeInfo triggers the fetching of fixed data about a node. Returns
// true if the request was sent successfully.
//
// Causes the node's data to be obtained from the Z-Wave network in the same way
// as if it had just been added. This method would normally be called
// automatically by OpenZWave, but if you know that a node has been changed,
// calling this method will force a refresh of all of the data held by the
// library. This can be especially useful for devices that were asleep when the
// application was first run. This is the same as the query state starting from
// the beginning.
func RefreshNodeInfo(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_refreshNodeInfo(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// RequestNodeState triggers the fetching of dynamic value data for a node.
// Returns true if the request was sent successfully.
//
// Causes the node's values to be requested from the Z-Wave network. This is the
// same as the query state starting from the associations state.
func RequestNodeState(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_requestNodeState(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// RequestNodeDynamic triggers the fetching of just the dynamic value data for a
// node. Returns true if the request was sent successfully.
//
// Causes the node's values to be requested from the Z-Wave network. This is the
// same as the query state starting from the dynamic state.
func RequestNodeDynamic(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_requestNodeDynamic(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// IsNodeListeningDevice returns true if the node is a listening device that
// does not go to sleep.
func IsNodeListeningDevice(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_isNodeListeningDevice(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// IsNodeFrequentListeningDevice returns true if the node is a frequent
// listening device that goes to sleep but can be woken up by a beam. Useful to
// determine node and controller consistency.
func IsNodeFrequentListeningDevice(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_isNodeFrequentListeningDevice(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// IsNodeBeamingDevice returns true if the node is a beam capable device.
func IsNodeBeamingDevice(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_isNodeBeamingDevice(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// IsNodeRoutingDevice returns true if the node is a routing device that passes
// messages to other nodes.
func IsNodeRoutingDevice(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_isNodeRoutingDevice(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// IsNodeSecurityDevice returns true if the node supports security features.
func IsNodeSecurityDevice(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_isNodeSecurityDevice(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeMaxBaudRate returns the maximum baud rate of a node's communications.
func GetNodeMaxBaudRate(homeID uint32, nodeID uint8) uint32 {
	return uint32(C.manager_getNodeMaxBaudRate(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeVersion returns the version number of a node.
func GetNodeVersion(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_getNodeVersion(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeSecurity returns the security byte of a node.
func GetNodeSecurity(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_getNodeSecurity(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// IsNodeZWavePlus returns true if this a ZWave+ Supported Node.
func IsNodeZWavePlus(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_isNodeZWavePlus(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeBasicType returns the basic type of a node.
func GetNodeBasicType(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_getNodeBasic(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeGenericType returns the generic type of a node.
func GetNodeGenericType(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_getNodeGeneric(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeSpecificType returns the specific type of a node.
func GetNodeSpecificType(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_getNodeSpecific(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeType returns a human-readable label describing the node.
//
// The label is taken from the Z-Wave specific, generic or basic type, depending
// on which of those values are specified by the node.
func GetNodeType(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeType(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetNodeNeighbours returns the bitmap of this node's neighbors.
//TODO func GetNodeNeighbours(homeID uint32, nodeID uint8) (uint32, uint8 nodeNeighbors)

// GetNodeManufacturerName returns the manufacturer name of a device.
//
// The manufacturer name would normally be handled by the Manufacturer Specific
// commmand class, taking the manufacturer ID reported by the device and using
// it to look up the name from the manufacturer_specific.xml file in the
// OpenZWave config folder. However, there are some devices that do not support
// the command class, so to enable the user to manually set the name, it is
// stored with the node data and accessed via this method rather than being
// reported via a command class Value object.
func GetNodeManufacturerName(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeManufacturerName(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetNodeProductName returns the product name of a device.
//
// The product name would normally be handled by the Manufacturer Specific
// commmand class, taking the product Type and ID reported by the device and
// using it to look up the name from the manufacturer_specific.xml file in the
// OpenZWave config folder. However, there are some devices that do not support
// the command class, so to enable the user to manually set the name, it is
// stored with the node data and accessed via this method rather than being
// reported via a command class Value object.
func GetNodeProductName(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeProductName(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetNodeName returns the name of a node.
//
// The node name is a user-editable label for the node that would normally be
// handled by the Node Naming commmand class, but many devices do not support
// it. So that a node can always be named, OpenZWave stores it with the node
// data, and provides access through this method and SetNodeName, rather than
// reporting it via a command class Value object. The maximum length of a node
// name is 16 characters.
func GetNodeName(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeName(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetNodeLocation returns the location of a node.
//
// The node location is a user-editable string that would normally be handled by
// the Node Naming commmand class, but many devices do not support it. So that a
// node can always report its location, OpenZWave stores it with the node data,
// and provides access through this method and SetNodeLocation, rather than
// reporting it via a command class Value object.
func GetNodeLocation(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeLocation(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetNodeManufacturerID returns the manufacturer ID of a device.
//
// The manufacturer ID is a four digit hex code and would normally be handled by
// the Manufacturer Specific commmand class, but not all devices support it.
// Although the value reported by this method will be an empty string if the
// command class is not supported and cannot be set by the user, the
// manufacturer ID is still stored with the node data (rather than being
// reported via a command class Value object) to retain a consistent approach
// with the other manufacturer specific data.
func GetNodeManufacturerID(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeManufacturerId(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetNodeProductType returns the product type of a device.
//
// The product type is a four digit hex code and would normally be handled by
// the Manufacturer Specific commmand class, but not all devices support it.
// Although the value reported by this method will be an empty string if the
// command class is not supported and cannot be set by the user, the product
// type is still stored with the node data (rather than being reported via a
// command class Value object) to retain a consistent approach with the other
// manufacturer specific data.
func GetNodeProductType(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeProductType(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetNodeProductID returns the product ID of a device.
//
// The product ID is a four digit hex code and would normally be handled by the
// Manufacturer Specific commmand class, but not all devices support it.
// Although the value reported by this method will be an empty string if the
// command class is not supported and cannot be set by the user, the product ID
// is still stored with the node data (rather than being reported via a command
// class Value object) to retain a consistent approach with the other
// manufacturer specific data.
func GetNodeProductID(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeProductId(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// SetNodeManufacturerName sets the manufacturer name of a device.
//
// The manufacturer name would normally be handled by the Manufacturer Specific
// commmand class, taking the manufacturer ID reported by the device and using
// it to look up the name from the manufacturer_specific.xml file in the
// OpenZWave config folder. However, there are some devices that do not support
// the command class, so to enable the user to manually set the name, it is
// stored with the node data and accessed via this method rather than being
// reported via a command class Value object.
func SetNodeManufacturerName(homeID uint32, nodeID uint8, manufacturerName string) {
	cString := C.CString(manufacturerName)
	C.manager_setNodeManufacturerName(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID), cString)
	C.free(unsafe.Pointer(cString))
}

// SetNodeProductName sets the product name of a device.
//
// The product name would normally be handled by the Manufacturer Specific
// commmand class, taking the product Type and ID reported by the device and
// using it to look up the name from the manufacturer_specific.xml file in the
// OpenZWave config folder. However, there are some devices that do not support
// the command class, so to enable the user to manually set the name, it is
// stored with the node data and accessed via this method rather than being
// reported via a command class Value object.
func SetNodeProductName(homeID uint32, nodeID uint8, productName string) {
	cString := C.CString(productName)
	C.manager_setNodeProductName(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID), cString)
	C.free(unsafe.Pointer(cString))
}

// SetNodeName sets the name of a node.
//
// The node name is a user-editable label for the node that would normally be
// handled by the Node Naming commmand class, but many devices do not support
// it. So that a node can always be named, OpenZWave stores it with the node
// data, and provides access through this method and GetNodeName, rather than
// reporting it via a command class Value object. If the device does support the
// Node Naming command class, the new name will be sent to the node. The maximum
// length of a node name is 16 characters.
func SetNodeName(homeID uint32, nodeID uint8, nodeName string) {
	cString := C.CString(nodeName)
	C.manager_setNodeName(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID), cString)
	C.free(unsafe.Pointer(cString))
}

// SetNodeLocation sets the location of a node.
//
// The node location is a user-editable string that would normally be handled by
// the Node Naming commmand class, but many devices do not support it. So that a
// node can always report its location, OpenZWave stores it with the node data,
// and provides access through this method and GetNodeLocation, rather than
// reporting it via a command class Value object. If the device does support the
// Node Naming command class, the new location will be sent to the node.
func SetNodeLocation(homeID uint32, nodeID uint8, location string) {
	cString := C.CString(location)
	C.manager_setNodeLocation(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID), cString)
	C.free(unsafe.Pointer(cString))
}

// SetNodeOn turns a node on.
//
// This is a helper method to simplify basic control of a node. It is the
// equivalent of changing the level reported by the node's Basic command class
// to 255, and will generate a ValueChanged notification from that class. This
// command will turn on the device at its last known level, if supported by the
// device, otherwise it will turn it on at 100%.
func SetNodeOn(homeID uint32, nodeID uint8) {
	C.manager_setNodeOn(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID))
}

// SetNodeOff turns a node off.
//
// This is a helper method to simplify basic control of a node. It is the
// equivalent of changing the level reported by the node's Basic command class
// to zero, and will generate a ValueChanged notification from that class.
func SetNodeOff(homeID uint32, nodeID uint8) {
	C.manager_setNodeOff(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID))
}

// SetNodeLevel sets the basic level of a node.
//
// This is a helper method to simplify basic control of a node. It is the
// equivalent of changing the value reported by the node's Basic command class
// and will generate a ValueChanged notification from that class.
func SetNodeLevel(homeID uint32, nodeID uint8, level uint8) {
	C.manager_setNodeLevel(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint8_t(level))
}

// IsNodeInfoReceived returns whether the node information has been received.
func IsNodeInfoReceived(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_isNodeInfoReceived(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeClassInformation returns true if the node has the defined class
// available or not, and the class name and version if available.
func GetNodeClassInformation(homeID uint32, nodeID uint8, commandClassID uint8) (bool, string, uint8) {
	cClassName := C.string_emptyString()
	var cClassVersion C.uint8_t
	result := bool(C.manager_getNodeClassInformation(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint8_t(commandClassID), cClassName, &cClassVersion))
	goClassName := C.GoString(cClassName.data)
	goClassVersion := uint8(cClassVersion)
	C.string_freeString(cClassName)
	return result, goClassName, goClassVersion
}

// IsNodeAwake returns true if the node is awake, otherwise false if it is
// asleep.
func IsNodeAwake(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_isNodeAwake(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// IsNodeFailed returns true if the node is working, otherwise false if it has
// failed.
func IsNodeFailed(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_isNodeFailed(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeQueryStage returns the node's query stage as a string.
func GetNodeQueryStage(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeQueryStage(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetNodeDeviceType returns the node device type as reported in the Z-Wave+
// Info report.
func GetNodeDeviceType(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_getNodeDeviceType(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeDeviceTypeString returns a string of the node device type as reported
// in the Z-Wave+ Info report.
func GetNodeDeviceTypeString(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeDeviceTypeString(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetNodeRole returns the node role as reported in the Z-Wave+ Info report.
func GetNodeRole(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_getNodeRole(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodeRoleString returns a string of the node role as reported in the
// Z-Wave+ Info report.
func GetNodeRoleString(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodeRoleString(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// GetNodePlusType returns the node PlusType as reported in the Z-Wave+ Info
// report.
func GetNodePlusType(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_getNodePlusType(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetNodePlusTypeString returns a string of the node PlusType as reported in
// the Z-Wave+ Info report.
func GetNodePlusTypeString(homeID uint32, nodeID uint8) string {
	cString := C.manager_getNodePlusTypeString(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// SetNodeConfigParam sets the value of a configurable parameter in a device.
// Returns true if the message setting was sent to the device.
//
// Some devices have various parameters that can be configured to control the
// device behaviour. These are not reported by the device over the Z-Wave
// network, but can usually be found in the device's user manual. This method
// returns immediately, without waiting for confirmation from the device that
// the change has been made.
func SetNodeConfigParam(homeID uint32, nodeID uint8, param uint8, value int32, size uint8) bool {
	return bool(C.manager_setConfigParam(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint8_t(param), C.int32_t(value), C.uint8_t(size)))
}

// SwitchAllOn Switch all devices on. All devices that support the SwitchAll command class will be turned on.
func SwitchAllOn(homeID uint32) {
	C.manager_switchAllOn(cmanager, C.uint32_t(homeID))
}

// SwitchAllOff Switch all devices off. All devices that support the SwitchAll command class will be turned off.
func SwitchAllOff(homeID uint32) {
	C.manager_switchAllOff(cmanager, C.uint32_t(homeID))
}

// RequestNodeConfigParam requests the value of a configurable parameter from a
// device.
//
// Some devices have various parameters that can be configured to control the
// device behaviour. These are not reported by the device over the Z-Wave
// network, but can usually be found in the device's user manual. This method
// requests the value of a parameter from the device, and then returns
// immediately, without waiting for a response. If the parameter index is valid
// for this device, and the device is awake, the value will eventually be
// reported via a ValueChanged notification callback. The ValueID reported in
// the callback will have an index set the same as _param and a command class
// set to the same value as returned by a call to
// Configuration::StaticGetCommandClassId.
func RequestNodeConfigParam(homeID uint32, nodeID uint8, param uint8) {
	C.manager_requestConfigParam(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint8_t(param))
}

// RequestNodeAllConfigParam requests the values of all known configurable
// parameters from a device.
func RequestNodeAllConfigParam(homeID uint32, nodeID uint8) {
	C.manager_requestAllConfigParams(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID))
}

// GetNodeStatistics Retrieve statistics per node.
//TODO func GetNodeStatistics(...) ...
