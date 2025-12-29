package main

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/strawberry-productions/ebitengine-template/physics"
	"github.com/strawberry-productions/ebitengine-template/render"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")

	render.NewText("hello").Draw(screen, &physics.Position{X: 25, Y: 25})
	render.NewImage("./assets/ebitengine.png").DrawSubImage(screen, &physics.Position{X: 25, Y: 50}, image.Rect(470, 100, 670, 280), math.Pi/8)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
