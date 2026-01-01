package main

import (
	"github.com/ethanmil/asteroid/physics"
	"github.com/ethanmil/asteroid/render"
)

type Sprite struct {
	image    *render.Image
	location *physics.Position
	size     *physics.Size
}

func NewSprite(sourceImage *render.Image, location *physics.Position, size *physics.Size) Sprite {
	return Sprite{
		image: sourceImage,
	}
}

type ShipKind string

const (
	ShipPlayer ShipKind = "player"
)

type Ship struct {
	Positon *physics.Position
	Size    *physics.Size

	sprite Sprite
}

func NewShip(position physics.Position, kind ShipKind) *Ship {
	return &Ship{
		position,
		getSize(kind),
		NewSprite(kind),
	}
}

func getSize(kind ShipKind) *physics.Size {
	switch kind {
	case ShipPlayer:
		return &physics.Size{Width: 8, Height: 8}
	}

	panic("unknown ship type")
}

func getImage(kind ShipKind) *render.Image {
	switch kind {
	case ShipPlayer:
		return render.NewImage()
	}

	panic("unknown ship type")
}
