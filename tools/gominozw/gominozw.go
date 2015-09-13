package main

import (
	"flag"
	"fmt"
	"github.com/jimjibone/goopenzwave"
	"os"
	"os/signal"
)

type NodeInfo struct {
	HomeID uint32
	NodeID uint8
	Node   *goopenzwave.Node
	Values map[uint64]*goopenzwave.ValueID
}

var Nodes = make(map[uint8]*NodeInfo)
var InitialQueryComplete = make(chan bool)

func processNotifications(m *goopenzwave.Manager) {
	var sentInitialQueryComplete bool
	for {
		select {
		case notification := <-m.Notifications:
			fmt.Println(notification)

			// Switch based on notification type.
			switch notification.Type {
			case goopenzwave.NotificationTypeNodeAdded:
				// Create a NodeInfo from the notification then add it to the
				// map.
				nodeinfo := &NodeInfo{
					HomeID: notification.HomeID,
					NodeID: notification.NodeID,
					Node:   goopenzwave.NewNode(notification.HomeID, notification.NodeID),
					Values: make(map[uint64]*goopenzwave.ValueID),
				}
				Nodes[nodeinfo.NodeID] = nodeinfo

			case goopenzwave.NotificationTypeNodeRemoved:
				// Remove the NodeInfo from the map.
				if _, found := Nodes[notification.NodeID]; found {
					delete(Nodes, notification.NodeID)
				}

			case goopenzwave.NotificationTypeValueAdded, goopenzwave.NotificationTypeValueChanged:
				// Find the NodeInfo in the map and add/change the ValueID.
				if node, found := Nodes[notification.NodeID]; found {
					node.Values[notification.ValueID.ID] = notification.ValueID
				}

			case goopenzwave.NotificationTypeValueRemoved:
				// Find the NodeInfo in the map and remove the ValueID.
				if node, found := Nodes[notification.NodeID]; found {
					if _, foundVal := node.Values[notification.ValueID.ID]; foundVal {
						delete(node.Values, notification.ValueID.ID)
					}
				}

			case goopenzwave.NotificationTypeAwakeNodesQueried, goopenzwave.NotificationTypeAllNodesQueried, goopenzwave.NotificationTypeAllNodesQueriedSomeDead:
				// The initial node query has completed.
				if sentInitialQueryComplete == false {
					InitialQueryComplete <- true
				}
			}
		}
	}
}

func main() {
	var controllerPath string
	flag.StringVar(&controllerPath, "controller", "/dev/ttyACM0", "the path to your controller device")
	flag.Parse()

	fmt.Println("gominozw example starting with openzwave version:", goopenzwave.GetManagerVersionAsString())

	// Setup the OpenZWave library.
	options := goopenzwave.CreateOptions("/usr/local/etc/openzwave/", "", "")
	options.AddOptionLogLevel("SaveLogLevel", goopenzwave.LogLevelNone)
	options.AddOptionLogLevel("QueueLogLevel", goopenzwave.LogLevelNone)
	options.AddOptionInt("DumpTrigger", 4)
	options.AddOptionInt("PollInterval", 500)
	options.AddOptionBool("IntervalBetweenPolls", true)
	options.AddOptionBool("ValidateValueChanges", true)
	options.Lock()

	// Start the library and listen for notifications.
	err := goopenzwave.Start(handleNotifcation)
	if err != nil {
		log.Fatalln("failed to start goopenzwave package:", err)
	}

	// Add a driver using the supplied controller path.
	err = goopenzwave.AddDriver(controllerPath)
	if err != nil {
		log.Fatalln("failed to add goopenzwave driver:", err)
	}

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
	goopenzwave.RemoveDriver(controllerPath)
	goopenzwave.Stop()
	goopenzwave.DestroyOptions()
}
