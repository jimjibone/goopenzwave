#include "gzw_helpers.h"
#include <stdlib.h>
#include <string.h>
#include <assert.h>

char* zwhelper_makeCString(std::string &str)
{
    char *cstr = (char*)malloc(sizeof(char) * str.size()+1);
    memset(cstr, 0, str.size()+1);
    if (str.size() > 0) {
    	strncpy(cstr, str.c_str(), str.size());
    }
    return cstr;
}
