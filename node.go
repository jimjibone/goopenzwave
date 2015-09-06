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
	node.ManufacturerID = manager.GetNodeManufacturerID(homeID, nodeID)
	node.ProductType = manager.GetNodeProductType(homeID, nodeID)
	node.ProductID = manager.GetNodeProductID(homeID, nodeID)
	return node
}

// RefeshNodeInfo Trigger the fetching of fixed data about a node. Causes the node's data to be obtained from the Z-Wave network in the same way as if it had just been added. This method would normally be called automatically by OpenZWave, but if you know that a node has been changed, calling this method will force a refresh of all of the data held by the library. This can be especially useful for devices that were asleep when the application was first run. This is the same as the query state starting from the beginning.
func (n *Node) RefeshNodeInfo() bool {
	manager := GetManager()
	return manager.RefreshNodeInfo(n.HomeID, n.NodeID)
}

// RequestNodeState Trigger the fetching of dynamic value data for a node. Causes the node's values to be requested from the Z-Wave network. This is the same as the query state starting from the associations state.
func (n *Node) RequestNodeState() bool {
	manager := GetManager()
	return manager.RequestNodeState(n.HomeID, n.NodeID)
}

// RequestNodeDynamic Trigger the fetching of just the dynamic value data for a node. Causes the node's values to be requested from the Z-Wave network. This is the same as the query state starting from the dynamic state.
func (n *Node) RequestNodeDynamic() bool {
	manager := GetManager()
	return manager.RequestNodeDynamic(n.HomeID, n.NodeID)
}

// IsNodeListeningDevice Get whether the node is a listening device that does not go to sleep.
func (n *Node) IsNodeListeningDevice() bool {
	manager := GetManager()
	return manager.IsNodeListeningDevice(n.HomeID, n.NodeID)
}

// IsNodeFrequentListeningDevice Get whether the node is a frequent listening device that goes to sleep but can be woken up by a beam. Useful to determine node and controller consistency.
func (n *Node) IsNodeFrequentListeningDevice() bool {
	manager := GetManager()
	return manager.IsNodeFrequentListeningDevice(n.HomeID, n.NodeID)
}

// IsNodeBeamingDevice Get whether the node is a beam capable device.
func (n *Node) IsNodeBeamingDevice() bool {
	manager := GetManager()
	return manager.IsNodeBeamingDevice(n.HomeID, n.NodeID)
}

// IsNodeRoutingDevice Get whether the node is a routing device that passes messages to other nodes.
func (n *Node) IsNodeRoutingDevice() bool {
	manager := GetManager()
	return manager.IsNodeRoutingDevice(n.HomeID, n.NodeID)
}

// IsNodeSecurityDevice Get the security attribute for a node. True if node supports security features.
func (n *Node) IsNodeSecurityDevice() bool {
	manager := GetManager()
	return manager.IsNodeSecurityDevice(n.HomeID, n.NodeID)
}

// GetNodeMaxBaudRate Get the maximum baud rate of a node's communications.
func (n *Node) GetNodeMaxBaudRate() uint32 {
	manager := GetManager()
	return manager.GetNodeMaxBaudRate(n.HomeID, n.NodeID)
}

// GetNodeVersion Get the version number of a node.
func (n *Node) GetNodeVersion() uint8 {
	manager := GetManager()
	return manager.GetNodeVersion(n.HomeID, n.NodeID)
}

// GetNodeSecurity Get the security byte of a node.
func (n *Node) GetNodeSecurity() uint8 {
	manager := GetManager()
	return manager.GetNodeSecurity(n.HomeID, n.NodeID)
}

// IsNodeZWavePlus Is this a ZWave+ Supported Node?
func (n *Node) IsNodeZWavePlus() bool {
	manager := GetManager()
	return manager.IsNodeZWavePlus(n.HomeID, n.NodeID)
}

// GetNodeBasic Get the basic type of a node.
func (n *Node) GetNodeBasic() uint8 {
	manager := GetManager()
	return manager.GetNodeBasic(n.HomeID, n.NodeID)
}

// GetNodeGeneric Get the generic type of a node.
func (n *Node) GetNodeGeneric() uint8 {
	manager := GetManager()
	return manager.GetNodeGeneric(n.HomeID, n.NodeID)
}

// GetNodeSpecific Get the specific type of a node.
func (n *Node) GetNodeSpecific() uint8 {
	manager := GetManager()
	return manager.GetNodeSpecific(n.HomeID, n.NodeID)
}

// GetNodeType Get a human-readable label describing the node The label is taken from the Z-Wave specific, generic or basic type, depending on which of those values are specified by the node.
func (n *Node) GetNodeType() string {
	manager := GetManager()
	return manager.GetNodeType(n.HomeID, n.NodeID)
}

