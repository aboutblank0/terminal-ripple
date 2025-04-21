package main

import (
	"aboutblank0/terminal-ripple/ripple"
	"aboutblank0/terminal-ripple/terminal"
)

type Game struct {
	Ripples []*ripple.Ripple
}

func (g *Game) Update(delta float64, input terminal.Input) {
	if input.Mouse != nil && input.Mouse.Button == 3 {
		newRipple := ripple.StartRipple(input.Mouse.X, input.Mouse.Y, terminal.GetRandomColor())
		g.Ripples = append(g.Ripples, newRipple)
	}

	ripple.UpdateRipples(g.Ripples, delta)
}

func (g *Game) Render(screen *terminal.Screen) {
	ripple.RenderRipples(g.Ripples, screen, terminal.BlackColor)
}

func main() {
	app, err := terminal.NewApp()
	if err != nil {
		panic(err)
	}

	game := new(Game)
	app.AddElement(game)

	app.Start()
}
