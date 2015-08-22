#include "manager.h"
#include <openzwave/Manager.h>

// Static public member functions.
manager_t manager_create()
{
	OpenZWave::Manager* man = OpenZWave::Manager::Create();
	return (manager_t)man;
}

manager_t manager_get()
{
	OpenZWave::Manager* man = OpenZWave::Manager::Get();
	return (manager_t)man;
}

void manager_destroy()
{
	OpenZWave::Manager::Destroy();
}

const char* manager_getVersionAsString()
{
	return OpenZWave::Manager::getVersionAsString().c_str();
}

const char* manager_getVersionLongAsString()
{
	return OpenZWave::Manager::getVersionLongAsString().c_str();
}
