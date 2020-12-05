package goopenzwave

import "fmt"

//go:generate stringer -trimprefix=ValueGenre_ -type=ValueGenre
//go:generate stringer -trimprefix=ValueType_ -type=ValueType

type ValueGenre uint8

const (
	ValueGenre_Basic  ValueGenre = iota // The 'level' as controlled by basic commands. Usually duplicated by another command class.
	ValueGenre_User                     // Basic values an ordinary user would be interested in.
	ValueGenre_Config                   // Device-specific configuration parameters. These cannot be automatically discovered via Z-Wave, and are usually described in the user manual instead.
	ValueGenre_System                   // Values of significance only to users who understand the Z-Wave protocol
)

type ValueType uint8

const (
	ValueType_Bool     ValueType = iota // Boolean, true or false
	ValueType_Byte                      // 8-bit unsigned value
	ValueType_Decimal                   // Represents a non-integer value as a string, to avoid floating point accuracy issues.
	ValueType_Int                       // 32-bit signed value
	ValueType_List                      // List from which one item can be selected
	ValueType_Schedule                  // Complex type used with the Climate Control Schedule command class
	ValueType_Short                     // 16-bit signed value
	ValueType_String                    // Text string
	ValueType_Button                    // A write-only value that is the equivalent of pressing a button to send a command to a device
	ValueType_Raw                       // A collection of bytes
	ValueType_BitSet                    // A collection of bits
)

type ValueID struct {
	homeid uint32
	id0    uint32
	id1    uint32
}

// ID0 Packing:
// Bits
// 24-31:	8 bits. Node ID of device
// 22-23:	2 bits. genre of value (see ValueGenre enum).
// 14-21:	8 bits. ID of command class that created and manages this value.
// 12-13:	Unused.
// 04-11:	8 bits. Instance of the Value
// 00-03:	4 bits. Type of value (bool, byte, string etc).

// ID1 Packing:
// Bits
// 16-31:	16 bits. Instance Index of the command class.

// HomeID of the driver that controls the node containing the value.
func (v ValueID) HomeID() uint32 {
	return v.homeid
}

// NodeID of the device reporting the value.
func (v ValueID) NodeID() uint8 {
	return uint8((v.id0 & 0xff000000) >> 24)
}

// Genre is the classification of the value to enable low level system or configuration parameters to be filtered out.
func (v ValueID) Genre() ValueGenre {
	return ValueGenre((v.id0 & 0x00c00000) >> 22)
}

// CommandClassID is the Z-Wave command class that created and manages this value. Knowledge of command classes is not required to use OpenZWave, but this information is exposed in case it is of interest.
func (v ValueID) CommandClassID() uint8 {
	return uint8((v.id0 & 0x003fc000) >> 14)
}

// Instance of this value. It is possible for there to be multiple instances of a command class, although currently it appears that only the SensorMultilevel command class ever does this. Knowledge of instances and command classes is not required to use OpenZWave, but this information is exposed in case it is of interest.
func (v ValueID) Instance() uint8 {
	return uint8((v.id0 & 0xff0) >> 4)
}

// Index is used to identify one of multiple values created and managed by a command class. In the case of configurable parameters (handled by the configuration command class), the index is the same as the parameter ID. Knowledge of command classes is not required to use OpenZWave, but this information is exposed in case it is of interest.
func (v ValueID) Index() uint16 {
	return uint16((v.id1 & 0xffff0000) >> 16)
}

// Type describes the data held by the value and enables the user to select the correct value accessor method in the Manager class.
func (v ValueID) Type() ValueType {
	return ValueType(v.id0 & 0x0000000f)
}

func (v ValueID) String() string {
	return fmt.Sprintf("{HomeID:%d, NodeID:%d, Genre:%s, CC:%d, Instance:%d, Idx:%d, Type:%s}", v.HomeID(), v.NodeID(), v.Genre(), v.CommandClassID(), v.Instance(), v.Index(), v.Type())
}

// Equal reports whether v and o represent the same openzwave value.
func (v ValueID) Equal(o ValueID) bool {
	return (v.homeid == o.homeid) && (v.id0 == o.id0) && (v.id1 == o.id1)
}

// SortIndex returns a deterministic index based on the ValueID command class,
// instance and index.
func (v ValueID) SortIndex() uint32 {
	// 8 bits CC ID
	// 8 bits Instance
	// 16 bits Instance Index
	return (uint32(v.CommandClassID())<<24)&0xff000000 | (uint32(v.Instance())<<16)&0x00ff0000 | uint32(v.Index())&0xffff0000
}
