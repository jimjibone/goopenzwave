#ifndef GOZWAVE_OPTIONS
#define GOZWAVE_OPTIONS

#include <stdint.h>
#include <stdbool.h>
#include <stddef.h>

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
    bool options_addOptionString(options_t o, const char* name, const char* value, bool append);
    bool options_getOptionAsBool(options_t o, const char* name, bool* value_out);
    bool options_getOptionAsInt(options_t o, const char* name, int32_t* value_out);
    bool options_getOptionAsString(options_t o, const char* name, char* value_out, size_t* value_size);
    // OptionType options_getOptionType(string const &_name);
    bool options_areLocked(options_t o);

#ifdef __cplusplus
}
#endif

#endif // define GOZWAVE_OPTIONS
