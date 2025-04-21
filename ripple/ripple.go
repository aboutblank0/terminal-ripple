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

func UpdateRipples(ripples []*Ripple, grid [][]terminal.Cell, deltaTime float64, originalColor terminal.BackgroundColor) {
	rippleSpeed := 20.0
	threshold := 0.8

	for _, ripple := range ripples {
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
					grid[y][x].Color = ripple.Color
					ripple.LastAffected = append(ripple.LastAffected, [2]int{x, y})
				}
			}
		}
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
