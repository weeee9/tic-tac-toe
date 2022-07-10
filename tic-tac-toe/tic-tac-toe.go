package tictactoe

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 600
	screenHeight = 600
	size         = 3
)

type Game struct {
	input      *Input
	board      *Board
	boardImage *ebiten.Image
}

func NewGame() *Game {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello world!")
	return &Game{
		input: NewInput(),
		board: NewBoard(size),
	}
}

// Update updates the game's logical state
// and is called every 1/60 second tick
func (g *Game) Update() error {
	g.input.Update()

	g.board.Update(g.input)

	return nil
}

// Draw rendering screen object
// and is called every frame depends on the display's refresh rate
func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		w, h := g.board.Size()
		g.boardImage = ebiten.NewImage(w, h)
	}
	screen.Fill(backgroundColor)
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
