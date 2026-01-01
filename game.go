package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/strawberry-productions/ebitengine-template/physics"
	"github.com/strawberry-productions/ebitengine-template/render"
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
