package game

import (
	"github.com/alphaaleph/wolfpack/assets/graphics"
)

var WolfpackApp *Game

// Game implements ebiten.Game interface.
type Game struct {
	modeLevel modeType
	//keys      []ebiten.Key
	destroyer graphics.SpriteCharacterObject   // boat used by the player to destroy uboats
	u103      graphics.SpriteCharacterObject   // the boss uboat, appears when all uboats are destroyed
	wolfpack  []graphics.SpriteCharacterObject // a slice of uboat
}

// NewGame creates a single instance of the game
func NewGame() *Game {
	if WolfpackApp == nil {
		WolfpackApp = &Game{}
	}
	return WolfpackApp
}
