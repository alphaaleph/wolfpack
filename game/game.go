package game

import (
	"github.com/alphaaleph/spiders/img"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 600
	ScreenHeight = 800
)

var SpidersApp *Game

// Game implements ebiten.Game interface.
type Game struct {
	modeLevel modeType
	keys      []ebiten.Key
	img.Player
	img.Bullet
	img.Spider
	img.Minion
	img.Poison
}

// NewGame creates a single instance of the game
func NewGame() *Game {
	if SpidersApp == nil {
		SpidersApp = &Game{}
	}
	return SpidersApp
}
