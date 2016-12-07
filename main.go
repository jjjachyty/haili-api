package main

import "github.com/lunny/tango"
import "haili/common"

func main() {
	t := tango.Classic()
	Router(t)
	t.Use(common.MyHandler())
	t.Run()
}
