package render

import (
	"bytes"
	"image/color"
	"log"

	"github.com/ethanmil/asteroid/physics"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	defaultFontSize = 12
	defaultFont     = getArcadeFont(defaultFontSize)
	defaultColor    = color.White
)

type Text struct {
	text string

	fontSize int
	fontFace *text.GoTextFaceSource
	col      color.Color
}

func NewText(text string) *Text {
	return &Text{
		text: text,

		fontSize: defaultFontSize,
		fontFace: defaultFont,
		col:      color.Color(defaultColor),
	}
}

func (t *Text) GetXOffsetToCenter() int {
	return t.fontSize * len(t.text) / 2
}

func (t *Text) Draw(screen *ebiten.Image, pos *physics.Position) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(pos.X), float64(pos.Y))
	op.ColorScale.ScaleWithColor(t.col)
	text.Draw(screen, t.text, &text.GoTextFace{
		Source: t.fontFace,
		Size:   float64(t.fontSize),
	}, op)
}

func getArcadeFont(size int) *text.GoTextFaceSource {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}
	return s
}
