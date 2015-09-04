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
