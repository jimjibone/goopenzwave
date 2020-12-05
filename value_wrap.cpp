#include "value_wrap.h"
#include "openzwave/Manager.h"

value_result* value_result_new()
{
    value_result* res = static_cast<value_result*>(malloc(sizeof(value_result)));
    memset(res, 0, sizeof(value_result));
    return res;
}

void value_result_free(value_result* res)
{
    if (res)
    {
        if (res->val_string) free(res->val_string);
        if (res->val_raw) free(res->val_raw);
        
        if (res->val_item_list)
        {
            for (uint8_t i = 0; i < res->val_item_list_len; ++i)
            {
                if (res->val_item_list[i])
                {
                    free(res->val_item_list[i]);
                }
            }
            free(res->val_item_list);
        }

        if (res->val_value_list) free(res->val_value_list);
        if (res->err_msg) free(res->err_msg);
        free(res);
    }
}

void value_result_set_err(value_result* res, OpenZWave::OZWException& e)
{
    if (res)
    {
        res->is_ok = false;
        res->is_err = true;

        if (res->err_msg) free(res->err_msg);
        res->err_msg = strdup(e.GetMsg().c_str());

        switch (e.GetType())
        {
        case OpenZWave::OZWException::OZWEXCEPTION_OPTIONS:
            res->err_type = OZWEXCEPTION_OPTIONS;
            break;
        case OpenZWave::OZWException::OZWEXCEPTION_CONFIG:
            res->err_type = OZWEXCEPTION_CONFIG;
            break;
        case OpenZWave::OZWException::OZWEXCEPTION_INVALID_HOMEID:
            res->err_type = OZWEXCEPTION_INVALID_HOMEID;
            break;
        case OpenZWave::OZWException::OZWEXCEPTION_INVALID_VALUEID:
            res->err_type = OZWEXCEPTION_INVALID_VALUEID;
            break;
        case OpenZWave::OZWException::OZWEXCEPTION_CANNOT_CONVERT_VALUEID:
            res->err_type = OZWEXCEPTION_CANNOT_CONVERT_VALUEID;
            break;
        case OpenZWave::OZWException::OZWEXCEPTION_SECURITY_FAILED:
            res->err_type = OZWEXCEPTION_SECURITY_FAILED;
            break;
        case OpenZWave::OZWException::OZWEXCEPTION_INVALID_NODEID:
            res->err_type = OZWEXCEPTION_INVALID_NODEID;
            break;
        }
    }
}

OpenZWave::ValueID build_valueid(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    return OpenZWave::ValueID(homeid, (((static_cast<uint64_t>(vid1) << 32) & 0xFFFFFFFF00000000) | (static_cast<uint64_t>(vid0) & 0x00000000FFFFFFFF)));
}

