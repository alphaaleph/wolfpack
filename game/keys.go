package game

import (
	"github.com/alphaaleph/wolfpack/assets/graphics"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"os"
)

// gameKeys checks which keys have been pressed and executes a command accordingly during the game mode
func gameKeys() {
	// Q or Alt-F4 - close the game window

	if inpututil.IsKeyJustPressed(ebiten.KeyQ) ||
		((inpututil.IsKeyJustPressed(ebiten.KeyAlt)) && inpututil.IsKeyJustPressed(ebiten.KeyF4)) {
		closeWindow()
	}

	// left arrow key moves the destroyer to the left
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		WolfpackApp.destroyer.Character(graphics.CharacterLeft)
		WolfpackApp.destroyer.DecX(WolfpackApp.destroyer.GetSpeed())
	}

	// right arrow key moves the destroyer to the right
	if ebiten.IsKeyPressed(ebiten.KeyRight) {

		WolfpackApp.destroyer.Character(graphics.CharacterRight)
		WolfpackApp.destroyer.IncX(WolfpackApp.destroyer.GetSpeed())
	}

	// space drops depth charges
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		WolfpackApp.destroyer.Fire(true)
	}
}

// function closeWindow closes the game when the Q key or Alt-F4 are pressed
func closeWindow() {
	os.Exit(0)
}
