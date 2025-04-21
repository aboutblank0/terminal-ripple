package ripple

import (
	"aboutblank0/terminal-ripple/terminal"
	"math"
)

type Ripple struct {
	CenterX, CenterY int
	Radius           float64
	ElapsedTime      float64
	LastAffected     [][2]int // List of (x, y) positions
	Color            terminal.BackgroundColor
}

func RenderRipples(ripples []*Ripple, screen *terminal.Screen, originalColor terminal.BackgroundColor) {
	threshold := 0.8
	for _, ripple := range ripples {
		for _, pos := range ripple.LastAffected {
			x, y := pos[0], pos[1]
			screen.Cells[y][x].Color = originalColor
		}

		// Reset previously affected cells
		ripple.LastAffected = ripple.LastAffected[:0] // Clear the list

		for y := range screen.Cells {
			for x := range screen.Cells[0] {
				dx := x - ripple.CenterX
				dy := y - ripple.CenterY

				scaledY := float64(dy) * 2 // shorter on y axis because terminal yes
				dist := math.Sqrt(float64(dx*dx) + scaledY*scaledY)

				r := math.Abs(dist - ripple.Radius)
				if r < threshold {
					screen.Cells[y][x].Color = ripple.Color
					ripple.LastAffected = append(ripple.LastAffected, [2]int{x, y})
				}
			}
		}
	}
}

func UpdateRipples(ripples []*Ripple, deltaTime float64) {
	rippleSpeed := 20.0

	for _, ripple := range ripples {
		ripple.ElapsedTime += deltaTime
		ripple.Radius = ripple.ElapsedTime * rippleSpeed
	}
}

func StartRipple(x, y int, color terminal.BackgroundColor) *Ripple {
	return &Ripple{
		CenterX:     x,
		CenterY:     y,
		Radius:      0,
		ElapsedTime: 0,
		Color:       color,
	}
}
