package terminal

import (
	"fmt"
)

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

const ESC = "\033"
func printAnsi(ansi string) { fmt.Printf("%s%s", ESC, ansi) }

func EraseScreen() { printAnsi("[2J") }

func SaveScreen() { printAnsi("[?47h") }
func RestoreScreen() { printAnsi("[?47l") }

func SaveCursor() { printAnsi("[s") }
func RestoreCursor() { printAnsi("[u") }
func MoveCursor(x, y int) { printAnsi(fmt.Sprintf("[%d;%dH", y, x)) }

func SetCursorInvisible() { printAnsi("[?25l") }
func SetCursorVisible() { printAnsi("[?25h") }

func SetForegroundColor(color BackgroundColor) { printAnsi(fmt.Sprintf("[38;5;%dm", color)) }
func SetBackgroundColor(color BackgroundColor) { printAnsi(fmt.Sprintf("[48;5;%dm", color)) }

func ResetAttributes() { printAnsi("[0m") }
