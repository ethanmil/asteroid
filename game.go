package main

import (
	"github.com/ethanmil/asteroid/physics"
	"github.com/ethanmil/asteroid/render"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	render.NewImage("./assets/bg.png").Draw(screen, &physics.Position{X: 0, Y: 0})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
