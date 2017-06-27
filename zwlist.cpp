#include "zwlist.h"
#include <stdlib.h>
#include <string.h>
#include <assert.h>

zwlist_t *zwlist_new()
{
    std::vector<std::string> *vec = new std::vector<std::string>();
    assert(vec);
    return (zwlist_t*)vec;
}

int zwlist_size(zwlist_t *list)
{
    assert(list);
    std::vector<std::string> *vec = (std::vector<std::string>*)list;
    return vec->size();
}

char* zwlist_at(zwlist_t *list, int pos)
{
    assert(list);
    std::vector<std::string> *vec = (std::vector<std::string>*)list;
    char* cstr = strdup(vec->at(pos).c_str());
    return cstr;
}

void zwlist_free(zwlist_t *list)
{
    assert(list);
    std::vector<std::string> *vec = (std::vector<std::string>*)list;
    delete vec;
}

zwlist_t* zwlist_copy(std::vector<std::string> &vec)
{
    std::vector<std::string> *out = new std::vector<std::string>(vec);
    assert(out);
    return (zwlist_t*)out;
}
