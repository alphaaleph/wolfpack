package game

import "github.com/alphaaleph/wolfpack/assets/ui"

// initialize images
func init() {

	// init the game
	WolfpackApp = NewGame()
	WolfpackApp.modeLevel = ModeTitle

	// init each sprite used in the game
	WolfpackApp.destroyer = ui.NewDestroyer()
	WolfpackApp.u103 = ui.NewU103()
	//TODO: WolfpackApp.wolfpack = ui.NewWolfpack()
}
