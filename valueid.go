package gozwave

// #cgo LDFLAGS: -lopenzwave -L/usr/local/lib
// #cgo CPPFLAGS: -I/usr/local/include -I/usr/local/include/openzwave
// #include "valueid.h"
// #include <stdlib.h>
import "C"

// ValueGenre defines a type for the valueid genre enum.
type ValueIdGenre int

// ValueType defines a type for the valueid type enum.
type ValueIdType int

const (
	ValueIdGenreBasic  = ValueIdGenre(C.valueid_genre_basic)
	ValueIdGenreUser   = ValueIdGenre(C.valueid_genre_user)
	ValueIdGenreConfig = ValueIdGenre(C.valueid_genre_config)
	ValueIdGenreSystem = ValueIdGenre(C.valueid_genre_system)
	ValueIdGenreCount  = ValueIdGenre(C.valueid_genre_count)

	ValueIdTypeBool     = ValueIdType(C.valueid_type_bool)
	ValueIdTypeByte     = ValueIdType(C.valueid_type_byte)
	ValueIdTypeDecimal  = ValueIdType(C.valueid_type_decimal)
	ValueIdTypeInt      = ValueIdType(C.valueid_type_int)
	ValueIdTypeList     = ValueIdType(C.valueid_type_list)
	ValueIdTypeSchedule = ValueIdType(C.valueid_type_schedule)
	ValueIdTypeShort    = ValueIdType(C.valueid_type_short)
	ValueIdTypeString   = ValueIdType(C.valueid_type_string)
	ValueIdTypeButton   = ValueIdType(C.valueid_type_button)
	ValueIdTypeRaw      = ValueIdType(C.valueid_type_raw)
	ValueIdTypeMax      = ValueIdType(C.valueid_type_max)
)

// ValueId is a container for the C++ OpenZWave library ValueId class.
type ValueId struct {
	valueid C.valueid_t
}

func (v *ValueId) GetHomeId() uint32 {
	return uint32(C.valueid_getHomeId(v.valueid))
}

func (v *ValueId) GetNodeId() uint8 {
	return uint8(C.valueid_getNodeId(v.valueid))
}

func (v *ValueId) GetGenre() ValueIdGenre {
	return ValueIdGenre(C.valueid_getGenre(v.valueid))
}

func (v *ValueId) GetCommandClassId() uint8 {
	return uint8(C.valueid_getCommandClassId(v.valueid))
}

func (v *ValueId) GetInstance() uint8 {
	return uint8(C.valueid_getInstance(v.valueid))
}

func (v *ValueId) GetIndex() uint8 {
	return uint8(C.valueid_getIndex(v.valueid))
}

func (v *ValueId) GetType() ValueIdType {
	return ValueIdType(C.valueid_getType(v.valueid))
}

func (v *ValueId) GetId() uint64 {
	return uint64(C.valueid_getId(v.valueid))
}
