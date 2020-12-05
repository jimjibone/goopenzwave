#ifndef GOOPENZWAVE_NODE_WRAP_H
#define GOOPENZWAVE_NODE_WRAP_H

#include <stdint.h>
#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    bool ok;
    char* className;
    uint8_t classVersion;
} manager_node_class_information;

bool manager_refresh_node_info(uint32_t const homeId, uint8_t const nodeId);
bool manager_request_node_state(uint32_t const homeId, uint8_t const nodeId);
bool manager_request_node_dynamic(uint32_t const homeId, uint8_t const nodeId);
void ozw_RequestAllConfigParams(uint32_t const homeId, uint8_t const nodeId);
bool manager_is_node_listening_device(uint32_t const homeId, uint8_t const nodeId);
bool manager_is_node_frequent_listening_device(uint32_t const homeId, uint8_t const nodeId);
bool manager_is_node_beaming_device(uint32_t const homeId, uint8_t const nodeId);
bool manager_is_node_routing_device(uint32_t const homeId, uint8_t const nodeId);
bool manager_is_node_security_device(uint32_t const homeId, uint8_t const nodeId);
uint32_t manager_get_node_max_baud_rate(uint32_t const homeId, uint8_t const nodeId);
uint8_t manager_get_node_version(uint32_t const homeId, uint8_t const nodeId);
uint8_t manager_get_node_security(uint32_t const homeId, uint8_t const nodeId);
bool manager_is_node_zwave_plus(uint32_t const homeId, uint8_t const nodeId);
uint8_t manager_get_node_basic(uint32_t const homeId, uint8_t const nodeId);
uint8_t manager_get_node_generic(uint32_t const homeId, uint8_t const nodeId);
uint8_t manager_get_node_specific(uint32_t const homeId, uint8_t const nodeId);
char* manager_get_node_type(uint32_t const homeId, uint8_t const nodeId);
uint32_t* manager_get_node_neighbors(uint32_t const homeId, uint8_t const nodeId);
void manager_syncronize_node_neighbors(uint32_t const homeId, uint8_t const nodeId);
char* manager_get_node_manufacturer_name(uint32_t const homeId, uint8_t const nodeId);
char* manager_get_node_product_name(uint32_t const homeId, uint8_t const nodeId);
char* manager_get_node_name(uint32_t const homeId, uint8_t const nodeId);
char* manager_get_node_location(uint32_t const homeId, uint8_t const nodeId);
char* manager_get_node_manufacturer_id(uint32_t const homeId, uint8_t const nodeId);
char* manager_get_node_product_type(uint32_t const homeId, uint8_t const nodeId);
char* manager_get_node_product_id(uint32_t const homeId, uint8_t const nodeId);
void manager_set_node_manufacturer_name(uint32_t const homeId, uint8_t const nodeId, const char* manufacturerName);
void manager_set_node_product_name(uint32_t const homeId, uint8_t const nodeId, const char* productName);
void manager_set_node_name(uint32_t const homeId, uint8_t const nodeId, const char* nodeName);
void manager_set_node_location(uint32_t const homeId, uint8_t const nodeId, const char* location);
bool manager_is_node_info_received(uint32_t const homeId, uint8_t const nodeId);
manager_node_class_information* manager_get_node_class_information(uint32_t const homeId, uint8_t const nodeId, uint8_t const commandClassId);
void manager_free_node_class_information(manager_node_class_information* ptr);
bool manager_is_node_awake(uint32_t const homeId, uint8_t const nodeId);
bool manager_is_node_failed(uint32_t const homeId, uint8_t const nodeId);
char* manager_get_node_query_stage(uint32_t const homeId, uint8_t const nodeId);
uint16_t manager_get_node_device_type(uint32_t const homeId, uint8_t const nodeId);
uint8_t manager_get_node_role(uint32_t const homeId, uint8_t const nodeId);
uint8_t manager_get_node_plus_type(uint32_t const homeId, uint8_t const nodeId);

#ifdef __cplusplus
}
#endif

#endif // GOOPENZWAVE_NODE_WRAP_H
