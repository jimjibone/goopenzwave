package goopenzwave

// #cgo pkg-config: libopenzwave
// #include <stdlib.h>
// #include <stdint.h>
// #include "value_wrap.h"
// #include "util.h"
import "C"
import (
	"errors"
	"unsafe"
)

var (
	ErrInvalidValueID = errors.New("invalid valueid")
	ErrInvalidHomeID  = errors.New("invalid homeid")
)

type ValueError struct {
	Type C.ozw_exception
	Msg  string
}

func (v ValueError) Error() string {
	return v.Msg
}

func newValueError(res *C.value_result) *ValueError {
	if res.is_err == true {
		return &ValueError{
			Type: res.err_type,
			Msg:  C.GoString(res.err_msg),
		}
	}
	return nil
}

// Gets the user-friendly label for the value. pos is the bit to get the label for if its a BitSet ValueID, -1 for no bitset.
func GetValueLabel(id ValueID, pos int32) (string, error) {
	res := C.ozw_GetValueLabel(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1), C.int32_t(pos))
	defer C.value_result_free(res)
	if res.is_err {
		return "", newValueError(res)
	}
	return C.GoString(res.val_string), nil
}

// Sets the user-friendly label for the value. pos is the bit to set the label for if its a BitSet ValueID, -1 for no bitset.
func SetValueLabel(id ValueID, label string, pos int32) error {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	res := C.ozw_SetValueLabel(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1), cstr, C.int32_t(pos))
	defer C.value_result_free(res)
	return newValueError(res)
}

// Gets the units that the value is measured in.
func GetValueUnits(id ValueID) (string, error) {
	res := C.ozw_GetValueUnits(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return "", newValueError(res)
	}
	return C.GoString(res.val_string), nil
}

// Sets the units that the value is measured in.
func SetValueUnits(id ValueID, units string) error {
	cstr := C.CString(units)
	defer C.free(unsafe.Pointer(cstr))
	res := C.ozw_SetValueUnits(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1), cstr)
	defer C.value_result_free(res)
	return newValueError(res)
}

// Gets a help string describing the value's purpose and usage. pos is the bit to get the help for if its a BitSet ValueID, -1 for no bitset.
func GetValueHelp(id ValueID, pos int32) (string, error) {
	res := C.ozw_GetValueHelp(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1), C.int32_t(pos))
	defer C.value_result_free(res)
	if res.is_err {
		return "", newValueError(res)
	}
	return C.GoString(res.val_string), nil
}

// Sets a help string describing the value's purpose and usage. pos is the bit to set the help for if its a BitSet ValueID.
func SetValueHelp(id ValueID, help string, pos int32) error {
	cstr := C.CString(help)
	defer C.free(unsafe.Pointer(cstr))
	res := C.ozw_SetValueHelp(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1), cstr, C.int32_t(pos))
	defer C.value_result_free(res)
	return newValueError(res)
}

// Gets the minimum that this value may contain.
func GetValueMin(id ValueID) (int32, error) {
	res := C.ozw_GetValueMin(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return 0, newValueError(res)
	}
	return int32(res.val_int), nil
}

// Gets the maximum that this value may contain.
func GetValueMax(id ValueID) (int32, error) {
	res := C.ozw_GetValueMax(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return 0, newValueError(res)
	}
	return int32(res.val_int), nil
}

// Test whether the value is read-only.
func IsValueReadOnly(id ValueID) (bool, error) {
	res := C.ozw_IsValueReadOnly(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return false, newValueError(res)
	}
	return bool(res.val_bool), nil
}

// Test whether the value is write-only.
func IsValueWriteOnly(id ValueID) (bool, error) {
	res := C.ozw_IsValueWriteOnly(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return false, newValueError(res)
	}
	return bool(res.val_bool), nil
}

// Test whether the value has been set.
func IsValueSet(id ValueID) (bool, error) {
	res := C.ozw_IsValueSet(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return false, newValueError(res)
	}
	return bool(res.val_bool), nil
}

// Test whether the value is currently being polled.
func IsValuePolled(id ValueID) (bool, error) {
	res := C.ozw_IsValuePolled(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return false, newValueError(res)
	}
	return bool(res.val_bool), nil
}

// Test whether the ValueID is valid.
func IsValueValid(id ValueID) (bool, error) {
	res := C.ozw_IsValueValid(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return false, newValueError(res)
	}
	return bool(res.val_bool), nil
}

