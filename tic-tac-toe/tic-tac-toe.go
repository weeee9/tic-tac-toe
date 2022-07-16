package tictactoe

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	// 3 x 3 board
	numBlocks    = 3
	screenWidth  = (blockSize * numBlocks) + (blockMargin * (numBlocks - 1))
	screenHeight = screenWidth
)

type Game struct {
	input      *Input
	board      *Board
	boardImage *ebiten.Image
}

func NewGame() *Game {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Tic Tac Toe")
	return &Game{
		input: NewInput(),
		board: NewBoard(),
	}
}

// Update updates the game's logical state
// and is called every 1/60 second tick
func (g *Game) Update() error {
	if g.isEnd() {
		return nil
	}

	g.input.Update()

	g.board.Update(g.input)

	return nil
}

// Draw rendering screen object
// and is called every frame depends on the display's refresh rate
func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		g.boardImage = newBoardImage()
	}

	g.board.Draw(g.boardImage)

	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Size()
	bw, bh := g.boardImage.Size()
	x := (sw - bw) / 2
	y := (sh - bh) / 2

	op.GeoM.Translate(float64(x), float64(y))

	screen.DrawImage(g.boardImage, op)
}

// Layout accepts an outside size, which is a window size on desktop
// and returns the game's logical screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) isEnd() bool {
	return g.board.isEnd()
}
