#include "node_wrap.h"
#include "openzwave/Manager.h"

bool manager_refresh_node_info(uint32_t const homeId, uint8_t const nodeId)
{
    return OpenZWave::Manager::Get()->RefreshNodeInfo(homeId, nodeId);
}

bool manager_request_node_state(uint32_t const homeId, uint8_t const nodeId)
{
    return OpenZWave::Manager::Get()->RequestNodeState(homeId, nodeId);
}

bool manager_request_node_dynamic(uint32_t const homeId, uint8_t const nodeId)
{
    return OpenZWave::Manager::Get()->RequestNodeDynamic(homeId, nodeId);
}

void ozw_RequestAllConfigParams(uint32_t const homeId, uint8_t const nodeId)
{
    OpenZWave::Manager::Get()->RequestAllConfigParams(homeId, nodeId);
}

bool manager_is_node_listening_device(uint32_t const homeId, uint8_t const nodeId)
{
    return OpenZWave::Manager::Get()->IsNodeListeningDevice(homeId, nodeId);
}

bool manager_is_node_frequent_listening_device(uint32_t const homeId, uint8_t const nodeId)
{
    return OpenZWave::Manager::Get()->IsNodeFrequentListeningDevice(homeId, nodeId);
}

bool manager_is_node_beaming_device(uint32_t const homeId, uint8_t const nodeId)
{
    return OpenZWave::Manager::Get()->IsNodeBeamingDevice(homeId, nodeId);
}

bool manager_is_node_routing_device(uint32_t const homeId, uint8_t const nodeId)
{
    return OpenZWave::Manager::Get()->IsNodeRoutingDevice(homeId, nodeId);
}

bool manager_is_node_security_device(uint32_t const homeId, uint8_t const nodeId)
{
    return OpenZWave::Manager::Get()->IsNodeSecurityDevice(homeId, nodeId);
}

uint32_t manager_get_node_max_baud_rate(uint32_t const homeId, uint8_t const nodeId)
{
    return OpenZWave::Manager::Get()->GetNodeMaxBaudRate(homeId, nodeId);
}

uint8_t manager_get_node_version(uint32_t const homeId, uint8_t const nodeId)
{
    return OpenZWave::Manager::Get()->GetNodeVersion(homeId, nodeId);
}

uint8_t manager_get_node_security(uint32_t const homeId, uint8_t const nodeId)
{
    return OpenZWave::Manager::Get()->GetNodeSecurity(homeId, nodeId);
}

bool manager_is_node_zwave_plus(uint32_t const homeId, uint8_t const nodeId)
{
    return OpenZWave::Manager::Get()->IsNodeZWavePlus(homeId, nodeId);
}

uint8_t manager_get_node_basic(uint32_t const homeId, uint8_t const nodeId)
{
    return OpenZWave::Manager::Get()->GetNodeBasic(homeId, nodeId);
}

uint8_t manager_get_node_generic(uint32_t const homeId, uint8_t const nodeId)
{
    return OpenZWave::Manager::Get()->GetNodeGeneric(homeId, nodeId);
}

uint8_t manager_get_node_specific(uint32_t const homeId, uint8_t const nodeId)
{
    return OpenZWave::Manager::Get()->GetNodeSpecific(homeId, nodeId);
}

char* manager_get_node_type(uint32_t const homeId, uint8_t const nodeId)
{
    std::string str = OpenZWave::Manager::Get()->GetNodeType(homeId, nodeId);
    return strdup(str.c_str());
}

uint32_t* manager_get_node_neighbors(uint32_t const homeId, uint8_t const nodeId)
{
    // returns number of neighbours.
    // array is always of size 29.
    // we will return a new array of uint32's, with the first being the size, followed by the result.
    uint8_t* neighbors = nullptr;
    uint32_t size = OpenZWave::Manager::Get()->GetNodeNeighbors(homeId, nodeId, &neighbors);
    uint32_t* res = static_cast<uint32_t*>(malloc((size+1) * sizeof(uint32_t)));
    res[0] = size;
    for (uint32_t i = 0; i < size; ++i) {
        res[i+1] = static_cast<uint32_t>(neighbors[i]);
    }
    return res;
}

void manager_syncronize_node_neighbors(uint32_t const homeId, uint8_t const nodeId)
{
    OpenZWave::Manager::Get()->SyncronizeNodeNeighbors(homeId, nodeId);
}

char* manager_get_node_manufacturer_name(uint32_t const homeId, uint8_t const nodeId)
{
    std::string str = OpenZWave::Manager::Get()->GetNodeManufacturerName(homeId, nodeId);
    return strdup(str.c_str());
}