// GetNodeNeighbors Get the bitmap of this node's neighbors.
// func (n *Node) GetNodeNeighbors() bool {
// 	manager := GetManager()
// 	return manager.GetNodeNeighbors(n.HomeID, n.NodeID)
// }

// GetNodeManufacturerName Get the manufacturer name of a device The manufacturer name would normally be handled by the Manufacturer Specific commmand class, taking the manufacturer ID reported by the device and using it to look up the name from the manufacturer_specific.xml file in the OpenZWave config folder. However, there are some devices that do not support the command class, so to enable the user to manually set the name, it is stored with the node data and accessed via this method rather than being reported via a command class Value object.
func (n *Node) GetNodeManufacturerName() string {
	manager := GetManager()
	return manager.GetNodeManufacturerName(n.HomeID, n.NodeID)
}

// GetNodeProductName Get the product name of a device The product name would normally be handled by the Manufacturer Specific commmand class, taking the product Type and ID reported by the device and using it to look up the name from the manufacturer_specific.xml file in the OpenZWave config folder. However, there are some devices that do not support the command class, so to enable the user to manually set the name, it is stored with the node data and accessed via this method rather than being reported via a command class Value object.
func (n *Node) GetNodeProductName() string {
	manager := GetManager()
	return manager.GetNodeProductName(n.HomeID, n.NodeID)
}

// GetNodeName Get the name of a node The node name is a user-editable label for the node that would normally be handled by the Node Naming commmand class, but many devices do not support it. So that a node can always be named, OpenZWave stores it with the node data, and provides access through this method and SetNodeName, rather than reporting it via a command class Value object. The maximum length of a node name is 16 characters.
func (n *Node) GetNodeName() string {
	manager := GetManager()
	return manager.GetNodeName(n.HomeID, n.NodeID)
}

// GetNodeLocation Get the location of a node The node location is a user-editable string that would normally be handled by the Node Naming commmand class, but many devices do not support it. So that a node can always report its location, OpenZWave stores it with the node data, and provides access through this method and SetNodeLocation, rather than reporting it via a command class Value object.
func (n *Node) GetNodeLocation() string {
	manager := GetManager()
	return manager.GetNodeLocation(n.HomeID, n.NodeID)
}

// GetNodeManufacturerID Get the manufacturer ID of a device The manufacturer ID is a four digit hex code and would normally be handled by the Manufacturer Specific commmand class, but not all devices support it. Although the value reported by this method will be an empty string if the command class is not supported and cannot be set by the user, the manufacturer ID is still stored with the node data (rather than being reported via a command class Value object) to retain a consistent approach with the other manufacturer specific data.
func (n *Node) GetNodeManufacturerID() string {
	manager := GetManager()
	return manager.GetNodeManufacturerID(n.HomeID, n.NodeID)
}

// GetNodeProductType Get the product type of a device The product type is a four digit hex code and would normally be handled by the Manufacturer Specific commmand class, but not all devices support it. Although the value reported by this method will be an empty string if the command class is not supported and cannot be set by the user, the product type is still stored with the node data (rather than being reported via a command class Value object) to retain a consistent approach with the other manufacturer specific data.
func (n *Node) GetNodeProductType() string {
	manager := GetManager()
	return manager.GetNodeProductType(n.HomeID, n.NodeID)
}

// GetNodeProductID Get the product ID of a device The product ID is a four digit hex code and would normally be handled by the Manufacturer Specific commmand class, but not all devices support it. Although the value reported by this method will be an empty string if the command class is not supported and cannot be set by the user, the product ID is still stored with the node data (rather than being reported via a command class Value object) to retain a consistent approach with the other manufacturer specific data.
func (n *Node) GetNodeProductID() string {
	manager := GetManager()
	return manager.GetNodeProductID(n.HomeID, n.NodeID)
}

// SetNodeManufacturerName Set the manufacturer name of a device The manufacturer name would normally be handled by the Manufacturer Specific commmand class, taking the manufacturer ID reported by the device and using it to look up the name from the manufacturer_specific.xml file in the OpenZWave config folder. However, there are some devices that do not support the command class, so to enable the user to manually set the name, it is stored with the node data and accessed via this method rather than being reported via a command class Value object.
func (n *Node) SetNodeManufacturerName(name string) {
	manager := GetManager()
	manager.SetNodeManufacturerName(n.HomeID, n.NodeID, name)
}

// SetNodeProductName Set the product name of a device The product name would normally be handled by the Manufacturer Specific commmand class, taking the product Type and ID reported by the device and using it to look up the name from the manufacturer_specific.xml file in the OpenZWave config folder. However, there are some devices that do not support the command class, so to enable the user to manually set the name, it is stored with the node data and accessed via this method rather than being reported via a command class Value object.
func (n *Node) SetNodeProductName(name string) {
	manager := GetManager()
	manager.SetNodeProductName(n.HomeID, n.NodeID, name)
}

