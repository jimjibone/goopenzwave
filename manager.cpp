#include "manager.h"
#include <openzwave/Manager.h>
#include <openzwave/Notification.h>
#include <openzwave/Defs.h>

#include <iostream>

//
// Construction.
//

manager_t manager_create()
{
	OpenZWave::Manager *man = OpenZWave::Manager::Create();
	return (manager_t)man;
}

manager_t manager_get()
{
	OpenZWave::Manager *man = OpenZWave::Manager::Get();
	return (manager_t)man;
}

void manager_destroy()
{
	OpenZWave::Manager::Destroy();
}

string_t* manager_getVersionAsString()
{
	std::string str = OpenZWave::Manager::getVersionAsString();
	return string_fromStdString(str);
}

string_t* manager_getVersionLongAsString()
{
	std::string str = OpenZWave::Manager::getVersionLongAsString();
	return string_fromStdString(str);
}

void manager_getVersion(uint16_t *major, uint16_t *minor)
{
	ozwversion ver = OpenZWave::Manager::getVersion();
	*major = ver._v >> 16 & 0x00FF;
	*minor = ver._v & 0x00FF;
}

//
// Configuration.
//

void manager_writeConfig(manager_t m, uint32_t homeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	man->WriteConfig(homeId);
}

options_t manager_getOptions(manager_t m)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return (options_t)man->GetOptions();
}

//
// Drivers.
//

bool manager_addDriver(manager_t m, const char *controllerPath)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	if (strcasecmp(controllerPath, "usb") == 0) {
		return man->AddDriver("HID Controller", OpenZWave::Driver::ControllerInterface_Hid);
	} else {
		return man->AddDriver(controllerPath);
	}
}

bool manager_removeDriver(manager_t m, const char *controllerPath)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	if (strcasecmp(controllerPath, "usb") == 0) {
		return man->RemoveDriver("HID Controller");
	} else {
		return man->RemoveDriver(controllerPath);
	}
}

uint8_t manager_getControllerNodeId(manager_t m, uint32_t homeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->GetControllerNodeId(homeId);
}

uint8_t manager_getSUCNodeId(manager_t m, uint32_t homeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->GetSUCNodeId(homeId);
}

bool manager_isPrimaryController(manager_t m, uint32_t homeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->IsPrimaryController(homeId);
}

bool manager_isStaticUpdateController(manager_t m, uint32_t homeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->IsStaticUpdateController(homeId);
}

bool manager_isBridgeController(manager_t m, uint32_t homeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->IsBridgeController(homeId);
}

string_t* manager_getLibraryVersion(manager_t m, uint32_t homeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str = man->GetLibraryVersion(homeId);
	return string_fromStdString(str);
}

string_t* manager_getLibraryTypeName(manager_t m, uint32_t homeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str = man->GetLibraryTypeName(homeId);
	return string_fromStdString(str);
}

int32_t manager_getSendQueueCount(manager_t m, uint32_t homeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->GetSendQueueCount(homeId);
}

void manager_logDriverStatistics(manager_t m, uint32_t homeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	man->LogDriverStatistics(homeId);
}

//TODO driver_controllerInterface_t manager_getControllerInterfaceType(manager_t m, uint32_t homeId)
// {
// 	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
// 	return man->GetControllerInterfaceType(homeId);
// }

string_t* manager_getControllerPath(manager_t m, uint32_t homeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str = man->GetControllerPath(homeId);
	return string_fromStdString(str);
}

//
// Polling Z-Wave devices.
//

int32_t manager_getPollInterval(manager_t m)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->GetPollInterval();
}

void manager_setPollInterval(manager_t m, int32_t milliseconds, bool intervalBetweenPolls)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	man->SetPollInterval(milliseconds, intervalBetweenPolls);
}

bool manager_enablePoll(manager_t m, valueid_t valueid, uint8_t intensity)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->EnablePoll(*val, intensity);
}

bool manager_disablePoll(manager_t m, valueid_t valueid)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->DisablePoll(*val);
}

bool manager_isPolled(manager_t m, valueid_t valueid)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->isPolled(*val);
}

