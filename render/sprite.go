package render

import (
	"github.com/ethanmil/asteroid/physics"
	"github.com/ethanmil/asteroid/render"
)

type Sprite struct {
	image    *render.Image
	location *physics.Position
	size     *physics.Size
}

func NewSprite(sourceImage *Image, location *physics.Position, size *physics.Size) Sprite {
	return Sprite{
		image: sourceImage,
	}
}
