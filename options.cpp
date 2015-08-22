#include "options.h"
#include <openzwave/Options.h>

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
    std::string _name(name);
    return opts->AddOptionBool(_name, value);
}

bool options_addOptionInt(options_t o, const char* name, int32_t value)
{
    OpenZWave::Options *opts = (OpenZWave::Options*)o;
    std::string _name(name);
    return opts->AddOptionInt(_name, value);
}

bool options_addOptionString(options_t o, const char* name, const char* value, bool append)
{
    OpenZWave::Options *opts = (OpenZWave::Options*)o;
    std::string _name(name);
    std::string _value(value);
    return opts->AddOptionString(_name, _value, append);
}

bool options_getOptionAsBool(options_t o, const char* name, bool* value_out)
{
    OpenZWave::Options *opts = (OpenZWave::Options*)o;
    std::string _name(name);
    return opts->GetOptionAsBool(_name, value_out);
}

bool options_getOptionAsInt(options_t o, const char* name, int32_t* value_out)
{
    OpenZWave::Options *opts = (OpenZWave::Options*)o;
    std::string _name(name);
    return opts->GetOptionAsInt(_name, value_out);
}

bool options_getOptionAsString(options_t o, const char* name, char* value_out, size_t* value_size)
{
    OpenZWave::Options *opts = (OpenZWave::Options*)o;
    std::string _name(name);
    std::string _value;
    bool result = opts->GetOptionAsString(_name, &_value);
    assert(_value.size() < *value_size);
    memcpy(value_out, _value.c_str(), _value.size());
    return result;
}

// OptionType options_getOptionType(string const &_name);

bool options_areLocked(options_t o)
{
    OpenZWave::Options *opts = (OpenZWave::Options*)o;
    return opts->AreLocked();
}
