package ship

import (
	"github.com/ethanmil/asteroid/physics"
	"github.com/ethanmil/asteroid/render"
)

type Kind string

const (
	Player Kind = "player"
)

type Ship struct {
	Positon *physics.Position
	Size    *physics.Size

	sprite render.Sprite
}

func NewShip(position *physics.Position, kind Kind) *Ship {
	return &Ship{
		position,
		getSize(kind),
		getSprite(kind),
	}
}

func getSize(kind Kind) *physics.Size {
	switch kind {
	case Player:
		return &physics.Size{Width: 8, Height: 8}
	}

	panic("unknown ship type")
}

func getSprite(kind Kind) *render.Sprite {
	switch kind {
	case Player:
		return render.NewSprite()
	}

	panic("unknown ship type")
}
