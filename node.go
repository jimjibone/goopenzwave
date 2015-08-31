package gozwave

import (
	"fmt"
)

type Node struct {
	HomeID           uint32
	NodeID           uint8
	BasicType        uint8
	GenericType      uint8
	SpecificType     uint8
	NodeType         string
	ManufacturerName string
	ProductName      string
	NodeName         string
	Location         string
	ManufacturerID   string
	ProductType      string
	ProductID        string
}

func NewNode(homeId uint32, nodeId uint8) *Node {
	node := &Node{
		HomeID: homeId,
		NodeID: nodeId,
	}
	manager := GetManager()
	node.BasicType = manager.GetNodeBasic(homeId, nodeId)
	node.GenericType = manager.GetNodeGeneric(homeId, nodeId)
	node.SpecificType = manager.GetNodeSpecific(homeId, nodeId)
	node.NodeType = manager.GetNodeType(homeId, nodeId)
	node.ManufacturerName = manager.GetNodeManufacturerName(homeId, nodeId)
	node.ProductName = manager.GetNodeProductName(homeId, nodeId)
	node.NodeName = manager.GetNodeName(homeId, nodeId)
	node.Location = manager.GetNodeLocation(homeId, nodeId)
	node.ManufacturerID = manager.GetNodeManufacturerId(homeId, nodeId)
	node.ProductType = manager.GetNodeProductType(homeId, nodeId)
	node.ProductID = manager.GetNodeProductId(homeId, nodeId)
	return node
}

func (n *Node) String() string {
	return fmt.Sprintf("Node{HomeID: %d, NodeID: %d, BasicType: %d, GenericType: %d, SpecificType: %d, NodeType: %q, ManufacturerName: %q, ProductName: %q, NodeName: %q, Location: %q, ManufacturerID: %q, ProductType: %q, ProductID: %q}",
		n.HomeID,
		n.NodeID,
		n.BasicType,
		n.GenericType,
		n.SpecificType,
		n.NodeType,
		n.ManufacturerName,
		n.ProductName,
		n.NodeName,
		n.Location,
		n.ManufacturerID,
		n.ProductType,
		n.ProductID,
	)
}
