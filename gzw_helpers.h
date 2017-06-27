#ifndef GOOPENZWAVE_STRING
#define GOOPENZWAVE_STRING

#include <stdint.h>

#ifdef __cplusplus
#include <string>
#include <vector>
extern "C" {
#endif

#ifdef __cplusplus
} // end extern "C"

/*
 *
 * C++ only functions.
 *
 */

char* zwhelper_makeCString(std::string &str);

#endif

#endif // define GOOPENZWAVE_STRING
