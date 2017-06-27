#include "zwbytes.h"
#include <stdlib.h>
#include <string.h>
#include <assert.h>

zwbytes_t* zwbytes_new()
{
    zwbytes_t *bytes = (zwbytes_t*)malloc(sizeof(zwbytes_t));
    bytes->data = NULL;
    bytes->size = 0;
    return bytes;
}

void zwbytes_reserve(zwbytes_t *bytes, size_t size)
{
    bytes->data = (uint8_t*)realloc(bytes->data, size * sizeof(uint8_t));
    bytes->size = size;
}

void zwbytes_set(zwbytes_t *bytes, size_t pos, uint8_t val)
{
    if (pos >= bytes->size) return;
    bytes->data[pos] = val;
}

uint8_t zwbytes_at(zwbytes_t *bytes, size_t pos)
{
    if (pos >= bytes->size) return 0;
    return bytes->data[pos];
}

void zwbytes_free(zwbytes_t *bytes)
{
    if (bytes->data != NULL) {
        free(bytes->data);
    }
    free(bytes);
}
