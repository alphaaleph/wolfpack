package game

import (
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

		// left arrow key moves the player to the left
	} else if repeatingKeyPressed(ebiten.KeyLeft) {
		SpidersApp.Player.X -= SpidersApp.Player.Speed

		// right arrow key moves the player to the right
	} else if repeatingKeyPressed(ebiten.KeyRight) {
		SpidersApp.Player.X += SpidersApp.Player.Speed

		// space shoot bullets
	} else if repeatingKeyPressed(ebiten.KeySpace) {

	}
}

// repeatingKeyPressed return true when key is pressed considering the repeat state.
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

// function closeWindow closes the game when the Q key or Alt-F4 are pressed
func closeWindow() {
	os.Exit(0)
}
