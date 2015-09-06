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

// ValueID contains all information available for a ValueID from the OpenZWave
// library. You should not create a new ValueID manually, but receive it from
// the gozwave package after a Notification has been received from the OpenZWave
// library.
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

type ValueIDStringID string

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

func (v *ValueID) StringID() ValueIDStringID {
	return ValueIDStringID(fmt.Sprintf("%d", v.ID))
}

func (v *ValueID) toC() C.valueid_t {
	return C.valueid_create(C.uint32_t(v.HomeID), C.uint64_t(v.ID))
}

func (v *ValueID) String() string {
	return fmt.Sprintf("ValueID{HomeID: %d, NodeID: %d, Genre: %s, CommandClassID: %d, Instance: %d, Index: %d, Type: %s, ID: %d}",
		v.HomeID, v.NodeID, v.Genre, v.CommandClassID, v.Instance, v.Index, v.Type, v.ID)
}

func (v *ValueID) InfoString() string {
	manager := GetManager()
	_, valueString := manager.GetValueAsString(v)
	return fmt.Sprintf("Value{ValueID: %d, NodeID: %d, Genre: %s, "+
		"CommandClassID: %d, Type: %d, RO: %t, WO: %t, "+
		"Set: %t, Polled: %t, Label: %q, Units: %q, Help: %q, "+
		"Min: %d, Max: %d, AsString: %q}",
		v.ID, v.NodeID, v.Genre, v.CommandClassID, v.Type,
		manager.IsValueReadOnly(v), manager.IsValueWriteOnly(v),
		manager.IsValueSet(v), manager.IsValuePolled(v),
		manager.GetValueLabel(v), manager.GetValueUnits(v),
		manager.GetValueHelp(v), manager.GetValueMin(v), manager.GetValueMax(v),
		valueString)
}

func (v *ValueID) GetValueLabel(valueid *ValueID) string {
	manager := GetManager()
	return manager.GetValueLabel(valueid)
}

func (v *ValueID) SetValueLabel(valueid *ValueID, label string) {
	manager := GetManager()
	manager.SetValueLabel(valueid, label)
}

func (v *ValueID) GetValueUnits(valueid *ValueID) string {
	manager := GetManager()
	return manager.GetValueUnits(valueid)
}

func (v *ValueID) SetValueUnits(valueid *ValueID, units string) {
	manager := GetManager()
	manager.SetValueUnits(valueid, units)
}

func (v *ValueID) GetValueHelp(valueid *ValueID) string {
	manager := GetManager()
	return manager.GetValueHelp(valueid)
}

func (v *ValueID) SetValueHelp(valueid *ValueID, help string) {
	manager := GetManager()
	manager.SetValueHelp(valueid, help)
}

func (v *ValueID) GetValueMin(valueid *ValueID) int32 {
	manager := GetManager()
	return manager.GetValueMin(valueid)
}

func (v *ValueID) GetValueMax(valueid *ValueID) int32 {
	manager := GetManager()
	return manager.GetValueMax(valueid)
}

func (v *ValueID) IsValueReadOnly(valueid *ValueID) bool {
	manager := GetManager()
	return manager.IsValueReadOnly(valueid)
}

func (v *ValueID) IsValueWriteOnly(valueid *ValueID) bool {
	manager := GetManager()
	return manager.IsValueWriteOnly(valueid)
}

func (v *ValueID) IsValueSet(valueid *ValueID) bool {
	manager := GetManager()
	return manager.IsValueSet(valueid)
}

func (v *ValueID) IsValuePolled(valueid *ValueID) bool {
	manager := GetManager()
	return manager.IsValuePolled(valueid)
}

func (v *ValueID) GetValueAsBool(valueid *ValueID) (bool, bool) {
	manager := GetManager()
	return manager.GetValueAsBool(valueid)
}

func (v *ValueID) GetValueAsByte(valueid *ValueID) (bool, byte) {
	manager := GetManager()
	return manager.GetValueAsByte(valueid)
}

func (v *ValueID) GetValueAsFloat(valueid *ValueID) (bool, float32) {
	manager := GetManager()
	return manager.GetValueAsFloat(valueid)
}

func (v *ValueID) GetValueAsInt(valueid *ValueID) (bool, int32) {
	manager := GetManager()
	return manager.GetValueAsInt(valueid)
}

func (v *ValueID) GetValueAsShort(valueid *ValueID) (bool, int16) {
	manager := GetManager()
	return manager.GetValueAsShort(valueid)
}

func (v *ValueID) GetValueAsString(valueid *ValueID) (bool, string) {
	manager := GetManager()
	return manager.GetValueAsString(valueid)
}

func (v *ValueID) GetValueAsRaw(valueid *ValueID) (bool, []byte) {
	manager := GetManager()
	return manager.GetValueAsRaw(valueid)
}

func (v *ValueID) GetValueListSelectionAsString(valueid *ValueID) (bool, string) {
	manager := GetManager()
	return manager.GetValueListSelectionAsString(valueid)
}

func (v *ValueID) GetValueListSelectionAsInt32(valueid *ValueID) (bool, int32) {
	manager := GetManager()
	return manager.GetValueListSelectionAsInt32(valueid)
}

func (v *ValueID) GetValueListItems(valueid *ValueID) (bool, []string) {
	manager := GetManager()
	return manager.GetValueListItems(valueid)
}

func (v *ValueID) GetValueFloatPrecision(valueid *ValueID) (bool, uint8) {
	manager := GetManager()
	return manager.GetValueFloatPrecision(valueid)
}

func (v *ValueID) SetValueBool(valueid *ValueID, value bool) bool {
	manager := GetManager()
	return manager.SetValueBool(valueid, value)
}

func (v *ValueID) SetValueUint8(valueid *ValueID, value uint8) bool {
	manager := GetManager()
	return manager.SetValueUint8(valueid, value)
}

func (v *ValueID) SetValueFloat(valueid *ValueID, value float32) bool {
	manager := GetManager()
	return manager.SetValueFloat(valueid, value)
}

func (v *ValueID) SetValueInt32(valueid *ValueID, value int32) bool {
	manager := GetManager()
	return manager.SetValueInt32(valueid, value)
}

func (v *ValueID) SetValueInt16(valueid *ValueID, value int16) bool {
	manager := GetManager()
	return manager.SetValueInt16(valueid, value)
}

func (v *ValueID) SetValueBytes(valueid *ValueID, value []byte) bool {
	manager := GetManager()
	return manager.SetValueBytes(valueid, value)
}

func (v *ValueID) SetValueString(valueid *ValueID, value string) bool {
	manager := GetManager()
	return manager.SetValueString(valueid, value)
}

func (v *ValueID) SetValueListSelection(valueid *ValueID, selectedItem string) bool {
	manager := GetManager()
	return manager.SetValueListSelection(valueid, selectedItem)
}

func (v *ValueID) RefreshValue(valueid *ValueID) bool {
	manager := GetManager()
	return manager.RefreshValue(valueid)
}

func (v *ValueID) SetChangeVerified(valueid *ValueID, verify bool) {
	manager := GetManager()
	manager.SetChangeVerified(valueid, verify)
}

func (v *ValueID) GetChangeVerified(valueid *ValueID) bool {
	manager := GetManager()
	return manager.GetChangeVerified(valueid)
}

func (v *ValueID) PressButton(valueid *ValueID) bool {
	manager := GetManager()
	return manager.PressButton(valueid)
}

func (v *ValueID) ReleaseButton(valueid *ValueID) bool {
	manager := GetManager()
	return manager.ReleaseButton(valueid)
}
