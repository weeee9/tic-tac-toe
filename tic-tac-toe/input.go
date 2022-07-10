package tictactoe

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type mouseState int

const (
	mouseStateNone mouseState = iota
	mouseStatePressed
)

type Input struct {
	mouseState mouseState

	mouseX int
	mouseY int
}

func NewInput() *Input {
	return &Input{
		mouseState: mouseStateNone,
	}
}

func (i *Input) Update() {
	switch i.mouseState {
	case mouseStateNone:
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()

			i.mouseX = x
			i.mouseY = y
			i.mouseState = mouseStatePressed
		}
	case mouseStatePressed:
		i.mouseState = mouseStateNone
	}
}

func (i *Input) Pressed() bool {
	return i.mouseState == mouseStatePressed
}

func (i *Input) Pos() (int, int) {
	return i.mouseX, i.mouseY
}
