package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
)

func main() {
	// set a game instance
	game := &Game{}

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("SPIDERS")

	// call ebiten to start the game loop
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
