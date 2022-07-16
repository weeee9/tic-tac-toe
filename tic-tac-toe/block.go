package tictactoe

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

const (
	blockSize   = 150
	blockMargin = 5
)

type Block struct {
	symbol symbol
	x      int
	y      int
}

func NewBlockImage() *ebiten.Image {
	blockImage := ebiten.NewImage(blockSize, blockSize)
	blockImage.Fill(blockColor)

	return blockImage
}

func (t *Block) Value() symbol {
	return t.symbol
}

func (t *Block) canSet() bool {
	return t.symbol == symbolNon
}

func (t *Block) Draw(boardImage *ebiten.Image) {
	i, j := t.x, t.y

	op := &ebiten.DrawImageOptions{}
	x := i*blockSize + (i)*blockMargin
	y := j*blockSize + (j)*blockMargin

	op.GeoM.Translate(float64(x), float64(y))
	boardImage.DrawImage(NewBlockImage(), op)

	str := t.symbol.String()

	bound, _ := font.BoundString(mplusBigFont, str)
	strWitdh := (bound.Max.X - bound.Min.X).Ceil()
	strHeight := (bound.Min.Y - bound.Max.Y).Ceil()

	xOffset := bound.Min.X.Ceil()
	yOffset := bound.Max.Y.Ceil()

	x = x + (blockSize-strWitdh)/2 - xOffset
	y = y + (blockSize-strHeight)/2 - yOffset

	text.Draw(boardImage, str, mplusBigFont, x, y, t.symbol.Color())
}

func NewBlock(symbol symbol, x, y int) *Block {
	return &Block{
		symbol: symbol,
		x:      x,
		y:      y,
	}
}
