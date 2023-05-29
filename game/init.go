package game

import "github.com/alphaaleph/wolfpack/assets/graphics"

// initialize images
func init() {

	// init the game
	WolfpackApp = NewGame()
	WolfpackApp.modeLevel = ModeTitle

	// init each sprite used in the game
	WolfpackApp.destroyer = graphics.NewDestroyer()
	WolfpackApp.u103 = graphics.NewU103()
	WolfpackApp.wolfpack = graphics.NewWolfpack()
}