void manager_setPollIntensity(manager_t m, valueid_t valueid, uint8_t intensity)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	man->SetPollIntensity(*val, intensity);
}

uint8_t manager_getPollIntensity(manager_t m, valueid_t valueid)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->DisablePoll(*val);
}

//
// Node information.
//

bool manager_refreshNodeInfo(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->RefreshNodeInfo(homeId, nodeId);
}

bool manager_requestNodeState(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->RequestNodeState(homeId, nodeId);
}

bool manager_requestNodeDynamic(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->RequestNodeDynamic(homeId, nodeId);
}

bool manager_isNodeListeningDevice(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->IsNodeListeningDevice(homeId, nodeId);
}

bool manager_isNodeFrequentListeningDevice(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->IsNodeFrequentListeningDevice(homeId, nodeId);
}

bool manager_isNodeBeamingDevice(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->IsNodeBeamingDevice(homeId, nodeId);
}

bool manager_isNodeRoutingDevice(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->IsNodeRoutingDevice(homeId, nodeId);
}

bool manager_isNodeSecurityDevice(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->IsNodeSecurityDevice(homeId, nodeId);
}

uint32_t manager_getNodeMaxBaudRate(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->GetNodeMaxBaudRate(homeId, nodeId);
}

uint8_t manager_getNodeVersion(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->GetNodeVersion(homeId, nodeId);
}

uint8_t manager_getNodeSecurity(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->GetNodeSecurity(homeId, nodeId);
}

bool manager_isNodeZWavePlus(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->IsNodeZWavePlus(homeId, nodeId);
}

uint8_t manager_getNodeBasic(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->GetNodeBasic(homeId, nodeId);
}

uint8_t manager_getNodeGeneric(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->GetNodeGeneric(homeId, nodeId);
}

uint8_t manager_getNodeSpecific(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->GetNodeSpecific(homeId, nodeId);
}

string_t* manager_getNodeType(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str = man->GetNodeType(homeId, nodeId);
	return string_fromStdString(str);
}

//TODO uint32_t manager_getNodeNeighbors(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t **nodeNeighbors)
// {
// 	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
// 	return man->GetNodeNeighbours(homeId, nodeId);
// }

string_t* manager_getNodeManufacturerName(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str = man->GetNodeManufacturerName(homeId, nodeId);
	return string_fromStdString(str);
}

string_t* manager_getNodeProductName(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str = man->GetNodeProductName(homeId, nodeId);
	return string_fromStdString(str);
}

string_t* manager_getNodeName(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str = man->GetNodeName(homeId, nodeId);
	return string_fromStdString(str);
}

string_t* manager_getNodeLocation(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str = man->GetNodeLocation(homeId, nodeId);
	return string_fromStdString(str);
}

string_t* manager_getNodeManufacturerId(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str = man->GetNodeManufacturerId(homeId, nodeId);
	return string_fromStdString(str);
}

string_t* manager_getNodeProductType(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str = man->GetNodeProductType(homeId, nodeId);
	return string_fromStdString(str);
}

string_t* manager_getNodeProductId(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str = man->GetNodeProductId(homeId, nodeId);
	return string_fromStdString(str);
}

void manager_setNodeManufacturerName(manager_t m, uint32_t homeId, uint8_t nodeId, const char* manufacturerName)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str(manufacturerName);
	man->SetNodeManufacturerName(homeId, nodeId, str);
}

void manager_setNodeProductName(manager_t m, uint32_t homeId, uint8_t nodeId, const char* productName)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str(productName);
	man->SetNodeProductName(homeId, nodeId, str);
}

void manager_setNodeName(manager_t m, uint32_t homeId, uint8_t nodeId, const char* nodeName)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str(nodeName);
	man->SetNodeName(homeId, nodeId, str);
}

void manager_setNodeLocation(manager_t m, uint32_t homeId, uint8_t nodeId, const char* location)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str(location);
	man->SetNodeLocation(homeId, nodeId, str);
}

void manager_setNodeOn(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	man->SetNodeOn(homeId, nodeId);
}

