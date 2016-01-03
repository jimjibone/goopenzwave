package goopenzwave

// #include "valueid.h"
// #include <stdlib.h>
import "C"
import (
	"fmt"
)

// ValueIDGenre defines a type for the valueid genre enum.
type ValueIDGenre int

const (
	ValueIDGenreBasic ValueIDGenre = iota
	ValueIDGenreUser
	ValueIDGenreConfig
	ValueIDGenreSystem
	ValueIDGenreCount
)

func (v ValueIDGenre) String() string {
	switch v {
	case ValueIDGenreBasic:
		return "Basic"
	case ValueIDGenreUser:
		return "User"
	case ValueIDGenreConfig:
		return "Config"
	case ValueIDGenreSystem:
		return "System"
	case ValueIDGenreCount:
		return "Count"
	}
	return "UNKNOWN"
}

// ValueIDType defines a type for the valueid type enum.
type ValueIDType int

const (
	ValueIDTypeBool ValueIDType = iota
	ValueIDTypeByte
	ValueIDTypeDecimal
	ValueIDTypeInt
	ValueIDTypeList
	ValueIDTypeSchedule
	ValueIDTypeShort
	ValueIDTypeString
	ValueIDTypeButton
	ValueIDTypeRaw
	ValueIDTypeMax
)

func (v ValueIDType) String() string {
	switch v {
	case ValueIDTypeBool:
		return "Bool"
	case ValueIDTypeByte:
		return "Byte"
	case ValueIDTypeDecimal:
		return "Decimal"
	case ValueIDTypeInt:
		return "Int"
	case ValueIDTypeList:
		return "List"
	case ValueIDTypeSchedule:
		return "Schedule"
	case ValueIDTypeShort:
		return "Short"
	case ValueIDTypeString:
		return "String"
	case ValueIDTypeButton:
		return "Button"
	case ValueIDTypeRaw: // also ValueIDTypeMax
		return "Raw/Max"
	}
	return "UNKNOWN"
}

// ValueID contains all appropriate information available for a ValueID from the
// OpenZWave library. You should not create a new ValueID manually, but receive
// it from the goopenzwave package after a Notification has been received from
// the OpenZWave library.
type ValueID struct {
	HomeID         uint32
	NodeID         uint8
	Genre          ValueIDGenre
	CommandClassID uint8
	Instance       uint8
	Index          uint8
	Type           ValueIDType
	ID             uint64
}

// buildValueID creates a new valueid.ValueID from the C valueid_t (and
// therefore C++ ValueID) value.
func buildValueID(v C.valueid_t) *ValueID {
	vid := &ValueID{
		HomeID:         uint32(C.valueid_getHomeId(v)),
		NodeID:         uint8(C.valueid_getNodeId(v)),
		CommandClassID: uint8(C.valueid_getCommandClassId(v)),
		Instance:       uint8(C.valueid_getInstance(v)),
		Index:          uint8(C.valueid_getIndex(v)),
		ID:             uint64(C.valueid_getId(v)),
	}

	switch C.valueid_getGenre(v) {
	case C.valueid_genre_basic:
		vid.Genre = ValueIDGenreBasic
	case C.valueid_genre_user:
		vid.Genre = ValueIDGenreUser
	case C.valueid_genre_config:
		vid.Genre = ValueIDGenreConfig
	case C.valueid_genre_system:
		vid.Genre = ValueIDGenreSystem
	case C.valueid_genre_count:
		vid.Genre = ValueIDGenreCount
	}

	switch C.valueid_getType(v) {
	case C.valueid_type_bool:
		vid.Type = ValueIDTypeBool
	case C.valueid_type_byte:
		vid.Type = ValueIDTypeByte
	case C.valueid_type_decimal:
		vid.Type = ValueIDTypeDecimal
	case C.valueid_type_int:
		vid.Type = ValueIDTypeInt
	case C.valueid_type_list:
		vid.Type = ValueIDTypeList
	case C.valueid_type_schedule:
		vid.Type = ValueIDTypeSchedule
	case C.valueid_type_short:
		vid.Type = ValueIDTypeShort
	case C.valueid_type_string:
		vid.Type = ValueIDTypeString
	case C.valueid_type_button:
		vid.Type = ValueIDTypeButton
	case C.valueid_type_raw:
		vid.Type = ValueIDTypeRaw
		// case C.valueid_type_max:
		// 	vid.Type = ValueIDTypeMax
	}

	return vid
}

