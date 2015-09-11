package goopenzwave

import (
	"fmt"
)

// Node contains all necessary information for a Node from the OpenZWave
// library. Create a new Node by using the `NewNode` function with the HomeID
// and NodeID as supplied from a Notification.
type Node struct {
	HomeID uint32
	NodeID uint8
}

// NewNode will create a new Node object filled with the data available from the
// Manager based on the homeID and nodeID.
func NewNode(homeID uint32, nodeID uint8) *Node {
	return &Node{
		HomeID: homeID,
		NodeID: nodeID,
	}
}

// String will return a string containing some useful information about the
// Node.
func (n *Node) String() string {
	manager := GetManager()
	return fmt.Sprintf("Node{HomeID: %d, NodeID: %d, BasicType: %d, "+
		"GenericType: %d, SpecificType: %d, NodeType: %q, "+
		"ManufacturerName: %q, ProductName: %q, NodeName: %q, Location: %q, "+
		"ManufacturerID: %q, ProductType: %q, ProductID: %q}",
		n.HomeID,
		n.NodeID,
		manager.GetNodeBasicType(n.HomeID, n.NodeID),
		manager.GetNodeGenericType(n.HomeID, n.NodeID),
		manager.GetNodeSpecificType(n.HomeID, n.NodeID),
		manager.GetNodeType(n.HomeID, n.NodeID),
		manager.GetNodeManufacturerName(n.HomeID, n.NodeID),
		manager.GetNodeProductName(n.HomeID, n.NodeID),
		manager.GetNodeName(n.HomeID, n.NodeID),
		manager.GetNodeLocation(n.HomeID, n.NodeID),
		manager.GetNodeManufacturerID(n.HomeID, n.NodeID),
		manager.GetNodeProductType(n.HomeID, n.NodeID),
		manager.GetNodeProductID(n.HomeID, n.NodeID),
	)
}

// RefeshInfo Trigger the fetching of fixed data about a node. Causes the node's data to be obtained from the Z-Wave network in the same way as if it had just been added. This method would normally be called automatically by OpenZWave, but if you know that a node has been changed, calling this method will force a refresh of all of the data held by the library. This can be especially useful for devices that were asleep when the application was first run. This is the same as the query state starting from the beginning.
func (n *Node) RefeshInfo() bool {
	manager := GetManager()
	return manager.RefreshNodeInfo(n.HomeID, n.NodeID)
}

// RequestState Trigger the fetching of dynamic value data for a node. Causes the node's values to be requested from the Z-Wave network. This is the same as the query state starting from the associations state.
func (n *Node) RequestState() bool {
	manager := GetManager()
	return manager.RequestNodeState(n.HomeID, n.NodeID)
}

// RequestDynamic Trigger the fetching of just the dynamic value data for a node. Causes the node's values to be requested from the Z-Wave network. This is the same as the query state starting from the dynamic state.
func (n *Node) RequestDynamic() bool {
	manager := GetManager()
	return manager.RequestNodeDynamic(n.HomeID, n.NodeID)
}

// IsListeningDevice Get whether the node is a listening device that does not go to sleep.
func (n *Node) IsListeningDevice() bool {
	manager := GetManager()
	return manager.IsNodeListeningDevice(n.HomeID, n.NodeID)
}

// IsFrequentListeningDevice Get whether the node is a frequent listening device that goes to sleep but can be woken up by a beam. Useful to determine node and controller consistency.
func (n *Node) IsFrequentListeningDevice() bool {
	manager := GetManager()
	return manager.IsNodeFrequentListeningDevice(n.HomeID, n.NodeID)
}

// IsBeamingDevice Get whether the node is a beam capable device.
func (n *Node) IsBeamingDevice() bool {
	manager := GetManager()
	return manager.IsNodeBeamingDevice(n.HomeID, n.NodeID)
}

// IsRoutingDevice Get whether the node is a routing device that passes messages to other nodes.
func (n *Node) IsRoutingDevice() bool {
	manager := GetManager()
	return manager.IsNodeRoutingDevice(n.HomeID, n.NodeID)
}

// IsSecurityDevice Get the security attribute for a node. True if node supports security features.
func (n *Node) IsSecurityDevice() bool {
	manager := GetManager()
	return manager.IsNodeSecurityDevice(n.HomeID, n.NodeID)
}

// GetMaxBaudRate Get the maximum baud rate of a node's communications.
func (n *Node) GetMaxBaudRate() uint32 {
	manager := GetManager()
	return manager.GetNodeMaxBaudRate(n.HomeID, n.NodeID)
}

// GetVersion Get the version number of a node.
func (n *Node) GetVersion() uint8 {
	manager := GetManager()
	return manager.GetNodeVersion(n.HomeID, n.NodeID)
}

// GetSecurity Get the security byte of a node.
func (n *Node) GetSecurity() uint8 {
	manager := GetManager()
	return manager.GetNodeSecurity(n.HomeID, n.NodeID)
}

// IsZWavePlus Is this a ZWave+ Supported Node?
func (n *Node) IsZWavePlus() bool {
	manager := GetManager()
	return manager.IsNodeZWavePlus(n.HomeID, n.NodeID)
}

