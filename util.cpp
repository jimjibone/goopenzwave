#include "util.h"

void* ptr_at(void **ptr, uint32_t idx)
{
    return ptr[idx];
}

char* str_at(char** ptr, uint32_t idx)
{
    return ptr[idx];
}

int32_t int32_at(int32_t* ptr, uint32_t idx)
{
    return ptr[idx];
}
