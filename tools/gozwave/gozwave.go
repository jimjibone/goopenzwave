package main

import (
	"flag"
	"fmt"
	"gitlab.com/jimjibone/gozwave"
	"os"
	"os/signal"
)

func processNotifications(m *gozwave.Manager) {
	for {
		select {
		case notification := <-m.Notifications:
			fmt.Println(notification)
		}
	}
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

	fmt.Println("finished")
}
