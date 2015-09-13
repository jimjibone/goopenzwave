package goopenzwave

// #cgo LDFLAGS: -lopenzwave -L/usr/local/lib
// #cgo CPPFLAGS: -I/usr/local/include -I/usr/local/include/openzwave
// #include "manager.h"
// #include "valueid.h"
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

// GetValueLabel returns the user-friendly label for the value.
func GetValueLabel(homeID uint32, valueID uint64) string {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	cstring := C.manager_getValueLabel(cmanager, cvalueid)
	gostring := C.GoString(cstring.data)
	C.string_freeString(cstring)
	return gostring
}

// SetValueLabel sets the user-friendly label for the value.
func SetValueLabel(homeID uint32, valueID uint64, value string) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	cstring := C.CString(value)
	C.manager_setValueLabel(cmanager, cvalueid, cstring)
	C.free(unsafe.Pointer(cstring))
}

// GetValueUnits returns the units that the value is measured in.
func GetValueUnits(homeID uint32, valueID uint64) string {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	cstring := C.manager_getValueUnits(cmanager, cvalueid)
	gostring := C.GoString(cstring.data)
	C.string_freeString(cstring)
	return gostring
}

// SetValueUnits sets the units that the value is measured in.
func SetValueUnits(homeID uint32, valueID uint64, value string) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	cstring := C.CString(value)
	C.manager_setValueUnits(cmanager, cvalueid, cstring)
	C.free(unsafe.Pointer(cstring))
}

// GetValueHelp returns a help string describing the value's purpose and usage.
func GetValueHelp(homeID uint32, valueID uint64) string {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	cstring := C.manager_getValueHelp(cmanager, cvalueid)
	gostring := C.GoString(cstring.data)
	C.string_freeString(cstring)
	return gostring
}

// SetValueHelp sets a help string describing the value's purpose and usage.
func SetValueHelp(homeID uint32, valueID uint64, value string) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	cstring := C.CString(value)
	C.manager_setValueHelp(cmanager, cvalueid, cstring)
	C.free(unsafe.Pointer(cstring))
}

// GetValueMin returns the minimum that this value may contain.
func GetValueMin(homeID uint32, valueID uint64) int32 {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	return int32(C.manager_getValueMin(cmanager, cvalueid))
}

// GetValueMax returns the maximum that this value may contain.
func GetValueMax(homeID uint32, valueID uint64) int32 {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	return int32(C.manager_getValueMax(cmanager, cvalueid))
}

// IsValueReadOnly returns true if the value is read-only.
func IsValueReadOnly(homeID uint32, valueID uint64) bool {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	return bool(C.manager_isValueReadOnly(cmanager, cvalueid))
}

// IsValueWriteOnly returns true if the value is write-only.
func IsValueWriteOnly(homeID uint32, valueID uint64) bool {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	return bool(C.manager_isValueWriteOnly(cmanager, cvalueid))
}

// IsValueSet returns true if the value has been set.
func IsValueSet(homeID uint32, valueID uint64) bool {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	return bool(C.manager_isValueSet(cmanager, cvalueid))
}

// IsValuePolled returns true if the value is currently being polled.
func IsValuePolled(homeID uint32, valueID uint64) bool {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	return bool(C.manager_isValuePolled(cmanager, cvalueid))
}

// GetValueAsBool returns the value as a bool. It will also return an error if
// the value is not a bool type.
func GetValueAsBool(homeID uint32, valueID uint64) (bool, error) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	var cbool C.bool
	ok := bool(C.manager_getValueAsBool(cmanager, cvalueid, &cbool))
	if ok == false {
		return bool(cbool), fmt.Errorf("value is not of bool type")
	}
	return bool(cbool), nil
}

// GetValueAsByte returns the value as an 8-bit unsigned integer. It will also
// return an error if the value is not of byte type.
func GetValueAsByte(homeID uint32, valueID uint64) (byte, error) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	var cbyte C.uint8_t
	ok := bool(C.manager_getValueAsByte(cmanager, cvalueid, &cbyte))
	if ok == false {
		return byte(cbyte), fmt.Errorf("value is not of byte type")
	}
	return byte(cbyte), nil
}

// GetValueAsFloat returns the value as a float. It will also return an error if
// the value is not a decimal type.
func GetValueAsFloat(homeID uint32, valueID uint64) (float32, error) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	var cfloat C.float
	ok := bool(C.manager_getValueAsFloat(cmanager, cvalueid, &cfloat))
	if ok == false {
		return float32(cfloat), fmt.Errorf("value is not of decimal type")
	}
	return float32(cfloat), nil
}

