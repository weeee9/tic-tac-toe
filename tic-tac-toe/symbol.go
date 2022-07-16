package tictactoe

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type symbol uint

const (
	symbolNon symbol = iota
	symbolX
	symbolO

	fontSize = 130
)

var (
	mplusBigFont font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72

	mplusBigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (s symbol) String() string {
	switch s {
	case symbolX:
		return "X"
	case symbolO:
		return "O"
	}
	return ""
}

func (s symbol) Color() color.Color {
	switch s {
	case symbolNon:
		return blockColor
	case symbolX:
		// return colorBlue
		return symbolColor
	case symbolO:
		// return colorRed
		return symbolColor
	default:
		panic("unexpected symbol")
	}
}
