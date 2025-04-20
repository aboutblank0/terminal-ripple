package terminal

import (
	"fmt"
)

const ESC = "\033"

type BackgroundColor int

const (
	BlackColor   BackgroundColor = 0
	RedColor                     = 1
	GreenColor                   = 2
	YellowColor                  = 3
	BlueColor                    = 4
	MagentaColor                 = 5
	CyanColor                    = 6
	WhiteColor                   = 7
	DefaultColor                 = 8
)

func PrintAnsi(ansi string) {
	fmt.Printf("%s%s", ESC, ansi)
}

func SetBackgroundColor(color BackgroundColor) {
	EraseScreen()
	s := fmt.Sprintf("[48;5;%vm", color)
	PrintAnsi(s)
}

func EraseScreen() {
	PrintAnsi("[2J")
}

