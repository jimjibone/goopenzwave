package main

import (
	"fmt"
	"gitlab.com/jimjibone/gozwave"
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

	gozwave.CreateManager()

	gozwave.DestroyManager()
	gozwave.DestroyOptions()
	fmt.Println("finished")
}
