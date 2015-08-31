package main

import (
	"flag"
	"fmt"
	"gitlab.com/jimjibone/gozwave"
	"os"
	"os/signal"
)

type NodeInfo struct {
	HomeID uint32
	NodeID uint8
	Node   *gozwave.Node
	Values []*gozwave.ValueID
}

var Nodes map[uint8]*NodeInfo

func processNotifications(m *gozwave.Manager) {
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
				}
				Nodes[nodeinfo.NodeID] = nodeinfo

			case gozwave.NotificationTypeNodeRemoved:
				// Remove the NodeInfo from the map.
				if _, found := Nodes[notification.NodeID]; found {
					delete(Nodes, notification.NodeID)
				}

			case gozwave.NotificationTypeValueAdded:
				// Find the NodeInfo in the map and add the ValueID to it.
				if node, found := Nodes[notification.NodeID]; found {
					node.Values = append(node.Values, notification.ValueID)
				}
			}
		}
	}
}

func init() {
	Nodes = make(map[uint8]*NodeInfo)
}

func main() {
	var controllerPath string
	flag.StringVar(&controllerPath, "controller", "/dev/ttyACM0", "the path to your controller device")
	flag.Parse()

	fmt.Println("gozwave example starting with openzwave version:", gozwave.GetManagerVersionAsString())

	// Setup the OpenZWave library.
	options := gozwave.CreateOptions("./config/", "", "")
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

	// Now wait for the user to quit.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	// All done now finish up.
	manager.RemoveDriver(controllerPath)
	manager.StopNotifications()
	gozwave.DestroyManager()
	gozwave.DestroyOptions()

	// Print out what we know about the network.
	fmt.Println("Nodes:")
	for id, node := range Nodes {
		fmt.Printf("\t%d: %s\n", id, node.Node)
	}

	fmt.Println("finished")
}