void manager_setNodeOff(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	man->SetNodeOff(homeId, nodeId);
}

void manager_setNodeLevel(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t level)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	man->SetNodeLevel(homeId, nodeId, level);
}

bool manager_isNodeInfoReceived(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->IsNodeInfoReceived(homeId, nodeId);
}

bool manager_getNodeClassInformation(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t commandClassId, string_t *o_name, uint8_t *o_version)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	bool result;
	if (o_name == NULL) {
		result = man->GetNodeClassInformation(homeId, nodeId, commandClassId, NULL, o_version);
	} else {
		std::string str;
		result = man->GetNodeClassInformation(homeId, nodeId, commandClassId, &str, o_version);
		string_copyStdString(o_name, str);
	}
	return result;
}

bool manager_isNodeAwake(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->IsNodeAwake(homeId, nodeId);
}

bool manager_isNodeFailed(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->IsNodeFailed(homeId, nodeId);
}

string_t* manager_getNodeQueryStage(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str = man->GetNodeQueryStage(homeId, nodeId);
	return string_fromStdString(str);
}

uint16_t manager_getNodeDeviceType(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->GetNodeDeviceType(homeId, nodeId);
}

string_t* manager_getNodeDeviceTypeString(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str = man->GetNodeDeviceTypeString(homeId, nodeId);
	return string_fromStdString(str);
}

uint8_t manager_getNodeRole(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->GetNodeRole(homeId, nodeId);
}

string_t* manager_getNodeRoleString(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str = man->GetNodeRoleString(homeId, nodeId);
	return string_fromStdString(str);
}

uint8_t manager_getNodePlusType(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->GetNodePlusType(homeId, nodeId);
}

string_t* manager_getNodePlusTypeString(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str = man->GetNodePlusTypeString(homeId, nodeId);
	return string_fromStdString(str);
}

//
// Values.
//

string_t* manager_getValueLabel(manager_t m, valueid_t valueid)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	std::string str = man->GetValueLabel(*val);
	return string_fromStdString(str);
}

void manager_setValueLabel(manager_t m, valueid_t valueid, const char* value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	std::string str(value);
	man->SetValueLabel(*val, str);
}

string_t* manager_getValueUnits(manager_t m, valueid_t valueid)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	std::string str = man->GetValueUnits(*val);
	return string_fromStdString(str);
}

void manager_setValueUnits(manager_t m, valueid_t valueid, const char* value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	std::string str(value);
	man->SetValueUnits(*val, str);
}

string_t* manager_getValueHelp(manager_t m, valueid_t valueid)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	std::string str = man->GetValueHelp(*val);
	return string_fromStdString(str);
}

void manager_setValueHelp(manager_t m, valueid_t valueid, const char* value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	std::string str(value);
	man->SetValueHelp(*val, str);
}

int32_t manager_getValueMin(manager_t m, valueid_t valueid)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->GetValueMin(*val);
}

int32_t manager_getValueMax(manager_t m, valueid_t valueid)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->GetValueMax(*val);
}

bool manager_isValueReadOnly(manager_t m, valueid_t valueid)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->IsValueReadOnly(*val);
}

bool manager_isValueWriteOnly(manager_t m, valueid_t valueid)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->IsValueWriteOnly(*val);
}

bool manager_isValueSet(manager_t m, valueid_t valueid)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->IsValueSet(*val);
}

bool manager_isValuePolled(manager_t m, valueid_t valueid)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->IsValuePolled(*val);
}

bool manager_getValueAsBool(manager_t m, valueid_t valueid, bool *o_value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->GetValueAsBool(*val, o_value);
}

bool manager_getValueAsByte(manager_t m, valueid_t valueid, uint8_t *o_value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->GetValueAsByte(*val, o_value);
}

bool manager_getValueAsFloat(manager_t m, valueid_t valueid, float *o_value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->GetValueAsFloat(*val, o_value);
}

bool manager_getValueAsInt(manager_t m, valueid_t valueid, int32_t *o_value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->GetValueAsInt(*val, o_value);
}

