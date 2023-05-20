package game

import (
	"github.com/alphaaleph/wolfpack/assets/graphics"
)

var WolfpackApp *Game

// Game implements ebiten.Game interface.
type Game struct {
	modeLevel modeType
	//keys      []ebiten.Key
	destroyer graphics.SpriteObject   // boat used by the player to destroy uboats
	u103      graphics.SpriteObject   // the boss uboat, appears when all uboats are destroyed
	wolfpack  []graphics.SpriteObject // a slice of uboat
}

// NewGame creates a single instance of the game
func NewGame() *Game {
	if WolfpackApp == nil {
		WolfpackApp = &Game{}
	}
	return WolfpackApp
}
