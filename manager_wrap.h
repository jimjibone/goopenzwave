#ifndef GOOPENZWAVE_MANAGER_WRAP_H
#define GOOPENZWAVE_MANAGER_WRAP_H

#include <stdint.h>
#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

void manager_create();
bool manager_add_driver(const char* port);
bool manager_remove_driver(const char* port);
void manager_destroy();
char* manager_get_version_as_string();
char* manager_get_version_long_as_string();
uint32_t manager_get_version(); /* major << 16 | minor */

#ifdef __cplusplus
}
#endif

#endif // GOOPENZWAVE_MANAGER_WRAP_H