// GetBasicType Get the basic type of a node.
func (n *Node) GetBasicType() uint8 {
	manager := GetManager()
	return manager.GetNodeBasicType(n.HomeID, n.NodeID)
}

// GetGenericType Get the generic type of a node.
func (n *Node) GetGenericType() uint8 {
	manager := GetManager()
	return manager.GetNodeGenericType(n.HomeID, n.NodeID)
}

// GetSpecificType Get the specific type of a node.
func (n *Node) GetSpecificType() uint8 {
	manager := GetManager()
	return manager.GetNodeSpecificType(n.HomeID, n.NodeID)
}

// GetType Get a human-readable label describing the node The label is taken from the Z-Wave specific, generic or basic type, depending on which of those values are specified by the node.
func (n *Node) GetType() string {
	manager := GetManager()
	return manager.GetNodeType(n.HomeID, n.NodeID)
}

// TODO: implement node.GetNeighbors
// GetNeighbors Get the bitmap of this node's neighbors.
// func (n *Node) GetNeighbors() bool {
// 	manager := GetManager()
// 	return manager.GetNodeNeighbors(n.HomeID, n.NodeID)
// }

// GetManufacturerName Get the manufacturer name of a device The manufacturer name would normally be handled by the Manufacturer Specific commmand class, taking the manufacturer ID reported by the device and using it to look up the name from the manufacturer_specific.xml file in the OpenZWave config folder. However, there are some devices that do not support the command class, so to enable the user to manually set the name, it is stored with the node data and accessed via this method rather than being reported via a command class Value object.
func (n *Node) GetManufacturerName() string {
	manager := GetManager()
	return manager.GetNodeManufacturerName(n.HomeID, n.NodeID)
}

// GetProductName Get the product name of a device The product name would normally be handled by the Manufacturer Specific commmand class, taking the product Type and ID reported by the device and using it to look up the name from the manufacturer_specific.xml file in the OpenZWave config folder. However, there are some devices that do not support the command class, so to enable the user to manually set the name, it is stored with the node data and accessed via this method rather than being reported via a command class Value object.
func (n *Node) GetProductName() string {
	manager := GetManager()
	return manager.GetNodeProductName(n.HomeID, n.NodeID)
}

// GetName Get the name of a node The node name is a user-editable label for the node that would normally be handled by the Node Naming commmand class, but many devices do not support it. So that a node can always be named, OpenZWave stores it with the node data, and provides access through this method and SetNodeName, rather than reporting it via a command class Value object. The maximum length of a node name is 16 characters.
func (n *Node) GetName() string {
	manager := GetManager()
	return manager.GetNodeName(n.HomeID, n.NodeID)
}

// GetLocation Get the location of a node The node location is a user-editable string that would normally be handled by the Node Naming commmand class, but many devices do not support it. So that a node can always report its location, OpenZWave stores it with the node data, and provides access through this method and SetNodeLocation, rather than reporting it via a command class Value object.
func (n *Node) GetLocation() string {
	manager := GetManager()
	return manager.GetNodeLocation(n.HomeID, n.NodeID)
}

// GetManufacturerID Get the manufacturer ID of a device The manufacturer ID is a four digit hex code and would normally be handled by the Manufacturer Specific commmand class, but not all devices support it. Although the value reported by this method will be an empty string if the command class is not supported and cannot be set by the user, the manufacturer ID is still stored with the node data (rather than being reported via a command class Value object) to retain a consistent approach with the other manufacturer specific data.
func (n *Node) GetManufacturerID() string {
	manager := GetManager()
	return manager.GetNodeManufacturerID(n.HomeID, n.NodeID)
}

// GetProductType Get the product type of a device The product type is a four digit hex code and would normally be handled by the Manufacturer Specific commmand class, but not all devices support it. Although the value reported by this method will be an empty string if the command class is not supported and cannot be set by the user, the product type is still stored with the node data (rather than being reported via a command class Value object) to retain a consistent approach with the other manufacturer specific data.
func (n *Node) GetProductType() string {
	manager := GetManager()
	return manager.GetNodeProductType(n.HomeID, n.NodeID)
}

// GetProductID Get the product ID of a device The product ID is a four digit hex code and would normally be handled by the Manufacturer Specific commmand class, but not all devices support it. Although the value reported by this method will be an empty string if the command class is not supported and cannot be set by the user, the product ID is still stored with the node data (rather than being reported via a command class Value object) to retain a consistent approach with the other manufacturer specific data.
func (n *Node) GetProductID() string {
	manager := GetManager()
	return manager.GetNodeProductID(n.HomeID, n.NodeID)
}

// SetManufacturerName Set the manufacturer name of a device The manufacturer name would normally be handled by the Manufacturer Specific commmand class, taking the manufacturer ID reported by the device and using it to look up the name from the manufacturer_specific.xml file in the OpenZWave config folder. However, there are some devices that do not support the command class, so to enable the user to manually set the name, it is stored with the node data and accessed via this method rather than being reported via a command class Value object.
func (n *Node) SetManufacturerName(name string) {
	manager := GetManager()
	manager.SetNodeManufacturerName(n.HomeID, n.NodeID, name)
}

