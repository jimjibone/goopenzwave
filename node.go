package gozwave

import (
	"fmt"
)

// Node contains all information available for a Node from the OpenZWave
// library. Create a new Node by using the `NewNode` function with the HomeID
// and NodeID as supplied from a Notification.
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

func (n *Node) String() string {
	return fmt.Sprintf("Node{HomeID: %d, NodeID: %d, BasicType: %d, "+
		"GenericType: %d, SpecificType: %d, NodeType: %q, "+
		"ManufacturerName: %q, ProductName: %q, NodeName: %q, Location: %q, "+
		"ManufacturerID: %q, ProductType: %q, ProductID: %q}",
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

// NewNode will create a new Node object filled with the data available from the
// Manager based on the homeID and nodeID.
func NewNode(homeID uint32, nodeID uint8) *Node {
	node := &Node{
		HomeID: homeID,
		NodeID: nodeID,
	}
	manager := GetManager()
	node.BasicType = manager.GetNodeBasic(homeID, nodeID)
	node.GenericType = manager.GetNodeGeneric(homeID, nodeID)
	node.SpecificType = manager.GetNodeSpecific(homeID, nodeID)
	node.NodeType = manager.GetNodeType(homeID, nodeID)
	node.ManufacturerName = manager.GetNodeManufacturerName(homeID, nodeID)
	node.ProductName = manager.GetNodeProductName(homeID, nodeID)
	node.NodeName = manager.GetNodeName(homeID, nodeID)
	node.Location = manager.GetNodeLocation(homeID, nodeID)
	node.ManufacturerID = manager.GetNodeManufacturerId(homeID, nodeID)
	node.ProductType = manager.GetNodeProductType(homeID, nodeID)
	node.ProductID = manager.GetNodeProductId(homeID, nodeID)
	return node
}