bool manager_getValueAsShort(manager_t m, valueid_t valueid, int16_t *o_value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->GetValueAsShort(*val, o_value);
}

bool manager_getValueAsString(manager_t m, valueid_t valueid, string_t *o_value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	std::string str;
	bool result = man->GetValueAsString(*val, &str);
	string_copyStdString(o_value, str);
	return result;
}

bool manager_getValueAsRaw(manager_t m, valueid_t valueid, bytes_t *o_value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	string_initBytes(o_value, 255);
	return man->GetValueAsRaw(*val, &(o_value->data), (uint8_t*)&(o_value->length));
}

bool manager_getValueListSelectionAsString(manager_t m, valueid_t valueid, string_t *o_value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	std::string str;
	bool result = man->GetValueListSelection(*val, &str);
	string_copyStdString(o_value, str);
	return result;
}

bool manager_getValueListSelectionAsInt32(manager_t m, valueid_t valueid, int32_t *o_value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->GetValueListSelection(*val, o_value);
}

bool manager_getValueListItems(manager_t m, valueid_t valueid, stringlist_t *o_value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	std::vector<std::string> list;
	bool result = man->GetValueListItems(*val, &list);
	string_copyStdStringList(o_value, list);
	return result;
}

bool manager_getValueFloatPrecision(manager_t m, valueid_t valueid, uint8_t *o_value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->GetValueFloatPrecision(*val, o_value);
}

bool manager_setValueBool(manager_t m, valueid_t valueid, bool value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->SetValue(*val, value);
}

bool manager_setValueUint8(manager_t m, valueid_t valueid, uint8_t value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->SetValue(*val, value);
}

bool manager_setValueFloat(manager_t m, valueid_t valueid, float value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->SetValue(*val, value);
}

bool manager_setValueInt32(manager_t m, valueid_t valueid, int32_t value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->SetValue(*val, value);
}

bool manager_setValueInt16(manager_t m, valueid_t valueid, int16_t value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->SetValue(*val, value);
}

bool manager_setValueBytes(manager_t m, valueid_t valueid, bytes_t *value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->SetValue(*val, value->data, (uint8_t)value->length);
}

bool manager_setValueString(manager_t m, valueid_t valueid, const char* value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	std::string str(value);
	return man->SetValue(*val, str);
}

bool manager_setValueListSelection(manager_t m, valueid_t valueid, const char* selectedItem)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	std::string str(selectedItem);
	return man->SetValueListSelection(*val, str);
}

bool manager_refreshValue(manager_t m, valueid_t valueid)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->RefreshValue(*val);
}

void manager_setChangeVerified(manager_t m, valueid_t valueid, bool verify)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	man->SetChangeVerified(*val, verify);
}

bool manager_getChangeVerified(manager_t m, valueid_t valueid)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->GetChangeVerified(*val);
}

bool manager_pressButton(manager_t m, valueid_t valueid)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->PressButton(*val);
}

bool manager_releaseButton(manager_t m, valueid_t valueid)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->ReleaseButton(*val);
}

//
// Climate control schedules.
//

uint8_t manager_getNumSwitchPoints(manager_t m, valueid_t valueid)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->GetNumSwitchPoints(*val);
}

bool manager_setSwitchPoint(manager_t m, valueid_t valueid, uint8_t hours, uint8_t minutes, int8_t setback)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->SetSwitchPoint(*val, hours, minutes, setback);
}

bool manager_removeSwitchPoint(manager_t m, valueid_t valueid, uint8_t hours, uint8_t minutes)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->RemoveSwitchPoint(*val, hours, minutes);
}

void manager_clearSwitchPoints(manager_t m, valueid_t valueid)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->ClearSwitchPoints(*val);
}

bool manager_getSwitchPoint(manager_t m, valueid_t valueid, uint8_t idx, uint8_t *o_hours, uint8_t *o_minutes, int8_t *o_setback)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->GetSwitchPoint(*val, idx, o_hours, o_minutes, o_setback);
}

//
// Switch all.
//

void manager_switchAllOn(manager_t m, uint32_t homeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	man->SwitchAllOn(homeId);
}

