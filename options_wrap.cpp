#include "options_wrap.h"
#include "openzwave/Options.h"

void options_create(const char* configPath, const char* userPath, const char* commandLine)
{
    OpenZWave::Options::Create(configPath, userPath, commandLine);
}

bool options_add_bool(const char* name, bool defaultval)
{
    return OpenZWave::Options::Get()->AddOptionBool(name, defaultval);
}

bool options_add_int(const char* name, int32_t defaultval)
{
    return OpenZWave::Options::Get()->AddOptionInt(name, defaultval);
}

bool options_add_string(const char* name, const char* defaultval, bool append)
{
    return OpenZWave::Options::Get()->AddOptionString(name, defaultval, append);
}

bool options_lock()
{
    return OpenZWave::Options::Get()->Lock();
}

bool options_destroy()
{
    return OpenZWave::Options::Destroy();
}
