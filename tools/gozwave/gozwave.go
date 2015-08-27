package main

import (
	"flag"
	"fmt"
	"gitlab.com/jimjibone/gozwave"
	"os"
	"os/signal"
)

type NodeInfo struct {
	HomeId uint32
	NodeId uint8
	// Polled bool
	Values []*gozwave.ValueID
}

//type NodeInfoCollector struct {
//	nodeInfo chan NodeInfo
//}
//
//func watcherFunc(notification *gozwave.Notification, userdata interface{}) {
//	notistring := notification.GetAsString()
//	fmt.Println("gozwave watcher called with notification:", notistring)
//
//	nodeInfo := NodeInfo{
//		HomeId: notification.GetHomeId(),
//		NodeId: notification.GetNodeId(),
//		// Polled: false,
//	}
//	nodeInfo.Values = append(nodeInfo.Values, notification.GetValueId())
//
//	fmt.Println("gozwave watcher NodeInfo:", nodeInfo)
//
//	switch nodeInfoCltr := userdata.(type) {
//	case NodeInfoCollector:
//		nodeInfoCltr.nodeInfo <- nodeInfo
//	default:
//		panic("userdata is an unexpected type")
//	}
//}
//
//func collectNodeInfo(nodeInfoCollector *NodeInfoCollector, nodes *map[uint8]*NodeInfo) {
//	for {
//		nodeInfo, ok := <-nodeInfoCollector.nodeInfo
//		if !ok {
//			return
//		}
//		fmt.Println("collectNodeInfo NodeInfo:", nodeInfo)
//		if node, found := (*nodes)[nodeInfo.NodeId]; found {
//			node.Values = append(node.Values, nodeInfo.Values...)
//		} else {
//			(*nodes)[nodeInfo.NodeId] = &nodeInfo
//		}
//	}
//}

func processNotifications(m *gozwave.Manager) {
	for {
		select {
		case notification := <-m.Notifications:
			fmt.Println("processNotifications:", notification)
		}
	}
}

func main() {
	var controllerPath string
	flag.StringVar(&controllerPath, "controller", "/dev/ttyACM0", "the path to your controller device")
	flag.Parse()

	fmt.Println("gozwave example starting with openzwave version:", gozwave.GetManagerVersionAsString())
	options := gozwave.CreateOptions("./config/", "", "")
	options.AddOptionInt("SaveLogLevel", 8)
	options.AddOptionInt("QueueLogLevel", 9)
	options.AddOptionInt("DumpTrigger", 4)
	options.AddOptionInt("PollInterval", 500)
	options.AddOptionBool("IntervalBetweenPolls", true)
	options.AddOptionBool("ValidateValueChanges", true)
	options.Lock()

	// nodeInfoCollector := NodeInfoCollector{
	// 	nodeInfo: make(chan NodeInfo),
	// }
	manager := gozwave.CreateManager()
	// watcher, ok := manager.AddWatcher(watcherFunc, nodeInfoCollector)
	// if ok == false {
	// 	fmt.Println("ERROR: failed to add watcher")
	// 	return
	// }
	err := manager.StartNotifications()
	if err != nil {
		fmt.Println("ERROR: failed to start notifications:", err)
	}

	manager.AddDriver(controllerPath)

	// nodes := make(map[uint8]*NodeInfo)

	// Collect the node info from the channel.
	// go collectNodeInfo(&nodeInfoCollector, &nodes)

	go processNotifications(manager)

	// Now wait for the user to quit.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	// All done now finish up.
	manager.RemoveDriver(controllerPath)
	// manager.RemoveWatcher(watcher)
	manager.StopNotifications()
	gozwave.DestroyManager()
	gozwave.DestroyOptions()
	fmt.Println("finished")
}