void manager_switchAllOff(manager_t m, uint32_t homeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	man->SwitchAllOff(homeId);
}


//
// Configuration parameters.
//

bool manager_setConfigParam(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t param, int32_t value, uint8_t size)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->SetConfigParam(homeId, nodeId, param, value, size);
}

void manager_requestConfigParam(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t param)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	man->RequestConfigParam(homeId, nodeId, param);
}

void manager_requestAllConfigParams(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	man->RequestAllConfigParams(homeId, nodeId);
}

//
// Groups.
//

uint8_t manager_getNumGroups(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->GetNumGroups(homeId, nodeId);
}

uint32_t manager_getAssociations(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t groupIdx, uint8_t **o_associations)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->GetAssociations(homeId, nodeId, groupIdx, o_associations);
}

//TODO uint32_t manager_getAssociations(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t groupIdx, instanceassociation_t **o_associations)
// {
// 	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
// 	return man->GetAssociations(homeId, nodeId, groupIdx, ...);
// }

uint8_t manager_getMaxAssociations(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t groupIdx)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->GetMaxAssociations(homeId, nodeId, groupIdx);
}

string_t* manager_getGroupLabel(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t groupIdx)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str = man->GetGroupLabel(homeId, nodeId, groupIdx);
	return string_fromStdString(str);
}

void manager_addAssociation(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t groupIdx, uint8_t targetNodeId, uint8_t instance)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	man->AddAssociation(homeId, nodeId, groupIdx, targetNodeId, instance);
}

void manager_removeAssociation(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t groupIdx, uint8_t targetNodeId, uint8_t instance)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	man->RemoveAssociation(homeId, nodeId, groupIdx, targetNodeId, instance);
}

//
// Notifications.
//

static void manager_notificationHandler(OpenZWave::Notification const* notification, void* userdata)
{
	// Note that OpenZWave will delete the notification object when it thinks we
	// are done with it. Probably when we return control to it.
	notification_t noti = (notification_t)notification;
	goNotificationCB(noti, userdata);
}

bool manager_addWatcher(manager_t m, void *userdata)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->AddWatcher(manager_notificationHandler, userdata);
}

bool manager_removeWatcher(manager_t m, void *userdata)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->RemoveWatcher(manager_notificationHandler, userdata);
}

//
// Controller commands.
//

void manager_resetController(manager_t m, uint32_t homeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	man->ResetController(homeId);
}

void manager_softReset(manager_t m, uint32_t homeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	man->SoftReset(homeId);
}

bool manager_cancelControllerCommand(manager_t m, uint32_t homeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->CancelControllerCommand(homeId);
}

//
// Network commands.
//

void manager_testNetworkNode(manager_t m, uint32_t homeId, uint8_t nodeId, uint32_t count)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	man->TestNetworkNode(homeId, nodeId, count);
}

void manager_testNetwork(manager_t m, uint32_t homeId, uint32_t count)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	man->TestNetwork(homeId, count);
}

void manager_healNetworkNode(manager_t m, uint32_t homeId, uint8_t nodeId, bool doRR)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	man->HealNetworkNode(homeId, nodeId, doRR);
}

void manager_healNetwork(manager_t m, uint32_t homeId, bool doRR)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	man->HealNetwork(homeId, doRR);
}

bool manager_addNode(manager_t m, uint32_t homeId, bool doSecurity)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->AddNode(homeId, doSecurity);
}

bool manager_removeNode(manager_t m, uint32_t homeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->RemoveNode(homeId);
}

bool manager_removeFailedNode(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->RemoveFailedNode(homeId, nodeId);
}

bool manager_hasNodeFailed(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->HasNodeFailed(homeId, nodeId);
}

bool manager_requestNodeNeighborUpdate(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->RequestNodeNeighborUpdate(homeId, nodeId);
}

bool manager_assignReturnRoute(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->AssignReturnRoute(homeId, nodeId);
}

bool manager_deleteAllReturnRoutes(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->DeleteAllReturnRoutes(homeId, nodeId);
}

