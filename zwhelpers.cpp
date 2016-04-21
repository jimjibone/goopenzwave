#include "zwhelpers.h"
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

/*
 *
 * C types and functions.
 *
 */

zwlist_t *zwlist_new()
{
    std::vector<std::string> *vec = new std::vector<std::string>();
    assert(vec);
    return (zwlist_t*)vec;
}

int zwlist_size(zwlist_t *list)
{
    assert(list);
    std::vector<std::string> *vec = (std::vector<std::string>*)list;
    return vec->size();
}

char* zwlist_at(zwlist_t *list, int pos)
{
    assert(list);
    std::vector<std::string> *vec = (std::vector<std::string>*)list;
    return zwhelper_makeCString(vec->at(pos));
}

void zwlist_free(zwlist_t *list)
{
    assert(list);
    std::vector<std::string> *vec = (std::vector<std::string>*)list;
    delete vec;
}

/*
 *
 * C++ only functions.
 *
 */

char* zwhelper_makeCString(std::string &str)
{
    char *cstr = (char*)malloc(sizeof(char) * str.size()+1);
    memset(cstr, 0, str.size()+1);
    if (str.size() > 0) {
    	strncpy(cstr, str.c_str(), str.size());
    }
    return cstr;
}

zwlist_t* zwlist_copy(std::vector<std::string> &vec)
{
    std::vector<std::string> *out = new std::vector<std::string>(vec);
    assert(out);
    return (zwlist_t*)out;
}
