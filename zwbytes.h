#ifndef ZWBYTES_H
#define ZWBYTES_H

#include <stdint.h>

#ifdef __cplusplus
#include <string>
#include <vector>
extern "C" {
#endif

/**
 * zwbytes_t is a container for byte arrays.
 */
typedef struct {
    uint8_t* data;
    size_t size;
} zwbytes_t;

/**
 * zwbytes_new creates and returns a new, empty zwbytes_t object. Make sure to
 * destroy it when finished with zwbytes_free.
 * @return An empty zwbytes_t object.
 */
zwbytes_t* zwbytes_new();

/**
 * zwbytes_reserve resizes the zwbytes_t, if necessary, to the new size.
 * @param bytes The zwbytes_t object to operate on.
 * @param size  The new size.
 */
void zwbytes_reserve(zwbytes_t *bytes, size_t size);

/**
 * zwbytes_set sets the value of the byte at the given position.
 * @param bytes The zwbytes_t object to operate on.
 * @param pos   The position of the byte to set.
 * @param val   The byte value to set.
 */
void zwbytes_set(zwbytes_t *bytes, size_t pos, uint8_t val);

/**
 * zwbytes_at returns the value of the byte at the given position.
 * @param  bytes The zwbytes_t object to operate on.
 * @param  pos   The position of the byte to get.
 * @return       The value of the byte.
 */
uint8_t zwbytes_at(zwbytes_t *bytes, size_t pos);

/**
 * zwbytes_free frees the zwbytes_t object.
 * @param list The zwbytes_t object to free.
 */
void zwbytes_free(zwbytes_t *bytes);

#ifdef __cplusplus
} // end extern "C"
#endif

#endif // define ZWBYTES_H
