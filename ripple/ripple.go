package ripple

import (
	"aboutblank0/terminal-ripple/terminal"
	"os"

	"golang.org/x/term"
)

type RippleApp struct {
	Screen   *terminal.Screen
	Running  bool
	Fd       int
	oldState *term.State
}

func NewApplication() (app *RippleApp, err error) {
	width, height, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}

	screen := terminal.NewScreen(width, height)
	return &RippleApp{
		Screen:   screen,
		Fd:       int(os.Stdin.Fd()),
		Running:  false,
		oldState: nil,
	}, nil
}

func (app *RippleApp) Start() error {
	err := app.makeRaw()
	if err != nil {
		return err
	}

	app.Running = true
	for app.Running {
		input, err := app.ReadInput()
		if err != nil {
			break
		}

		app.handleInput(input)
	}
	return nil
}

func (app *RippleApp) Stop() {
	app.Running = false
	app.restore()
}

func (app *RippleApp) ReadInput() ([]byte, error) {
	buf := make([]byte, 1)
	_, err := os.Stdin.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (app *RippleApp) handleInput(input []byte) {
	if len(input) == 0 {
		return
	}

	switch input[0] {
	case 'q', 3: // 3 - Ctrl+c
		app.Stop()
	}
}

func (app *RippleApp) makeRaw() error {
	terminal.PrintAnsi("[?47h") // Save Screen
	terminal.PrintAnsi("[s")    //Save Cursor
	terminal.PrintAnsi("[?25l") //Cursor Invis

	oldState, err := term.MakeRaw(app.Fd)
	if err != nil {
		return err
	}

	app.oldState = oldState
	return nil
}

func (app *RippleApp) restore() error {
	err := term.Restore(app.Fd, app.oldState)

	terminal.PrintAnsi("[?47l") //Restore Screen
	terminal.PrintAnsi("[u")    //Restore Cursor
	terminal.PrintAnsi("[?25h") //Curosr Visible

	return err
}
