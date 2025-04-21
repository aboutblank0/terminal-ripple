package terminal

import (
	"fmt"
)

type Screen struct {
	Enabled bool
	Width   int
	Height  int
	Cells   [][]Cell
}

type Cell struct {
	Color   BackgroundColor
	Content string
}

func newScreen(width, height int) *Screen {
	screen := new(Screen)
	screen.Width = width
	screen.Height = height

	cols := make([][]Cell, height)
	for y := range height {
		cols[y] = make([]Cell, width)
		for x := range cols[y] {
			cols[y][x] = Cell{Color: BlackColor, Content: string(' ')}
		}
	}

	screen.Cells = cols
	return screen
}

// TODO: Fix how many times this is called
// Possible fix would be to draw in order of COLORS
func (s *Screen) render() {
	MoveCursor(0, 0)
	for y := range s.Cells {
		for _, cell := range s.Cells[y] {
			SetBackgroundColor(cell.Color)
			fmt.Print(cell.Content)
		}
	}

	ResetAttributes()
}

func (s *Screen) enable() {
	SaveCursor()
	SaveScreen()
	SetCursorInvisible()
	EnableMouseTracking()
}

func (s *Screen) disable() {
	RestoreScreen()
	RestoreCursor()
	SetCursorVisible()
	DisableMouseTracking()
}
