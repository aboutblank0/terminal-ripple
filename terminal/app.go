package terminal

import (
	"os"
	"time"

	"golang.org/x/term"
)

type TerminalApp struct {
	Running bool
	Screen  *Screen
	Fd      int
	Elements []AppElement
}

type AppElement interface {
	Update(screen *Screen, delta float64, input Input)
}

func NewApp() (*TerminalApp, error) {
	fd := int(os.Stdin.Fd())
	width, height, err := term.GetSize(fd)
	if err != nil {
		return nil, err
	}

	screen := newScreen(width, height)
	return &TerminalApp{
		Running: false,
		Screen:  screen,
		Fd:      fd,
		Elements: make([]AppElement, 0),
	}, nil
}

func (app *TerminalApp) AddElement(el AppElement) {
	app.Elements = append(app.Elements, el)
}

func (app *TerminalApp) Start() {
	app.Screen.enable()
	defer app.Screen.disable()

	//Raw terminal
	oldState, err := term.MakeRaw(app.Fd)
	if err != nil {
		panic(err)
	}
	defer term.Restore(app.Fd, oldState)

	//Run loop
	run(app)
}

func run(app *TerminalApp) {
	app.Running = true

	inputCh := getInputChannel()

	last := time.Now()
	for app.Running {
		input := getInput(inputCh)

		now := time.Now()
		delta := now.Sub(last).Seconds()
		last = now

		update(app, delta, input)

		app.Screen.render()
	}
}


func update(app *TerminalApp, delta float64, input Input) {
	if input.Key == 'q' {
		app.Running = false
	}

	for _, element := range app.Elements {
		element.Update(app.Screen, delta, input)
	}
}
