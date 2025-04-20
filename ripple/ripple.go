package ripple

import (
	"aboutblank0/terminal-ripple/terminal"
	"math"
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

	//Screen
	screen := NewScreen(width, height)
	screen.Enable()
	defer screen.Disable()

	//Raw terminal
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer term.Restore(fd, oldState)

	//Run loop
	run(screen)
}

var running bool = false

func run(screen *Screen) {
	running = true

	inputCh := getInputChannel()

	last := time.Now()
	for running {

		input := getInput(inputCh)

		now := time.Now()
		delta := now.Sub(last).Seconds()
		last = now

		update(screen, delta, input)

		screen.Render()
	}
}

func update(screen *Screen, delta float64, input Input) {
	if input.Key == 'q' {
		running = false
	}

	if input.Mouse != nil && input.Mouse.Button == 3 {
		startRipple(input.Mouse.X, input.Mouse.Y)
	}

	updateRipples(screen.Cells, delta, terminal.BlackColor)
}

type Ripple struct {
	CenterX, CenterY int
	Radius           float64
	ElapsedTime      float64
	LastAffected     [][2]int // List of (x, y) positions
}

func updateRipples(grid [][]Cell, deltaTime float64, originalColor terminal.BackgroundColor) {

	rippleSpeed := 20.0
	threshold := 0.8

	for i := range ripples {
		ripple := &ripples[i]

		// Reset previously affected cells
		for _, pos := range ripple.LastAffected {
			x, y := pos[0], pos[1]
			grid[y][x].Color = originalColor
		}
		ripple.LastAffected = ripple.LastAffected[:0] // Clear the list

		// Update ripple
		ripple.ElapsedTime += deltaTime
		ripple.Radius = ripple.ElapsedTime * rippleSpeed

		// Apply new ring
		for y := range grid {
			for x := range grid[0] {
				dx := x - ripple.CenterX
				dy := y - ripple.CenterY

				scaledY := float64(dy) * 2 // shorter on y axis because terminal yes
				dist := math.Sqrt(float64(dx*dx) + scaledY*scaledY)

				r := math.Abs(dist - ripple.Radius)
				if r < threshold {
					grid[y][x].Color = terminal.BlueColor
					ripple.LastAffected = append(ripple.LastAffected, [2]int{x, y})
				}
			}
		}
	}
}

var ripples = make([]Ripple, 0)
func startRipple(x, y int) {
	ripples = append(ripples, Ripple{
		CenterX:     x,
		CenterY:     y,
		Radius:      0,
		ElapsedTime: 0,
	})
}
