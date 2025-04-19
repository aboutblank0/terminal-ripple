package main

import (
	"aboutblank0/simple-gui/simplegui"
	"time"
)


func main() {
	app, err := simplegui.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Start()
	time.Sleep(time.Second * 2)
	app.Stop()
}
