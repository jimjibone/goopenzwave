package main

import (
	"fmt"
	"os"
	"os/signal"

	zwave "github.com/jimjibone/goopenzwave"
	"github.com/urfave/cli/v2"
)

var runCommand = &cli.Command{
	Name:        "run",
	Description: "runs goopenzwave",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "controller",
			Aliases:  []string{"p"},
			Usage:    `path to the controller device (e.g. "/dev/ttyUSB0" on Linux or "\\.\COM3" on Windows)`,
			Required: true,
		},
		&cli.StringFlag{
			Name:    "config-dir",
			Aliases: []string{"c"},
			Usage:   "path to the openzwave config folder",
			Value:   "./config/",
		},
		&cli.StringFlag{
			Name:    "user-dir",
			Aliases: []string{"u"},
			Usage:   "path to the openzwave user folder",
			Value:   "./user/",
		},
	},
	Action: func(c *cli.Context) error {
		zwave.OptionsCreate(c.String("config-dir"), c.String("user-dir"), "")
		zwave.OptionsAddInt("SaveLogLevel", 4)
		zwave.OptionsAddInt("QueueLogLevel", 4)
		zwave.OptionsLock()
		defer zwave.OptionsDestroy()

		zwave.ManagerCreate()
		defer zwave.ManagerDestroy()

		driverfailed := make(chan struct{})
		var nodes []*Node
		findNode := func(homeid uint32, nodeid uint8) *Node {
			for _, n := range nodes {
				if n.HomeID == homeid && n.NodeID == nodeid {
					return n
				}
			}
			return nil
		}
		removeValueID := func(n *Node, val zwave.ValueID) {
			for i, v := range n.ValueIDs {
				if v.Equal(val) {
					// https://github.com/golang/go/wiki/SliceTricks
					n.ValueIDs = append(n.ValueIDs[:i], n.ValueIDs[i+1:]...)
					return
				}
			}
		}
		removeNode := func(homeid uint32, nodeid uint8) {
			for _, n := range nodes {
				if n.HomeID == homeid && n.NodeID == nodeid {
					// // https://github.com/golang/go/wiki/SliceTricks
					// if i < len(nodes)-1 {
					// 	copy(nodes[i:], nodes[i+1:])
					// }
					// nodes[len(nodes)-1] = nil
					// nodes = nodes[:len(nodes)-1]
					// return
					n.Removed = true
				}
			}
		}
		zwave.ManagerAddWatcher(func(n zwave.Notification) {
			if n.Type != zwave.NotificationType_ValueChanged {
				fmt.Println("Notification:", n.String())
			}
			switch n.Type {
			case zwave.NotificationType_ValueAdded: // A new node value has been added to OpenZWave's list. These notifications occur after a node has been discovered, and details of its command classes have been received. Each command class may generate one or more values depending on the complexity of the item being represented.
				if node := findNode(n.ValueID.HomeID(), n.ValueID.NodeID()); node != nil {
					node.ValueIDs = append(node.ValueIDs, n.ValueID)
					fmt.Println("    Node value added:", node, n.ValueID)
				}

			case zwave.NotificationType_ValueRemoved: // A node value has been removed from OpenZWave's list. This only occurs when a node is removed.
				if node := findNode(n.ValueID.HomeID(), n.ValueID.NodeID()); node != nil {
					removeValueID(node, n.ValueID)
					fmt.Println("    Node value removed:", node, n.ValueID)
				}

			case zwave.NotificationType_ValueChanged: // A node value has been updated from the Z-Wave network and it is different from the previous value.
			case zwave.NotificationType_ValueRefreshed: // A node value has been updated from the Z-Wave network.
			case zwave.NotificationType_Group: // The associations for the node have changed. The application should rebuild any group information it holds about the node.
			case zwave.NotificationType_NodeNew: // A new node has been found (not already stored in zwcfg*.xml file)
				if node := findNode(n.ValueID.HomeID(), n.ValueID.NodeID()); node == nil {
					node := &Node{
						HomeID: n.ValueID.HomeID(),
						NodeID: n.ValueID.NodeID(),
						Polled: false,
					}
					nodes = append(nodes, node)
					fmt.Println("    Node new:", node)
				} else {
					fmt.Println("    Node already new:", node)
				}

			case zwave.NotificationType_NodeAdded: // A new node has been added to OpenZWave's list. This may be due to a device being added to the Z-Wave network, or because the application is initializing itself.
				if node := findNode(n.ValueID.HomeID(), n.ValueID.NodeID()); node == nil {
					node := &Node{
						HomeID: n.ValueID.HomeID(),
						NodeID: n.ValueID.NodeID(),
						Polled: false,
					}
					nodes = append(nodes, node)
					fmt.Println("    Node added:", node)
				} else {
					fmt.Println("    Node already added:", node)
				}

			case zwave.NotificationType_NodeRemoved: // A node has been removed from OpenZWave's list. This may be due to a device being removed from the Z-Wave network, or because the application is closing.
				if node := findNode(n.ValueID.HomeID(), n.ValueID.NodeID()); node != nil {
					fmt.Println("    Node removed:", node)
					removeNode(n.ValueID.HomeID(), n.ValueID.NodeID())
				}

			case zwave.NotificationType_NodeProtocolInfo: // Basic node information has been received, such as whether the node is a listening device, a routing device and its baud rate and basic, generic and specific types. It is after this notification that you can call Manager::GetNodeType to obtain a label containing the device description.
			case zwave.NotificationType_NodeNaming: // One of the node names has changed (name, manufacturer, product).
				if node := findNode(n.ValueID.HomeID(), n.ValueID.NodeID()); node != nil {
					node.Name = zwave.GetNodeName(node.HomeID, node.NodeID)
					node.ProductName = zwave.GetNodeProductName(node.HomeID, node.NodeID)
					node.ManufacturerName = zwave.GetNodeManufacturerName(node.HomeID, node.NodeID)
					fmt.Println("    Node naming:", node)
				}

			case zwave.NotificationType_NodeEvent: // A node has triggered an event. This is commonly caused when a node sends a Basic_Set command to the controller. The event value is stored in the notification.
				if node := findNode(n.ValueID.HomeID(), n.ValueID.NodeID()); node != nil {
					fmt.Println("    Node event:", node, n.GetEvent())
				}

			case zwave.NotificationType_PollingDisabled: // Polling of a node has been successfully turned off by a call to Manager::DisablePoll
				if node := findNode(n.ValueID.HomeID(), n.ValueID.NodeID()); node != nil {
					node.Polled = false
					fmt.Println("    Node polling disabled:", node)
				}

			case zwave.NotificationType_PollingEnabled: // Polling of a node has been successfully turned on by a call to Manager::EnablePoll
				if node := findNode(n.ValueID.HomeID(), n.ValueID.NodeID()); node != nil {
					node.Polled = true
					fmt.Println("    Node polling enabled:", node)
				}

			case zwave.NotificationType_SceneEvent: // Scene Activation Set received (Depreciated in 1.8)
			case zwave.NotificationType_CreateButton: // Handheld controller button event created
			case zwave.NotificationType_DeleteButton: // Handheld controller button event deleted
			case zwave.NotificationType_ButtonOn: // Handheld controller button on pressed event
			case zwave.NotificationType_ButtonOff: // Handheld controller button off pressed event
			case zwave.NotificationType_DriverReady: // A driver for a PC Z-Wave controller has been added and is ready to use. The notification will contain the controller's Home ID, which is needed to call most of the Manager methods.
				fmt.Println("    Driver ready")

			case zwave.NotificationType_DriverFailed: // Driver failed to load
				// Exit the app now... or try again
				fmt.Println("    Driver failed")
				close(driverfailed)

			case zwave.NotificationType_DriverReset: // All nodes and values for this driver have been removed. This is sent instead of potentially hundreds of individual node and value notifications.
			case zwave.NotificationType_EssentialNodeQueriesComplete: // The queries on a node that are essential to its operation have been completed. The node can now handle incoming messages.
			case zwave.NotificationType_NodeQueriesComplete: // All the initialization queries on a node have been completed.
			case zwave.NotificationType_AwakeNodesQueried: // All awake nodes have been queried, so client application can expected complete data for these nodes.
				// Ready to start doing things from this point
				fmt.Println("    Awake nodes queried")
				close(driverfailed)

			case zwave.NotificationType_AllNodesQueriedSomeDead: // All nodes have been queried but some dead nodes found.
				// Ready to start doing things from this point
				fmt.Println("    All nodes queried some dead")
				close(driverfailed)

			case zwave.NotificationType_AllNodesQueried: // All nodes have been queried, so client application can expected complete data.
				// Ready to start doing things from this point
				fmt.Println("    All nodes queried")
				close(driverfailed)

			case zwave.NotificationType_Notification: // An error has occurred that we need to report.
			case zwave.NotificationType_DriverRemoved: // The Driver is being removed. (either due to Error or by request) Do Not Call Any Driver Related Methods after receiving this call
			case zwave.NotificationType_ControllerCommand: // When Controller Commands are executed, Notifications of Success/Failure etc are communicated via this Notification Notification::GetEvent returns Driver::ControllerCommand and Notification::GetNotification returns Driver::ControllerState
			case zwave.NotificationType_NodeReset: // The Device has been reset and thus removed from the NodeList in OZW
			case zwave.NotificationType_UserAlerts: // Warnings and Notifications Generated by the library that should be displayed to the user (eg, out of date config files)
			case zwave.NotificationType_ManufacturerSpecificDBReady: // The ManufacturerSpecific Database Is Ready
			}
		})

		zwave.ManagerAddDriver(c.String("controller"))
		defer zwave.ManagerRemoveDriver(c.String("controller"))

		fmt.Println("version:", zwave.ManagerVersionString())
		fmt.Println("version:", zwave.ManagerVersionLongString())
		major, minor := zwave.ManagerVersion()
		fmt.Printf("version: major: %d, minor: %d\n", major, minor)

		// wait for ctrl-c
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)
		select {
		case <-sig:
		case <-driverfailed:
		}
		signal.Reset(os.Interrupt)

		fmt.Printf("nodes: %d\n", len(nodes))
		for _, node := range nodes {
			fmt.Printf("node: %s, name: %s, product: %s, manufacturer: %s\n",
				node,
				zwave.GetNodeName(node.HomeID, node.NodeID),
				zwave.GetNodeProductName(node.HomeID, node.NodeID),
				zwave.GetNodeManufacturerName(node.HomeID, node.NodeID),
			)
			for _, vid := range node.ValueIDs {
				valid, err := zwave.IsValueValid(vid)
				if err != nil {
					panic(err)
				}
				if valid {
					name, err := zwave.GetValueLabel(vid, -1)
					if err != nil {
						panic(err)
					}
					val, err := zwave.GetValueAsString(vid)
					if err != nil {
						panic(err)
					}
					units, err := zwave.GetValueUnits(vid)
					if err != nil {
						panic(err)
					}
					var list []string
					if vid.Type() == zwave.ValueType_List {
						list, err = zwave.GetValueListItems(vid)
						if err != nil {
							panic(err)
						}
					}
					help, err := zwave.GetValueHelp(vid, -1)
					if err != nil {
						panic(err)
					}
					fmt.Printf("    valueid: %s\n", vid)
					fmt.Printf("             name: %s\n", name)
					fmt.Printf("             val: %s (%s)\n", val, units)
					fmt.Printf("             list: %q\n", list)
					fmt.Printf("             help: %s\n", help)
				} else {
					fmt.Printf("    valueid: %s, INVALID\n", vid)
				}
			}
		}

		return nil
	},
}