// GetValueAsInt returns the value as a 32-bit signed integer. It will also
// return an error if the value is not of 32-bit signed integer type.
func GetValueAsInt(homeID uint32, valueID uint64) (int32, error) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	var cint C.int32_t
	ok := bool(C.manager_getValueAsInt(cmanager, cvalueid, &cint))
	if ok == false {
		return int32(cint), fmt.Errorf("value is not of 32-bit signed integer type")
	}
	return int32(cint), nil
}

// GetValueAsShort returns the value as a 16-bit signed integer. It will also
// return an error if the value is not of 16-bit signed integer type.
func GetValueAsShort(homeID uint32, valueID uint64) (int16, error) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	var cshort C.int16_t
	ok := bool(C.manager_getValueAsShort(cmanager, cvalueid, &cshort))
	if ok == false {
		return int16(cshort), fmt.Errorf("value is not of 16-bit signed integer type")
	}
	return int16(cshort), nil
}

// GetValueAsString returns the value as a string, regardless of its actual
// type.
func GetValueAsString(homeID uint32, valueID uint64) string {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	cstring := C.string_emptyString()
	_ = bool(C.manager_getValueAsString(cmanager, cvalueid, cstring))
	gostring := C.GoString(cstring.data)
	C.string_freeString(cstring)
	return gostring
}

// GetValueAsRaw returns the value as a raw byte slice. It will also return an
// error if the value is not of raw type.
func GetValueAsRaw(homeID uint32, valueID uint64) ([]byte, error) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	cbytes := C.string_emptyBytes()
	ok := bool(C.manager_getValueAsRaw(cmanager, cvalueid, cbytes))
	if ok == false {
		return nil, fmt.Errorf("value is not of raw type")
	}
	gobytes := make([]byte, int(cbytes.length))
	for i := 0; i < int(cbytes.length); i++ {
		gobytes[i] = byte(C.string_byteAt(cbytes, C.size_t(i)))
	}
	return gobytes, nil
}

// GetValueListSelectionAsString returns selected item from a list as a string.
// It will also return an error if the value is not of list type.
func GetValueListSelectionAsString(homeID uint32, valueID uint64) (string, error) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	cstring := C.string_emptyString()
	ok := bool(C.manager_getValueListSelectionAsString(cmanager, cvalueid, cstring))
	gostring := C.GoString(cstring.data)
	C.string_freeString(cstring)
	if ok == false {
		return gostring, fmt.Errorf("value is not of list type")
	}
	return gostring, nil
}

// GetValueListSelectionAsInt32 returns selected item from a list as an integer.
// It will also return an error if the value is not of list type.
func GetValueListSelectionAsInt32(homeID uint32, valueID uint64) (int32, error) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	var cint C.int32_t
	ok := bool(C.manager_getValueListSelectionAsInt32(cmanager, cvalueid, &cint))
	if ok == false {
		return int32(cint), fmt.Errorf("value is not of list type")
	}
	return int32(cint), nil
}

// GetValueListItems returns the list of items from a list value. It will also
// return an error if the value is not of list type.
func GetValueListItems(homeID uint32, valueID uint64) ([]string, error) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	cstringlist := C.string_emptyStringList()
	ok := bool(C.manager_getValueListItems(cmanager, cvalueid, cstringlist))
	gostringlist := make([]string, int(cstringlist.length))
	if ok == false {
		return nil, fmt.Errorf("value is not of list type")
	}
	for i := 0; i < int(cstringlist.length); i++ {
		cstring := C.string_stringAt(cstringlist, C.size_t(i))
		gostringlist[i] = C.GoString(cstring.data)
	}
	C.string_freeStringList(cstringlist)
	return gostringlist, nil
}

// GetValueFloatPrecision returns the float value's precision. It will also
// return an error if the value is not of decimal type.
func GetValueFloatPrecision(homeID uint32, valueID uint64) (uint8, error) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	var cprecision C.uint8_t
	ok := bool(C.manager_getValueFloatPrecision(cmanager, cvalueid, &cprecision))
	if ok == false {
		return uint8(cprecision), fmt.Errorf("value is not of decimal type")
	}
	return uint8(cprecision), nil
}

// SetValueBool sets the state of a bool. It will return an error if the value
// is not of bool type.
//
// Due to the possibility of a device being asleep, the command is assumed to
// succeed, and the value held by the node is updated directly. This will be
// reverted by a future status message from the device if the Z-Wave message
// actually failed to get through. Notification callbacks will be sent in both
// cases.
func SetValueBool(homeID uint32, valueID uint64, value bool) error {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	ok := bool(C.manager_setValueBool(cmanager, cvalueid, C.bool(value)))
	if ok == false {
		return fmt.Errorf("value is not of bool type")
	}
	return nil
}

