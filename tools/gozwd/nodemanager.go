package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/jimjibone/goopenzwave"
	"sync"
	"time"
)

type NodeInfo struct {
	HomeID uint32            `json:"home_id"`
	NodeID uint8             `json:"node_id"`
	Node   *goopenzwave.Node `json:"node"` //TODO: should not store Name etc. but provide getters and setters for these values.
	Values Values            `json:"values"`
	Time   time.Time         `json:"update-time"`
}

type NodeInfoID string

func (n *NodeInfo) ID() NodeInfoID {
	return CreateNodeInfoID(n.HomeID, n.NodeID)
}

func CreateNodeInfoID(homeID uint32, nodeID uint8) NodeInfoID {
	return NodeInfoID(fmt.Sprintf("%d:%d", homeID, nodeID))
}

func (n *NodeInfo) Summary() NodeSummary {
	return NodeSummary{
		NodeInfoID:       n.ID(),
		HomeID:           n.HomeID,
		NodeID:           n.NodeID,
		BasicType:        n.Node.GetBasicType(),
		GenericType:      n.Node.GetGenericType(),
		SpecificType:     n.Node.GetSpecificType(),
		NodeType:         n.Node.GetType(),
		ManufacturerName: n.Node.GetManufacturerName(),
		ProductName:      n.Node.GetProductName(),
		NodeName:         n.Node.GetName(),
		Location:         n.Node.GetLocation(),
		ManufacturerID:   n.Node.GetManufacturerID(),
		ProductType:      n.Node.GetProductType(),
		ProductID:        n.Node.GetProductID(),
		Values:           n.Values.Summary(),
	}
}

type Values map[string]*goopenzwave.ValueID

func (v *Values) Summary() map[string]ValueSummary {
	summaries := make(map[string]ValueSummary)
	for _, valueid := range *v {
		summary := ValueSummary{
			ValueID:        valueid.ID,
			NodeID:         valueid.NodeID,
			Genre:          valueid.Genre,
			CommandClassID: valueid.CommandClassID,
			Type:           valueid.Type,
			ReadOnly:       valueid.IsReadOnly(),
			WriteOnly:      valueid.IsWriteOnly(),
			Set:            valueid.IsSet(),
			Polled:         valueid.IsPolled(),
			Label:          valueid.GetLabel(),
			Units:          valueid.GetUnits(),
			Help:           valueid.GetHelp(),
			Min:            valueid.GetMin(),
			Max:            valueid.GetMax(),
			AsString:       valueid.GetAsString(),
		}
		summaries[valueid.IDString()] = summary
	}
	return summaries
}

type NodeInfos map[NodeInfoID]*NodeInfo

type NodeSummary struct {
	NodeInfoID       NodeInfoID              `json:"node_info_id"`
	HomeID           uint32                  `json:"home_id"`
	NodeID           uint8                   `json:"node_id"`
	BasicType        uint8                   `json:"basic_type"`
	GenericType      uint8                   `json:"generic_type"`
	SpecificType     uint8                   `json:"specific_type"`
	NodeType         string                  `json:"node_type"`
	ManufacturerName string                  `json:"manufacturer_name"`
	ProductName      string                  `json:"product_name"`
	NodeName         string                  `json:"node_name"`
	Location         string                  `json:"location"`
	ManufacturerID   string                  `json:"manufacturer_id"`
	ProductType      string                  `json:"product_type"`
	ProductID        string                  `json:"product_id"`
	Values           map[string]ValueSummary `json:"values"`
}

type ValueSummary struct {
	ValueID        uint64                   `json:"value_id"`
	NodeID         uint8                    `json:"node_id"`
	Genre          goopenzwave.ValueIDGenre `json:"genre"`
	CommandClassID uint8                    `json:"command_class_id"`
	Type           goopenzwave.ValueIDType  `json:"type"`
	ReadOnly       bool                     `json:"read_only"`
	WriteOnly      bool                     `json:"write_only"`
	Set            bool                     `json:"set"`
	Polled         bool                     `json:"polled"`
	Label          string                   `json:"label"`
	Units          string                   `json:"units"`
	Help           string                   `json:"help"`
	Min            int32                    `json:"min"`
	Max            int32                    `json:"max"`
	AsString       string                   `json:"string"`
}

type NodeInfoIDMessage struct {
	NodeInfoID NodeInfoID `json:"node_info_id"`
}

