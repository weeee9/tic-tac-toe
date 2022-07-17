package tictactoe

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = boardSize
	screenHeight = boardSize + infoBarHeight
)

type Game struct {
	input        *Input
	board        *Board
	boardImage   *ebiten.Image
	infoBar      *InfoBar
	infoBarImage *ebiten.Image
}

func NewGame() *Game {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Tic Tac Toe")
	return &Game{
		input:   NewInput(),
		board:   NewBoard(),
		infoBar: NewInfoBar(),
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

	g.infoBar.Update(g.board)

	return nil
}

// Draw rendering screen object
// and is called every frame depends on the display's refresh rate
func (g *Game) Draw(screen *ebiten.Image) {
	screenWidth, _ := screen.Size()

	if g.boardImage == nil {
		g.boardImage = newBoardImage()
	}

	g.board.Draw(g.boardImage)

	drawBoardOpt := &ebiten.DrawImageOptions{}
	boardWidth, boardHeight := g.boardImage.Size()
	boardX := (screenWidth - boardWidth) / 2
	boardY := 0

	drawBoardOpt.GeoM.Translate(float64(boardX), float64(boardY))
	screen.DrawImage(g.boardImage, drawBoardOpt)

	if g.infoBarImage == nil {
		g.infoBarImage = newInfoBarImage()
	}

	g.infoBar.Draw(g.infoBarImage)
	drawInfoBarOpt := &ebiten.DrawImageOptions{}
	barWidth, _ := g.infoBarImage.Size()
	barX := (screenWidth - barWidth) / 2
	// bar is under board, so the y starts from board Y
	barY := boardHeight + blockMargin
	drawInfoBarOpt.GeoM.Translate(float64(barX), float64(barY))

	screen.DrawImage(g.infoBarImage, drawInfoBarOpt)
}

// Layout accepts an outside size, which is a window size on desktop
// and returns the game's logical screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) isEnd() bool {
	return g.board.isEnd()
}
