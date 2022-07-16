package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	tictactoe "github.com/weeee9/tic-tac-toe/tic-tac-toe"
)

type Game struct{}

func main() {
	if err := ebiten.RunGame(tictactoe.NewGame()); err != nil {
		log.Fatal(err)
	}
}
