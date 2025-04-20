package main

import "aboutblank0/terminal-ripple/ripple"


func main() {
	app, err := ripple.NewApplication()
	if err != nil {
		panic(err)
	}

	err = app.Start()
	if err != nil {
		panic(err)
	}
}
