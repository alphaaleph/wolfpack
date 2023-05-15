package game

import (
	"github.com/alphaaleph/wolfpack/assets/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

var WolfpackApp *Game

// Game implements ebiten.Game interface.
type Game struct {
	modeLevel modeType
	keys      []ebiten.Key
	destroyer ui.Sprite   // boat used by the player to destroy uboats
	charge    ui.Sprite   // dept charges shot by the destroyer
	u103      ui.Sprite   // the boss uboat, appears when all uboats are destroyed
	wolfpack  []ui.Sprite // a slice of uboat
	torpedo   ui.Sprite   // the torpedos shot by uboats and the u103 wolf
}

// NewGame creates a single instance of the game
func NewGame() *Game {
	if WolfpackApp == nil {
		WolfpackApp = &Game{}
	}
	return WolfpackApp
}