bool manager_sendNodeInformation(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->SendNodeInformation(homeId, nodeId);
}

bool manager_createNewPrimary(manager_t m, uint32_t homeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->CreateNewPrimary(homeId);
}

bool manager_receiveConfiguration(manager_t m, uint32_t homeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->ReceiveConfiguration(homeId);
}

bool manager_replaceFailedNode(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->ReplaceFailedNode(homeId, nodeId);
}

bool manager_transferPrimaryRole(manager_t m, uint32_t homeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->TransferPrimaryRole(homeId);
}

bool manager_requestNetworkUpdate(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->RequestNetworkUpdate(homeId, nodeId);
}

bool manager_replicationSend(manager_t m, uint32_t homeId, uint8_t nodeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->ReplicationSend(homeId, nodeId);
}

bool manager_createButton(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t buttonid)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->CreateButton(homeId, nodeId, buttonid);
}

bool manager_deleteButton(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t buttonid)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->DeleteButton(homeId, nodeId, buttonid);
}

//
// Scene commands.
//

uint8_t manager_getNumScenes(manager_t m)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->GetNumScenes();
}

uint8_t manager_getAllScenes(manager_t m, uint8_t **sceneIds)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->GetAllScenes(sceneIds);
}

void manager_removeAllScenes(manager_t m, uint32_t homeId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	man->RemoveAllScenes(homeId);
}

uint8_t manager_createScene(manager_t m)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->CreateScene();
}

bool manager_removeScene(manager_t m, uint8_t sceneId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->RemoveScene(sceneId);
}

bool manager_addSceneValueBool(manager_t m, uint8_t sceneId, valueid_t valueid, bool value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->AddSceneValue(sceneId, *val, value);
}

bool manager_addSceneValueUint8(manager_t m, uint8_t sceneId, valueid_t valueid, uint8_t value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->AddSceneValue(sceneId, *val, value);
}

bool manager_addSceneValueFloat(manager_t m, uint8_t sceneId, valueid_t valueid, float value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->AddSceneValue(sceneId, *val, value);
}

bool manager_addSceneValueInt32(manager_t m, uint8_t sceneId, valueid_t valueid, int32_t value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->AddSceneValue(sceneId, *val, value);
}

bool manager_addSceneValueInt16(manager_t m, uint8_t sceneId, valueid_t valueid, int16_t value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->AddSceneValue(sceneId, *val, value);
}

bool manager_addSceneValueString(manager_t m, uint8_t sceneId, valueid_t valueid, const char* value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	std::string str(value);
	return man->AddSceneValue(sceneId, *val, str);
}

bool manager_addSceneValueListSelectionString(manager_t m, uint8_t sceneId, valueid_t valueid, const char* value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	std::string str(value);
	return man->AddSceneValueListSelection(sceneId, *val, str);
}

bool manager_addSceneValueListSelectionInt32(manager_t m, uint8_t sceneId, valueid_t valueid, int32_t value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->AddSceneValueListSelection(sceneId, *val, value);
}

bool manager_removeSceneValue(manager_t m, uint8_t sceneId, valueid_t valueid)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->RemoveSceneValue(sceneId, *val);
}

//TODO int manager_sceneGetValues(manager_t m, uint8_t sceneId, valueidlist_t *o_value)
// {
// 	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
// 	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
// 	std::vector<OpenZWave::ValueID> values;
// 	int result = man->SceneGetValues(sceneId, *val, &values);
// 	valueid_copyValueIDList(o_value, values);
// 	return result;
// }

bool manager_sceneGetValueAsBool(manager_t m, uint8_t sceneId, valueid_t valueid, bool *o_value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->SceneGetValueAsBool(sceneId, *val, o_value);
}

bool manager_sceneGetValueAsByte(manager_t m, uint8_t sceneId, valueid_t valueid, uint8_t *o_value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->SceneGetValueAsByte(sceneId, *val, o_value);
}

bool manager_sceneGetValueAsFloat(manager_t m, uint8_t sceneId, valueid_t valueid, float *o_value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->SceneGetValueAsFloat(sceneId, *val, o_value);
}

