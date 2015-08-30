#ifndef GOZWAVE_OPTIONS
#define GOZWAVE_OPTIONS

#include <stdint.h>
#include <stdbool.h>
#include <stddef.h>
#include "string_helpers.h"
#include "loglevel.h"

#ifdef __cplusplus
extern "C" {
#endif

    // Types.
    typedef void* options_t;

    // Static public member functions.
    options_t options_create(const char* configPath, const char* userPath, const char* commandLine);
    bool options_destroy();
	options_t options_get();

    // Public member functions.
    bool options_lock(options_t o);
    bool options_addOptionBool(options_t o, const char* name, bool value);
    bool options_addOptionInt(options_t o, const char* name, int32_t value);
    bool options_addOptionLogLevel(options_t o, const char* name, loglevel_t level);
    bool options_addOptionString(options_t o, const char* name, const char* value, bool append);
    bool options_getOptionAsBool(options_t o, const char* name, bool* o_value);
    bool options_getOptionAsInt(options_t o, const char* name, int32_t* o_value);
    bool options_getOptionAsString(options_t o, const char* name, string_t* o_value);
//TODO  OptionType options_getOptionType(string const &_name);
    bool options_areLocked(options_t o);

#ifdef __cplusplus
}
#endif

#endif // define GOZWAVE_OPTIONS
