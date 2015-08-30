#ifndef GOZWAVE_STRING
#define GOZWAVE_STRING

#ifdef __cplusplus
#include <string>
#include <vector>
extern "C" {
#endif

/**
 * Container for C strings.
 */
typedef struct {
    char *data;
    unsigned int length;
} string_t;

/**
 * Container for byte arrays.
 */
typedef struct {
    uint8_t *data;
    unsigned int length;
} bytes_t;

/**
 * Container for a list of C strings.
 */
typedef struct {
    string_t **list;
    unsigned int length;
} stringlist_t;

string_t* string_emptyString();
void string_initString(string_t *string, unsigned int size);
void string_freeString(string_t *string);

bytes_t* string_emptyBytes();
void string_initBytes(bytes_t *bytes, unsigned int size);
void string_setByteAt(bytes_t *bytes, uint8_t value, unsigned int position);
uint8_t string_byteAt(bytes_t *bytes, unsigned int position);
void string_freeBytes(bytes_t *bytes);

stringlist_t* string_emptyStringList();
string_t* string_stringAt(stringlist_t *list, unsigned int position);
void string_freeStringList(stringlist_t *strings);

#ifdef __cplusplus
} // end extern "C"
void string_copyStdString(string_t *cstr, std::string &string);
std::string string_toStdString(string_t *string);
string_t* string_fromStdString(std::string &string);
void string_copyStdStringList(stringlist_t *clist, std::vector<std::string> &strings);
#endif

#endif // define GOZWAVE_STRING
