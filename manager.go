package gozwave

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
	manager       C.manager_t
	Notifications chan *Notification
}

type ManagerVersion struct {
	Major int
	Minor int
}

//
// Construction.
//

func CreateManager() *Manager {
	m := &Manager{}
	m.manager = C.manager_create()
	m.Notifications = make(chan *Notification, 10)
	return m
}

func GetManager() *Manager {
	m := &Manager{}
	m.manager = C.manager_get()
	return m
}

func DestroyManager() {
	C.manager_destroy()
}

func GetManagerVersionAsString() string {
	cString := C.manager_getVersionAsString()
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

func GetManagerVersionLongAsString() string {
	cString := C.manager_getVersionLongAsString()
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

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

func (m *Manager) WriteConfig(homeId uint32) {
	C.manager_writeConfig(m.manager, C.uint32_t(homeId))
}

func (m *Manager) GetOptions() *Options {
	return &Options{
		options: C.manager_getOptions(m.manager),
	}
}

//
// Drivers.
//

func (m *Manager) AddDriver(controllerPath string) bool {
	cControllerPath := C.CString(controllerPath)
	result := bool(C.manager_addDriver(m.manager, cControllerPath))
	C.free(unsafe.Pointer(cControllerPath))
	return result
}

func (m *Manager) RemoveDriver(controllerPath string) bool {
	cControllerPath := C.CString(controllerPath)
	result := bool(C.manager_removeDriver(m.manager, cControllerPath))
	C.free(unsafe.Pointer(cControllerPath))
	return result
}

func (m *Manager) GetControllerNodeId(homeId uint32) uint8 {
	return uint8(C.manager_getControllerNodeId(m.manager, C.uint32_t(homeId)))
}

func (m *Manager) GetSUCNodeId(homeId uint32) uint8 {
	return uint8(C.manager_getSUCNodeId(m.manager, C.uint32_t(homeId)))
}

func (m *Manager) IsPrimaryController(homeId uint32) bool {
	return bool(C.manager_isPrimaryController(m.manager, C.uint32_t(homeId)))
}

func (m *Manager) IsStaticUpdateController(homeId uint32) bool {
	return bool(C.manager_isStaticUpdateController(m.manager, C.uint32_t(homeId)))
}

func (m *Manager) IsBridgeController(homeId uint32) bool {
	return bool(C.manager_isBridgeController(m.manager, C.uint32_t(homeId)))
}

func (m *Manager) GetLibraryVersion(homeId uint32) string {
	cString := C.manager_getLibraryVersion(m.manager, C.uint32_t(homeId))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

func (m *Manager) GetLibraryTypeName(homeId uint32) string {
	cString := C.manager_getLibraryTypeName(m.manager, C.uint32_t(homeId))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

func (m *Manager) GetSendQueueCount(homeId uint32) int32 {
	return int32(C.manager_getSendQueueCount(m.manager, C.uint32_t(homeId)))
}

func (m *Manager) LogDriverStatistics(homeId uint32) {
	C.manager_logDriverStatistics(m.manager, C.uint32_t(homeId))
}

//TODO func (m *Manager) GetControllerInterfaceType(homeId uint32) ...

func (m *Manager) GetControllerPath(homeId uint32) string {
	cString := C.manager_getControllerPath(m.manager, C.uint32_t(homeId))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

//
// Polling Z-Wave devices.
//

func (m *Manager) GetPollInterval() int32 {
	return int32(C.manager_getPollInterval(m.manager))
}

func (m *Manager) SetPollInterval(milliseconds int32, intervalBetweenPolls bool) {
	C.manager_setPollInterval(m.manager, C.int32_t(milliseconds), C.bool(intervalBetweenPolls))
}

func (m *Manager) EnablePoll(valueid *ValueID, intensity uint8) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_enablePoll(m.manager, cValueid, C.uint8_t(intensity)))
}

func (m *Manager) DisablePoll(valueid *ValueID) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_disablePoll(m.manager, cValueid))
}

func (m *Manager) IsPolled(valueid *ValueID) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_isPolled(m.manager, cValueid))
}

func (m *Manager) SetPollIntensity(valueid *ValueID, intensity uint8) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	C.manager_setPollIntensity(m.manager, cValueid, C.uint8_t(intensity))
}