// Gets a the value of a Bit from a BitSet ValueID. pos is the bit you want to test for.
func GetValueAsBitSet(id ValueID, pos uint8) (bool, error) {
	res := C.ozw_GetValueAsBitSet(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1), C.uint8_t(pos))
	defer C.value_result_free(res)
	if res.is_err {
		return false, newValueError(res)
	}
	return bool(res.val_bool), nil
}

// Gets a value as a bool.
func GetValueAsBool(id ValueID) (bool, error) {
	res := C.ozw_GetValueAsBool(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return false, newValueError(res)
	}
	return bool(res.val_bool), nil
}

// Gets a value as an 8-bit unsigned integer.
func GetValueAsByte(id ValueID) (byte, error) {
	res := C.ozw_GetValueAsByte(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return 0, newValueError(res)
	}
	return byte(res.val_byte), nil
}

// Gets a value as a float.
func GetValueAsFloat(id ValueID) (float32, error) {
	res := C.ozw_GetValueAsFloat(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return 0.0, newValueError(res)
	}
	return float32(res.val_float), nil
}

// Gets a value as a 32-bit signed integer.
func GetValueAsInt(id ValueID) (int32, error) {
	res := C.ozw_GetValueAsInt(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return 0, newValueError(res)
	}
	return int32(res.val_int), nil
}

// Gets a value as a 16-bit signed integer.
func GetValueAsShort(id ValueID) (int16, error) {
	res := C.ozw_GetValueAsShort(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return 0, newValueError(res)
	}
	return int16(res.val_short), nil
}

// Gets a value as a string.
func GetValueAsString(id ValueID) (string, error) {
	res := C.ozw_GetValueAsString(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return "", newValueError(res)
	}
	return C.GoString(res.val_string), nil
}

// Gets a value as a collection of bytes.
func GetValueAsRaw(id ValueID) ([]byte, error) {
	res := C.ozw_GetValueAsRaw(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return nil, newValueError(res)
	}
	return C.GoBytes(unsafe.Pointer(res.val_raw), C.int(res.val_raw_len)), nil
}

// Gets the selected item from a list (as a string).
func GetValueListSelectionString(id ValueID) (string, error) {
	res := C.ozw_GetValueListSelectionString(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return "", newValueError(res)
	}
	return C.GoString(res.val_string), nil
}

// Gets the selected item from a list (as an integer).
func GetValueListSelectionInt(id ValueID) (int32, error) {
	res := C.ozw_GetValueListSelectionInt(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return 0, newValueError(res)
	}
	return int32(res.val_int), nil
}

// Gets the list of items from a list value.
func GetValueListItems(id ValueID) ([]string, error) {
	res := C.ozw_GetValueListItems(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return nil, newValueError(res)
	}
	size := uint32(res.val_item_list_len)
	var items []string
	for i := uint32(0); i < size; i++ {
		items = append(items, C.GoString(C.str_at(res.val_item_list, C.uint32_t(i))))
	}
	return items, nil
}

// Gets the list of values from a list value.
func GetValueListValues(id ValueID) ([]int32, error) {
	res := C.ozw_GetValueListValues(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return nil, newValueError(res)
	}
	size := uint32(res.val_value_list_len)
	var items []int32
	for i := uint32(0); i < size; i++ {
		items = append(items, int32(C.int32_at(res.val_value_list, C.uint32_t(i))))
	}
	return items, nil
}

// Gets a float value's precision.
func GetValueFloatPrecision(id ValueID) (uint8, error) {
	res := C.ozw_GetValueFloatPrecision(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return 0, newValueError(res)
	}
	return uint8(res.val_byte), nil
}

// Sets the state of a bit in a BitSet ValueID.
func SetValueBitSet(id ValueID, pos uint8, value bool) error {
	res := C.ozw_SetValueBitSet(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1), C.uint8_t(pos), C.bool(value))
	defer C.value_result_free(res)
	if res.is_err {
		return newValueError(res)
	}
	return nil
}

// Sets the state of a bool.
func SetValueBool(id ValueID, value bool) error {
	res := C.ozw_SetValueBool(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1), C.bool(value))
	defer C.value_result_free(res)
	if res.is_err {
		return newValueError(res)
	}
	return nil
}

// Sets the value of a byte.
func SetValueByte(id ValueID, value byte) error {
	res := C.ozw_SetValueByte(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1), C.uint8_t(value))
	defer C.value_result_free(res)
	if res.is_err {
		return newValueError(res)
	}
	return nil
}

