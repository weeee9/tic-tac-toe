package tictactoe

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	numBlocks = 3
	// board has 3 x 3 blocks and 2 frame line in each of x, y axis
	boardSize = (blockSize * numBlocks) + (blockMargin * (numBlocks - 1))
)

func newBoardImage() *ebiten.Image {
	image := ebiten.NewImage(boardSize, boardSize)
	image.Fill(colorBlack)
	return image
}

type Board struct {
	round  symbol
	blocks [numBlocks][numBlocks]*Block
	winner symbol
	end    bool
}

func NewBoard() *Board {
	b := &Board{
		blocks: [numBlocks][numBlocks]*Block{},
	}

	// initial blocks
	for y := 0; y < numBlocks; y++ {
		for x := 0; x < numBlocks; x++ {
			t := NewBlock(symbolNon, x, y)
			b.blocks[y][x] = t
		}
	}

	b.setNextRound()

	return b
}

func (b *Board) Update(input *Input) {
	if !input.Pressed() {
		return
	}

	if !b.mousePositionInBoard(input) {
		return
	}

	if !b.canSet(input) {
		return
	}

	b.setBlock(input, b.round)

	b.checkState()
}

// Draw renders the board
// we draw blocks on the board and use the gap between blocks as frame line
func (b *Board) Draw(boardImage *ebiten.Image) {
	boardImage.Fill(frameColor)

	for y := range b.blocks {
		for x := range b.blocks[y] {
			b.blockAt(x, y).Draw(boardImage)
		}
	}
}

func (b *Board) Round() symbol {
	return b.round
}

func (b *Board) blockAt(x, y int) *Block {
	return b.blocks[y][x]
}

func (b *Board) setNextRound() {
	if b.round == symbolX {
		b.round = symbolO
		return
	}

	b.round = symbolX
}

func (b *Board) setBlock(input *Input, symbol symbol) {
	blockX, blockY := b.getBlockIndexFromInput(input)
	b.blocks[blockY][blockX].symbol = symbol
}

func (b *Board) size() (int, int) {
	x := boardSize
	y := x

	return x, y
}

func (b *Board) canSet(input *Input) bool {
	blockX, blockY := b.getBlockIndexFromInput(input)

	return b.blockAt(blockX, blockY).canSet()
}

func (b *Board) checkState() {
	// win or draw state
	if checkWin(b.round, b.blocks) {
		b.winner = b.round
		b.end = true
		return
	}

	// if checkDraw(b.blocks) {
	// }

	b.setNextRound()
}

func (b *Board) isEnd() bool {
	return b.end
}

func (b *Board) mousePositionInBoard(input *Input) bool {
	mouseX, mouseY := input.Pos()

	width, height := b.size()
	return (mouseX > 0 && mouseX < width) &&
		(mouseY > 0 && mouseY < height)
}

func (b *Board) getBlockIndexFromInput(input *Input) (int, int) {
	x, y := input.Pos()
	return b.getBlockIndexFromXY(x, y)
}

func (b *Board) getBlockIndexFromXY(x, y int) (int, int) {
	return x / blockSize, y / blockSize
}
