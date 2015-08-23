package main

import (
	"flag"
	"fmt"
	"gitlab.com/jimjibone/gozwave"
	"os"
	"os/signal"
)

func watcherFunc(notification *gozwave.Notification, userdata interface{}) {
	notistring := notification.GetAsString()
	fmt.Println("gozwave watcher called with userdata:", userdata, "notification:", notistring)
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

	manager := gozwave.CreateManager()
	watcher, ok := manager.AddWatcher(watcherFunc, "beefs")
	if ok == false {
		fmt.Println("ERROR: failed to add watcher")
		return
	}

	manager.AddDriver(controllerPath)

	// Now wait for the user to quit.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	// All done now finish up.
	manager.RemoveDriver("/dev/tty.usbmodem411")
	manager.RemoveWatcher(watcher)
	gozwave.DestroyManager()
	gozwave.DestroyOptions()
	fmt.Println("finished")
}
