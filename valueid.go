package gozwave

// #cgo LDFLAGS: -lopenzwave -L/usr/local/lib
// #cgo CPPFLAGS: -I/usr/local/include -I/usr/local/include/openzwave
// #include "valueid.h"
// #include <stdlib.h>
import "C"
import (
	"fmt"
)

// ValueIDGenre defines a type for the valueid genre enum.
type ValueIDGenre int

// ValueIDType defines a type for the valueid type enum.
type ValueIDType int

const (
	ValueIDGenreBasic  = ValueIDGenre(C.valueid_genre_basic)
	ValueIDGenreUser   = ValueIDGenre(C.valueid_genre_user)
	ValueIDGenreConfig = ValueIDGenre(C.valueid_genre_config)
	ValueIDGenreSystem = ValueIDGenre(C.valueid_genre_system)
	ValueIDGenreCount  = ValueIDGenre(C.valueid_genre_count)

	ValueIDTypeBool     = ValueIDType(C.valueid_type_bool)
	ValueIDTypeByte     = ValueIDType(C.valueid_type_byte)
	ValueIDTypeDecimal  = ValueIDType(C.valueid_type_decimal)
	ValueIDTypeInt      = ValueIDType(C.valueid_type_int)
	ValueIDTypeList     = ValueIDType(C.valueid_type_list)
	ValueIDTypeSchedule = ValueIDType(C.valueid_type_schedule)
	ValueIDTypeShort    = ValueIDType(C.valueid_type_short)
	ValueIDTypeString   = ValueIDType(C.valueid_type_string)
	ValueIDTypeButton   = ValueIDType(C.valueid_type_button)
	ValueIDTypeRaw      = ValueIDType(C.valueid_type_raw)
	ValueIDTypeMax      = ValueIDType(C.valueid_type_max)
)

func (v ValueIDGenre) String() string {
	switch v {
	case ValueIDGenreBasic:
		return "GenreBasic"
	case ValueIDGenreUser:
		return "GenreUser"
	case ValueIDGenreConfig:
		return "GenreConfig"
	case ValueIDGenreSystem:
		return "GenreSystem"
	case ValueIDGenreCount:
		return "GenreCount"
	}
	return "UNKNOWN"
}

func (v ValueIDType) String() string {
	switch v {
	case ValueIDTypeBool:
		return "TypeBool"
	case ValueIDTypeByte:
		return "TypeByte"
	case ValueIDTypeDecimal:
		return "TypeDecimal"
	case ValueIDTypeInt:
		return "TypeInt"
	case ValueIDTypeList:
		return "TypeList"
	case ValueIDTypeSchedule:
		return "TypeSchedule"
	case ValueIDTypeShort:
		return "TypeShort"
	case ValueIDTypeString:
		return "TypeString"
	case ValueIDTypeButton:
		return "TypeButton"
	case ValueIDTypeRaw: // also ValueIDTypeMax
		return "Type(Raw|Max)"
	}
	return "UNKNOWN"
}

// ValueIDStringID simply a string representing the ValueID's ID.
type ValueIDStringID string

// ValueID contains all appropriate information available for a ValueID from the
// OpenZWave library. You should not create a new ValueID manually, but receive
// it from the gozwave package after a Notification has been received from the
// OpenZWave library.
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

func buildValueID(v C.valueid_t) *ValueID {
	valueid := &ValueID{
		HomeID:         uint32(C.valueid_getHomeId(v)),
		NodeID:         uint8(C.valueid_getNodeId(v)),
		Genre:          ValueIDGenre(C.valueid_getGenre(v)),
		CommandClassID: uint8(C.valueid_getCommandClassId(v)),
		Instance:       uint8(C.valueid_getInstance(v)),
		Index:          uint8(C.valueid_getIndex(v)),
		Type:           ValueIDType(C.valueid_getType(v)),
		ID:             uint64(C.valueid_getId(v)),
	}
	return valueid
}

func (v *ValueID) toC() C.valueid_t {
	return C.valueid_create(C.uint32_t(v.HomeID), C.uint64_t(v.ID))
}

// StringID will create a string representation of the ID for use as a key.
func (v *ValueID) StringID() ValueIDStringID {
	return ValueIDStringID(fmt.Sprintf("%d", v.ID))
}

func (v *ValueID) String() string {
	manager := GetManager()
	_, valueString := manager.GetValueAsString(v)
	return fmt.Sprintf("ValueID{Label: %q, String: %q, Units: %q, RO: %t, WO: %t, Genre: %s, CommandClassID: %d, Instance: %d, Index: %d, Type: %s, HomeID: %d, ID: %d}",
		manager.GetValueLabel(v),
		valueString,
		manager.GetValueUnits(v),
		manager.IsValueReadOnly(v),
		manager.IsValueWriteOnly(v),
		v.Genre,
		v.CommandClassID,
		v.Instance,
		v.Index,
		v.Type,
		v.HomeID,
		v.ID)
}

