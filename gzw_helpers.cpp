#include "gzw_helpers.h"
#include <stdlib.h>
#include <string.h>
#include <assert.h>

bytes_t* string_emptyBytes()
{
    bytes_t *bytes = (bytes_t*)malloc(sizeof(bytes_t));
    bytes->data = NULL;
    bytes->length = 0;
    return bytes;
}

void string_initBytes(bytes_t *bytes, size_t size)
{
    if (bytes->data != NULL) {
        free(bytes->data);
    }
    bytes->length = 0;

    if (size == 0) {
        bytes->data = NULL;
        bytes->length = 0;
    } else {
        bytes->data = (uint8_t*)malloc(sizeof(uint8_t) * size);
        bytes->length = size;
    }
}

void string_setByteAt(bytes_t *bytes, uint8_t value, size_t position)
{
    bytes->data[position] = value;
}

uint8_t string_byteAt(bytes_t *bytes, size_t position)
{
    return bytes->data[position];
}

void string_freeBytes(bytes_t *bytes)
{
    if (bytes->data != NULL) {
        free(bytes->data);
    }
    free(bytes);
}

char* zwhelper_makeCString(std::string &str)
{
    char *cstr = (char*)malloc(sizeof(char) * str.size()+1);
    memset(cstr, 0, str.size()+1);
    if (str.size() > 0) {
    	strncpy(cstr, str.c_str(), str.size());
    }
    return cstr;
}