// IDString will create a string representation of the ID for use as a key.
func (v *ValueID) IDString() string {
	return fmt.Sprintf("%d", v.ID)
}

func (v *ValueID) String() string {
	return fmt.Sprintf("{Label: %q, String: %q, Units: %q, RO: %t, WO: %t, Genre: %s, CommandClassID: %d, Instance: %d, Index: %d, Type: %s, HomeID: %d, ID: %d}",
		v.GetLabel(),
		v.GetAsString(),
		v.GetUnits(),
		v.IsReadOnly(),
		v.IsWriteOnly(),
		v.Genre,
		v.CommandClassID,
		v.Instance,
		v.Index,
		v.Type,
		v.HomeID,
		v.ID)
}

// GetLabel returns the user-friendly label for the value.
func (v *ValueID) GetLabel() string {
	return GetValueLabel(v.HomeID, v.ID)
}

// SetLabel sets the user-friendly label for the value.
func (v *ValueID) SetLabel(label string) {
	SetValueLabel(v.HomeID, v.ID, label)
}

// GetUnits returns the units that the value is measured in.
func (v *ValueID) GetUnits() string {
	return GetValueUnits(v.HomeID, v.ID)
}

// SetUnits sets the units that the value is measured in.
func (v *ValueID) SetUnits(units string) {
	SetValueUnits(v.HomeID, v.ID, units)
}

// GetHelp returns a help string describing the value's purpose and usage.
func (v *ValueID) GetHelp() string {
	return GetValueHelp(v.HomeID, v.ID)
}

// SetHelp sets a help string describing the value's purpose and usage.
func (v *ValueID) SetHelp(help string) {
	SetValueHelp(v.HomeID, v.ID, help)
}

// GetMin returns the minimum that this value may contain.
func (v *ValueID) GetMin() int32 {
	return GetValueMin(v.HomeID, v.ID)
}

// GetMax returns the maximum that this value may contain.
func (v *ValueID) GetMax() int32 {
	return GetValueMax(v.HomeID, v.ID)
}

// IsReadOnly returns true if the value is read-only.
func (v *ValueID) IsReadOnly() bool {
	return IsValueReadOnly(v.HomeID, v.ID)
}

// IsWriteOnly returns true if the value is write-only.
func (v *ValueID) IsWriteOnly() bool {
	return IsValueWriteOnly(v.HomeID, v.ID)
}

// IsSet returns true if the value has been set.
func (v *ValueID) IsSet() bool {
	return IsValueSet(v.HomeID, v.ID)
}

// IsPolled returns true if the value is currently being polled.
func (v *ValueID) IsPolled() bool {
	return IsValuePolled(v.HomeID, v.ID)
}

// GetAsBool returns the value as a bool. It will also return an error if the
// value is not a bool type.
func (v *ValueID) GetAsBool() (bool, error) {
	return GetValueAsBool(v.HomeID, v.ID)
}

// GetAsByte returns the value as an 8-bit unsigned integer. It will also
// return an error if the value is not of byte type.
func (v *ValueID) GetAsByte() (byte, error) {
	return GetValueAsByte(v.HomeID, v.ID)
}

// GetAsFloat returns the value as a float. It will also return an error if
// the value is not a decimal type.
func (v *ValueID) GetAsFloat() (float32, error) {
	return GetValueAsFloat(v.HomeID, v.ID)
}

// GetAsInt returns the value as a 32-bit signed integer. It will also
// return an error if the value is not of 32-bit signed integer type.
func (v *ValueID) GetAsInt() (int32, error) {
	return GetValueAsInt(v.HomeID, v.ID)
}

