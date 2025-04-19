package simplegui

import (
	"os"

	"golang.org/x/term"
)

type Application struct {
	Screen *Screen
}

func NewApplication()(app *Application, err error) {
	width, height, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}

	screen := NewScreen(width, height)
	return &Application{
		Screen: screen,
	}, nil
}

func (app *Application) Start() {
	saveState()

	SetBackgroundColor(YellowColor)
	app.Screen.Render()
}

func (app *Application) Stop() {
	restoreState()
}

func saveState() {
	SendAnsi("[?47h") // Save Screen
	SendAnsi("[s")    //Save Cursor
}

func restoreState() {
	SendAnsi("[?47l") //Restore Screen
	SendAnsi("[u")    //Restore Cursor
}