func (m *Manager) GetPollIntensity(valueid *ValueID) uint8 {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return uint8(C.manager_getPollIntensity(m.manager, cValueid))
}

//
// Node information.
//

func (m *Manager) RefreshNodeInfo(homeId uint32, nodeId uint8) bool {
	return bool(C.manager_refreshNodeInfo(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) RequestNodeState(homeId uint32, nodeId uint8) bool {
	return bool(C.manager_requestNodeState(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) RequestNodeDynamic(homeId uint32, nodeId uint8) bool {
	return bool(C.manager_requestNodeDynamic(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) IsNodeListeningDevice(homeId uint32, nodeId uint8) bool {
	return bool(C.manager_isNodeListeningDevice(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) IsNodeFrequentListeningDevice(homeId uint32, nodeId uint8) bool {
	return bool(C.manager_isNodeFrequentListeningDevice(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) IsNodeBeamingDevice(homeId uint32, nodeId uint8) bool {
	return bool(C.manager_isNodeBeamingDevice(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) IsNodeRoutingDevice(homeId uint32, nodeId uint8) bool {
	return bool(C.manager_isNodeRoutingDevice(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) IsNodeSecurityDevice(homeId uint32, nodeId uint8) bool {
	return bool(C.manager_isNodeSecurityDevice(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) GetNodeMaxBaudRate(homeId uint32, nodeId uint8) uint32 {
	return uint32(C.manager_getNodeMaxBaudRate(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) GetNodeVersion(homeId uint32, nodeId uint8) uint8 {
	return uint8(C.manager_getNodeVersion(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) GetNodeSecurity(homeId uint32, nodeId uint8) uint8 {
	return uint8(C.manager_getNodeSecurity(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) IsNodeZWavePlus(homeId uint32, nodeId uint8) bool {
	return bool(C.manager_isNodeZWavePlus(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) GetNodeBasic(homeId uint32, nodeId uint8) uint8 {
	return uint8(C.manager_getNodeBasic(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) GetNodeGeneric(homeId uint32, nodeId uint8) uint8 {
	return uint8(C.manager_getNodeGeneric(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) GetNodeSpecific(homeId uint32, nodeId uint8) uint8 {
	return uint8(C.manager_getNodeSpecific(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) GetNodeType(homeId uint32, nodeId uint8) string {
	cString := C.manager_getNodeType(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

//TODO func (m *Manager) GetNodeNeighbours(homeId uint32, nodeId uint8) (uint32, uint8 nodeNeighbors)

func (m *Manager) GetNodeManufacturerName(homeId uint32, nodeId uint8) string {
	cString := C.manager_getNodeManufacturerName(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

func (m *Manager) GetNodeProductName(homeId uint32, nodeId uint8) string {
	cString := C.manager_getNodeProductName(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

func (m *Manager) GetNodeName(homeId uint32, nodeId uint8) string {
	cString := C.manager_getNodeName(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

func (m *Manager) GetNodeLocation(homeId uint32, nodeId uint8) string {
	cString := C.manager_getNodeLocation(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

func (m *Manager) GetNodeManufacturerId(homeId uint32, nodeId uint8) string {
	cString := C.manager_getNodeManufacturerId(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

func (m *Manager) GetNodeProductType(homeId uint32, nodeId uint8) string {
	cString := C.manager_getNodeProductType(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

func (m *Manager) GetNodeProductId(homeId uint32, nodeId uint8) string {
	cString := C.manager_getNodeProductId(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

func (m *Manager) SetNodeManufacturerName(homeId uint32, nodeId uint8, manufacturerName string) {
	cString := C.CString(manufacturerName)
	C.manager_setNodeManufacturerName(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId), cString)
	C.free(unsafe.Pointer(cString))
}

func (m *Manager) SetNodeProductName(homeId uint32, nodeId uint8, productName string) {
	cString := C.CString(productName)
	C.manager_setNodeProductName(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId), cString)
	C.free(unsafe.Pointer(cString))
}

func (m *Manager) SetNodeName(homeId uint32, nodeId uint8, nodeName string) {
	cString := C.CString(nodeName)
	C.manager_setNodeName(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId), cString)
	C.free(unsafe.Pointer(cString))
}

func (m *Manager) SetNodeLocation(homeId uint32, nodeId uint8, location string) {
	cString := C.CString(location)
	C.manager_setNodeLocation(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId), cString)
	C.free(unsafe.Pointer(cString))
}

func (m *Manager) SetNodeOn(homeId uint32, nodeId uint8) {
	C.manager_setNodeOn(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId))
}

func (m *Manager) SetNodeOff(homeId uint32, nodeId uint8) {
	C.manager_setNodeOff(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId))
}

func (m *Manager) SetNodeLevel(homeId uint32, nodeId uint8, level uint8) {
	C.manager_setNodeLevel(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId), C.uint8_t(level))
}

func (m *Manager) IsNodeInfoReceived(homeId uint32, nodeId uint8) bool {
	return bool(C.manager_isNodeInfoReceived(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) GetNodeClassInformation(homeId uint32, nodeId uint8, commandClassId uint8) (bool, string, uint8) {
	cClassName := C.string_emptyString()
	var cClassVersion C.uint8_t
	result := bool(C.manager_getNodeClassInformation(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId), C.uint8_t(commandClassId), cClassName, &cClassVersion))
	goClassName := C.GoString(cClassName.data)
	goClassVersion := uint8(cClassVersion)
	C.string_freeString(cClassName)
	return result, goClassName, goClassVersion
}

func (m *Manager) IsNodeAwake(homeId uint32, nodeId uint8) bool {
	return bool(C.manager_isNodeAwake(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) IsNodeFailed(homeId uint32, nodeId uint8) bool {
	return bool(C.manager_isNodeFailed(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) GetNodeQueryStage(homeId uint32, nodeId uint8) string {
	cString := C.manager_getNodeQueryStage(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

func (m *Manager) GetNodeDeviceType(homeId uint32, nodeId uint8) uint8 {
	return uint8(C.manager_getNodeDeviceType(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) GetNodeDeviceTypeString(homeId uint32, nodeId uint8) string {
	cString := C.manager_getNodeDeviceTypeString(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

func (m *Manager) GetNodeRole(homeId uint32, nodeId uint8) uint8 {
	return uint8(C.manager_getNodeRole(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) GetNodeRoleString(homeId uint32, nodeId uint8) string {
	cString := C.manager_getNodeRoleString(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

func (m *Manager) GetNodePlusType(homeId uint32, nodeId uint8) uint8 {
	return uint8(C.manager_getNodePlusType(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) GetNodePlusTypeString(homeId uint32, nodeId uint8) string {
	cString := C.manager_getNodePlusTypeString(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

//
// Values.
//

func (m *Manager) GetValueLabel(valueid *ValueID) string {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.manager_getValueLabel(m.manager, cValueid)
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

func (m *Manager) SetValueLabel(valueid *ValueID, value string) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.CString(value)
	C.manager_setValueLabel(m.manager, cValueid, cString)
	C.free(unsafe.Pointer(cString))
}

func (m *Manager) GetValueUnits(valueid *ValueID) string {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.manager_getValueUnits(m.manager, cValueid)
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

func (m *Manager) SetValueUnits(valueid *ValueID, value string) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.CString(value)
	C.manager_setValueUnits(m.manager, cValueid, cString)
	C.free(unsafe.Pointer(cString))
}

func (m *Manager) GetValueHelp(valueid *ValueID) string {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.manager_getValueHelp(m.manager, cValueid)
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

func (m *Manager) SetValueHelp(valueid *ValueID, value string) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.CString(value)
	C.manager_setValueHelp(m.manager, cValueid, cString)
	C.free(unsafe.Pointer(cString))
}

func (m *Manager) GetValueMin(valueid *ValueID) int32 {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return int32(C.manager_getValueMin(m.manager, cValueid))
}

func (m *Manager) GetValueMax(valueid *ValueID) int32 {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return int32(C.manager_getValueMax(m.manager, cValueid))
}

func (m *Manager) IsValueReadOnly(valueid *ValueID) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_isValueReadOnly(m.manager, cValueid))
}

func (m *Manager) IsValueWriteOnly(valueid *ValueID) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_isValueWriteOnly(m.manager, cValueid))
}

func (m *Manager) IsValueSet(valueid *ValueID) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_isValueSet(m.manager, cValueid))
}

func (m *Manager) IsValuePolled(valueid *ValueID) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_isValuePolled(m.manager, cValueid))
}

func (m *Manager) GetValueAsBool(valueid *ValueID) (bool, bool) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cBool C.bool
	result := bool(C.manager_getValueAsBool(m.manager, cValueid, &cBool))
	return result, bool(cBool)
}

func (m *Manager) GetValueAsByte(valueid *ValueID) (bool, byte) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cByte C.uint8_t
	result := bool(C.manager_getValueAsByte(m.manager, cValueid, &cByte))
	return result, byte(cByte)
}

func (m *Manager) GetValueAsFloat(valueid *ValueID) (bool, float32) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cFloat C.float
	result := bool(C.manager_getValueAsFloat(m.manager, cValueid, &cFloat))
	return result, float32(cFloat)
}

func (m *Manager) GetValueAsInt(valueid *ValueID) (bool, int32) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cInt C.int32_t
	result := bool(C.manager_getValueAsInt(m.manager, cValueid, &cInt))
	return result, int32(cInt)
}

func (m *Manager) GetValueAsShort(valueid *ValueID) (bool, int16) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cShort C.int16_t
	result := bool(C.manager_getValueAsShort(m.manager, cValueid, &cShort))
	return result, int16(cShort)
}

func (m *Manager) GetValueAsString(valueid *ValueID) (bool, string) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.string_emptyString()
	result := bool(C.manager_getValueAsString(m.manager, cValueid, cString))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return result, goString
}

func (m *Manager) GetValueAsRaw(valueid *ValueID) (bool, []byte) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cBytes := C.string_emptyBytes()
	result := bool(C.manager_getValueAsRaw(m.manager, cValueid, cBytes))
	goBytes := make([]byte, int(cBytes.length))
	for i := 0; i < int(cBytes.length); i++ {
		goBytes[i] = byte(C.string_byteAt(cBytes, C.size_t(i)))
	}
	return result, goBytes
}

func (m *Manager) GetValueListSelectionAsString(valueid *ValueID) (bool, string) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.string_emptyString()
	result := bool(C.manager_getValueListSelectionAsString(m.manager, cValueid, cString))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return result, goString
}

func (m *Manager) GetValueListSelectionAsInt32(valueid *ValueID) (bool, int32) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cInt C.int32_t
	result := bool(C.manager_getValueListSelectionAsInt32(m.manager, cValueid, &cInt))
	return result, int32(cInt)
}

func (m *Manager) GetValueListItems(valueid *ValueID) (bool, []string) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cStringList := C.string_emptyStringList()
	result := bool(C.manager_getValueListItems(m.manager, cValueid, cStringList))
	goStringList := make([]string, int(cStringList.length))
	for i := 0; i < int(cStringList.length); i++ {
		cString := C.string_stringAt(cStringList, C.size_t(i))
		goStringList[i] = C.GoString(cString.data)
	}
	C.string_freeStringList(cStringList)
	return result, goStringList
}

func (m *Manager) GetValueFloatPrecision(valueid *ValueID) (bool, uint8) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cPrecision C.uint8_t
	result := bool(C.manager_getValueFloatPrecision(m.manager, cValueid, &cPrecision))
	return result, uint8(cPrecision)
}

func (m *Manager) SetValueBool(valueid *ValueID, value bool) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setValueBool(m.manager, cValueid, C.bool(value)))
}

func (m *Manager) SetValueUint8(valueid *ValueID, value uint8) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setValueUint8(m.manager, cValueid, C.uint8_t(value)))
}

func (m *Manager) SetValueFloat(valueid *ValueID, value float32) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setValueFloat(m.manager, cValueid, C.float(value)))
}

func (m *Manager) SetValueInt32(valueid *ValueID, value int32) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setValueInt32(m.manager, cValueid, C.int32_t(value)))
}

func (m *Manager) SetValueInt16(valueid *ValueID, value int16) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setValueInt16(m.manager, cValueid, C.int16_t(value)))
}

func (m *Manager) SetValueBytes(valueid *ValueID, value []byte) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cBytes := C.string_emptyBytes()
	C.string_initBytes(cBytes, C.size_t(len(value)))
	for i := range value {
		C.string_setByteAt(cBytes, C.uint8_t(value[i]), C.size_t(i))
	}
	result := bool(C.manager_setValueBytes(m.manager, cValueid, cBytes))
	C.string_freeBytes(cBytes)
	return result
}

func (m *Manager) SetValueString(valueid *ValueID, value string) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.CString(value)
	result := bool(C.manager_setValueString(m.manager, cValueid, cString))
	C.free(unsafe.Pointer(cString))
	return result
}

func (m *Manager) SetValueListSelection(valueid *ValueID, selectedItem string) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.CString(selectedItem)
	result := bool(C.manager_setValueListSelection(m.manager, cValueid, cString))
	C.free(unsafe.Pointer(cString))
	return result
}

func (m *Manager) RefreshValue(valueid *ValueID) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_refreshValue(m.manager, cValueid))
}

func (m *Manager) SetChangeVerified(valueid *ValueID, verify bool) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	C.manager_setChangeVerified(m.manager, cValueid, C.bool(verify))
}

func (m *Manager) GetChangeVerified(valueid *ValueID) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_getChangeVerified(m.manager, cValueid))
}

func (m *Manager) PressButton(valueid *ValueID) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_pressButton(m.manager, cValueid))
}

func (m *Manager) ReleaseButton(valueid *ValueID) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_releaseButton(m.manager, cValueid))
}

//
// Climate control schedules.
//

func (m *Manager) GetNumSwitchPoints(valueid *ValueID) uint8 {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return uint8(C.manager_getNumSwitchPoints(m.manager, cValueid))
}

func (m *Manager) SetSwitchPoint(valueid *ValueID, hours, minutes uint8, setback int8) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setSwitchPoint(m.manager, cValueid, C.uint8_t(hours), C.uint8_t(minutes), C.int8_t(setback)))
}

func (m *Manager) RemoveSwitchPoint(valueid *ValueID, hours, minutes uint8) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_removeSwitchPoint(m.manager, cValueid, C.uint8_t(hours), C.uint8_t(minutes)))
}

func (m *Manager) ClearSwitchPoints(valueid *ValueID) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	C.manager_clearSwitchPoints(m.manager, cValueid)
}

func (m *Manager) GetSwitchPoint(valueid *ValueID, idx uint8) (bool, uint8, uint8, int8) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cHours C.uint8_t
	var cMinutes C.uint8_t
	var cSetback C.int8_t
	result := bool(C.manager_getSwitchPoint(m.manager, cValueid, C.uint8_t(idx), &cHours, &cMinutes, &cSetback))
	return result, uint8(cHours), uint8(cMinutes), int8(cSetback)
}

//
// Switch all.
//

func (m *Manager) SwitchAllOn(homeId uint32) {
	C.manager_switchAllOn(m.manager, C.uint32_t(homeId))
}

func (m *Manager) SwitchAllOff(homeId uint32) {
	C.manager_switchAllOff(m.manager, C.uint32_t(homeId))
}

//
// Configuration parameters.
//

func (m *Manager) SetConfigParam(homeId uint32, nodeId uint8, param uint8, value int32, size uint8) bool {
	return bool(C.manager_setConfigParam(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId), C.uint8_t(param), C.int32_t(value), C.uint8_t(size)))
}

func (m *Manager) RequestConfigParam(homeId uint32, nodeId uint8, param uint8) {
	C.manager_requestConfigParam(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId), C.uint8_t(param))
}

func (m *Manager) RequestAllConfigParam(homeId uint32, nodeId uint8) {
	C.manager_requestAllConfigParams(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId))
}

//
// Groups.
//

func (m *Manager) GetNumGroups(homeId uint32, nodeId uint8) uint8 {
	return uint8(C.manager_getNumGroups(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

//TODO func (m *Manager) GetAssociations(homeId uint32, nodeId uint8, groupIdx uint8) ...
//TODO func (m *Manager) GetAssociations(homeId uint32, nodeId uint8, groupIdx uint8) ...

func (m *Manager) GetMaxAssociations(homeId uint32, nodeId uint8, groupIdx uint8) uint8 {
	return uint8(C.manager_getMaxAssociations(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId), C.uint8_t(groupIdx)))
}

func (m *Manager) GetGroupLabel(homeId uint32, nodeId uint8, groupIdx uint8) string {
	cString := C.manager_getGroupLabel(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId), C.uint8_t(groupIdx))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

func (m *Manager) AddAssociation(homeId uint32, nodeId uint8, groupIdx uint8, targetNodeId uint8, instance uint8) {
	C.manager_addAssociation(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId), C.uint8_t(groupIdx), C.uint8_t(targetNodeId), C.uint8_t(instance))
}

func (m *Manager) RemoveAssociation(homeId uint32, nodeId uint8, groupIdx uint8, targetNodeId uint8, instance uint8) {
	C.manager_removeAssociation(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId), C.uint8_t(groupIdx), C.uint8_t(targetNodeId), C.uint8_t(instance))
}

//
// Notifications.
//

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

func (m *Manager) StartNotifications() error {
	themanager := unsafe.Pointer(m)
	result := C.manager_addWatcher(m.manager, themanager)
	if result {
		return nil
	}
	return fmt.Errorf("failed to add watcher")
}

func (m *Manager) StopNotifications() error {
	themanager := unsafe.Pointer(m)
	result := C.manager_removeWatcher(m.manager, themanager)
	if result {
		return nil
	}
	return fmt.Errorf("failed to remove watcher")
}

//
// Controller commands.
//

func (m *Manager) ResetController(homeId uint32) {
	C.manager_resetController(m.manager, C.uint32_t(homeId))
}

func (m *Manager) softReset(homeId uint32) {
	C.manager_softReset(m.manager, C.uint32_t(homeId))
}

func (m *Manager) CancelControllerCommand(homeId uint32) {
	C.manager_cancelControllerCommand(m.manager, C.uint32_t(homeId))
}

//
// Network commands.
//

func (m *Manager) TestNetworkNode(homeId uint32, nodeId uint8, count uint32) {
	C.manager_testNetworkNode(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId), C.uint32_t(count))
}

func (m *Manager) TestNetwork(homeId uint32, count uint32) {
	C.manager_testNetwork(m.manager, C.uint32_t(homeId), C.uint32_t(count))
}

func (m *Manager) HealNetworkNode(homeId uint32, nodeId uint8, doRR bool) {
	C.manager_healNetworkNode(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId), C.bool(doRR))
}

func (m *Manager) HealNetwork(homeId uint32, doRR bool) {
	C.manager_healNetwork(m.manager, C.uint32_t(homeId), C.bool(doRR))
}

func (m *Manager) AddNode(homeId uint32, doSecurity bool) bool {
	return bool(C.manager_addNode(m.manager, C.uint32_t(homeId), C.bool(doSecurity)))
}

func (m *Manager) RemoveNode(homeId uint32) bool {
	return bool(C.manager_removeNode(m.manager, C.uint32_t(homeId)))
}

func (m *Manager) RemoveFailedNode(homeId uint32, nodeId uint8) bool {
	return bool(C.manager_removeFailedNode(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) HasNodeFailed(homeId uint32, nodeId uint8) bool {
	return bool(C.manager_hasNodeFailed(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) RequestNodeNeighborUpdate(homeId uint32, nodeId uint8) bool {
	return bool(C.manager_requestNodeNeighborUpdate(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) AssignReturnRoute(homeId uint32, nodeId uint8) bool {
	return bool(C.manager_assignReturnRoute(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) DeleteAllReturnRoutes(homeId uint32, nodeId uint8) bool {
	return bool(C.manager_deleteAllReturnRoutes(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) SendNodeInformation(homeId uint32, nodeId uint8) bool {
	return bool(C.manager_sendNodeInformation(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) CreateNewPrimary(homeId uint32) bool {
	return bool(C.manager_createNewPrimary(m.manager, C.uint32_t(homeId)))
}

func (m *Manager) ReceiveConfiguration(homeId uint32) bool {
	return bool(C.manager_receiveConfiguration(m.manager, C.uint32_t(homeId)))
}
func (m *Manager) ReplaceFailedNode(homeId uint32, nodeId uint8) bool {
	return bool(C.manager_replaceFailedNode(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) TransferPrimaryRole(homeId uint32) bool {
	return bool(C.manager_transferPrimaryRole(m.manager, C.uint32_t(homeId)))
}

func (m *Manager) RequestNetworkUpdate(homeId uint32, nodeId uint8) bool {
	return bool(C.manager_requestNetworkUpdate(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) ReplicationSend(homeId uint32, nodeId uint8) bool {
	return bool(C.manager_replicationSend(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId)))
}

func (m *Manager) CreateButton(homeId uint32, nodeId uint8, buttonId uint8) bool {
	return bool(C.manager_createButton(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId), C.uint8_t(buttonId)))
}

func (m *Manager) DeleteButton(homeId uint32, nodeId uint8, buttonId uint8) bool {
	return bool(C.manager_deleteButton(m.manager, C.uint32_t(homeId), C.uint8_t(nodeId), C.uint8_t(buttonId)))
}

//
// Scene commands.
//

func (m *Manager) GetNumScenes() uint8 {
	return uint8(C.manager_getNumScenes(m.manager))
}

//TODO func (m *Manager) GetAllScenes(...) ...

func (m *Manager) RemoveAllScenes(homeId uint32) {
	C.manager_removeAllScenes(m.manager, C.uint32_t(homeId))
}

func (m *Manager) CreateScene() uint8 {
	return uint8(C.manager_createScene(m.manager))
}

func (m *Manager) RemoveScene(sceneId uint8) bool {
	return bool(C.manager_removeScene(m.manager, C.uint8_t(sceneId)))
}

func (m *Manager) AddSceneValueBool(sceneId uint8, valueid *ValueID, value bool) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_addSceneValueBool(m.manager, C.uint8_t(sceneId), cValueid, C.bool(value)))
}

func (m *Manager) AddSceneValueUint8(sceneId uint8, valueid *ValueID, value uint8) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_addSceneValueUint8(m.manager, C.uint8_t(sceneId), cValueid, C.uint8_t(value)))
}

func (m *Manager) AddSceneValueFloat(sceneId uint8, valueid *ValueID, value float32) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_addSceneValueFloat(m.manager, C.uint8_t(sceneId), cValueid, C.float(value)))
}

func (m *Manager) AddSceneValueInt32(sceneId uint8, valueid *ValueID, value int32) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_addSceneValueInt32(m.manager, C.uint8_t(sceneId), cValueid, C.int32_t(value)))
}

func (m *Manager) AddSceneValueInt16(sceneId uint8, valueid *ValueID, value int16) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_addSceneValueInt16(m.manager, C.uint8_t(sceneId), cValueid, C.int16_t(value)))
}

func (m *Manager) AddSceneValueString(sceneId uint8, valueid *ValueID, value string) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.CString(value)
	result := bool(C.manager_addSceneValueString(m.manager, C.uint8_t(sceneId), cValueid, cString))
	C.free(unsafe.Pointer(cString))
	return result
}

func (m *Manager) AddSceneValueListSelectionString(sceneId uint8, valueid *ValueID, value string) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.CString(value)
	result := bool(C.manager_addSceneValueListSelectionString(m.manager, C.uint8_t(sceneId), cValueid, cString))
	C.free(unsafe.Pointer(cString))
	return result
}

func (m *Manager) AddSceneValueListSelectionInt32(sceneId uint8, valueid *ValueID, value int32) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_addSceneValueListSelectionInt32(m.manager, C.uint8_t(sceneId), cValueid, C.int32_t(value)))
}

func (m *Manager) GetSceneValueAsBool(sceneId uint8, valueid *ValueID) (bool, bool) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cBool C.bool
	result := bool(C.manager_sceneGetValueAsBool(m.manager, C.uint8_t(sceneId), cValueid, &cBool))
	return result, bool(cBool)
}

func (m *Manager) GetSceneValueAsByte(sceneId uint8, valueid *ValueID) (bool, byte) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cByte C.uint8_t
	result := bool(C.manager_sceneGetValueAsByte(m.manager, C.uint8_t(sceneId), cValueid, &cByte))
	return result, byte(cByte)
}

func (m *Manager) GetSceneValueAsFloat(sceneId uint8, valueid *ValueID) (bool, float32) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cFloat C.float
	result := bool(C.manager_sceneGetValueAsFloat(m.manager, C.uint8_t(sceneId), cValueid, &cFloat))
	return result, float32(cFloat)
}

func (m *Manager) GetSceneValueAsInt(sceneId uint8, valueid *ValueID) (bool, int32) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cInt C.int32_t
	result := bool(C.manager_sceneGetValueAsInt(m.manager, C.uint8_t(sceneId), cValueid, &cInt))
	return result, int32(cInt)
}

func (m *Manager) GetSceneValueAsShort(sceneId uint8, valueid *ValueID) (bool, int16) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cShort C.int16_t
	result := bool(C.manager_sceneGetValueAsShort(m.manager, C.uint8_t(sceneId), cValueid, &cShort))
	return result, int16(cShort)
}

func (m *Manager) GetSceneValueAsString(sceneId uint8, valueid *ValueID) (bool, string) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.string_emptyString()
	result := bool(C.manager_sceneGetValueAsString(m.manager, C.uint8_t(sceneId), cValueid, cString))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return result, goString
}

