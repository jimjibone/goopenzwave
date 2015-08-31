#include "string_helpers.h"
#include <stdlib.h>

string_t* string_emptyString()
{
    string_t *string = (string_t*)malloc(sizeof(string_t));
    string->data = NULL;
    string->length = 0;
    return string;
}

void string_initString(string_t *string, size_t size)
{
    if (string->data != NULL) {
        delete string->data;
    }
    string->length = 0;

    if (size == 0) {
        string->data = NULL;
        string->length = 0;
    } else {
        string->data = (char*)malloc(sizeof(char) * size);
        string->length = size;
    }
}

void string_freeString(string_t *string)
{
    if (string->data != NULL) {
        delete string->data;
    }
    delete string;
}

bytes_t* string_emptyBytes()
{
    bytes_t *bytes = (bytes_t*)malloc(sizeof(bytes_t));
    bytes->data = NULL;
    bytes->length = 0;
    return bytes;
}

void string_initBytes(bytes_t *bytes, size_t size)
{
    if (bytes->data != NULL) {
        delete bytes->data;
    }
    bytes->length = 0;

    if (size == 0) {
        bytes->data = NULL;
        bytes->length = 0;
    } else {
        bytes->data = (uint8_t*)malloc(sizeof(uint8_t) * size);
        bytes->length = size;
    }
}

void string_setByteAt(bytes_t *bytes, uint8_t value, size_t position)
{
    bytes->data[position] = value;
}

uint8_t string_byteAt(bytes_t *bytes, size_t position)
{
    return bytes->data[position];
}

void string_freeBytes(bytes_t *bytes)
{
    if (bytes->data != NULL) {
        delete bytes->data;
    }
    delete bytes;
}

stringlist_t* string_emptyStringList()
{
    stringlist_t *list = (stringlist_t*)malloc(sizeof(stringlist_t));
    list->length = 0;
    return list;
}

string_t* string_stringAt(stringlist_t *list, size_t position)
{
    return list->list[position];
}

void string_freeStringList(stringlist_t *list)
{
    for (int i = 0; i < list->length; i++) {
        string_freeString(list->list[i]);
    }
    if (list->list != NULL) {
        delete list->list;
    }
    delete list;
}

void string_copyStdString(string_t *cstr, std::string &string)
{
    // Free existing string from cstr.
    if (cstr->data != NULL) {
        delete cstr->data;
    }

    // Copy string into a new cstr->data.
    cstr->length = string.size() + 1;
    cstr->data = (char*)malloc(sizeof(char) * cstr->length);
    string.copy(cstr->data, string.size(), 0);
    cstr->data[cstr->length] = '\0';
}

std::string string_toStdString(string_t *string)
{
    std::string sstr(string->data);
    return sstr;
}

string_t* string_fromStdString(std::string &string)
{
    string_t *cstr = (string_t*)malloc(sizeof(string_t));
    cstr->length = string.size() + 1;
    cstr->data = (char*)malloc(sizeof(char) * cstr->length);
    string.copy(cstr->data, string.size(), 0);
    cstr->data[cstr->length] = '\0';
    return cstr;
}

void string_copyStdStringList(stringlist_t *clist, std::vector<std::string> &list)
{
    // Free all existing strings in the clist.
    for (int i = 0; i < clist->length; i++) {
        string_freeString(clist->list[i]);
    }
    if (clist->list != NULL) {
        delete clist->list;
    }

    // Prepare the list.
    clist->length = list.size();
    clist->list = (string_t**)malloc(sizeof(string_t*) * clist->length);

    // Copy each string from the list into the clist.
    for(std::vector<std::string>::size_type i = 0; i != list.size(); i++) {
        // Create a new string_t from the std::string and add it to the list.
        clist->list[i] = string_fromStdString(list[i]);
    }
}