bool manager_sceneGetValueAsInt(manager_t m, uint8_t sceneId, valueid_t valueid, int32_t *o_value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->SceneGetValueAsInt(sceneId, *val, o_value);
}

bool manager_sceneGetValueAsShort(manager_t m, uint8_t sceneId, valueid_t valueid, int16_t *o_value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->SceneGetValueAsShort(sceneId, *val, o_value);
}

bool manager_sceneGetValueAsString(manager_t m, uint8_t sceneId, valueid_t valueid, string_t *o_value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	std::string str;
	bool result = man->SceneGetValueAsString(sceneId, *val, &str);
	string_copyStdString(o_value, str);
	return result;
}

bool manager_sceneGetValueListSelectionString(manager_t m, uint8_t sceneId, valueid_t valueid, string_t *o_value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	std::string str;
	bool result = man->SceneGetValueListSelection(sceneId, *val, &str);
	string_copyStdString(o_value, str);
	return result;
}

bool manager_sceneGetValueListSelectionInt32(manager_t m, uint8_t sceneId, valueid_t valueid, int32_t *o_value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->SceneGetValueListSelection(sceneId, *val, o_value);
}

bool manager_setSceneValueBool(manager_t m, uint8_t sceneId, valueid_t valueid, bool value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->SetSceneValue(sceneId, *val, value);
}

bool manager_setSceneValueUint8(manager_t m, uint8_t sceneId, valueid_t valueid, uint8_t value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->SetSceneValue(sceneId, *val, value);
}

bool manager_setSceneValueFloat(manager_t m, uint8_t sceneId, valueid_t valueid, float value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->SetSceneValue(sceneId, *val, value);
}

bool manager_setSceneValueInt32(manager_t m, uint8_t sceneId, valueid_t valueid, int32_t value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->SetSceneValue(sceneId, *val, value);
}

bool manager_setSceneValueInt16(manager_t m, uint8_t sceneId, valueid_t valueid, int16_t value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->SetSceneValue(sceneId, *val, value);
}

bool manager_setSceneValueString(manager_t m, uint8_t sceneId, valueid_t valueid, const char* value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	std::string str(value);
	return man->SetSceneValue(sceneId, *val, str);
}

bool manager_setSceneValueListSelectionString(manager_t m, uint8_t sceneId, valueid_t valueid, const char* value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	std::string str(value);
	return man->SetSceneValueListSelection(sceneId, *val, str);
}

bool manager_setSceneValueListSelectionInt32(manager_t m, uint8_t sceneId, valueid_t valueid, int32_t value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	OpenZWave::ValueID *val = (OpenZWave::ValueID*)valueid;
	return man->SetSceneValueListSelection(sceneId, *val, value);
}

string_t* manager_getSceneLabel(manager_t m, uint8_t sceneId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str = man->GetSceneLabel(sceneId);
	return string_fromStdString(str);
}

void manager_setSceneLabel(manager_t m, uint8_t sceneId, const char* value)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	std::string str(value);
	man->SetSceneLabel(sceneId, str);
}

bool manager_sceneExists(manager_t m, uint8_t sceneId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->SceneExists(sceneId);
}

bool manager_activateScene(manager_t m, uint8_t sceneId)
{
	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
	return man->ActivateScene(sceneId);
}

//
// Statistics retreival interface.
//

//TODO void manager_getDriverStatistics(manager_t m, uint32_t homeId, driver_driverdata_t *data)
// {
// 	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
// 	OpenZWave::Driver::DriverData *driverData = (OpenZWave::Driver::DriverData*)data;
// 	man->GetDriverStatistics(homeId, driverData);
// }
//
//TODO void manager_getNodeStatistics(manager_t m, uint32_t homeId, uint8_t nodeId, node_nodedata_t *data)
// {
// 	OpenZWave::Manager *man = (OpenZWave::Manager*)m;
// 	OpenZWave::Node::NodeData *nodeData = (OpenZWave::Node::NodeData*)data;
// 	man->GetNodeStatistics(homeId, nodeId, nodeData);
// }