// SetValueUint8 sets the value of a byte. It will return an error if the value
// is not of byte type.
//
// Due to the possibility of a device being asleep, the command is assumed to
// succeed, and the value held by the node is updated directly. This will be
// reverted by a future status message from the device if the Z-Wave message
// actually failed to get through. Notification callbacks will be sent in both
// cases.
func SetValueUint8(homeID uint32, valueID uint64, value uint8) error {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	ok := bool(C.manager_setValueUint8(cmanager, cvalueid, C.uint8_t(value)))
	if ok == false {
		return fmt.Errorf("value is not of byte type")
	}
	return nil
}

// SetValueFloat sets the value of a decimal. It will return an error if the
// value is not of decimal type.
//
// It is usually better to handle decimal values using strings rather than
// floats, to avoid floating point accuracy issues. Due to the possibility of a
// device being asleep, the command is assumed to succeed, and the value held by
// the node is updated directly. This will be reverted by a future status
// message from the device if the Z-Wave message actually failed to get through.
// Notification callbacks will be sent in both cases.
func SetValueFloat(homeID uint32, valueID uint64, value float32) error {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	ok := bool(C.manager_setValueFloat(cmanager, cvalueid, C.float(value)))
	if ok == false {
		return fmt.Errorf("value is not of decimal type")
	}
	return nil
}

// SetValueInt32 sets the value of a 32-bit signed integer. It will return an
// error if the value is not of 32-bit signed integer type.
//
// Due to the possibility of a device being asleep, the command is assumed to
// succeed, and the value held by the node is updated directly. This will be
// reverted by a future status message from the device if the Z-Wave message
// actually failed to get through. Notification callbacks will be sent in both
// cases.
func SetValueInt32(homeID uint32, valueID uint64, value int32) error {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	ok := bool(C.manager_setValueInt32(cmanager, cvalueid, C.int32_t(value)))
	if ok == false {
		return fmt.Errorf("value is not of 32-bit signed integer type")
	}
	return nil
}

// SetValueInt16 sets the value of a 16-bit signed integer. It will return an
// error if the value is not of 16-bit signed integer type.
//
// Due to the possibility of a device being asleep, the command is assumed to
// succeed, and the value held by the node is updated directly. This will be
// reverted by a future status message from the device if the Z-Wave message
// actually failed to get through. Notification callbacks will be sent in both
// cases.
func SetValueInt16(homeID uint32, valueID uint64, value int16) error {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	ok := bool(C.manager_setValueInt16(cmanager, cvalueid, C.int16_t(value)))
	if ok == false {
		return fmt.Errorf("value is not of 16-bit signed integer type")
	}
	return nil
}

// SetValueBytes sets the value of a raw value. It will return an error if the
// value is not of raw type.
//
// Due to the possibility of a device being asleep, the command is assumed to
// succeed, and the value held by the node is updated directly. This will be
// reverted by a future status message from the device if the Z-Wave message
// actually failed to get through. Notification callbacks will be sent in both
// cases.
func SetValueBytes(homeID uint32, valueID uint64, value []byte) error {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	cbytes := C.string_emptyBytes()
	C.string_initBytes(cbytes, C.size_t(len(value)))
	for i := range value {
		C.string_setByteAt(cbytes, C.uint8_t(value[i]), C.size_t(i))
	}
	ok := bool(C.manager_setValueBytes(cmanager, cvalueid, cbytes))
	C.string_freeBytes(cbytes)
	if ok == false {
		return fmt.Errorf("value is not of raw type")
	}
	return nil
}

// SetValueString sets the value from a string, regardless of type. It will
// return an error if the value could not be parsed into the correct type for
// the value.
//
// Due to the possibility of a device being asleep, the command is assumed to
// succeed, and the value held by the node is updated directly. This will be
// reverted by a future status message from the device if the Z-Wave message
// actually failed to get through. Notification callbacks will be sent in both
// cases.
func SetValueString(homeID uint32, valueID uint64, value string) error {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	cstring := C.CString(value)
	ok := bool(C.manager_setValueString(cmanager, cvalueid, cstring))
	C.free(unsafe.Pointer(cstring))
	if ok == false {
		return fmt.Errorf("could not parse string into correct type for value")
	}
	return nil
}

// SetValueListSelection sets the selected item in a list. It will return an
// error if the value is not of list type or if the selection is not in the
// list.
//
// Due to the possibility of a device being asleep, the command is assumed to
// succeed, and the value held by the node is updated directly. This will be
// reverted by a future status message from the device if the Z-Wave message
// actually failed to get through. Notification callbacks will be sent in both
// cases.
func SetValueListSelection(homeID uint32, valueID uint64, selection string) error {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	cstring := C.CString(selection)
	ok := bool(C.manager_setValueListSelection(cmanager, cvalueid, cstring))
	C.free(unsafe.Pointer(cstring))
	if ok == false {
		return fmt.Errorf("value is not of list type or selection is not in the list")
	}
	return nil
}

