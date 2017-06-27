#ifndef ZWLIST_H
#define ZWLIST_H

#include <stdint.h>

#ifdef __cplusplus
#include <string>
#include <vector>
extern "C" {
#endif

/**
 * zwlist_t is a C-friendly wrapper around a std::vector<std::string>*.
 */
typedef void* zwlist_t;

/**
 * zwlist_new creates and returns a new, empty zwlist_t object. Make sure to
 * destroy it when finished with zwlist_free.
 * @return An empty zwlist_t object.
 */
zwlist_t* zwlist_new();

/**
 * zwlist_size returns the size of the zwlist_t object.
 * @param  list The zwlist_t object to operate on.
 * @return      The number of items in the zwlist_t.
 */
size_t zwlist_size(zwlist_t *list);

/**
 * zwlist_at returns a copy of the item from the given position in the list. The
 * caller must free the returned string when finished.
 * @param  list The zwlist_t object to operate on.
 * @param  pos  The position of the item to return within the list.
 * @return      A copy of the item string. Free when finished.
 */
char* zwlist_at(zwlist_t *list, size_t pos);

/**
 * zwlist_free frees the zwlist_t object.
 * @param list The zwlist_t object to free.
 */
void zwlist_free(zwlist_t *list);

#ifdef __cplusplus
} // end extern "C"

/**
 * zwlist_copy copies the vector and returns it. Make sure to destroy it when
 * finished with zwlist_free.
 * @param  vec The vector to copy.
 * @return     A new zwlist_t object containing a copy of vec.
 */
zwlist_t* zwlist_copy(std::vector<std::string> &vec);

#endif

#endif // define ZWLIST_H
