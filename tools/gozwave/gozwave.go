package main

import (
	"fmt"
	"gitlab.com/jimjibone/gozwave"
	"os"
	"os/signal"
)

func main() {
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
	manager.AddWatcher()
	// func() {
	// 	fmt.Println("watcher callback was called")
	// })

	manager.AddDriver("/dev/tty.usbmodem411")

	// Now wait for the user to quit.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	manager.RemoveDriver("/dev/tty.usbmodem411")
	manager.RemoveWatcher()
	gozwave.DestroyManager()
	gozwave.DestroyOptions()
	fmt.Println("finished")
}