// SetNodeName Set the name of a node The node name is a user-editable label for the node that would normally be handled by the Node Naming commmand class, but many devices do not support it. So that a node can always be named, OpenZWave stores it with the node data, and provides access through this method and GetNodeName, rather than reporting it via a command class Value object. If the device does support the Node Naming command class, the new name will be sent to the node. The maximum length of a node name is 16 characters.
func (n *Node) SetNodeName(name string) {
	manager := GetManager()
	manager.SetNodeName(n.HomeID, n.NodeID, name)
}

// SetNodeLocation Set the location of a node The node location is a user-editable string that would normally be handled by the Node Naming commmand class, but many devices do not support it. So that a node can always report its location, OpenZWave stores it with the node data, and provides access through this method and GetNodeLocation, rather than reporting it via a command class Value object. If the device does support the Node Naming command class, the new location will be sent to the node.
func (n *Node) SetNodeLocation(location string) {
	manager := GetManager()
	manager.SetNodeLocation(n.HomeID, n.NodeID, location)
}

// SetNodeOn Turns a node on This is a helper method to simplify basic control of a node. It is the equivalent of changing the level reported by the node's Basic command class to 255, and will generate a ValueChanged notification from that class. This command will turn on the device at its last known level, if supported by the device, otherwise it will turn it on at 100%.
func (n *Node) SetNodeOn() {
	manager := GetManager()
	manager.SetNodeOn(n.HomeID, n.NodeID)
}

// SetNodeOff Turns a node off This is a helper method to simplify basic control of a node. It is the equivalent of changing the level reported by the node's Basic command class to zero, and will generate a ValueChanged notification from that class.
func (n *Node) SetNodeOff() {
	manager := GetManager()
	manager.SetNodeOff(n.HomeID, n.NodeID)
}

// SetNodeLevel Sets the basic level of a node This is a helper method to simplify basic control of a node. It is the equivalent of changing the value reported by the node's Basic command class and will generate a ValueChanged notification from that class.
func (n *Node) SetNodeLevel(level uint8) {
	manager := GetManager()
	manager.SetNodeLevel(n.HomeID, n.NodeID, level)
}

// IsNodeInfoReceived Get whether the node information has been received.
func (n *Node) IsNodeInfoReceived() bool {
	manager := GetManager()
	return manager.IsNodeInfoReceived(n.HomeID, n.NodeID)
}

// GetNodeClassInformation Get whether the node has the defined class available or not.
func (n *Node) GetNodeClassInformation(commandClassID uint8) (bool, string, uint8) {
	manager := GetManager()
	return manager.GetNodeClassInformation(n.HomeID, n.NodeID, commandClassID)
}

// IsNodeAwake Get whether the node is awake or asleep.
func (n *Node) IsNodeAwake() bool {
	manager := GetManager()
	return manager.IsNodeAwake(n.HomeID, n.NodeID)
}

// IsNodeFailed Get whether the node is working or has failed.
func (n *Node) IsNodeFailed() bool {
	manager := GetManager()
	return manager.IsNodeFailed(n.HomeID, n.NodeID)
}

// GetNodeQueryStage Get whether the node's query stage as a string.
func (n *Node) GetNodeQueryStage() string {
	manager := GetManager()
	return manager.GetNodeQueryStage(n.HomeID, n.NodeID)
}

// GetNodeDeviceType Get the node device type as reported in the Z-Wave+ Info report.
func (n *Node) GetNodeDeviceType() uint8 {
	manager := GetManager()
	return manager.GetNodeDeviceType(n.HomeID, n.NodeID)
}

// GetNodeDeviceTypeString Get the node device type as reported in the Z-Wave+ Info report.
func (n *Node) GetNodeDeviceTypeString() string {
	manager := GetManager()
	return manager.GetNodeDeviceTypeString(n.HomeID, n.NodeID)
}

// GetNodeRole Get the node device type as reported in the Z-Wave+ Info report.
func (n *Node) GetNodeRole() uint8 {
	manager := GetManager()
	return manager.GetNodeRole(n.HomeID, n.NodeID)
}

// GetNodeRoleString Get the node role as reported in the Z-Wave+ Info report.
func (n *Node) GetNodeRoleString() string {
	manager := GetManager()
	return manager.GetNodeRoleString(n.HomeID, n.NodeID)
}

// GetNodePlusType Get the node PlusType as reported in the Z-Wave+ Info report.
func (n *Node) GetNodePlusType() uint8 {
	manager := GetManager()
	return manager.GetNodePlusType(n.HomeID, n.NodeID)
}

// GetNodePlusTypeString Get the node PlusType as reported in the Z-Wave+ Info report.
func (n *Node) GetNodePlusTypeString() string {
	manager := GetManager()
	return manager.GetNodePlusTypeString(n.HomeID, n.NodeID)
}
