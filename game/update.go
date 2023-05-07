package game

import "github.com/hajimehoshi/ebiten/v2/inpututil"

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {

	// check the correct game mode
	//switch g.modeLevel {
	//case ModeTitle:
	//	g.modeLevel = ModeGame //TODO: make sure to add title screen before switching to game mode
	//case ModeGame:
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	gameKeys()
	//case ModeOver:
	//}
	return nil
}