// GetLabel Gets the user-friendly label for the value.
func (v *ValueID) GetLabel() string {
	manager := GetManager()
	return manager.GetValueLabel(v)
}

// SetLabel Sets the user-friendly label for the value.
func (v *ValueID) SetLabel(label string) {
	manager := GetManager()
	manager.SetValueLabel(v, label)
}

// GetUnits Gets the units that the value is measured in.
func (v *ValueID) GetUnits() string {
	manager := GetManager()
	return manager.GetValueUnits(v)
}

// SetUnits Sets the units that the value is measured in.
func (v *ValueID) SetUnits(units string) {
	manager := GetManager()
	manager.SetValueUnits(v, units)
}

// GetHelp Gets a help string describing the value's purpose and usage.
func (v *ValueID) GetHelp() string {
	manager := GetManager()
	return manager.GetValueHelp(v)
}

// SetHelp Sets a help string describing the value's purpose and usage.
func (v *ValueID) SetHelp(help string) {
	manager := GetManager()
	manager.SetValueHelp(v, help)
}

// GetMin Gets the minimum that this value may contain.
func (v *ValueID) GetMin() int32 {
	manager := GetManager()
	return manager.GetValueMin(v)
}

// GetMax Gets the maximum that this value may contain.
func (v *ValueID) GetMax() int32 {
	manager := GetManager()
	return manager.GetValueMax(v)
}

// IsReadOnly Test whether the value is read-only.
func (v *ValueID) IsReadOnly() bool {
	manager := GetManager()
	return manager.IsValueReadOnly(v)
}

// IsWriteOnly Test whether the value is write-only.
func (v *ValueID) IsWriteOnly() bool {
	manager := GetManager()
	return manager.IsValueWriteOnly(v)
}

// IsSet Test whether the value has been set.
func (v *ValueID) IsSet() bool {
	manager := GetManager()
	return manager.IsValueSet(v)
}

// IsPolled Test whether the value is currently being polled.
func (v *ValueID) IsPolled() bool {
	manager := GetManager()
	return manager.IsValuePolled(v)
}

// GetAsBool Gets a value as a bool.
func (v *ValueID) GetAsBool() (bool, bool) {
	manager := GetManager()
	return manager.GetValueAsBool(v)
}

// GetAsByte Gets a value as an 8-bit unsigned integer.
func (v *ValueID) GetAsByte() (bool, byte) {
	manager := GetManager()
	return manager.GetValueAsByte(v)
}

// GetAsFloat Gets a value as a float.
func (v *ValueID) GetAsFloat() (bool, float32) {
	manager := GetManager()
	return manager.GetValueAsFloat(v)
}

// GetAsInt Gets a value as a 32-bit signed integer.
func (v *ValueID) GetAsInt() (bool, int32) {
	manager := GetManager()
	return manager.GetValueAsInt(v)
}

// GetAsShort Gets a value as a 16-bit signed integer.
func (v *ValueID) GetAsShort() (bool, int16) {
	manager := GetManager()
	return manager.GetValueAsShort(v)
}

// GetAsString Gets a value as a string. Creates a string representation of a value, regardless of type.
func (v *ValueID) GetAsString() (bool, string) {
	manager := GetManager()
	return manager.GetValueAsString(v)
}

// GetAsRaw Gets a value as a collection of bytes.
func (v *ValueID) GetAsRaw() (bool, []byte) {
	manager := GetManager()
	return manager.GetValueAsRaw(v)
}

// GetListSelectionAsString Gets the selected item from a list (as a string).
func (v *ValueID) GetListSelectionAsString() (bool, string) {
	manager := GetManager()
	return manager.GetValueListSelectionAsString(v)
}

// GetListSelectionAsInt32 Gets the selected item from a list (as an integer).
func (v *ValueID) GetListSelectionAsInt32() (bool, int32) {
	manager := GetManager()
	return manager.GetValueListSelectionAsInt32(v)
}

// GetListItems Gets the list of items from a list value.
func (v *ValueID) GetListItems() (bool, []string) {
	manager := GetManager()
	return manager.GetValueListItems(v)
}

// GetFloatPrecision Gets a float value's precision.
func (v *ValueID) GetFloatPrecision() (bool, uint8) {
	manager := GetManager()
	return manager.GetValueFloatPrecision(v)
}

