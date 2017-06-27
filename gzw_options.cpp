#include "gzw_options.h"
#include <openzwave/Options.h>
#include <string.h>
#include <stdlib.h>

//
// Static public member functions.
//

options_t options_create(const char* configPath, const char* userPath, const char* commandLine)
{
    OpenZWave::Options *opts = OpenZWave::Options::Create(configPath, userPath, commandLine);
    return (options_t)opts;
}

bool options_destroy()
{
    return OpenZWave::Options::Destroy();
}

options_t options_get()
{
    OpenZWave::Options *opts = OpenZWave::Options::Get();
    return (options_t)opts;
}

//
// Public member functions.
//

bool options_lock(options_t o)
{
    OpenZWave::Options *opts = (OpenZWave::Options*)o;
    return opts->Lock();
}

bool options_addOptionBool(options_t o, const char* name, bool value)
{
    OpenZWave::Options *opts = (OpenZWave::Options*)o;
    std::string inStr(name);
    return opts->AddOptionBool(inStr, value);
}

bool options_addOptionInt(options_t o, const char* name, int32_t value)
{
    OpenZWave::Options *opts = (OpenZWave::Options*)o;
    std::string inStr(name);
    return opts->AddOptionInt(inStr, value);
}

bool options_addOptionLogLevel(options_t o, const char* name, loglevel_t level)
{
    OpenZWave::Options *opts = (OpenZWave::Options*)o;
    std::string inStr(name);
    return opts->AddOptionInt(inStr, loglevel_toLogLevel(level));
}

bool options_addOptionString(options_t o, const char* name, const char* value, bool append)
{
    OpenZWave::Options *opts = (OpenZWave::Options*)o;
    std::string inStr(name);
    std::string _value(value);
    return opts->AddOptionString(inStr, _value, append);
}

bool options_getOptionAsBool(options_t o, const char* name, bool* o_value)
{
    OpenZWave::Options *opts = (OpenZWave::Options*)o;
    std::string inStr(name);
    return opts->GetOptionAsBool(inStr, o_value);
}

bool options_getOptionAsInt(options_t o, const char* name, int32_t* o_value)
{
    OpenZWave::Options *opts = (OpenZWave::Options*)o;
    std::string inStr(name);
    return opts->GetOptionAsInt(inStr, o_value);
}

bool options_getOptionAsString(options_t o, const char* name, char **o_value)
{
    OpenZWave::Options *opts = (OpenZWave::Options*)o;
    std::string inStr(name);
    std::string str;
    bool result = opts->GetOptionAsString(inStr, &str);
    if (*o_value) {
        free(*o_value);
    }
    *o_value = zwhelper_makeCString(str);
    return result;
}

//TODO  OptionType options_getOptionType(string const &inStr);

bool options_areLocked(options_t o)
{
    OpenZWave::Options *opts = (OpenZWave::Options*)o;
    return opts->AreLocked();
}