// GetAsShort returns the value as a 16-bit signed integer. It will also
// return an error if the value is not of 16-bit signed integer type.
func (v *ValueID) GetAsShort() (int16, error) {
	return GetValueAsShort(v.HomeID, v.ID)
}

// GetAsString returns the value as a string, regardless of its actual
// type.
func (v *ValueID) GetAsString() string {
	return GetValueAsString(v.HomeID, v.ID)
}

// GetAsRaw returns the value as a raw byte slice. It will also return an
// error if the value is not of raw type.
func (v *ValueID) GetAsRaw() ([]byte, error) {
	return GetValueAsRaw(v.HomeID, v.ID)
}

// GetListSelectionAsString returns selected item from a list as a string.
// It will also return an error if the value is not of list type.
func (v *ValueID) GetListSelectionAsString() (string, error) {
	return GetValueListSelectionAsString(v.HomeID, v.ID)
}

// GetListSelectionAsInt32 returns selected item from a list as an integer.
// It will also return an error if the value is not of list type.
func (v *ValueID) GetListSelectionAsInt32() (int32, error) {
	return GetValueListSelectionAsInt32(v.HomeID, v.ID)
}

// GetListItems returns the list of items from a list value. It will also
// return an error if the value is not of list type.
func (v *ValueID) GetListItems() ([]string, error) {
	return GetValueListItems(v.HomeID, v.ID)
}

// GetFloatPrecision returns the float value's precision. It will also
// return an error if the value is not of decimal type.
func (v *ValueID) GetFloatPrecision() (uint8, error) {
	return GetValueFloatPrecision(v.HomeID, v.ID)
}

// SetBool sets the state of a bool. It will return an error if the value
// is not of bool type.
//
// Due to the possibility of a device being asleep, the command is assumed to
// succeed, and the value held by the node is updated directly. This will be
// reverted by a future status message from the device if the Z-Wave message
// actually failed to get through. Notification callbacks will be sent in both
// cases.
func (v *ValueID) SetBool(value bool) error {
	return SetValueBool(v.HomeID, v.ID, value)
}

// SetUint8 sets the value of a byte. It will return an error if the value
// is not of byte type.
//
// Due to the possibility of a device being asleep, the command is assumed to
// succeed, and the value held by the node is updated directly. This will be
// reverted by a future status message from the device if the Z-Wave message
// actually failed to get through. Notification callbacks will be sent in both
// cases.
func (v *ValueID) SetUint8(value uint8) error {
	return SetValueUint8(v.HomeID, v.ID, value)
}

// SetFloat sets the value of a decimal. It will return an error if the
// value is not of decimal type.
//
// It is usually better to handle decimal values using strings rather than
// floats, to avoid floating point accuracy issues. Due to the possibility of a
// device being asleep, the command is assumed to succeed, and the value held by
// the node is updated directly. This will be reverted by a future status
// message from the device if the Z-Wave message actually failed to get through.
// Notification callbacks will be sent in both cases.
func (v *ValueID) SetFloat(value float32) error {
	return SetValueFloat(v.HomeID, v.ID, value)
}

// SetInt32 sets the value of a 32-bit signed integer. It will return an
// error if the value is not of 32-bit signed integer type.
//
// Due to the possibility of a device being asleep, the command is assumed to
// succeed, and the value held by the node is updated directly. This will be
// reverted by a future status message from the device if the Z-Wave message
// actually failed to get through. Notification callbacks will be sent in both
// cases.
func (v *ValueID) SetInt32(value int32) error {
	return SetValueInt32(v.HomeID, v.ID, value)
}

// SetInt16 sets the value of a 16-bit signed integer. It will return an
// error if the value is not of 16-bit signed integer type.
//
// Due to the possibility of a device being asleep, the command is assumed to
// succeed, and the value held by the node is updated directly. This will be
// reverted by a future status message from the device if the Z-Wave message
// actually failed to get through. Notification callbacks will be sent in both
// cases.
func (v *ValueID) SetInt16(value int16) error {
	return SetValueInt16(v.HomeID, v.ID, value)
}

