package goopenzwave

// #cgo pkg-config: libopenzwave
// #include <stdlib.h>
// #include <stdint.h>
// #include "node_wrap.h"
// #include "util.h"
import "C"
import "unsafe"

// Refresh a Node and Reload it into OZW.
//
// Causes the node's Supported CommandClasses and Capabilities to be obtained
// from the Z-Wave network. This method would normally be called automatically
// by OpenZWave, but if you know that a node's capabilities or command classes
// has been changed, calling this method will force a refresh of that information.
// This call shouldn't be needed except in special circumstances.
// Returns true if the request was sent successfully.
func RefreshNodeInfo(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_refresh_node_info(C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// Trigger the fetching of dynamic value data for a node.
//
// Causes the node's values to be requested from the Z-Wave network. This is the
// same as the query state starting from the associations state.
// Returns true if the request was sent successfully.
func RequestNodeState(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_request_node_state(C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// Trigger the fetching of just the dynamic value data for a node.
//
// Causes the node's values to be requested from the Z-Wave network. This is the
// same as the query state starting from the dynamic state.
// Returns true if the request was sent successfully.
func RequestNodeDynamic(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_request_node_dynamic(C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// Request the values of all known configurable parameters from a device.
func RequestNodeAllConfigParams(homeID uint32, nodeID uint8) {
	C.ozw_RequestAllConfigParams(C.uint32_t(homeID), C.uint8_t(nodeID))
}

// Get whether the node is a listening device that does not go to sleep.
//
// Returns true if the request was sent successfully.
func IsNodeListeningDevice(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_is_node_listening_device(C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// Get whether the node is a frequent listening device that goes to sleep but can be woken up by a beam. Useful to determine node and controller consistency.
//
// Returns true if the request was sent successfully.
func IsNodeFrequentListeningDevice(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_is_node_frequent_listening_device(C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// Get whether the node is a beam capable device.
//
// Returns true if the request was sent successfully.
func IsNodeBeamingDevice(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_is_node_beaming_device(C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// Get whether the node is a routing device that passes messages to other nodes
//
// Returns true if the request was sent successfully.
func IsNodeRoutingDevice(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_is_node_routing_device(C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// Get the security attribute for a node. True if node supports security features.
//
// Returns true if the request was sent successfully.
func IsNodeSecurityDevice(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_is_node_security_device(C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// Get the maximum baud rate of a node's communications.
//
// Returns the baud rate in bits per second.
func GetNodeMaxBaudRate(homeID uint32, nodeID uint8) uint32 {
	return uint32(C.manager_get_node_max_baud_rate(C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// Get the version number of a node.
//
// Returns the node's version number.
func GetNodeVersion(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_get_node_version(C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// Get the security byte of a node
//
// Returns the node's security byte.
func GetNodeSecurity(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_get_node_security(C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// Is this a ZWave+ Supported Node?
//
// Returns If this node is a Z-Wave Plus Node.
func IsNodeZWavePlus(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_is_node_zwave_plus(C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// Get the basic type of a node.
//
// Returns the node's basic type.
func GetNodeBasic(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_get_node_basic(C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// Get the generic type of a node. Set instance to 0 if not required.
//
// Returns the node's generic type.
//
// TODO: utilise instance argument in v1.6+
func GetNodeGeneric(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_get_node_generic(C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// Get the specific type of a node. Set instance to 0 if not required.
//
// Returns the node's specific type.
//
// TODO: utilise instance argument in v1.6+
func GetNodeSpecific(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_get_node_specific(C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// Get a human-readable label describing the node.
//
// The label is taken from the Z-Wave specific, generic or basic type, depending on which of those values are specified by the node.
// Returns A string containing the label text.
func GetNodeType(homeID uint32, nodeID uint8) string {
	cstr := C.manager_get_node_type(C.uint32_t(homeID), C.uint8_t(nodeID))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// Get the bitmap of this node's neighbors. See SyncronizeNodeNeighbors.
//
// Returns an array of uint8s to hold the neighbor bitmap.
func GetNodeNeighbors(homeID uint32, nodeID uint8) []uint8 {
	cres := C.manager_get_node_neighbors(C.uint32_t(homeID), C.uint8_t(nodeID))
	defer C.free(unsafe.Pointer(cres))
	getval := func(i uint32) uint32 {
		return *(*uint32)(C.ptr_at((*unsafe.Pointer)(unsafe.Pointer(cres)), C.uint32_t(i)))
	}
	size := getval(0)
	var res []uint8
	for i := uint32(0); i < size; i++ {
		res = append(res, uint8(getval(i+1)))
	}
	return res
}

// Update the List of Neighbors on a particular node.
//
// This retrieves the latest copy of the Neighbor lists for a particular node and should be called
// before calling GetNodeNeighbors to ensure OZW returns the most recent list of Neighbors.
func SyncronizeNodeNeighbors(homeID uint32, nodeID uint8) {
	C.manager_syncronize_node_neighbors(C.uint32_t(homeID), C.uint8_t(nodeID))
}

// Get the manufacturer name of a device.
//
// The manufacturer name would normally be handled by the Manufacturer Specific command class,
// taking the manufacturer ID reported by the device and using it to look up the name from the
// manufacturer_specific.xml file in the OpenZWave config folder.
// However, there are some devices that do not support the command class, so to enable the user
// to manually set the name, it is stored with the node data and accessed via this method rather
// than being reported via a command class Value object.
//
// Returns a string containing the node's manufacturer name.
//
// See SetNodeManufacturerName, GetNodeProductName, SetNodeProductName.
func GetNodeManufacturerName(homeID uint32, nodeID uint8) string {
	cstr := C.manager_get_node_manufacturer_name(C.uint32_t(homeID), C.uint8_t(nodeID))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// Get the product name of a device.
//
// The product name would normally be handled by the Manufacturer Specific command class,
// taking the product Type and ID reported by the device and using it to look up the name from the
// manufacturer_specific.xml file in the OpenZWave config folder.
// However, there are some devices that do not support the command class, so to enable the user
// to manually set the name, it is stored with the node data and accessed via this method rather
// than being reported via a command class Value object.
//
// Returns a string containing the node's product name.
//
// See SetNodeProductName, GetNodeManufacturerName, SetNodeManufacturerName.
func GetNodeProductName(homeID uint32, nodeID uint8) string {
	cstr := C.manager_get_node_product_name(C.uint32_t(homeID), C.uint8_t(nodeID))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// Get the name of a node.
//
// The node name is a user-editable label for the node that would normally be handled by the
// Node Naming command class, but many devices do not support it.  So that a node can always
// be named, OpenZWave stores it with the node data, and provides access through this method
// and SetNodeName, rather than reporting it via a command class Value object.
// The maximum length of a node name is 16 characters.
//
// Returns a string containing the node's name.
//
// See SetNodeName, GetNodeLocation, SetNodeLocation.
func GetNodeName(homeID uint32, nodeID uint8) string {
	cstr := C.manager_get_node_name(C.uint32_t(homeID), C.uint8_t(nodeID))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// Get the location of a node.
//
// The node location is a user-editable string that would normally be handled by the Node Naming
// command class, but many devices do not support it.  So that a node can always report its
// location, OpenZWave stores it with the node data, and provides access through this method
// and SetNodeLocation, rather than reporting it via a command class Value object.
//
// Returns a string containing the node's location.
//
// See SetNodeLocation, GetNodeName, SetNodeName.
func GetNodeLocation(homeID uint32, nodeID uint8) string {
	cstr := C.manager_get_node_location(C.uint32_t(homeID), C.uint8_t(nodeID))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// Get the manufacturer ID of a device.
//
// The manufacturer ID is a four digit hex code and would normally be handled by the Manufacturer
// Specific command class, but not all devices support it.  Although the value reported by this
// method will be an empty string if the command class is not supported and cannot be set by the
// user, the manufacturer ID is still stored with the node data (rather than being reported via a
// command class Value object) to retain a consistent approach with the other manufacturer specific data.
//
// Returns A string containing the node's manufacturer ID, or an empty string if the manufacturer
// specific command class is not supported by the device.
//
// See GetNodeProductType, GetNodeProductId, GetNodeManufacturerName, GetNodeProductName.
//
// TODO: Change the return to uint16 in 2.0 time frame
func GetNodeManufacturerID(homeID uint32, nodeID uint8) string {
	cstr := C.manager_get_node_manufacturer_id(C.uint32_t(homeID), C.uint8_t(nodeID))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// Get the product type of a device.
//
// The product type is a four digit hex code and would normally be handled by the Manufacturer Specific
// command class, but not all devices support it.  Although the value reported by this method will
// be an empty string if the command class is not supported and cannot be set by the user, the product
// type is still stored with the node data (rather than being reported via a command class Value object)
// to retain a consistent approach with the other manufacturer specific data.
//
// Returns a string containing the node's product type, or an empty string if the manufacturer
// specific command class is not supported by the device.
//
// See GetNodeManufacturerId, GetNodeProductId, GetNodeManufacturerName, GetNodeProductName.
//
// TODO: Change the return to uint16 in 2.0 time frame
func GetNodeProductType(homeID uint32, nodeID uint8) string {
	cstr := C.manager_get_node_product_type(C.uint32_t(homeID), C.uint8_t(nodeID))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// Get the product ID of a device.
//
// The product ID is a four digit hex code and would normally be handled by the Manufacturer Specific
// command class, but not all devices support it.  Although the value reported by this method will
// be an empty string if the command class is not supported and cannot be set by the user, the product
// ID is still stored with the node data (rather than being reported via a command class Value object)
// to retain a consistent approach with the other manufacturer specific data.
//
// Returns a string containing the node's product ID, or an empty string if the manufacturer
// specific command class is not supported by the device.
//
// See GetNodeManufacturerId, GetNodeProductType, GetNodeManufacturerName, GetNodeProductName.
//
// TODO: Change the return to uint16 in 2.0 time frame
func GetNodeProductId(homeID uint32, nodeID uint8) string {
	cstr := C.manager_get_node_product_type(C.uint32_t(homeID), C.uint8_t(nodeID))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// Set the manufacturer name of a device.
//
// The manufacturer name would normally be handled by the Manufacturer Specific command class,
// taking the manufacturer ID reported by the device and using it to look up the name from the
// manufacturer_specific.xml file in the OpenZWave config folder.
// However, there are some devices that do not support the command class, so to enable the user
// to manually set the name, it is stored with the node data and accessed via this method rather
// than being reported via a command class Value object.
//
// See GetNodeManufacturerName, GetNodeProductName, SetNodeProductName.
func SetNodeManufacturerName(homeID uint32, nodeID uint8, manufacturerName string) {
	cstr := C.CString(manufacturerName)
	defer C.free(unsafe.Pointer(cstr))
	C.manager_set_node_manufacturer_name(C.uint32_t(homeID), C.uint8_t(nodeID), cstr)
}

// Set the product name of a device.
//
// The product name would normally be handled by the Manufacturer Specific command class,
// taking the product Type and ID reported by the device and using it to look up the name from the
// manufacturer_specific.xml file in the OpenZWave config folder.
// However, there are some devices that do not support the command class, so to enable the user
// to manually set the name, it is stored with the node data and accessed via this method rather
// than being reported via a command class Value object.
//
// See GetNodeProductName, GetNodeManufacturerName, SetNodeManufacturerName.
func SetNodeProductName(homeID uint32, nodeID uint8, productName string) {
	cstr := C.CString(productName)
	defer C.free(unsafe.Pointer(cstr))
	C.manager_set_node_product_name(C.uint32_t(homeID), C.uint8_t(nodeID), cstr)
}

// Set the name of a node.
//
// The node name is a user-editable label for the node that would normally be handled by the
// Node Naming command class, but many devices do not support it.  So that a node can always
// be named, OpenZWave stores it with the node data, and provides access through this method
// and GetNodeName, rather than reporting it via a command class Value object.
// If the device does support the Node Naming command class, the new name will be sent to the node.
// The maximum length of a node name is 16 characters.
//
// See GetNodeName, GetNodeLocation, SetNodeLocation.
func SetNodeName(homeID uint32, nodeID uint8, nodeName string) {
	cstr := C.CString(nodeName)
	defer C.free(unsafe.Pointer(cstr))
	C.manager_set_node_name(C.uint32_t(homeID), C.uint8_t(nodeID), cstr)
}

// Set the location of a node.
//
// The node location is a user-editable string that would normally be handled by the Node Naming
// command class, but many devices do not support it.  So that a node can always report its
// location, OpenZWave stores it with the node data, and provides access through this method
// and GetNodeLocation, rather than reporting it via a command class Value object.
// If the device does support the Node Naming command class, the new location will be sent to the node.
//
// See GetNodeLocation, GetNodeName, SetNodeName.
func SetNodeLocation(homeID uint32, nodeID uint8, location string) {
	cstr := C.CString(location)
	defer C.free(unsafe.Pointer(cstr))
	C.manager_set_node_location(C.uint32_t(homeID), C.uint8_t(nodeID), cstr)
}

// Get whether the node information has been received.
//
// Returns true if the node information has been received yet.
func IsNodeInfoReceived(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_is_node_info_received(C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// Get whether the node has the defined class available or not.
//
// Returns true if the node does have the class instantiated, will return name & version.
func GetNodeClassInformation(homeID uint32, nodeID uint8, commandClassId uint8) (className string, classVersion uint8, hasClass bool) {
	res := C.manager_get_node_class_information(C.uint32_t(homeID), C.uint8_t(nodeID), C.uint8_t(commandClassId))
	defer C.manager_free_node_class_information(res)
	return C.GoString(res.className), uint8(res.classVersion), bool(res.ok)
}

// Get whether the node is awake or asleep.
//
// Returns true if the node is awake.
func IsNodeAwake(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_is_node_awake(C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// Get whether the node is working or has failed.
//
// Returns true if the node has failed and is no longer part of the network.
func IsNodeFailed(homeID uint32, nodeID uint8) bool {
	return bool(C.manager_is_node_failed(C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// Get whether the node's query stage as a string.
//
// Returns name of current query stage as a string..
func GetNodeQueryStage(homeID uint32, nodeID uint8) string {
	cstr := C.manager_get_node_query_stage(C.uint32_t(homeID), C.uint8_t(nodeID))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// Get the node device type as reported in the Z-Wave+ Info report.
//
// Returns the node's DeviceType.
func GetNodeDeviceType(homeID uint32, nodeID uint8) uint16 {
	return uint16(C.manager_get_node_device_type(C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// Get the node role as reported in the Z-Wave+ Info report.
//
// Returns the node's user icon..
func GetNodeRole(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_get_node_role(C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// Get the node PlusType as reported in the Z-Wave+ Info report.
//
// Returns the node's PlusType.
func GetNodePlusType(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_get_node_plus_type(C.uint32_t(homeID), C.uint8_t(nodeID)))
}
