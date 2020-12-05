#ifndef GOOPENZWAVE_UTIL_H
#define GOOPENZWAVE_UTIL_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

void* ptr_at(void **ptr, uint32_t idx);
char* str_at(char** ptr, uint32_t idx);
int32_t int32_at(int32_t* ptr, uint32_t idx);

#ifdef __cplusplus
}
#endif

#endif // GOOPENZWAVE_UTIL_H
