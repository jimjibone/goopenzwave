package gozwave

import (
	"fmt"
)

type Value struct {
	ValueID       *ValueID
	IsReadOnly    bool
	IsWriteOnly   bool
	IsSet         bool
	IsPolled      bool
	Label         string
	Units         string
	Help          string
	PollIntensity uint8
	Min           int32
	Max           int32
	AsString      string
}

func buildValue(valueid *ValueID) *Value {
	manager := GetManager()
	value := &Value{
		ValueID:       valueid,
		IsReadOnly:    manager.IsValueReadOnly(valueid),
		IsWriteOnly:   manager.IsValueWriteOnly(valueid),
		IsSet:         manager.IsValueSet(valueid),
		IsPolled:      manager.IsValuePolled(valueid),
		Label:         manager.GetValueLabel(valueid),
		Units:         manager.GetValueUnits(valueid),
		Help:          manager.GetValueHelp(valueid),
		PollIntensity: manager.GetPollIntensity(valueid),
		Min:           manager.GetValueMin(valueid),
		Max:           manager.GetValueMax(valueid),
	}
	_, value.AsString = manager.GetValueAsString(valueid)
	return value
}

func (v *Value) String() string {
	return fmt.Sprintf("Value{ValueID: %s, IsReadOnly: %t, IsWriteOnly: %t, IsSet: %t, IsPolled: %t, Label: %q, Units: %q, Help: %q, PollIntensity: %d, Min: %d, Max: %d, AsString: %q}",
		v.ValueID, v.IsReadOnly, v.IsWriteOnly, v.IsSet, v.IsPolled, v.Label, v.Units, v.Help, v.PollIntensity, v.Min, v.Max, v.AsString)
}
