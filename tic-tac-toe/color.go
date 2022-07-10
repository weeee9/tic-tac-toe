package tictactoe

import "image/color"

var (
	backgroundColor = color.RGBA{0xfa, 0xf8, 0xef, 0xff}
	frameColor      = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
)

func blockColor(symbol symbol) color.Color {
	switch symbol {
	case symbolNon:
		return color.RGBA{230, 230, 230, 255}
	case symbolX:
		return color.RGBA{40, 205, 40, 180}
	case symbolO:
		return color.RGBA{255, 0, 0, 180}
	default:
		panic("unexpected symbol")
	}
}

func blockBackgroundColor() color.Color {
	return color.RGBA{230, 230, 230, 255}
}
