#ifndef GOOPENZWAVE_MANAGER
#define GOOPENZWAVE_MANAGER

#include <stdint.h>
#include <stdbool.h>
#include <stddef.h>
#include "zwhelpers.h"
#include "options.h"
#include "notification.h"

#ifdef __cplusplus
extern "C" {
#endif

	// Types.
	typedef void* manager_t;

	//
	// Construction.
	//

	manager_t manager_create();
	manager_t manager_get();
	void manager_destroy();
	char* manager_getVersionAsString(); /*!< C string must be freed. */
	char* manager_getVersionLongAsString(); /*!< C string must be freed. */
	void manager_getVersion(uint16_t *major, uint16_t *minor);

	//
	// Configuration.
	//

	void manager_writeConfig(manager_t m, uint32_t homeId);
	options_t manager_getOptions(manager_t m);

	//
	// Drivers.
	//

	bool manager_addDriver(manager_t m, const char* controllerPath);
	bool manager_removeDriver(manager_t m, const char* controllerPath);
	uint8_t manager_getControllerNodeId(manager_t m, uint32_t homeId);
	uint8_t manager_getSUCNodeId(manager_t m, uint32_t homeId);
	bool manager_isPrimaryController(manager_t m, uint32_t homeId);
	bool manager_isStaticUpdateController(manager_t m, uint32_t homeId);
	bool manager_isBridgeController(manager_t m, uint32_t homeId);
	char* manager_getLibraryVersion(manager_t m, uint32_t homeId);  /*!< C string must be freed. */
	char* manager_getLibraryTypeName(manager_t m, uint32_t homeId); /*!< C string must be freed. */
	int32_t manager_getSendQueueCount(manager_t m, uint32_t homeId);
	void manager_logDriverStatistics(manager_t m, uint32_t homeId);
//TODO driver_controllerInterface_t manager_getControllerInterfaceType(manager_t m, uint32_t homeId);
	char* manager_getControllerPath(manager_t m, uint32_t homeId);  /*!< C string must be freed. */

	//
	// Polling Z-Wave devices.
	//

	int32_t manager_getPollInterval(manager_t m);
	void manager_setPollInterval(manager_t m, int32_t milliseconds, bool intervalBetweenPolls);
	bool manager_enablePoll(manager_t m, valueid_t valueid, uint8_t intensity);
	bool manager_disablePoll(manager_t m, valueid_t valueid);
	bool manager_isPolled(manager_t m, valueid_t valueid);
	void manager_setPollIntensity(manager_t m, valueid_t valueid, uint8_t intensity);
	uint8_t manager_getPollIntensity(manager_t m, valueid_t valueid);

	//
	// Node information.
	//

	bool manager_refreshNodeInfo(manager_t m, uint32_t homeId, uint8_t nodeId);
	bool manager_requestNodeState(manager_t m, uint32_t homeId, uint8_t nodeId);
	bool manager_requestNodeDynamic(manager_t m, uint32_t homeId, uint8_t nodeId);
	bool manager_isNodeListeningDevice(manager_t m, uint32_t homeId, uint8_t nodeId);
	bool manager_isNodeFrequentListeningDevice(manager_t m, uint32_t homeId, uint8_t nodeId);
	bool manager_isNodeBeamingDevice(manager_t m, uint32_t homeId, uint8_t nodeId);
	bool manager_isNodeRoutingDevice(manager_t m, uint32_t homeId, uint8_t nodeId);
	bool manager_isNodeSecurityDevice(manager_t m, uint32_t homeId, uint8_t nodeId);
	uint32_t manager_getNodeMaxBaudRate(manager_t m, uint32_t homeId, uint8_t nodeId);
	uint8_t manager_getNodeVersion(manager_t m, uint32_t homeId, uint8_t nodeId);
	uint8_t manager_getNodeSecurity(manager_t m, uint32_t homeId, uint8_t nodeId);
	bool manager_isNodeZWavePlus(manager_t m, uint32_t homeId, uint8_t nodeId);
	uint8_t manager_getNodeBasic(manager_t m, uint32_t homeId, uint8_t nodeId);
	uint8_t manager_getNodeGeneric(manager_t m, uint32_t homeId, uint8_t nodeId);
	uint8_t manager_getNodeSpecific(manager_t m, uint32_t homeId, uint8_t nodeId);
	char* manager_getNodeType(manager_t m, uint32_t homeId, uint8_t nodeId); /*!< C string must be freed. */
//TODO uint32_t manager_getNodeNeighbors(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t **nodeNeighbors);
	char* manager_getNodeManufacturerName(manager_t m, uint32_t homeId, uint8_t nodeId); /*!< C string must be freed. */
	char* manager_getNodeProductName(manager_t m, uint32_t homeId, uint8_t nodeId); /*!< C string must be freed. */
	char* manager_getNodeName(manager_t m, uint32_t homeId, uint8_t nodeId); /*!< C string must be freed. */
	char* manager_getNodeLocation(manager_t m, uint32_t homeId, uint8_t nodeId); /*!< C string must be freed. */
	char* manager_getNodeManufacturerId(manager_t m, uint32_t homeId, uint8_t nodeId); /*!< C string must be freed. */
	char* manager_getNodeProductType(manager_t m, uint32_t homeId, uint8_t nodeId); /*!< C string must be freed. */
	char* manager_getNodeProductId(manager_t m, uint32_t homeId, uint8_t nodeId); /*!< C string must be freed. */
	void manager_setNodeManufacturerName(manager_t m, uint32_t homeId, uint8_t nodeId, const char* manufacturerName);
	void manager_setNodeProductName(manager_t m, uint32_t homeId, uint8_t nodeId, const char* productName);
	void manager_setNodeName(manager_t m, uint32_t homeId, uint8_t nodeId, const char* nodeName);
	void manager_setNodeLocation(manager_t m, uint32_t homeId, uint8_t nodeId, const char* location);
	void manager_setNodeOn(manager_t m, uint32_t homeId, uint8_t nodeId);
	void manager_setNodeOff(manager_t m, uint32_t homeId, uint8_t nodeId);
	void manager_setNodeLevel(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t level);
	bool manager_isNodeInfoReceived(manager_t m, uint32_t homeId, uint8_t nodeId);
	bool manager_getNodeClassInformation(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t commandClassId, char **o_name, uint8_t *o_version);
	bool manager_isNodeAwake(manager_t m, uint32_t homeId, uint8_t nodeId);
	bool manager_isNodeFailed(manager_t m, uint32_t homeId, uint8_t nodeId);
	char* manager_getNodeQueryStage(manager_t m, uint32_t homeId, uint8_t nodeId); /*!< C string must be freed. */
	uint16_t manager_getNodeDeviceType(manager_t m, uint32_t homeId, uint8_t nodeId);
	char* manager_getNodeDeviceTypeString(manager_t m, uint32_t homeId, uint8_t nodeId); /*!< C string must be freed. */
	uint8_t manager_getNodeRole(manager_t m, uint32_t homeId, uint8_t nodeId);
	char* manager_getNodeRoleString(manager_t m, uint32_t homeId, uint8_t nodeId); /*!< C string must be freed. */
	uint8_t manager_getNodePlusType(manager_t m, uint32_t homeId, uint8_t nodeId);
	char* manager_getNodePlusTypeString(manager_t m, uint32_t homeId, uint8_t nodeId); /*!< C string must be freed. */

	//
	// Values.
	//

	char* manager_getValueLabel(manager_t m, valueid_t valueid);
	void manager_setValueLabel(manager_t m, valueid_t valueid, const char* value);
	char* manager_getValueUnits(manager_t m, valueid_t valueid);
	void manager_setValueUnits(manager_t m, valueid_t valueid, const char* value);
	char* manager_getValueHelp(manager_t m, valueid_t valueid);
	void manager_setValueHelp(manager_t m, valueid_t valueid, const char* value);
	int32_t manager_getValueMin(manager_t m, valueid_t valueid);
	int32_t manager_getValueMax(manager_t m, valueid_t valueid);
	bool manager_isValueReadOnly(manager_t m, valueid_t valueid);
	bool manager_isValueWriteOnly(manager_t m, valueid_t valueid);
	bool manager_isValueSet(manager_t m, valueid_t valueid);
	bool manager_isValuePolled(manager_t m, valueid_t valueid);
	bool manager_getValueAsBool(manager_t m, valueid_t valueid, bool *o_value);
	bool manager_getValueAsByte(manager_t m, valueid_t valueid, uint8_t *o_value);
	bool manager_getValueAsFloat(manager_t m, valueid_t valueid, float *o_value);
	bool manager_getValueAsInt(manager_t m, valueid_t valueid, int32_t *o_value);
	bool manager_getValueAsShort(manager_t m, valueid_t valueid, int16_t *o_value);
	bool manager_getValueAsString(manager_t m, valueid_t valueid, char **o_value);
	bool manager_getValueAsRaw(manager_t m, valueid_t valueid, bytes_t *o_value);
	bool manager_getValueListSelectionAsString(manager_t m, valueid_t valueid, char **o_value);
	bool manager_getValueListSelectionAsInt32(manager_t m, valueid_t valueid, int32_t *o_value);
	bool manager_getValueListItems(manager_t m, valueid_t valueid, zwlist_t **o_value);
	bool manager_getValueFloatPrecision(manager_t m, valueid_t valueid, uint8_t *o_value);
	bool manager_setValueBool(manager_t m, valueid_t valueid, bool value);
	bool manager_setValueUint8(manager_t m, valueid_t valueid, uint8_t value);
	bool manager_setValueFloat(manager_t m, valueid_t valueid, float value);
	bool manager_setValueInt32(manager_t m, valueid_t valueid, int32_t value);
	bool manager_setValueInt16(manager_t m, valueid_t valueid, int16_t value);
	bool manager_setValueBytes(manager_t m, valueid_t valueid, bytes_t *value);
	bool manager_setValueString(manager_t m, valueid_t valueid, const char* value);
	bool manager_setValueListSelection(manager_t m, valueid_t valueid, const char* selectedItem);
	bool manager_refreshValue(manager_t m, valueid_t valueid);
	void manager_setChangeVerified(manager_t m, valueid_t valueid, bool verify);
	bool manager_getChangeVerified(manager_t m, valueid_t valueid);
	bool manager_pressButton(manager_t m, valueid_t valueid);
	bool manager_releaseButton(manager_t m, valueid_t valueid);

	//
	// Climate control schedules.
	//

	uint8_t manager_getNumSwitchPoints(manager_t m, valueid_t valueid);
	bool manager_setSwitchPoint(manager_t m, valueid_t valueid, uint8_t hours, uint8_t minutes, int8_t setback);
	bool manager_removeSwitchPoint(manager_t m, valueid_t valueid, uint8_t hours, uint8_t minutes);
	void manager_clearSwitchPoints(manager_t m, valueid_t valueid);
	bool manager_getSwitchPoint(manager_t m, valueid_t valueid, uint8_t idx, uint8_t *o_hours, uint8_t *o_minutes, int8_t *o_setback);

	//
	// Switch all.
	//

	void manager_switchAllOn(manager_t m, uint32_t homeId);
	void manager_switchAllOff(manager_t m, uint32_t homeId);

	//
	// Configuration parameters.
	//

	bool manager_setConfigParam(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t param, int32_t value, uint8_t size);
	void manager_requestConfigParam(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t param);
	void manager_requestAllConfigParams(manager_t m, uint32_t homeId, uint8_t nodeId);

	//
	// Groups.
	//

	uint8_t manager_getNumGroups(manager_t m, uint32_t homeId, uint8_t nodeId);
	uint32_t manager_getAssociations(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t groupIdx, uint8_t **o_associations);
//TODO uint32_t manager_getAssociations(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t groupIdx, instanceassociation_t **o_associations);
	uint8_t manager_getMaxAssociations(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t groupIdx);
	char* manager_getGroupLabel(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t groupIdx);
	void manager_addAssociation(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t groupIdx, uint8_t targetNodeId, uint8_t instance);
	void manager_removeAssociation(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t groupIdx, uint8_t targetNodeId, uint8_t instance);

	//
	// Notifications.
	//

	extern void goNotificationCB(notification_t notification, void *userdata); /*!< Must be implemented in cgo. */
	bool manager_addWatcher(manager_t m, void *userdata);
	bool manager_removeWatcher(manager_t m, void *userdata);

	//
	// Controller commands.
	//

	void manager_resetController(manager_t m, uint32_t homeId);
	void manager_softReset(manager_t m, uint32_t homeId);
	bool manager_cancelControllerCommand(manager_t m, uint32_t homeId);

	//
	// Network commands.
	//

	void manager_testNetworkNode(manager_t m, uint32_t homeId, uint8_t nodeId, uint32_t count);
	void manager_testNetwork(manager_t m, uint32_t homeId, uint32_t count);
	void manager_healNetworkNode(manager_t m, uint32_t homeId, uint8_t nodeId, bool doRR);
	void manager_healNetwork(manager_t m, uint32_t homeId, bool doRR);
	bool manager_addNode(manager_t m, uint32_t homeId, bool doSecurity);
	bool manager_removeNode(manager_t m, uint32_t homeId);
	bool manager_removeFailedNode(manager_t m, uint32_t homeId, uint8_t nodeId);
	bool manager_hasNodeFailed(manager_t m, uint32_t homeId, uint8_t nodeId);
	bool manager_requestNodeNeighborUpdate(manager_t m, uint32_t homeId, uint8_t nodeId);
	bool manager_assignReturnRoute(manager_t m, uint32_t homeId, uint8_t nodeId);
	bool manager_deleteAllReturnRoutes(manager_t m, uint32_t homeId, uint8_t nodeId);
	bool manager_sendNodeInformation(manager_t m, uint32_t homeId, uint8_t nodeId);
	bool manager_createNewPrimary(manager_t m, uint32_t homeId);
	bool manager_receiveConfiguration(manager_t m, uint32_t homeId);
	bool manager_replaceFailedNode(manager_t m, uint32_t homeId, uint8_t nodeId);
	bool manager_transferPrimaryRole(manager_t m, uint32_t homeId);
	bool manager_requestNetworkUpdate(manager_t m, uint32_t homeId, uint8_t nodeId);
	bool manager_replicationSend(manager_t m, uint32_t homeId, uint8_t nodeId);
	bool manager_createButton(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t buttonid);
	bool manager_deleteButton(manager_t m, uint32_t homeId, uint8_t nodeId, uint8_t buttonid);

	//
	// Scene commands.
	//

	uint8_t manager_getNumScenes(manager_t m);
	uint8_t manager_getAllScenes(manager_t m, uint8_t **sceneIds);
	void manager_removeAllScenes(manager_t m, uint32_t homeId);
	uint8_t manager_createScene(manager_t m);
	bool manager_removeScene(manager_t m, uint8_t sceneId);
	bool manager_addSceneValueBool(manager_t m, uint8_t sceneId, valueid_t valueid, bool value);
	bool manager_addSceneValueUint8(manager_t m, uint8_t sceneId, valueid_t valueid, uint8_t value);
	bool manager_addSceneValueFloat(manager_t m, uint8_t sceneId, valueid_t valueid, float value);
	bool manager_addSceneValueInt32(manager_t m, uint8_t sceneId, valueid_t valueid, int32_t value);
	bool manager_addSceneValueInt16(manager_t m, uint8_t sceneId, valueid_t valueid, int16_t value);
	bool manager_addSceneValueString(manager_t m, uint8_t sceneId, valueid_t valueid, const char* value);
	bool manager_addSceneValueListSelectionString(manager_t m, uint8_t sceneId, valueid_t valueid, const char* value);
	bool manager_addSceneValueListSelectionInt32(manager_t m, uint8_t sceneId, valueid_t valueid, int32_t value);
	bool manager_removeSceneValue(manager_t m, uint8_t sceneId, valueid_t valueid);
//TODO int manager_sceneGetValues(manager_t m, uint8_t sceneId, valueidlist_t *o_value);
	bool manager_sceneGetValueAsBool(manager_t m, uint8_t sceneId, valueid_t valueid, bool *o_value);
	bool manager_sceneGetValueAsByte(manager_t m, uint8_t sceneId, valueid_t valueid, uint8_t *o_value);
	bool manager_sceneGetValueAsFloat(manager_t m, uint8_t sceneId, valueid_t valueid, float *o_value);
	bool manager_sceneGetValueAsInt(manager_t m, uint8_t sceneId, valueid_t valueid, int32_t *o_value);
	bool manager_sceneGetValueAsShort(manager_t m, uint8_t sceneId, valueid_t valueid, int16_t *o_value);
	bool manager_sceneGetValueAsString(manager_t m, uint8_t sceneId, valueid_t valueid, char **o_value);
	bool manager_sceneGetValueListSelectionString(manager_t m, uint8_t sceneId, valueid_t valueid, char **o_value);
	bool manager_sceneGetValueListSelectionInt32(manager_t m, uint8_t sceneId, valueid_t valueid, int32_t *o_value);
	bool manager_setSceneValueBool(manager_t m, uint8_t sceneId, valueid_t valueid, bool value);
	bool manager_setSceneValueUint8(manager_t m, uint8_t sceneId, valueid_t valueid, uint8_t value);
	bool manager_setSceneValueFloat(manager_t m, uint8_t sceneId, valueid_t valueid, float value);
	bool manager_setSceneValueInt32(manager_t m, uint8_t sceneId, valueid_t valueid, int32_t value);
	bool manager_setSceneValueInt16(manager_t m, uint8_t sceneId, valueid_t valueid, int16_t value);
	bool manager_setSceneValueString(manager_t m, uint8_t sceneId, valueid_t valueid, const char* value);
	bool manager_setSceneValueListSelectionString(manager_t m, uint8_t sceneId, valueid_t valueid, const char* value);
	bool manager_setSceneValueListSelectionInt32(manager_t m, uint8_t sceneId, valueid_t valueid, int32_t value);
	char* manager_getSceneLabel(manager_t m, uint8_t sceneId);
	void manager_setSceneLabel(manager_t m, uint8_t sceneId, const char* value);
	bool manager_sceneExists(manager_t m, uint8_t sceneId);
	bool manager_activateScene(manager_t m, uint8_t sceneId);

	//
	// Statistics retreival interface.
	//

//TODO void manager_getDriverStatistics(manager_t m, uint32_t homeId, driver_driverdata_t *data);
//TODO void manager_getNodeStatistics(manager_t m, uint32_t homeId, uint8_t nodeId, node_nodedata_t *data);

#ifdef __cplusplus
}
#endif

#endif // define GOOPENZWAVE_MANAGER
