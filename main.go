package main

import (
	"fmt"
	game "github.com/alphaaleph/spiders/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("S P I D E RS")

	// call ebiten to start the game loop
	if err := ebiten.RunGame(game.SpidersApp); err != nil {
		fmt.Println("initializing ebiten failed:", err)
	}
}
