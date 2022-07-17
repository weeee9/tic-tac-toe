package tictactoe

import (
	"image/color"
)

type symbol uint

const (
	symbolNon symbol = iota
	symbolX
	symbolO

	fontSize = 130
)

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