// RefreshValue refreshes the specified value from the Z-Wave network. It will
// return true if the driver and node were found, otherwise false.
//
// A call to this function causes the library to send a message to the network
// to retrieve the current value of the specified ValueID (just like a poll,
// except only one-time, not recurring).
func RefreshValue(homeID uint32, valueID uint64) bool {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	return bool(C.manager_refreshValue(cmanager, cvalueid))
}

// SetChangeVerified sets a flag indicating whether value changes noted upon a
// refresh should be verified. If so, the library will immediately refresh the
// value a second time whenever a change is observed. This helps to filter out
// spurious data reported occasionally by some devices.
func SetChangeVerified(homeID uint32, valueID uint64, verify bool) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	C.manager_setChangeVerified(cmanager, cvalueid, C.bool(verify))
}

// GetChangeVerified returns true if value changes upon a refresh should be
// verified. If so, the library will immediately refresh the value a second time
// whenever a change is observed. This helps to filter out spurious data
// reported occasionally by some devices.
func GetChangeVerified(homeID uint32, valueID uint64) bool {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	return bool(C.manager_getChangeVerified(cmanager, cvalueid))
}

// PressButton starts an activity in a device. It will return an error if the
// value is not of button type.
//
// Since buttons are write-only values that do not report a state, no
// notification callbacks are sent.
func PressButton(homeID uint32, valueID uint64) error {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	ok := bool(C.manager_pressButton(cmanager, cvalueid))
	if ok == false {
		return fmt.Errorf("value is not of button type")
	}
	return nil
}

// ReleaseButton stops an activity in a device. It will return an error if the
// value is not of button type.
//
// Since buttons are write-only values that do not report a state, no
// notification callbacks are sent.
func ReleaseButton(homeID uint32, valueID uint64) error {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	ok := bool(C.manager_releaseButton(cmanager, cvalueid))
	if ok == false {
		return fmt.Errorf("value is not of button type")
	}
	return nil
}

//
// Also, climate control schedules.
//

// GetNumSwitchPoints returns the number of switch points defined in a schedule.
// It will return zero if the value if not of schedule type.
func GetNumSwitchPoints(homeID uint32, valueID uint64) (uint8, error) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	result := uint8(C.manager_getNumSwitchPoints(cmanager, cvalueid))
	if result == 0 {
		return result, fmt.Errorf("value is not of schedule type")
	}
	return result, nil
}

// SetSwitchPoint sets a switch point in the schedule. It will return an error
// if the value is not of schedule type.
//
// Inserts a new switch point into the schedule, unless a switch point already
// exists at the specified time in which case that switch point is updated with
// the new setback value instead. A maximum of nine switch points can be set in
// the schedule.
func SetSwitchPoint(homeID uint32, valueID uint64, hours, minutes uint8, setback int8) error {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	ok := bool(C.manager_setSwitchPoint(cmanager, cvalueid, C.uint8_t(hours), C.uint8_t(minutes), C.int8_t(setback)))
	if ok == false {
		return fmt.Errorf("value is not of schedule type")
	}
	return nil
}

// RemoveSwitchPoint removes a switch point from the schedule. It will return an
// error if the value is not of schedule type or there is no switch point with
// the specified time values.
func RemoveSwitchPoint(homeID uint32, valueID uint64, hours, minutes uint8) error {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	ok := bool(C.manager_removeSwitchPoint(cmanager, cvalueid, C.uint8_t(hours), C.uint8_t(minutes)))
	if ok == false {
		return fmt.Errorf("value is not of schedule type or no switch point found with specified time values")
	}
	return nil
}

// ClearSwitchPoints clears all switch points from the schedule.
func ClearSwitchPoints(homeID uint32, valueID uint64) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	C.manager_clearSwitchPoints(cmanager, cvalueid)
}

// GetSwitchPoint returns switch point data from the schedule. It will also
// return an error if the value is not of schedule type.
//
// It retrieves the time and setback values from a switch point in the schedule.
func GetSwitchPoint(homeID uint32, valueID uint64, idx uint8) (uint8, uint8, int8, error) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	var chours C.uint8_t
	var cminutes C.uint8_t
	var csetback C.int8_t
	ok := bool(C.manager_getSwitchPoint(cmanager, cvalueid, C.uint8_t(idx), &chours, &cminutes, &csetback))
	if ok == false {
		return uint8(chours), uint8(cminutes), int8(csetback), fmt.Errorf("value is not of schedule type")
	}
	return uint8(chours), uint8(cminutes), int8(csetback), nil
}