// SetBytes sets the value of a raw value. It will return an error if the
// value is not of raw type.
//
// Due to the possibility of a device being asleep, the command is assumed to
// succeed, and the value held by the node is updated directly. This will be
// reverted by a future status message from the device if the Z-Wave message
// actually failed to get through. Notification callbacks will be sent in both
// cases.
func (v *ValueID) SetBytes(value []byte) error {
	return SetValueBytes(v.HomeID, v.ID, value)
}

// SetString sets the value from a string, regardless of type. It will
// return an error if the value could not be parsed into the correct type for
// the value.
//
// Due to the possibility of a device being asleep, the command is assumed to
// succeed, and the value held by the node is updated directly. This will be
// reverted by a future status message from the device if the Z-Wave message
// actually failed to get through. Notification callbacks will be sent in both
// cases.
func (v *ValueID) SetString(value string) error {
	return SetValueString(v.HomeID, v.ID, value)
}

// SetListSelection sets the selected item in a list. It will return an
// error if the value is not of list type or if the selection is not in the
// list.
//
// Due to the possibility of a device being asleep, the command is assumed to
// succeed, and the value held by the node is updated directly. This will be
// reverted by a future status message from the device if the Z-Wave message
// actually failed to get through. Notification callbacks will be sent in both
// cases.
func (v *ValueID) SetListSelection(selectedItem string) error {
	return SetValueListSelection(v.HomeID, v.ID, selectedItem)
}

// Refresh refreshes the specified value from the Z-Wave network. It will
// return true if the driver and node were found, otherwise false.
//
// A call to this function causes the library to send a message to the network
// to retrieve the current value of the specified ValueID (just like a poll,
// except only one-time, not recurring).
func (v *ValueID) Refresh() bool {
	return RefreshValue(v.HomeID, v.ID)
}

// SetChangeVerified sets a flag indicating whether value changes noted upon a
// refresh should be verified. If so, the library will immediately refresh the
// value a second time whenever a change is observed. This helps to filter out
// spurious data reported occasionally by some devices.
func (v *ValueID) SetChangeVerified(verify bool) {
	SetChangeVerified(v.HomeID, v.ID, verify)
}

// GetChangeVerified returns true if value changes upon a refresh should be
// verified. If so, the library will immediately refresh the value a second time
// whenever a change is observed. This helps to filter out spurious data
// reported occasionally by some devices.
func (v *ValueID) GetChangeVerified() bool {
	return GetChangeVerified(v.HomeID, v.ID)
}

// PressButton starts an activity in a device. It will return an error if the
// value is not of button type.
//
// Since buttons are write-only values that do not report a state, no
// notification callbacks are sent.
func (v *ValueID) PressButton() error {
	return PressButton(v.HomeID, v.ID)
}

// ReleaseButton stops an activity in a device. It will return an error if the
// value is not of button type.
//
// Since buttons are write-only values that do not report a state, no
// notification callbacks are sent.
func (v *ValueID) ReleaseButton() error {
	return ReleaseButton(v.HomeID, v.ID)
}

// EnablePoll enables the polling of a device's state. Returns true if polling
// was enabled.
func (v *ValueID) EnablePoll(intensity uint8) bool {
	return EnablePoll(v.HomeID, v.ID, intensity)
}

// DisablePoll disables the polling of a device's state. Returns true if polling
// was disabled.
func (v *ValueID) DisablePoll() bool {
	return DisablePoll(v.HomeID, v.ID)
}

// SetPollIntensity sets the frequency of polling.
//
//  - 0 = none
//  - 1 = every time through the list
//  - 2 = every other time
//  - etc.
func (v *ValueID) SetPollIntensity(intensity uint8) {
	SetPollIntensity(v.HomeID, v.ID, intensity)
}

// GetPollIntensity returns the polling intensity of a device's state.
func (v *ValueID) GetPollIntensity() uint8 {
	return GetPollIntensity(v.HomeID, v.ID)
}
