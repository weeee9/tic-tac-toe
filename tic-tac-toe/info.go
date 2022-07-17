package tictactoe

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	infoBarWidth  = boardSize
	infoBarHeight = (blockSize * 2) / 3
)

func newInfoBarImage() *ebiten.Image {
	image := ebiten.NewImage(infoBarWidth, infoBarHeight)
	image.Fill(colorBlack)
	return image
}

type InfoBar struct {
	round  symbol
	winner symbol
	end    bool
}

func NewInfoBar() *InfoBar {
	return &InfoBar{
		round: symbolNon,
	}
}

func (info *InfoBar) Update(board *Board) {
	if board.isEnd() {
		info.winner = board.winner
		info.end = true
		return
	}
	info.round = board.round
}

func (info *InfoBar) Draw(infoBarImage *ebiten.Image) {
	msg := info.String()

	infoBarImage.Clear()

	x := 25
	y := infoBarHeight / 2

	text.Draw(infoBarImage, msg, mplusNormalFont, x, y, symbolColor)
}

func (info *InfoBar) String() string {
	if info.end {
		return fmt.Sprintf("Player %v win!", info.winner)
	}

	return fmt.Sprintf("Round: %v", info.round)
}
