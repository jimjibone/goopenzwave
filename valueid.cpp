#include "valueid.h"
#include <openzwave/value_classes/ValueID.h>

//
// Public member functions.
//

uint32_t valueid_getHomeId(valueid_t v)
{
    OpenZWave::ValueID *valid = (OpenZWave::ValueID*)v;
    return valid->GetHomeId();
}

uint8_t valueid_getNodeId(valueid_t v)
{
    OpenZWave::ValueID *valid = (OpenZWave::ValueID*)v;
    return valid->GetNodeId();
}

valueid_genre valueid_getGenre(valueid_t v)
{
    OpenZWave::ValueID *valid = (OpenZWave::ValueID*)v;
    valueid_genre val_genre;
    switch (valid->GetGenre()) {
        case OpenZWave::ValueID::ValueGenre_Basic:
            val_genre = valueid_genre_basic;
            break;
        case OpenZWave::ValueID::ValueGenre_User:
            val_genre = valueid_genre_user;
            break;
        case OpenZWave::ValueID::ValueGenre_Config:
            val_genre = valueid_genre_config;
            break;
        case OpenZWave::ValueID::ValueGenre_System:
            val_genre = valueid_genre_system;
            break;
        case OpenZWave::ValueID::ValueGenre_Count:
            val_genre = valueid_genre_count;
            break;
    }
    return val_genre;
}

uint8_t valueid_getCommandClassId(valueid_t v)
{
    OpenZWave::ValueID *valid = (OpenZWave::ValueID*)v;
    return valid->GetCommandClassId();
}

uint8_t valueid_getInstance(valueid_t v)
{
    OpenZWave::ValueID *valid = (OpenZWave::ValueID*)v;
    return valid->GetInstance();
}

uint8_t valueid_getIndex(valueid_t v)
{
    OpenZWave::ValueID *valid = (OpenZWave::ValueID*)v;
    return valid->GetIndex();
}

valueid_type valueid_getType(valueid_t v)
{
    OpenZWave::ValueID *valid = (OpenZWave::ValueID*)v;
    valueid_type val_type;
    switch (valid->GetType()) {
        case OpenZWave::ValueID::ValueType_Bool:
            val_type = valueid_type_bool;
            break;
        case OpenZWave::ValueID::ValueType_Byte:
            val_type = valueid_type_byte;
            break;
        case OpenZWave::ValueID::ValueType_Decimal:
            val_type = valueid_type_decimal;
            break;
        case OpenZWave::ValueID::ValueType_Int:
            val_type = valueid_type_int;
            break;
        case OpenZWave::ValueID::ValueType_List:
            val_type = valueid_type_list;
            break;
        case OpenZWave::ValueID::ValueType_Schedule:
            val_type = valueid_type_schedule;
            break;
        case OpenZWave::ValueID::ValueType_Short:
            val_type = valueid_type_short;
            break;
        case OpenZWave::ValueID::ValueType_String:
            val_type = valueid_type_string;
            break;
        case OpenZWave::ValueID::ValueType_Button:
            val_type = valueid_type_button;
            break;
        case OpenZWave::ValueID::ValueType_Raw:
            val_type = valueid_type_raw;
            break;
        // case OpenZWave::ValueID::ValueType_Max:
        //     val_type = valueid_type_max;
        //     break;
    }
    return val_type;
}

uint64_t valueid_getId(valueid_t v)
{
    OpenZWave::ValueID *valid = (OpenZWave::ValueID*)v;
    return valid->GetId();
}
