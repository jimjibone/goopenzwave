package main

import (
	"fmt"
	"gitlab.com/jimjibone/gozwave"
)

func main() {
	fmt.Println("gozwave example starting")
	manager := gozwave.NewManager()
	manager.Bar()
	manager.Free()
	fmt.Println("finished")
}