func (m *Manager) GetSceneValueListSelectionString(sceneId uint8, valueid *ValueID) (bool, string) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.string_emptyString()
	result := bool(C.manager_sceneGetValueListSelectionString(m.manager, C.uint8_t(sceneId), cValueid, cString))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return result, goString
}

func (m *Manager) GetSceneValueListSelectionInt32(sceneId uint8, valueid *ValueID) (bool, int32) {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	var cInt C.int32_t
	result := bool(C.manager_sceneGetValueListSelectionInt32(m.manager, C.uint8_t(sceneId), cValueid, &cInt))
	return result, int32(cInt)
}

func (m *Manager) SetSceneValueBool(sceneId uint8, valueid *ValueID, value bool) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setSceneValueBool(m.manager, C.uint8_t(sceneId), cValueid, C.bool(value)))
}

func (m *Manager) SetSceneValueUint8(sceneId uint8, valueid *ValueID, value uint8) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setSceneValueUint8(m.manager, C.uint8_t(sceneId), cValueid, C.uint8_t(value)))
}

func (m *Manager) SetSceneValueFloat(sceneId uint8, valueid *ValueID, value float32) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setSceneValueFloat(m.manager, C.uint8_t(sceneId), cValueid, C.float(value)))
}

func (m *Manager) SetSceneValueInt32(sceneId uint8, valueid *ValueID, value int32) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setSceneValueInt32(m.manager, C.uint8_t(sceneId), cValueid, C.int32_t(value)))
}

