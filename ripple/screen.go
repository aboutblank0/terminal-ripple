package ripple

import (
	"aboutblank0/terminal-ripple/terminal"
	"fmt"
)

type Screen struct {
	Enabled bool
	Width  int
	Height int
	Cells  [][]*Cell
}

type Cell struct {
	Color terminal.BackgroundColor
	Content string
}

func NewScreen(width, height int) *Screen {
	screen := new(Screen)
	screen.Width = width
	screen.Height = height

	cols := make([][]*Cell, height)
	for y := range height {
		cols[y] = make([]*Cell, width)
		for x := range cols[y] {
			cols[y][x] = &Cell{ Color: terminal.BlackColor, Content: string(' ')}
		}
	}

	screen.Cells = cols
	return screen
}

func (s *Screen) Render() {
	terminal.MoveCursor(0, 0)
	for y := range s.Cells {
		for _, cell := range s.Cells[y] {
			terminal.SetBackgroundColor(cell.Color)
			fmt.Print(cell.Content)
		}
	}

	terminal.ResetAttributes()
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