var (
	nodeinfos            = make(NodeInfos)
	running              = false
	stop                 = make(chan bool)
	initialQueryComplete = false
)

func NodeManagerRun(controllerPath string, wg *sync.WaitGroup) error {
	// Setup the OpenZWave library options§.
	options := goopenzwave.CreateOptions("/usr/local/etc/openzwave/", "", "")
	options.AddOptionLogLevel("SaveLogLevel", goopenzwave.LogLevelNone)
	options.AddOptionLogLevel("QueueLogLevel", goopenzwave.LogLevelNone)
	options.AddOptionInt("DumpTrigger", 4)
	options.AddOptionInt("PollInterval", 500)
	options.AddOptionBool("IntervalBetweenPolls", true)
	options.AddOptionBool("ValidateValueChanges", true)
	options.Lock()

	// Start the library and listen for notifications.
	err := goopenzwave.Start(handleNotification)
	if err != nil {
		log.Fatalln("failed to start goopenzwave package:", err)
	}

	err = goopenzwave.AddDriver(controllerPath)
	if err != nil {
		log.Fatalln("failed to add goopenzwave driver:", err)
	}

	// For when we are finished...
	defer func() {
		// All done now finish up.
		err := goopenzwave.RemoveDriver(controllerPath)
		if err != nil {
			log.Fatalln("failed to remove goopenzwave driver:", err)
		}
		err = goopenzwave.Stop()
		if err != nil {
			log.Fatalln("failed to stop goopenzwave package:", err)
		}
		goopenzwave.DestroyOptions()
		wg.Done()
	}()

	// Now continuously listen for the stop signal.
	running = true
	for {
		select {
		case <-stop:
			running = false
			return nil
		}
	}
}

func NodeManagerStop() {
	if running {
		stop <- true
	}
}

func NodeManagerGetNodes() []NodeSummary {
	var summaries []NodeSummary
	for _, nodeinfo := range nodeinfos {
		summaries = append(summaries, nodeinfo.Summary())
	}
	return summaries
}

func NodeManagerUpdateNode(nodesummary NodeSummary) error {
	// Compare the received node summary with the current values from the
	// Manager.
	nodeinfo, found := nodeinfos[nodesummary.NodeInfoID]
	if !found {
		return fmt.Errorf("could not find node (%s) in the node list", nodesummary.NodeInfoID)
	}

	updated := false

	// Node.Name
	if oldvalue := nodeinfo.Node.GetName(); nodesummary.NodeName != oldvalue {
		log.WithFields(log.Fields{
			"node":     nodesummary.NodeInfoID,
			"previous": oldvalue,
			"new":      nodesummary.NodeName,
		}).Infoln("setting new name for node")
		nodeinfo.Node.SetName(nodesummary.NodeName)
		updated = true
	}

	// Node.Location
	if oldvalue := nodeinfo.Node.GetLocation(); nodesummary.Location != oldvalue {
		log.WithFields(log.Fields{
			"node":     nodesummary.NodeInfoID,
			"previous": oldvalue,
			"new":      nodesummary.Location,
		}).Infoln("setting new location for node")
		nodeinfo.Node.SetLocation(nodesummary.Location)
		updated = true
	}

	// Node.Values
	for oldvalueidstringid, oldvalueid := range nodeinfo.Values {
		newvalue, found := nodesummary.Values[oldvalueidstringid]
		if found {
			oldvalueString := oldvalueid.GetAsString()
			if newvalue.AsString != oldvalueString {
				log.WithFields(log.Fields{
					"node":     nodesummary.NodeInfoID,
					"value":    newvalue.Label,
					"previous": oldvalueString,
					"new":      newvalue.AsString,
				}).Infoln("setting new value for node's value")
				err := oldvalueid.SetString(newvalue.AsString)
				if err != nil {
					log.WithFields(log.Fields{
						"node":     nodesummary.NodeInfoID,
						"value":    newvalue.Label,
						"previous": oldvalueString,
						"new":      newvalue.AsString,
						"error":    err,
					}).Errorln("failed to set value as string")
				}
				updated = true
			}
		}
	}

	if updated == false {
		log.WithFields(log.Fields{
			"node":     nodesummary.NodeInfoID,
			"previous": nodeinfo,
			"new":      nodesummary,
		}).Warnln("new node is identical to existing")
	}

	// NodeInfoID       NodeInfoID     `json:"node_info_id"`
	// HomeID           uint32         `json:"home_id"`
	// NodeID           uint8          `json:"node_id"`
	// BasicType        uint8          `json:"basic_type"`
	// GenericType      uint8          `json:"generic_type"`
	// SpecificType     uint8          `json:"specific_type"`
	// NodeType         string         `json:"node_type"`
	// ManufacturerName string         `json:"manufacturer_name"`
	// ProductName      string         `json:"product_name"`
	// NodeName         string         `json:"node_name"`
	// Location         string         `json:"location"`
	// ManufacturerID   string         `json:"manufacturer_id"`
	// ProductType      string         `json:"product_type"`
	// ProductID        string         `json:"product_id"`
	// Values           []ValueSummary `json:"values"`

	// TODO: return an error if one of the updates failed.
	return nil
}