char* manager_get_node_product_name(uint32_t const homeId, uint8_t const nodeId)
{
    std::string str = OpenZWave::Manager::Get()->GetNodeProductName(homeId, nodeId);
    return strdup(str.c_str());
}

char* manager_get_node_name(uint32_t const homeId, uint8_t const nodeId)
{
    std::string str = OpenZWave::Manager::Get()->GetNodeName(homeId, nodeId);
    return strdup(str.c_str());
}

char* manager_get_node_location(uint32_t const homeId, uint8_t const nodeId)
{
    std::string str = OpenZWave::Manager::Get()->GetNodeLocation(homeId, nodeId);
    return strdup(str.c_str());
}

char* manager_get_node_manufacturer_id(uint32_t const homeId, uint8_t const nodeId)
{
    std::string str = OpenZWave::Manager::Get()->GetNodeManufacturerId(homeId, nodeId);
    return strdup(str.c_str());
}

char* manager_get_node_product_type(uint32_t const homeId, uint8_t const nodeId)
{
    std::string str = OpenZWave::Manager::Get()->GetNodeProductType(homeId, nodeId);
    return strdup(str.c_str());
}

char* manager_get_node_product_id(uint32_t const homeId, uint8_t const nodeId)
{
    std::string str = OpenZWave::Manager::Get()->GetNodeProductId(homeId, nodeId);
    return strdup(str.c_str());
}

void manager_set_node_manufacturer_name(uint32_t const homeId, uint8_t const nodeId, const char* manufacturerName)
{
    OpenZWave::Manager::Get()->SetNodeManufacturerName(homeId, nodeId, manufacturerName);
}

void manager_set_node_product_name(uint32_t const homeId, uint8_t const nodeId, const char* productName)
{
    OpenZWave::Manager::Get()->SetNodeProductName(homeId, nodeId, productName);
}

void manager_set_node_name(uint32_t const homeId, uint8_t const nodeId, const char* nodeName)
{
    OpenZWave::Manager::Get()->SetNodeName(homeId, nodeId, nodeName);
}

void manager_set_node_location(uint32_t const homeId, uint8_t const nodeId, const char* location)
{
    OpenZWave::Manager::Get()->SetNodeLocation(homeId, nodeId, location);
}

bool manager_is_node_info_received(uint32_t const homeId, uint8_t const nodeId)
{
    return OpenZWave::Manager::Get()->IsNodeInfoReceived(homeId, nodeId);
}

manager_node_class_information* manager_get_node_class_information(uint32_t const homeId, uint8_t const nodeId, uint8_t const commandClassId)
{
    std::string classNameTmp;
    uint8_t classVersionTmp;
    bool ok = OpenZWave::Manager::Get()->GetNodeClassInformation(homeId, nodeId, commandClassId, &classNameTmp, &classVersionTmp);
    manager_node_class_information* res = static_cast<manager_node_class_information*>(malloc(sizeof(manager_node_class_information)));
    res->ok = ok;
    res->className = strdup(classNameTmp.c_str());
    res->classVersion = classVersionTmp;
    return res;
}

void manager_free_node_class_information(manager_node_class_information* ptr)
{
    if (ptr)
    {
        free(ptr->className);
        free(ptr);
    }
}

bool manager_is_node_awake(uint32_t const homeId, uint8_t const nodeId)
{
    return OpenZWave::Manager::Get()->IsNodeAwake(homeId, nodeId);
}

bool manager_is_node_failed(uint32_t const homeId, uint8_t const nodeId)
{
    return OpenZWave::Manager::Get()->IsNodeFailed(homeId, nodeId);
}

char* manager_get_node_query_stage(uint32_t const homeId, uint8_t const nodeId)
{
    std::string str = OpenZWave::Manager::Get()->GetNodeQueryStage(homeId, nodeId);
    return strdup(str.c_str());
}

uint16_t manager_get_node_device_type(uint32_t const homeId, uint8_t const nodeId)
{
    return OpenZWave::Manager::Get()->GetNodeDeviceType(homeId, nodeId);
}

uint8_t manager_get_node_role(uint32_t const homeId, uint8_t const nodeId)
{
    return OpenZWave::Manager::Get()->GetNodeRole(homeId, nodeId);
}

uint8_t manager_get_node_plus_type(uint32_t const homeId, uint8_t const nodeId)
{
    return OpenZWave::Manager::Get()->GetNodePlusType(homeId, nodeId);
}
