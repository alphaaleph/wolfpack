package main

import (
	"fmt"
	game "github.com/alphaaleph/wolfpack/game"
	"github.com/alphaaleph/wolfpack/util"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(util.ScreenWidth, util.ScreenHeight)
	ebiten.SetWindowTitle("W O L F P A C K")

	// call ebiten to start the game loop
	if err := ebiten.RunGame(game.WolfpackApp); err != nil {
		fmt.Println("initializing ebiten failed:", err)
	}
}
