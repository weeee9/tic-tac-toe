package tictactoe

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	size   int
	round  symbol
	blocks [3][3]*Block
	end    bool
}

func NewBoard(size int) *Board {
	b := &Board{
		size:   size,
		blocks: [3][3]*Block{},
	}

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			t := NewBlock(symbolNon, x, y)
			b.blocks[y][x] = t
		}
	}

	return b
}

func (b *Board) blockAt(x, y int) *Block {
	return b.blocks[y][x]
}

func (b *Board) Next() {
	if b.round == symbolX {
		b.round = symbolO
		return
	}

	b.round = symbolX
}

func (b *Board) SetBlock(x, y int, symbol symbol) {
	if !b.blockAt(x, y).CanSet() {
		return
	}
	b.blocks[y][x].current.symbol = symbol
	b.Next()
}

func (b *Board) Size() (int, int) {
	x := b.size*blockSize + (b.size+1)*blockMargin
	y := x

	return x, y
}

func (b *Board) Draw(boardImage *ebiten.Image) {
	boardImage.Fill(frameColor)
	for j := 0; j < b.size; j++ {
		for i := 0; i < b.size; i++ {
			op := &ebiten.DrawImageOptions{}
			x := i*blockSize + (i+1)*blockMargin
			y := j*blockSize + (j+1)*blockMargin
			op.GeoM.Translate(float64(x), float64(y))
			op.ColorM.ScaleWithColor(blockBackgroundColor())
			boardImage.DrawImage(blockImage, op)
		}
	}
	for y := range b.blocks {
		for x := range b.blocks[y] {
			b.blocks[y][x].Draw(boardImage)
		}
	}
}

func (b *Board) checkState() {
	// win or draw state
	if 0 == 1 {
		b.end = true
	}

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			if b.blocks[y][x].current.symbol == symbolNon {
				return
			}
		}
	}
	b.end = true
}

func (b *Board) mouseInBoard(x, y int) bool {
	width, height := b.Size()

	boardMargin := (screenWidth - width) / 2

	return (x < width || x > boardMargin) &&
		(y < height || y > boardMargin)
}

func (b *Board) getIndexFromXY(x, y int) (int, int) {
	width, height := b.Size()

	boardLeftRightMargin := (screenWidth - width) / 2

	boardTopBottomMargin := (screenWidth - height) / 2
	x -= boardLeftRightMargin
	y -= boardTopBottomMargin

	return x / blockSize, y / blockSize
}

func (b *Board) Update(input *Input) {
	if b.end {
		return
	}

	if !input.Pressed() {
		return
	}

	x, y := input.Pos()
	if !b.mouseInBoard(x, y) {
		return
	}

	boardX, boardY := b.getIndexFromXY(x, y)
	b.SetBlock(boardX, boardY, b.round)

	b.checkState()
}
