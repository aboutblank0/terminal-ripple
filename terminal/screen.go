package terminal

import "fmt"

type Screen struct {
	Width  int
	Height int
	Cells  [][]rune
}

func NewScreen(width, height int) *Screen {
	screen := new(Screen)
	screen.Width = width
	screen.Height = height

	cols := make([][]rune, height)
	for x := range height {
		cols[x] = make([]rune, width)
		for i := range cols[x] {
			cols[x][i] = ' '
		}
	}

	screen.Cells = cols
	return screen
}

func (s *Screen) Render() {
	for _, rows := range s.Cells {
		for _, char := range rows {
			fmt.Printf("%c", char)
		}
	}
}
