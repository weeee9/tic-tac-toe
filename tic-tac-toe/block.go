package tictactoe

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	blockSize   = 150
	blockMargin = 10
)

var (
	blockImage = ebiten.NewImage(blockSize, blockSize)

	mplusBigFont font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72

	mplusBigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

type Block struct {
	current blockData
}

type blockData struct {
	symbol symbol
	x      int
	y      int
}

func (t *Block) Value() symbol {
	return t.current.symbol
}

func (t *Block) CanSet() bool {
	return t.current.symbol == symbolNon
}

func (t *Block) Draw(boardImage *ebiten.Image) {
	i, j := t.current.x, t.current.y
	op := &ebiten.DrawImageOptions{}
	x := i*blockSize + (i+1)*blockMargin
	y := j*blockSize + (j+1)*blockMargin

	op.GeoM.Translate(float64(x), float64(y))
	op.ColorM.ScaleWithColor(blockBackgroundColor())
	boardImage.DrawImage(blockImage, op)

	var str string
	switch t.current.symbol {
	case symbolX:
		str = "X"
	case symbolO:
		str = "O"
	}

	f := mplusBigFont
	bound, _ := font.BoundString(f, str)
	w := (bound.Max.X - bound.Min.X).Ceil()
	h := (bound.Max.Y - bound.Min.Y).Ceil()
	x = x + (blockSize-w)/2
	y = y + (blockSize-h)/2 + h
	text.Draw(boardImage, str, f, x, y, blockColor(t.current.symbol))
}

func NewBlock(symbol symbol, x, y int) *Block {
	return &Block{
		current: blockData{
			symbol: symbol,
			x:      x,
			y:      y,
		},
	}
}