value_result* ozw_GetValueLabel(uint32_t homeid, uint32_t vid0, uint32_t vid1, int32_t pos)
{
    value_result* res = value_result_new();
    try {
        // string GetValueLabel(ValueID const& _id, int32 _pos = -1);
        std::string str = OpenZWave::Manager::Get()->GetValueLabel(build_valueid(homeid, vid0, vid1), pos);
        res->is_ok = true;
        res->val_string = strdup(str.c_str());
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_SetValueLabel(uint32_t homeid, uint32_t vid0, uint32_t vid1, const char* value, int32_t pos)
{
    value_result* res = value_result_new();
    try {
        // void SetValueLabel(ValueID const& _id, string const& _value, int32 _pos = -1);
        OpenZWave::Manager::Get()->SetValueLabel(build_valueid(homeid, vid0, vid1), value, pos);
        res->is_ok = true;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_GetValueUnits(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // string GetValueUnits(ValueID const& _id);
        std::string str = OpenZWave::Manager::Get()->GetValueUnits(build_valueid(homeid, vid0, vid1));
        res->is_ok = true;
        res->val_string = strdup(str.c_str());
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_SetValueUnits(uint32_t homeid, uint32_t vid0, uint32_t vid1, const char* value)
{
    value_result* res = value_result_new();
    try {
        // void SetValueUnits(ValueID const& _id, string const& _value);
        OpenZWave::Manager::Get()->SetValueUnits(build_valueid(homeid, vid0, vid1), value);
        res->is_ok = true;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_GetValueHelp(uint32_t homeid, uint32_t vid0, uint32_t vid1, int32_t pos)
{
    value_result* res = value_result_new();
    try {
        // string GetValueHelp(ValueID const& _id, int32 _pos = -1);
        std::string str = OpenZWave::Manager::Get()->GetValueHelp(build_valueid(homeid, vid0, vid1), pos);
        res->is_ok = true;
        res->val_string = strdup(str.c_str());
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_SetValueHelp(uint32_t homeid, uint32_t vid0, uint32_t vid1, const char* value, int32_t pos)
{
    value_result* res = value_result_new();
    try {
        // void SetValueHelp(ValueID const& _id, string const& _value, int32 _pos = -1);
        OpenZWave::Manager::Get()->SetValueHelp(build_valueid(homeid, vid0, vid1), value, pos);
        res->is_ok = true;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_GetValueMin(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // int32 GetValueMin(ValueID const& _id);
        res->val_int = OpenZWave::Manager::Get()->GetValueMin(build_valueid(homeid, vid0, vid1));
        res->is_ok = true;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_GetValueMax(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // int32 GetValueMax(ValueID const& _id);
        res->val_int = OpenZWave::Manager::Get()->GetValueMax(build_valueid(homeid, vid0, vid1));
        res->is_ok = true;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_IsValueReadOnly(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool IsValueReadOnly(ValueID const& _id);
        res->val_bool = OpenZWave::Manager::Get()->IsValueReadOnly(build_valueid(homeid, vid0, vid1));
        res->is_ok = true;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_IsValueWriteOnly(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool IsValueWriteOnly(ValueID const& _id);
        res->val_bool = OpenZWave::Manager::Get()->IsValueWriteOnly(build_valueid(homeid, vid0, vid1));
        res->is_ok = true;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_IsValueSet(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool IsValueSet(ValueID const& _id);
        res->val_bool = OpenZWave::Manager::Get()->IsValueSet(build_valueid(homeid, vid0, vid1));
        res->is_ok = true;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_IsValuePolled(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool IsValuePolled(ValueID const& _id);
        res->val_bool = OpenZWave::Manager::Get()->IsValuePolled(build_valueid(homeid, vid0, vid1));
        res->is_ok = true;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_IsValueValid(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool IsValueValid(ValueID const& _id);
        res->val_bool = OpenZWave::Manager::Get()->IsValueValid(build_valueid(homeid, vid0, vid1));
        res->is_ok = true;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_GetValueAsBitSet(uint32_t homeid, uint32_t vid0, uint32_t vid1, uint8_t pos)
{
    value_result* res = value_result_new();
    try {
        // bool GetValueAsBitSet(ValueID const& _id, uint8 _pos, bool* o_value);
        bool out;
        res->is_ok = OpenZWave::Manager::Get()->GetValueAsBitSet(build_valueid(homeid, vid0, vid1), pos, &out);
        res->val_bool = out;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_GetValueAsBool(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool GetValueAsBool(ValueID const& _id, bool* o_value);
        bool out;
        res->is_ok = OpenZWave::Manager::Get()->GetValueAsBool(build_valueid(homeid, vid0, vid1), &out);
        res->val_bool = out;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_GetValueAsByte(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool GetValueAsByte(ValueID const& _id, uint8* o_value);
        uint8_t out;
        res->is_ok = OpenZWave::Manager::Get()->GetValueAsByte(build_valueid(homeid, vid0, vid1), &out);
        res->val_byte = out;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_GetValueAsFloat(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool GetValueAsFloat(ValueID const& _id, float* o_value);
        float out;
        res->is_ok = OpenZWave::Manager::Get()->GetValueAsFloat(build_valueid(homeid, vid0, vid1), &out);
        res->val_float = out;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_GetValueAsInt(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool GetValueAsInt(ValueID const& _id, int32* o_value);
        int32_t out;
        res->is_ok = OpenZWave::Manager::Get()->GetValueAsInt(build_valueid(homeid, vid0, vid1), &out);
        res->val_int = out;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_GetValueAsShort(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool GetValueAsShort(ValueID const& _id, int16* o_value);
        int16_t out;
        res->is_ok = OpenZWave::Manager::Get()->GetValueAsShort(build_valueid(homeid, vid0, vid1), &out);
        res->val_short = out;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_GetValueAsString(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool GetValueAsString(ValueID const& _id, string* o_value);
        std::string out;
        res->is_ok = OpenZWave::Manager::Get()->GetValueAsString(build_valueid(homeid, vid0, vid1), &out);
        res->val_string = strdup(out.c_str());
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_GetValueAsRaw(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool GetValueAsRaw(ValueID const& _id, uint8** o_value, uint8* o_length);
        uint8_t* out = nullptr;
        uint8_t outlen = 0;
        res->is_ok = OpenZWave::Manager::Get()->GetValueAsRaw(build_valueid(homeid, vid0, vid1), &out, &outlen);
        res->val_raw = out;
        res->val_raw_len = outlen;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_GetValueListSelectionString(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool GetValueListSelection(ValueID const& _id, string* o_value);
        std::string out;
        res->is_ok = OpenZWave::Manager::Get()->GetValueListSelection(build_valueid(homeid, vid0, vid1), &out);
        res->val_string = strdup(out.c_str());
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_GetValueListSelectionInt(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool GetValueListSelection(ValueID const& _id, int32* o_value);
        int32_t out;
        res->is_ok = OpenZWave::Manager::Get()->GetValueListSelection(build_valueid(homeid, vid0, vid1), &out);
        res->val_int = out;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_GetValueListItems(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool GetValueListItems(ValueID const& _id, vector<string>* o_value);
        std::vector<std::string> out;
        res->is_ok = OpenZWave::Manager::Get()->GetValueListItems(build_valueid(homeid, vid0, vid1), &out);
        res->val_item_list_len = out.size();
        res->val_item_list = static_cast<char**>(malloc(out.size() * sizeof(char*)));
        for (uint8_t i = 0; i < res->val_item_list_len; ++i)
        {
            res->val_item_list[i] = strdup(out[i].c_str());
        }
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_GetValueListValues(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool GetValueListValues(ValueID const& _id, vector<int32>* o_value);
        std::vector<int32> out;
        res->is_ok = OpenZWave::Manager::Get()->GetValueListValues(build_valueid(homeid, vid0, vid1), &out);
        res->val_value_list_len = out.size();
        res->val_value_list = static_cast<int32_t*>(malloc(out.size() * sizeof(int32_t)));
        for (uint8_t i = 0; i < res->val_value_list_len; ++i)
        {
            res->val_value_list[i] = out[i];
        }
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_GetValueFloatPrecision(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool GetValueFloatPrecision(ValueID const& _id, uint8* o_value);
        uint8_t out;
        res->is_ok = OpenZWave::Manager::Get()->GetValueFloatPrecision(build_valueid(homeid, vid0, vid1), &out);
        res->val_byte = out;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_SetValueBitSet(uint32_t homeid, uint32_t vid0, uint32_t vid1, uint8_t pos, bool value)
{
    value_result* res = value_result_new();
    try {
        // bool SetValue(ValueID const& _id, uint8 _pos, bool const _value);
        res->is_ok = OpenZWave::Manager::Get()->SetValue(build_valueid(homeid, vid0, vid1), pos, value);
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_SetValueBool(uint32_t homeid, uint32_t vid0, uint32_t vid1, bool value)
{
    value_result* res = value_result_new();
    try {
        // bool SetValue(ValueID const& _id, bool const _value);
        res->is_ok = OpenZWave::Manager::Get()->SetValue(build_valueid(homeid, vid0, vid1), value);
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_SetValueByte(uint32_t homeid, uint32_t vid0, uint32_t vid1, uint8 value)
{
    value_result* res = value_result_new();
    try {
        // bool SetValue(ValueID const& _id, uint8 const _value);
        res->is_ok = OpenZWave::Manager::Get()->SetValue(build_valueid(homeid, vid0, vid1), value);
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_SetValueFloat(uint32_t homeid, uint32_t vid0, uint32_t vid1, float value)
{
    value_result* res = value_result_new();
    try {
        // bool SetValue(ValueID const& _id, float const _value);
        res->is_ok = OpenZWave::Manager::Get()->SetValue(build_valueid(homeid, vid0, vid1), value);
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_SetValueInt(uint32_t homeid, uint32_t vid0, uint32_t vid1, int32 value)
{
    value_result* res = value_result_new();
    try {
        // bool SetValue(ValueID const& _id, int32 const _value);
        res->is_ok = OpenZWave::Manager::Get()->SetValue(build_valueid(homeid, vid0, vid1), value);
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_SetValueShort(uint32_t homeid, uint32_t vid0, uint32_t vid1, int16 value)
{
    value_result* res = value_result_new();
    try {
        // bool SetValue(ValueID const& _id, int16 const _value);
        res->is_ok = OpenZWave::Manager::Get()->SetValue(build_valueid(homeid, vid0, vid1), value);
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_SetValueRaw(uint32_t homeid, uint32_t vid0, uint32_t vid1, uint8_t const* value, uint8_t length)
{
    value_result* res = value_result_new();
    try {
        // bool SetValue(ValueID const& _id, uint8 const* _value, uint8 const _length);
        res->is_ok = OpenZWave::Manager::Get()->SetValue(build_valueid(homeid, vid0, vid1), value, length);
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_SetValueString(uint32_t homeid, uint32_t vid0, uint32_t vid1, const char* value)
{
    value_result* res = value_result_new();
    try {
        // bool SetValue(ValueID const& _id, string const& _value);
        res->is_ok = OpenZWave::Manager::Get()->SetValue(build_valueid(homeid, vid0, vid1), std::string(value));
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_SetValueListSelection(uint32_t homeid, uint32_t vid0, uint32_t vid1, const char* selectedItem)
{
    value_result* res = value_result_new();
    try {
        // bool SetValueListSelection(ValueID const& _id, string const& _selectedItem);
        res->is_ok = OpenZWave::Manager::Get()->SetValueListSelection(build_valueid(homeid, vid0, vid1), std::string(selectedItem));
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_RefreshValue(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool RefreshValue(ValueID const& _id);
        res->is_ok = OpenZWave::Manager::Get()->RefreshValue(build_valueid(homeid, vid0, vid1));
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_SetChangeVerified(uint32_t homeid, uint32_t vid0, uint32_t vid1, bool verify)
{
    value_result* res = value_result_new();
    try {
        // void SetChangeVerified(ValueID const& _id, bool _verify);
        OpenZWave::Manager::Get()->SetChangeVerified(build_valueid(homeid, vid0, vid1), verify);
        res->is_ok = true;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_GetChangeVerified(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // void GetChangeVerified(ValueID const& _id);
        res->val_bool = OpenZWave::Manager::Get()->GetChangeVerified(build_valueid(homeid, vid0, vid1));
        res->is_ok = true;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_PressButton(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool PressButton(ValueID const& _id);
        res->is_ok = OpenZWave::Manager::Get()->PressButton(build_valueid(homeid, vid0, vid1));
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_ReleaseButton(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool ReleaseButton(ValueID const& _id);
        res->is_ok = OpenZWave::Manager::Get()->ReleaseButton(build_valueid(homeid, vid0, vid1));
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_SetBitMask(uint32_t homeid, uint32_t vid0, uint32_t vid1, uint32_t mask)
{
    value_result* res = value_result_new();
    try {
        // bool SetBitMask(ValueID const& _id, uint32 _mask);
        res->is_ok = OpenZWave::Manager::Get()->SetBitMask(build_valueid(homeid, vid0, vid1), mask);
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_GetBitMask(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool GetBitMask(ValueID const& _id, int32* o_mask);
        int32_t out;
        res->is_ok = OpenZWave::Manager::Get()->GetBitMask(build_valueid(homeid, vid0, vid1), &out);
        res->val_int = out;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}

value_result* ozw_GetBitSetSize(uint32_t homeid, uint32_t vid0, uint32_t vid1)
{
    value_result* res = value_result_new();
    try {
        // bool GetBitSetSize(ValueID const& _id, uint8* o_size);
        uint8_t out;
        res->is_ok = OpenZWave::Manager::Get()->GetBitSetSize(build_valueid(homeid, vid0, vid1), &out);
        res->val_byte = out;
    } catch (OpenZWave::OZWException& e) {
        value_result_set_err(res, e);
    }
    return res;
}
