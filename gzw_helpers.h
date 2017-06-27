#ifndef GOOPENZWAVE_STRING
#define GOOPENZWAVE_STRING

#include <stdint.h>

#ifdef __cplusplus
#include <string>
#include <vector>
extern "C" {
#endif

/*
 *
 * C types and functions.
 *
 */

/**
 * Container for byte arrays.
 */
typedef struct {
    uint8_t *data;
    size_t length;
} bytes_t;

bytes_t* string_emptyBytes();
void string_initBytes(bytes_t *bytes, size_t size);
void string_setByteAt(bytes_t *bytes, uint8_t value, size_t position);
uint8_t string_byteAt(bytes_t *bytes, size_t position);
void string_freeBytes(bytes_t *bytes);

/**
 * C friendly representation of std::vector<std::string>*.
 */
typedef void* zwlist_t;

zwlist_t* zwlist_new();
int zwlist_size(zwlist_t *list);
char* zwlist_at(zwlist_t *list, int pos);
void zwlist_free(zwlist_t *list);

#ifdef __cplusplus
} // end extern "C"

/*
 *
 * C++ only functions.
 *
 */

/**
 * Create and return a copy of the std::string as a char*. The returned pointer
 * must be freed!
 * @param  str The string to copy.
 * @return     The new C string.
 */
char* zwhelper_makeCString(std::string &str);

zwlist_t* zwlist_copy(std::vector<std::string> &vec);

#endif

#endif // define GOOPENZWAVE_STRING