func NodeManagerToggleNode(nodeinfoid NodeInfoIDMessage) error {
	nodeinfo, found := nodeinfos[nodeinfoid.NodeInfoID]
	if !found {
		return fmt.Errorf("could not find node (%s) in the node list", nodeinfoid.NodeInfoID)
	}

	_ = nodeinfo

	log.Warnln("not toggling state of node")
	return nil
}

func handleNotification(notification *goopenzwave.Notification) {
	// Switch based on notification type.
	switch notification.Type {
	case goopenzwave.NotificationTypeValueAdded:
		// A new node value has been added to OpenZWave's list. These
		// notifications occur after a node has been discovered, and details of
		// its command classes have been received. Each command class may
		// generate one or more values depending on the complexity of the item
		// being represented.
		log.WithFields(log.Fields{
			"Home":     notification.HomeID,
			"Node":     notification.NodeID,
			"Genre":    notification.ValueID.Genre,
			"Class":    notification.ValueID.CommandClassID,
			"Instance": notification.ValueID.Instance,
			"Index":    notification.ValueID.Index,
			"Type":     notification.ValueID.Type,
		}).Infoln("Notification: Value Added")

		// Add the value to the correct node.
		nodeinfoid := CreateNodeInfoID(notification.HomeID, notification.NodeID)
		if nodeinfo, found := nodeinfos[nodeinfoid]; found {
			nodeinfo.Values[notification.ValueID.IDString()] = notification.ValueID
			nodeinfo.Time = time.Now()

			// Broadcast to all clients that the node has updated.
			message := OutputMessage{
				Topic:   "node-updated",
				Payload: nodeinfo.Summary(),
			}
			clients.Broadcast(message)
		}

	case goopenzwave.NotificationTypeValueRemoved:
		// A node value has been removed from OpenZWave's list. This only occurs
		// when a node is removed.
		log.WithFields(log.Fields{
			"Home":     notification.HomeID,
			"Node":     notification.NodeID,
			"Genre":    notification.ValueID.Genre,
			"Class":    notification.ValueID.CommandClassID,
			"Instance": notification.ValueID.Instance,
			"Index":    notification.ValueID.Index,
			"Type":     notification.ValueID.Type,
		}).Infoln("Notification: Value Removed")

		// Remove the value from the node.
		nodeinfoid := CreateNodeInfoID(notification.HomeID, notification.NodeID)
		if nodeinfo, found := nodeinfos[nodeinfoid]; found {
			if _, foundVal := nodeinfo.Values[notification.ValueID.IDString()]; foundVal {
				delete(nodeinfo.Values, notification.ValueID.IDString())
			}
			nodeinfo.Time = time.Now()

			// Broadcast to all clients that the node has updated.
			message := OutputMessage{
				Topic:   "node-updated",
				Payload: nodeinfo.Summary(),
			}
			clients.Broadcast(message)
		}

	case goopenzwave.NotificationTypeValueChanged:
		// A node value has been updated from the Z-Wave network and it is
		// different from the previous value.
		log.WithFields(log.Fields{
			"Home":     notification.HomeID,
			"Node":     notification.NodeID,
			"Genre":    notification.ValueID.Genre,
			"Class":    notification.ValueID.CommandClassID,
			"Instance": notification.ValueID.Instance,
			"Index":    notification.ValueID.Index,
			"Type":     notification.ValueID.Type,
		}).Infoln("Notification: Value Changed")

		// Change the value of the correct node.
		nodeinfoid := CreateNodeInfoID(notification.HomeID, notification.NodeID)
		if nodeinfo, found := nodeinfos[nodeinfoid]; found {
			nodeinfo.Values[notification.ValueID.IDString()] = notification.ValueID
			nodeinfo.Time = time.Now()

			// Broadcast to all clients that the node has updated.
			message := OutputMessage{
				Topic:   "node-updated",
				Payload: nodeinfo.Summary(),
			}
			clients.Broadcast(message)
		}

	case goopenzwave.NotificationTypeValueRefreshed:
		// A node value has been updated from the Z-Wave network.
		log.WithFields(log.Fields{
			"Home":     notification.HomeID,
			"Node":     notification.NodeID,
			"Genre":    notification.ValueID.Genre,
			"Class":    notification.ValueID.CommandClassID,
			"Instance": notification.ValueID.Instance,
			"Index":    notification.ValueID.Index,
			"Type":     notification.ValueID.Type,
		}).Infoln("Notification: Value Refreshed")

		// Update the update time of the correct node.
		nodeinfoid := CreateNodeInfoID(notification.HomeID, notification.NodeID)
		if nodeinfo, found := nodeinfos[nodeinfoid]; found {
			nodeinfo.Time = time.Now()
		}

	case goopenzwave.NotificationTypeGroup:
		// TODO this...

		//------------------------------------------------------------------------------

	case goopenzwave.NotificationTypeNodeAdded:
		// Create a NodeInfo from the notification then add it to the
		// map.
		nodeinfo := &NodeInfo{
			HomeID: notification.HomeID,
			NodeID: notification.NodeID,
			Node:   goopenzwave.NewNode(notification.HomeID, notification.NodeID),
			Values: make(map[string]*goopenzwave.ValueID),
		}
		nodeinfos[nodeinfo.ID()] = nodeinfo

	case goopenzwave.NotificationTypeNodeRemoved:
		// Remove the NodeInfo from the map.
		nodeinfoid := CreateNodeInfoID(notification.HomeID, notification.NodeID)
		if _, found := nodeinfos[nodeinfoid]; found {
			delete(nodeinfos, nodeinfoid)
		}

	case goopenzwave.NotificationTypeNodeNaming:
		// Find the NodeInfo in the map and broadcast to all clients that the
		// node has updated.
		nodeinfoid := CreateNodeInfoID(notification.HomeID, notification.NodeID)
		if nodeinfo, found := nodeinfos[nodeinfoid]; found {
			message := OutputMessage{
				Topic:   "node-updated",
				Payload: nodeinfo.Summary(),
			}
			clients.Broadcast(message)
		}

	// case goopenzwave.NotificationTypeValueAdded, goopenzwave.NotificationTypeValueChanged:
	// 	// Find the NodeInfo in the map and add/change the ValueID.
	// 	nodeinfoid := CreateNodeInfoID(notification.HomeID, notification.NodeID)
	// 	if nodeinfo, found := nodeinfos[nodeinfoid]; found {
	// 		nodeinfo.Values[notification.ValueID.IDString()] = notification.ValueID
	// 	}

	// 	// Broadcast to all clients that the node has updated.
	// 	if nodeinfo, found := nodeinfos[nodeinfoid]; found {
	// 		message := OutputMessage{
	// 			Topic:   "node-updated",
	// 			Payload: nodeinfo.Summary(),
	// 		}
	// 		clients.Broadcast(message)
	// 	}

	// case goopenzwave.NotificationTypeValueRemoved:
	// 	// Find the NodeInfo in the map and remove the ValueID.
	// 	nodeinfoid := CreateNodeInfoID(notification.HomeID, notification.NodeID)
	// 	if node, found := nodeinfos[nodeinfoid]; found {
	// 		if _, foundVal := node.Values[notification.ValueID.IDString()]; foundVal {
	// 			delete(node.Values, notification.ValueID.IDString())
	// 		}
	// 	}

	case goopenzwave.NotificationTypeAwakeNodesQueried, goopenzwave.NotificationTypeAllNodesQueried, goopenzwave.NotificationTypeAllNodesQueriedSomeDead:
		// The initial node query has completed.
		initialQueryComplete = true
		// TODO: broadcast all node info to connected clients.

	default:
		log.WithFields(log.Fields{
			"notification": notification,
		}).Warnln("unhandled notification received")
	}
}