// SetBool Sets the state of a bool. Due to the possibility of a device being asleep, the command is assumed to suceed, and the value held by the node is updated directly. This will be reverted by a future status message from the device if the Z-Wave message actually failed to get through. Notification callbacks will be sent in both cases.
func (v *ValueID) SetBool(value bool) bool {
	manager := GetManager()
	return manager.SetValueBool(v, value)
}

// SetUint8 Sets the value of a byte. Due to the possibility of a device being asleep, the command is assumed to suceed, and the value held by the node is updated directly. This will be reverted by a future status message from the device if the Z-Wave message actually failed to get through. Notification callbacks will be sent in both cases.
func (v *ValueID) SetUint8(value uint8) bool {
	manager := GetManager()
	return manager.SetValueUint8(v, value)
}

// SetFloat Sets the value of a decimal. It is usually better to handle decimal values using strings rather than floats, to avoid floating point accuracy issues. Due to the possibility of a device being asleep, the command is assumed to succeed, and the value held by the node is updated directly. This will be reverted by a future status message from the device if the Z-Wave message actually failed to get through. Notification callbacks will be sent in both cases.
func (v *ValueID) SetFloat(value float32) bool {
	manager := GetManager()
	return manager.SetValueFloat(v, value)
}

// SetInt32 Sets the value of a 32-bit signed integer. Due to the possibility of a device being asleep, the command is assumed to suceed, and the value held by the node is updated directly. This will be reverted by a future status message from the device if the Z-Wave message actually failed to get through. Notification callbacks will be sent in both cases.
func (v *ValueID) SetInt32(value int32) bool {
	manager := GetManager()
	return manager.SetValueInt32(v, value)
}

// SetInt16 Sets the value of a 16-bit signed integer. Due to the possibility of a device being asleep, the command is assumed to suceed, and the value held by the node is updated directly. This will be reverted by a future status message from the device if the Z-Wave message actually failed to get through. Notification callbacks will be sent in both cases.
func (v *ValueID) SetInt16(value int16) bool {
	manager := GetManager()
	return manager.SetValueInt16(v, value)
}

// SetBytes Sets the value of a collection of bytes. Due to the possibility of a device being asleep, the command is assumed to suceed, and the value held by the node is updated directly. This will be reverted by a future status message from the device if the Z-Wave message actually failed to get through. Notification callbacks will be sent in both cases.
func (v *ValueID) SetBytes(value []byte) bool {
	manager := GetManager()
	return manager.SetValueBytes(v, value)
}

// SetString Sets the value from a string, regardless of type. Due to the possibility of a device being asleep, the command is assumed to suceed, and the value held by the node is updated directly. This will be reverted by a future status message from the device if the Z-Wave message actually failed to get through. Notification callbacks will be sent in both cases.
func (v *ValueID) SetString(value string) bool {
	manager := GetManager()
	return manager.SetValueString(v, value)
}

// SetListSelection Sets the selected item in a list. Due to the possibility of a device being asleep, the command is assumed to suceed, and the value held by the node is updated directly. This will be reverted by a future status message from the device if the Z-Wave message actually failed to get through. Notification callbacks will be sent in both cases.
func (v *ValueID) SetListSelection(selectedItem string) bool {
	manager := GetManager()
	return manager.SetValueListSelection(v, selectedItem)
}

// Refresh Refreshes the specified value from the Z-Wave network. A call to this function causes the library to send a message to the network to retrieve the current value of the specified ValueID (just like a poll, except only one-time, not recurring).
func (v *ValueID) Refresh() bool {
	manager := GetManager()
	return manager.RefreshValue(v)
}

// SetChangeVerified Sets a flag indicating whether value changes noted upon a refresh should be verified. If so, the library will immediately refresh the value a second time whenever a change is observed. This helps to filter out spurious data reported occasionally by some devices.
func (v *ValueID) SetChangeVerified(verify bool) {
	manager := GetManager()
	manager.SetChangeVerified(v, verify)
}

// GetChangeVerified determine if value changes upon a refresh should be verified. If so, the library will immediately refresh the value a second time whenever a change is observed. This helps to filter out spurious data reported occasionally by some devices.
func (v *ValueID) GetChangeVerified() bool {
	manager := GetManager()
	return manager.GetChangeVerified(v)
}

// PressButton Starts an activity in a device. Since buttons are write-only values that do not report a state, no notification callbacks are sent.
func (v *ValueID) PressButton() bool {
	manager := GetManager()
	return manager.PressButton(v)
}

// ReleaseButton Stops an activity in a device. Since buttons are write-only values that do not report a state, no notification callbacks are sent.
func (v *ValueID) ReleaseButton() bool {
	manager := GetManager()
	return manager.ReleaseButton(v)
}
