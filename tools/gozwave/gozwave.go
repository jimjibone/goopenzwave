package main

import (
	"fmt"
	"gitlab.com/jimjibone/gozwave"
)

func main() {
	fmt.Println("gozwave example starting")
	foo := gozwave.New()
	foo.Bar()
	foo.Free()
	fmt.Println("finished")
}