func (m *Manager) SetSceneValueInt16(sceneId uint8, valueid *ValueID, value int16) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setSceneValueInt16(m.manager, C.uint8_t(sceneId), cValueid, C.int16_t(value)))
}

func (m *Manager) SetSceneValueString(sceneId uint8, valueid *ValueID, value string) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.CString(value)
	result := bool(C.manager_setSceneValueString(m.manager, C.uint8_t(sceneId), cValueid, cString))
	C.free(unsafe.Pointer(cString))
	return result
}

func (m *Manager) SetSceneValueListSelectionString(sceneId uint8, valueid *ValueID, value string) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	cString := C.CString(value)
	result := bool(C.manager_setSceneValueListSelectionString(m.manager, C.uint8_t(sceneId), cValueid, cString))
	C.free(unsafe.Pointer(cString))
	return result
}

func (m *Manager) SetSceneValueListSelectionInt32(sceneId uint8, valueid *ValueID, value int32) bool {
	cValueid := valueid.toC()
	defer C.valueid_free(cValueid)
	return bool(C.manager_setSceneValueListSelectionInt32(m.manager, C.uint8_t(sceneId), cValueid, C.int32_t(value)))
}

func (m *Manager) GetSceneLabel(sceneId uint8) string {
	cString := C.manager_getSceneLabel(m.manager, C.uint8_t(sceneId))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

func (m *Manager) SetSceneLabel(sceneId uint8, value string) {
	cString := C.CString(value)
	C.manager_setSceneLabel(m.manager, C.uint8_t(sceneId), cString)
	C.free(unsafe.Pointer(cString))
}

func (m *Manager) SceneExists(sceneId uint8) bool {
	return bool(C.manager_sceneExists(m.manager, C.uint8_t(sceneId)))
}

func (m *Manager) ActivateScene(sceneId uint8) bool {
	return bool(C.manager_activateScene(m.manager, C.uint8_t(sceneId)))
}

//
// Statistics retreival interface.
//

//TODO func (m *Manager) GetDriverStatistics(...) ...
//TODO func (m *Manager) GetNodeStatistics(...) ...
