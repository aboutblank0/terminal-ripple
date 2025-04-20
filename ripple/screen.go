package ripple

import (
	"aboutblank0/terminal-ripple/terminal"
	"fmt"
	"math/rand"
	"os"
)

type Screen struct {
	Enabled bool
	Width  int
	Height int
	Cells  [][]rune
}

func NewScreen(width, height int) *Screen {
	screen := new(Screen)
	screen.Width = width
	screen.Height = height

	cols := make([][]rune, height)
	for y := range height {
		cols[y] = make([]rune, width)
		for x := range cols[y] {
			cols[y][x] = ' '
		}
	}

	screen.Cells = cols
	return screen
}

func (s *Screen) SetRandomCell() {
	randColor := rand.Intn(terminal.DefaultColor)
	randY := rand.Intn(s.Height)
	randX := rand.Intn(s.Width)

	terminal.MoveCursor(randX, randY)
	terminal.SetBackgroundColor(terminal.BackgroundColor(randColor))
	fmt.Print(" ")
	os.Stdout.Sync()
}

func (s *Screen) Enable() {
	terminal.SaveCursor()
	terminal.SaveScreen()
	terminal.SetCursorInvisible()
}

func (s *Screen) Disable() {
	terminal.RestoreScreen()
	terminal.RestoreCursor()
	terminal.SetCursorVisible()
}
