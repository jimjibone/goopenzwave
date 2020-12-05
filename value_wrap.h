#ifndef GOOPENZWAVE_VALUE_WRAP_H
#define GOOPENZWAVE_VALUE_WRAP_H

#include <stdint.h>
#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef enum {
    OZWEXCEPTION_OPTIONS,
    OZWEXCEPTION_CONFIG,
    OZWEXCEPTION_INVALID_HOMEID,
    OZWEXCEPTION_INVALID_VALUEID,
    OZWEXCEPTION_CANNOT_CONVERT_VALUEID,
    OZWEXCEPTION_SECURITY_FAILED,
    OZWEXCEPTION_INVALID_NODEID
} ozw_exception;

typedef struct {
    bool is_ok;
    bool val_bool;
    uint8_t val_byte;
    int16_t val_short;
    int32_t val_int;
    float val_float;
    char* val_string;
    uint8_t* val_raw;
    uint8_t val_raw_len;
    char** val_item_list;
    uint8_t val_item_list_len;
    int32_t* val_value_list;
    uint8_t val_value_list_len;

    bool is_err;
    ozw_exception err_type;
    char* err_msg;
} value_result;

value_result* value_result_new();
void value_result_free(value_result* res);

value_result* ozw_GetValueLabel(uint32_t homeid, uint32_t vid0, uint32_t vid1, int32_t pos); // val_string
value_result* ozw_SetValueLabel(uint32_t homeid, uint32_t vid0, uint32_t vid1, const char* value, int32_t pos);
value_result* ozw_GetValueUnits(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_string
value_result* ozw_SetValueUnits(uint32_t homeid, uint32_t vid0, uint32_t vid1, const char* value);
value_result* ozw_GetValueHelp(uint32_t homeid, uint32_t vid0, uint32_t vid1, int32_t pos); // val_string
value_result* ozw_SetValueHelp(uint32_t homeid, uint32_t vid0, uint32_t vid1, const char* value, int32_t pos);
value_result* ozw_GetValueMin(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_int
value_result* ozw_GetValueMax(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_int
value_result* ozw_IsValueReadOnly(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_bool
value_result* ozw_IsValueWriteOnly(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_bool
value_result* ozw_IsValueSet(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_bool
value_result* ozw_IsValuePolled(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_bool
value_result* ozw_IsValueValid(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_bool
value_result* ozw_GetValueAsBitSet(uint32_t homeid, uint32_t vid0, uint32_t vid1, uint8_t pos); // val_bool
value_result* ozw_GetValueAsBool(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_bool
value_result* ozw_GetValueAsByte(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_byte
value_result* ozw_GetValueAsFloat(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_float
value_result* ozw_GetValueAsInt(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_int
value_result* ozw_GetValueAsShort(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_short
value_result* ozw_GetValueAsString(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_string
value_result* ozw_GetValueAsRaw(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_raw
value_result* ozw_GetValueListSelectionString(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_string
value_result* ozw_GetValueListSelectionInt(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_int
value_result* ozw_GetValueListItems(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_item_list
value_result* ozw_GetValueListValues(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_value_list
value_result* ozw_GetValueFloatPrecision(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_byte
value_result* ozw_SetValueBitSet(uint32_t homeid, uint32_t vid0, uint32_t vid1, uint8_t pos, bool value);
value_result* ozw_SetValueBool(uint32_t homeid, uint32_t vid0, uint32_t vid1, bool value);
value_result* ozw_SetValueByte(uint32_t homeid, uint32_t vid0, uint32_t vid1, uint8_t value);
value_result* ozw_SetValueFloat(uint32_t homeid, uint32_t vid0, uint32_t vid1, float value);
value_result* ozw_SetValueInt(uint32_t homeid, uint32_t vid0, uint32_t vid1, int32_t value);
value_result* ozw_SetValueShort(uint32_t homeid, uint32_t vid0, uint32_t vid1, int16_t value);
value_result* ozw_SetValueRaw(uint32_t homeid, uint32_t vid0, uint32_t vid1, uint8_t const* value, uint8_t length);
value_result* ozw_SetValueString(uint32_t homeid, uint32_t vid0, uint32_t vid1, const char* value);
value_result* ozw_SetValueListSelection(uint32_t homeid, uint32_t vid0, uint32_t vid1, const char* selectedItem);
value_result* ozw_RefreshValue(uint32_t homeid, uint32_t vid0, uint32_t vid1);
value_result* ozw_SetChangeVerified(uint32_t homeid, uint32_t vid0, uint32_t vid1, bool verify);
value_result* ozw_GetChangeVerified(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_bool
value_result* ozw_PressButton(uint32_t homeid, uint32_t vid0, uint32_t vid1);
value_result* ozw_ReleaseButton(uint32_t homeid, uint32_t vid0, uint32_t vid1);
value_result* ozw_SetBitMask(uint32_t homeid, uint32_t vid0, uint32_t vid1, uint32_t mask);
value_result* ozw_GetBitMask(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_int
value_result* ozw_GetBitSetSize(uint32_t homeid, uint32_t vid0, uint32_t vid1); // val_byte

#ifdef __cplusplus
}
#endif

#endif // GOOPENZWAVE_VALUE_WRAP_H
