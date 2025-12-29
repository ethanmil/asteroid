package render

import (
	"image"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/strawberry-productions/ebitengine-template/physics"

	_ "image/png"
)

type Image struct {
	image *ebiten.Image
}

func NewImage(path string) *Image {
	image, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal("failed to load ", err)
	}

	return &Image{image}
}

func (i *Image) Draw(screen *ebiten.Image, pos *physics.Position) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(pos.X), float64(pos.Y))

	screen.DrawImage(i.image, &op)
}

func (i *Image) DrawWithTransparency(screen *ebiten.Image, pos *physics.Position, a float64) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(pos.X), float64(pos.Y))
	op.ColorM.ChangeHSV(0, 0, 1)

	screen.DrawImage(i.image, &op)
}

func (i *Image) DrawWithSaturation(screen *ebiten.Image, pos *physics.Position, s float64) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(pos.X), float64(pos.Y))
	op.ColorM.ChangeHSV(0, s, 1)

	screen.DrawImage(i.image, &op)
}

func (i *Image) DrawSubImage(screen *ebiten.Image, pos *physics.Position, rect image.Rectangle, angle float64) {
	var (
		subImage = i.image.SubImage(rect).(*ebiten.Image)
		op       = ebiten.DrawImageOptions{}
	)

	if angle != 0 {
		var (
			deltaX = int(math.Abs(float64(rect.Max.X - rect.Min.X)))
			deltaY = int(math.Abs(float64(rect.Max.Y - rect.Min.Y)))
		)

		op.GeoM.Translate(-float64(deltaX)/2, -float64(deltaY)/2)
		op.GeoM.Rotate(angle)
		op.GeoM.Translate(float64(deltaX)/2, float64(deltaY)/2)
	}
	op.GeoM.Translate(float64(pos.X), float64(pos.Y))

	screen.DrawImage(subImage, &op)
}