// SetProductName Set the product name of a device The product name would normally be handled by the Manufacturer Specific commmand class, taking the product Type and ID reported by the device and using it to look up the name from the manufacturer_specific.xml file in the OpenZWave config folder. However, there are some devices that do not support the command class, so to enable the user to manually set the name, it is stored with the node data and accessed via this method rather than being reported via a command class Value object.
func (n *Node) SetProductName(name string) {
	manager := GetManager()
	manager.SetNodeProductName(n.HomeID, n.NodeID, name)
}

// SetName Set the name of a node The node name is a user-editable label for the node that would normally be handled by the Node Naming commmand class, but many devices do not support it. So that a node can always be named, OpenZWave stores it with the node data, and provides access through this method and GetNodeName, rather than reporting it via a command class Value object. If the device does support the Node Naming command class, the new name will be sent to the node. The maximum length of a node name is 16 characters.
func (n *Node) SetName(name string) {
	manager := GetManager()
	manager.SetNodeName(n.HomeID, n.NodeID, name)
}

// SetLocation Set the location of a node The node location is a user-editable string that would normally be handled by the Node Naming commmand class, but many devices do not support it. So that a node can always report its location, OpenZWave stores it with the node data, and provides access through this method and GetNodeLocation, rather than reporting it via a command class Value object. If the device does support the Node Naming command class, the new location will be sent to the node.
func (n *Node) SetLocation(location string) {
	manager := GetManager()
	manager.SetNodeLocation(n.HomeID, n.NodeID, location)
}

// SetOn Turns a node on This is a helper method to simplify basic control of a node. It is the equivalent of changing the level reported by the node's Basic command class to 255, and will generate a ValueChanged notification from that class. This command will turn on the device at its last known level, if supported by the device, otherwise it will turn it on at 100%.
func (n *Node) SetOn() {
	manager := GetManager()
	manager.SetNodeOn(n.HomeID, n.NodeID)
}

// SetOff Turns a node off This is a helper method to simplify basic control of a node. It is the equivalent of changing the level reported by the node's Basic command class to zero, and will generate a ValueChanged notification from that class.
func (n *Node) SetOff() {
	manager := GetManager()
	manager.SetNodeOff(n.HomeID, n.NodeID)
}

// SetLevel Sets the basic level of a node This is a helper method to simplify basic control of a node. It is the equivalent of changing the value reported by the node's Basic command class and will generate a ValueChanged notification from that class.
func (n *Node) SetLevel(level uint8) {
	manager := GetManager()
	manager.SetNodeLevel(n.HomeID, n.NodeID, level)
}

// IsInfoReceived Get whether the node information has been received.
func (n *Node) IsInfoReceived() bool {
	manager := GetManager()
	return manager.IsNodeInfoReceived(n.HomeID, n.NodeID)
}

// GetClassInformation Get whether the node has the defined class available or not.
func (n *Node) GetClassInformation(commandClassID uint8) (bool, string, uint8) {
	manager := GetManager()
	return manager.GetNodeClassInformation(n.HomeID, n.NodeID, commandClassID)
}

// IsAwake Get whether the node is awake or asleep.
func (n *Node) IsAwake() bool {
	manager := GetManager()
	return manager.IsNodeAwake(n.HomeID, n.NodeID)
}

// IsFailed Get whether the node is working or has failed.
func (n *Node) IsFailed() bool {
	manager := GetManager()
	return manager.IsNodeFailed(n.HomeID, n.NodeID)
}

// GetQueryStage Get whether the node's query stage as a string.
func (n *Node) GetQueryStage() string {
	manager := GetManager()
	return manager.GetNodeQueryStage(n.HomeID, n.NodeID)
}

// GetDeviceType Get the node device type as reported in the Z-Wave+ Info report.
func (n *Node) GetDeviceType() uint8 {
	manager := GetManager()
	return manager.GetNodeDeviceType(n.HomeID, n.NodeID)
}

// GetDeviceTypeString Get the node device type as reported in the Z-Wave+ Info report.
func (n *Node) GetDeviceTypeString() string {
	manager := GetManager()
	return manager.GetNodeDeviceTypeString(n.HomeID, n.NodeID)
}

// GetRole Get the node device type as reported in the Z-Wave+ Info report.
func (n *Node) GetRole() uint8 {
	manager := GetManager()
	return manager.GetNodeRole(n.HomeID, n.NodeID)
}

// GetRoleString Get the node role as reported in the Z-Wave+ Info report.
func (n *Node) GetRoleString() string {
	manager := GetManager()
	return manager.GetNodeRoleString(n.HomeID, n.NodeID)
}

// GetPlusType Get the node PlusType as reported in the Z-Wave+ Info report.
func (n *Node) GetPlusType() uint8 {
	manager := GetManager()
	return manager.GetNodePlusType(n.HomeID, n.NodeID)
}

// GetPlusTypeString Get the node PlusType as reported in the Z-Wave+ Info report.
func (n *Node) GetPlusTypeString() string {
	manager := GetManager()
	return manager.GetNodePlusTypeString(n.HomeID, n.NodeID)
}
