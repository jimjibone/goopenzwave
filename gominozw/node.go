package main

import (
	"fmt"

	zwave "github.com/jimjibone/goopenzwave"
)

type Node struct {
	HomeID  uint32
	NodeID  uint8
	Polled  bool
	Removed bool

	Name             string
	ProductName      string
	ManufacturerName string

	ValueIDs []zwave.ValueID
}

func (n Node) String() string {
	var name, product, manufacturer string
	if n.Name != "" {
		name = ", name: " + n.Name
	}
	if n.ProductName != "" {
		product = ", product: " + n.ProductName
	}
	if n.ManufacturerName != "" {
		manufacturer = ", manufacturer: " + n.ManufacturerName
	}
	return fmt.Sprintf("{homeid: %d, nodeid: %d, polled: %t, removed: %t%s%s%s, valueids: %d}",
		n.HomeID,
		n.NodeID,
		n.Polled,
		n.Removed,
		name,
		product,
		manufacturer,
		len(n.ValueIDs),
	)
}

type NodeValue struct {
	ValueID zwave.ValueID
	Label   string
	Value   string
	Units   string
	List    []string
	Help    string
}

func (n NodeValue) String() string {
	return fmt.Sprintf("{id: %s, label: %s, value: %s, units: %s, list: %q, help: %s}", n.ValueID, n.Label, n.Value, n.Units, n.List, n.Help)
}
