#ifndef GOZWAVE_VALUEID
#define GOZWAVE_VALUEID

#include <stdint.h>
#include <stdbool.h>
#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

    // Types.
    typedef void* valueid_t;

    // enum valueid_genre
    typedef enum {
        valueid_genre_basic = 0,
        valueid_genre_user,
        valueid_genre_config,
        valueid_genre_system,
        valueid_genre_count
    } valueid_genre;

    // enum valueid_type
    typedef enum {
        valueid_type_bool = 0,
        valueid_type_byte,
        valueid_type_decimal,
        valueid_type_int,
        valueid_type_list,
        valueid_type_schedule,
        valueid_type_short,
        valueid_type_string,
        valueid_type_button,
        valueid_type_raw,
        valueid_type_max = valueid_type_raw
    } valueid_type;

    // Public member functions.
    uint32_t valueid_getHomeId(valueid_t n);
    uint8_t valueid_getNodeId(valueid_t n);
    valueid_genre valueid_getGenre(valueid_t n);
    uint8_t valueid_getCommandClassId(valueid_t n);
    uint8_t valueid_getInstance(valueid_t n);
    uint8_t valueid_getIndex(valueid_t n);
    valueid_type valueid_getType(valueid_t n);
    uint64_t valueid_getId(valueid_t n);

    // Go helper functions.
    valueid_t valueid_create(uint32_t homeId, uint64_t id);
    void valueid_free(valueid_t valueid);

#ifdef __cplusplus
}
#endif

#endif // define GOZWAVE_VALUEID
