package gozwave

// #cgo LDFLAGS: -lopenzwave -L/usr/local/lib
// #cgo CPPFLAGS: -I/usr/local/include -I/usr/local/include/openzwave
// #include "valueid.h"
// #include <stdlib.h>
import "C"

// ValueGenre defines a type for the valueid genre enum.
type ValueIDGenre int

// ValueType defines a type for the valueid type enum.
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

// ValueID is a container for the C++ OpenZWave library ValueID class.
type ValueID struct {
	valueid C.valueid_t
}

func (v *ValueID) GetHomeId() uint32 {
	return uint32(C.valueid_getHomeId(v.valueid))
}

func (v *ValueID) GetNodeId() uint8 {
	return uint8(C.valueid_getNodeId(v.valueid))
}

func (v *ValueID) GetGenre() ValueIDGenre {
	return ValueIDGenre(C.valueid_getGenre(v.valueid))
}

func (v *ValueID) GetCommandClassId() uint8 {
	return uint8(C.valueid_getCommandClassId(v.valueid))
}

func (v *ValueID) GetInstance() uint8 {
	return uint8(C.valueid_getInstance(v.valueid))
}

func (v *ValueID) GetIndex() uint8 {
	return uint8(C.valueid_getIndex(v.valueid))
}

func (v *ValueID) GetType() ValueIDType {
	return ValueIDType(C.valueid_getType(v.valueid))
}

func (v *ValueID) GetId() uint64 {
	return uint64(C.valueid_getId(v.valueid))
}
