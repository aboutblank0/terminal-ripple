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

	screen := NewScreen(width, height)

	screen.Enable()
	defer screen.Disable()

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer term.Restore(fd, oldState)

	run(screen)
}

var running bool = false

func run(screen *Screen) {
	running = true

	inputChan := make(chan byte)
	go readInputLoop(inputChan)

	for running {
		handleInput(inputChan)

		screen.Render()

		//Shleep (~60 fps)
		time.Sleep(16 * time.Millisecond)
	}
}

func handleInput(inputChan <-chan byte) {
	select {
	case b := <-inputChan:
		switch b {
		case 'q', 3:
			running = false
		}
	default:
	}
}

var b [1]byte
func readInputLoop(ch chan<- byte) {
	for {
		_, err := os.Stdin.Read(b[:])
		if err == nil {
			ch <- b[0]
		}
	}
}
