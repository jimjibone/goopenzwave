package main

import (
	"flag"
	"fmt"
	"gitlab.com/jimjibone/gozwave"
	"os"
	"os/signal"
	"time"
)

type NodeInfo struct {
	HomeID uint32
	NodeID uint8
	Node   *gozwave.Node
	Values map[uint64]*gozwave.ValueID
	//TODO use (and create) Value type with embedded ValueID.
}

var Nodes = make(map[uint8]*NodeInfo)
var InitialQueryComplete = make(chan bool)

func processNotifications(m *gozwave.Manager) {
	var sentInitialQueryComplete bool
	for {
		select {
		case notification := <-m.Notifications:
			fmt.Println(notification)

			// Switch based on notification type.
			switch notification.Type {
			case gozwave.NotificationTypeNodeAdded:
				// Create a NodeInfo from the notification then add it to the
				// map.
				nodeinfo := &NodeInfo{
					HomeID: notification.HomeID,
					NodeID: notification.NodeID,
					Node:   gozwave.NewNode(notification.HomeID, notification.NodeID),
					Values: make(map[uint64]*gozwave.ValueID),
				}
				Nodes[nodeinfo.NodeID] = nodeinfo

			case gozwave.NotificationTypeNodeRemoved:
				// Remove the NodeInfo from the map.
				if _, found := Nodes[notification.NodeID]; found {
					delete(Nodes, notification.NodeID)
				}

			case gozwave.NotificationTypeValueAdded, gozwave.NotificationTypeValueChanged:
				// Find the NodeInfo in the map and add/change the ValueID.
				if node, found := Nodes[notification.NodeID]; found {
					node.Values[notification.ValueID.ID] = notification.ValueID
				}

			case gozwave.NotificationTypeValueRemoved:
				// Find the NodeInfo in the map and remove the ValueID.
				if node, found := Nodes[notification.NodeID]; found {
					if _, foundVal := node.Values[notification.ValueID.ID]; foundVal {
						delete(node.Values, notification.ValueID.ID)
					}
				}

			case gozwave.NotificationTypeAwakeNodesQueried, gozwave.NotificationTypeAllNodesQueried, gozwave.NotificationTypeAllNodesQueriedSomeDead:
				// The initial node query has completed.
				fmt.Println("about to send InitialQueryComplete flag...")
				if sentInitialQueryComplete == false {
					InitialQueryComplete <- true
				}
				fmt.Println("...done")
			}
		}
	}
}

//func init() {
//	Nodes = make(map[uint8]*NodeInfo)
//}

func main() {
	var controllerPath string
	flag.StringVar(&controllerPath, "controller", "/dev/ttyACM0", "the path to your controller device")
	flag.Parse()

	fmt.Println("gozwave example starting with openzwave version:", gozwave.GetManagerVersionAsString())

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
		fmt.Println("ERROR: failed to start notifications:", err)
	}
	manager.AddDriver(controllerPath)

	// Now listen to the many notifications.
	go processNotifications(manager)

	// Wait here until the initial node query has completed.
	<-InitialQueryComplete
	fmt.Println("Finished initial scan, now setting up polling...")

	// Now we will enable polling for a variable. In this simple example, it
	// has been hardwired to poll COMMAND_CLASS_BASIC on each node that
	// supports this setting.
	for _, node := range Nodes {
		// Skip the controller (most likely node 1).
		if node.NodeID == 1 {
			continue
		}

		// For each value for this node, set up polling.
		for i := range node.Values {
			valueid := node.Values[i]

			if valueid.CommandClassID == 0x20 {
				// Enable polling with "intensity" of 2. Though, this is irrelevant with only one value polled.
				manager.EnablePoll(valueid, 2)
			}
		}
	}

	fmt.Println("Initial scan complete. Now polling for updates...")
	fmt.Println("Hit ctrl-c to quit")

	if lightNode, found := Nodes[2]; found {
		time.Sleep(5 * time.Second)
		fmt.Println("swithcing light")
		for _, valueid := range lightNode.Values {
			if manager.GetValueLabel(valueid) == "Level" {
				manager.SetValueUint8(valueid, 255)
				fmt.Println("set value Level 255")
			}
		}
	}

	if lightNode, found := Nodes[2]; found {
		time.Sleep(5 * time.Second)
		fmt.Println("swithcing light colour")
		for _, valueid := range lightNode.Values {
			if manager.GetValueLabel(valueid) == "Color" {
				manager.SetValueString(valueid, "#00FF000000")
				fmt.Println("set value Color Lime")
			}
		}
	}

	if lightNode, found := Nodes[2]; found {
		time.Sleep(5 * time.Second)
		fmt.Println("swithcing light colour")
		for _, valueid := range lightNode.Values {
			if manager.GetValueLabel(valueid) == "Color" {
				manager.SetValueString(valueid, "#00000000FF")
				fmt.Println("set value Color CW")
			}
		}
	}

	if lightNode, found := Nodes[2]; found {
		time.Sleep(5 * time.Second)
		fmt.Println("swithcing light colour")
		for _, valueid := range lightNode.Values {
			if manager.GetValueLabel(valueid) == "Color" {
				manager.SetValueString(valueid, "#000000FF00")
				fmt.Println("set value Color WW")
			}
		}
	}

	if lightNode, found := Nodes[2]; found {
		time.Sleep(5 * time.Second)
		fmt.Println("swithcing light")
		for _, valueid := range lightNode.Values {
			if manager.GetValueLabel(valueid) == "Level" {
				manager.SetValueUint8(valueid, 0)
				fmt.Println("set value Level 0")
			}
		}
	}

	// Now wait for the user to quit.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	// Print out what we know about the network.
	fmt.Println("Nodes:")
	for id, node := range Nodes {
		fmt.Printf("\t%d: Node: %s Values:\n", id, node.Node)
		for i := range node.Values {
			fmt.Printf("\t\t%d: %s\n", i, node.Values[i].InfoString())
		}
	}

	// All done now finish up.
	manager.RemoveDriver(controllerPath)
	manager.StopNotifications()
	gozwave.DestroyManager()
	gozwave.DestroyOptions()

	fmt.Println("finished")
}
