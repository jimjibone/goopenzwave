#ifndef GOOPENZWAVE_OPTIONS_WRAP_H
#define GOOPENZWAVE_OPTIONS_WRAP_H

#include <stdint.h>
#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

void options_create(const char* configPath, const char* userPath, const char* commandLine);
bool options_add_bool(const char* name, bool defaultval);
bool options_add_int(const char* name, int32_t defaultval);
bool options_add_string(const char* name, const char* defaultval, bool append);
bool options_lock();
bool options_destroy();

#ifdef __cplusplus
}
#endif

#endif // GOOPENZWAVE_OPTIONS_WRAP_H