// Sets the value of a decimal.
func SetValueFloat(id ValueID, value float32) error {
	res := C.ozw_SetValueFloat(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1), C.float(value))
	defer C.value_result_free(res)
	if res.is_err {
		return newValueError(res)
	}
	return nil
}

// Sets the value of a 32-bit signed integer.
func SetValueInt(id ValueID, value int32) error {
	res := C.ozw_SetValueInt(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1), C.int32_t(value))
	defer C.value_result_free(res)
	if res.is_err {
		return newValueError(res)
	}
	return nil
}

// Sets the value of a 16-bit signed integer.
func SetValueShort(id ValueID, value int16) error {
	res := C.ozw_SetValueShort(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1), C.int16_t(value))
	defer C.value_result_free(res)
	if res.is_err {
		return newValueError(res)
	}
	return nil
}

// Sets the value of a collection of bytes.
func SetValueRaw(id ValueID, value []byte) error {
	cvalue := C.CBytes(value)
	defer C.free(unsafe.Pointer(cvalue))
	res := C.ozw_SetValueRaw(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1), (*C.uchar)(cvalue), C.uint8_t(len(value)))
	defer C.value_result_free(res)
	if res.is_err {
		return newValueError(res)
	}
	return nil
}

// Sets the value from a string, regardless of type.
func SetValueString(id ValueID, value string) error {
	cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(cvalue))
	res := C.ozw_SetValueString(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1), cvalue)
	defer C.value_result_free(res)
	if res.is_err {
		return newValueError(res)
	}
	return nil
}

// Sets the selected item in a list.
func SetValueListSelection(id ValueID, selectedItem string) error {
	cvalue := C.CString(selectedItem)
	defer C.free(unsafe.Pointer(cvalue))
	res := C.ozw_SetValueListSelection(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1), cvalue)
	defer C.value_result_free(res)
	if res.is_err {
		return newValueError(res)
	}
	return nil
}

// Refreshes the specified value from the Z-Wave network.
func RefreshValue(id ValueID) error {
	res := C.ozw_RefreshValue(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return newValueError(res)
	}
	return nil
}

// Sets a flag indicating whether value changes noted upon a refresh should be verified. If so, the
// library will immediately refresh the value a second time whenever a change is observed. This helps to filter
// out spurious data reported occasionally by some devices.
func SetChangeVerified(id ValueID, verify bool) error {
	res := C.ozw_SetChangeVerified(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1), C.bool(verify))
	defer C.value_result_free(res)
	if res.is_err {
		return newValueError(res)
	}
	return nil
}

// Determine if value changes upon a refresh should be verified. If so, the
// library will immediately refresh the value a second time whenever a change is observed. This helps to filter
// out spurious data reported occasionally by some devices.
func GetChangeVerified(id ValueID) (bool, error) {
	res := C.ozw_GetChangeVerified(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return false, newValueError(res)
	}
	return bool(res.val_bool), nil
}

// Starts an activity in a device.
// Since buttons are write-only values that do not report a state, no notification callbacks are sent.
func PressButton(id ValueID) error {
	res := C.ozw_PressButton(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return newValueError(res)
	}
	return nil
}

// Stops an activity in a device.
// Since buttons are write-only values that do not report a state, no notification callbacks are sent.
func ReleaseButton(id ValueID) error {
	res := C.ozw_ReleaseButton(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return newValueError(res)
	}
	return nil
}

// Sets the Valid BitMask for a BitSet ValueID
func SetBitMask(id ValueID, mask uint32) error {
	res := C.ozw_SetBitMask(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1), C.uint32_t(mask))
	defer C.value_result_free(res)
	if res.is_err {
		return newValueError(res)
	}
	return nil
}

// Gets the Valid BitMask for a BitSet ValueID.
func GetBitMask(id ValueID) (uint32, error) {
	res := C.ozw_GetBitMask(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return 0, newValueError(res)
	}
	return uint32(res.val_int), nil
}

// Gets the size of a BitMask ValueID - Either 1, 2 or 4.
func GetBitSetSize(id ValueID) (uint8, error) {
	res := C.ozw_GetBitSetSize(C.uint32_t(id.homeid), C.uint32_t(id.id0), C.uint32_t(id.id1))
	defer C.value_result_free(res)
	if res.is_err {
		return 0, newValueError(res)
	}
	return uint8(res.val_byte), nil
}
