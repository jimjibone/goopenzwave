package main

import (
	"flag"
	"fmt"
	"gitlab.com/jimjibone/gozwave"
	"os"
	"os/signal"
)

var allNotifications []gozwave.Notification

func processNotifications(m *gozwave.Manager) {
	for {
		select {
		case notification := <-m.Notifications:
			fmt.Printf("processNotifications: %+v\n", notification)
			allNotifications = append(allNotifications, notification)
		}
	}
}

func main() {
	var controllerPath string
	flag.StringVar(&controllerPath, "controller", "/dev/ttyACM0", "the path to your controller device")
	flag.Parse()

	fmt.Println("gozwave example starting with openzwave version:", gozwave.GetManagerVersionAsString())

	options := gozwave.CreateOptions("./config/", "", "")
	options.AddOptionLogLevel("SaveLogLevel", gozwave.LogLevelNone)
	options.AddOptionLogLevel("QueueLogLevel", gozwave.LogLevelNone)
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

	fmt.Println("all notifications:")
	for i := range allNotifications {
		fmt.Printf("\t%d: %+v\n", i, allNotifications[i])
	}

	fmt.Println("finished")
}
