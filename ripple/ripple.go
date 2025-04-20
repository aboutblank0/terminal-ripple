package ripple

import (
	"os"
	"time"

	"golang.org/x/term"
)

func Start() {
	fd := int(os.Stdin.Fd())
	width, height, err := term.GetSize(fd)
	if err != nil {
		panic(err)
	}
	
	//Screen
	screen := NewScreen(width, height)
	screen.Enable()
	defer screen.Disable()
	
	//Raw terminal
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer term.Restore(fd, oldState)
	
	//Run loop
	run(screen)
}

var running bool = false

func run(screen *Screen) {
	running = true

	inputCh := getInputChannel()

	last := time.Now()
	for running {

		input := getInput(inputCh)

		now := time.Now()
		delta := now.Sub(last).Seconds()
		last = now

		update(delta, input)

		screen.Render()
	}
}

func update(_ float64, input Input) {
	if input.Key == 'q' {
		running = false
	}
}
