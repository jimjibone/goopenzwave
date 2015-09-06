package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"gitlab.com/jimjibone/gozwave"
	"sync"
)

type NodeInfo struct {
	HomeID uint32        `json:"home_id"`
	NodeID uint8         `json:"node_id"`
	Node   *gozwave.Node `json:"node"` //TODO: should not store Name etc. but provide getters and setters for these values.
	Values Values        `json:"values"`
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

type Values map[gozwave.ValueIDStringID]*gozwave.ValueID

func (v *Values) Summary() map[gozwave.ValueIDStringID]ValueSummary {
	manager := gozwave.GetManager()
	summaries := make(map[gozwave.ValueIDStringID]ValueSummary)
	for _, valueid := range *v {
		_, valueString := manager.GetValueAsString(valueid)
		summary := ValueSummary{
			ValueID:        valueid.ID,
			NodeID:         valueid.NodeID,
			Genre:          valueid.Genre,
			CommandClassID: valueid.CommandClassID,
			Type:           valueid.Type,
			ReadOnly:       manager.IsValueReadOnly(valueid),
			WriteOnly:      manager.IsValueWriteOnly(valueid),
			Set:            manager.IsValueSet(valueid),
			Polled:         manager.IsValuePolled(valueid),
			Label:          manager.GetValueLabel(valueid),
			Units:          manager.GetValueUnits(valueid),
			Help:           manager.GetValueHelp(valueid),
			Min:            manager.GetValueMin(valueid),
			Max:            manager.GetValueMax(valueid),
			AsString:       valueString,
		}
		summaries[valueid.StringID()] = summary
	}
	return summaries
}

type NodeInfos map[NodeInfoID]*NodeInfo

type NodeSummary struct {
	NodeInfoID       NodeInfoID                               `json:"node_info_id"`
	HomeID           uint32                                   `json:"home_id"`
	NodeID           uint8                                    `json:"node_id"`
	BasicType        uint8                                    `json:"basic_type"`
	GenericType      uint8                                    `json:"generic_type"`
	SpecificType     uint8                                    `json:"specific_type"`
	NodeType         string                                   `json:"node_type"`
	ManufacturerName string                                   `json:"manufacturer_name"`
	ProductName      string                                   `json:"product_name"`
	NodeName         string                                   `json:"node_name"`
	Location         string                                   `json:"location"`
	ManufacturerID   string                                   `json:"manufacturer_id"`
	ProductType      string                                   `json:"product_type"`
	ProductID        string                                   `json:"product_id"`
	Values           map[gozwave.ValueIDStringID]ValueSummary `json:"values"`
}

type ValueSummary struct {
	ValueID        uint64               `json:"value_id"`
	NodeID         uint8                `json:"node_id"`
	Genre          gozwave.ValueIDGenre `json:"genre"`
	CommandClassID uint8                `json:"command_class_id"`
	Type           gozwave.ValueIDType  `json:"type"`
	ReadOnly       bool                 `json:"read_only"`
	WriteOnly      bool                 `json:"write_only"`
	Set            bool                 `json:"set"`
	Polled         bool                 `json:"polled"`
	Label          string               `json:"label"`
	Units          string               `json:"units"`
	Help           string               `json:"help"`
	Min            int32                `json:"min"`
	Max            int32                `json:"max"`
	AsString       string               `json:"string"`
}

var (
	nodeinfos            = make(NodeInfos)
	running              = false
	stop                 = make(chan bool)
	initialQueryComplete = false
)

func NodeManagerRun(controllerPath string, wg *sync.WaitGroup) error {
	// Setup the OpenZWave library.
	options := gozwave.CreateOptions("/usr/local/etc/openzwave/", "", "")
	options.AddOptionLogLevel("SaveLogLevel", gozwave.LogLevelNone)
	options.AddOptionLogLevel("QueueLogLevel", gozwave.LogLevelNone)
	options.AddOptionInt("DumpTrigger", 4)
	options.AddOptionInt("PollInterval", 500)
	options.AddOptionBool("IntervalBetweenPolls", true)
	options.AddOptionBool("ValidateValueChanges", true)
	options.Lock()

	// Start the library and listen for notifications.
	manager := gozwave.CreateManager()
	err := manager.StartNotifications()
	if err != nil {
		return fmt.Errorf("failed to start notifications: %s", err)
	}
	manager.AddDriver(controllerPath)

	// For when we are finished...
	defer func() {
		// All done now finish up.
		manager.RemoveDriver(controllerPath)
		manager.StopNotifications()
		gozwave.DestroyManager()
		gozwave.DestroyOptions()
		wg.Done()
	}()

	// Now continuously listen for notifications or the stop signal.
	running = true
	for {
		select {
		case <-stop:
			running = false
			return nil

		case notification := <-manager.Notifications:
			err = handleNotifcation(notification)
			if err != nil {
				return fmt.Errorf("failed to handle notification: %s", err)
			}
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
			_, oldvalueString := oldvalueid.GetAsString()
			if newvalue.AsString != oldvalueString {
				log.WithFields(log.Fields{
					"node":     nodesummary.NodeInfoID,
					"value":    newvalue.Label,
					"previous": oldvalueString,
					"new":      newvalue.AsString,
				}).Infoln("setting new value for node's value")
				ok := oldvalueid.SetString(newvalue.AsString)
				if !ok {
					log.WithFields(log.Fields{
						"node":     nodesummary.NodeInfoID,
						"value":    newvalue.Label,
						"previous": oldvalueString,
						"new":      newvalue.AsString,
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

	return nil
}

func handleNotifcation(notification *gozwave.Notification) error {
	log.Infoln("NodeManager: received notification:", notification)

	// Switch based on notification type.
	switch notification.Type {
	case gozwave.NotificationTypeNodeAdded:
		// Create a NodeInfo from the notification then add it to the
		// map.
		nodeinfo := &NodeInfo{
			HomeID: notification.HomeID,
			NodeID: notification.NodeID,
			Node:   gozwave.NewNode(notification.HomeID, notification.NodeID),
			Values: make(map[gozwave.ValueIDStringID]*gozwave.ValueID),
		}
		nodeinfos[nodeinfo.ID()] = nodeinfo

	case gozwave.NotificationTypeNodeRemoved:
		// Remove the NodeInfo from the map.
		nodeinfoid := CreateNodeInfoID(notification.HomeID, notification.NodeID)
		if _, found := nodeinfos[nodeinfoid]; found {
			delete(nodeinfos, nodeinfoid)
		}

	case gozwave.NotificationTypeNodeNaming:
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

	case gozwave.NotificationTypeValueAdded, gozwave.NotificationTypeValueChanged:
		// Find the NodeInfo in the map and add/change the ValueID.
		nodeinfoid := CreateNodeInfoID(notification.HomeID, notification.NodeID)
		if nodeinfo, found := nodeinfos[nodeinfoid]; found {
			nodeinfo.Values[notification.ValueID.StringID()] = notification.ValueID
		}

		// Broadcast to all clients that the node has updated.
		if nodeinfo, found := nodeinfos[nodeinfoid]; found {
			message := OutputMessage{
				Topic:   "node-updated",
				Payload: nodeinfo.Summary(),
			}
			clients.Broadcast(message)
		}

	case gozwave.NotificationTypeValueRemoved:
		// Find the NodeInfo in the map and remove the ValueID.
		nodeinfoid := CreateNodeInfoID(notification.HomeID, notification.NodeID)
		if node, found := nodeinfos[nodeinfoid]; found {
			if _, foundVal := node.Values[notification.ValueID.StringID()]; foundVal {
				delete(node.Values, notification.ValueID.StringID())
			}
		}

	case gozwave.NotificationTypeAwakeNodesQueried, gozwave.NotificationTypeAllNodesQueried, gozwave.NotificationTypeAllNodesQueriedSomeDead:
		// The initial node query has completed.
		initialQueryComplete = true
		// TODO: broadcast all node info to connected clients.
	}

	// TODO: return an error at some point.
	return nil
}
