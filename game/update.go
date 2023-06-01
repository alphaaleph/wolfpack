package game

import (
	"fmt"
	"github.com/alphaaleph/wolfpack/assets/graphics"
)

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {

	// check the correct game mode
	switch g.modeLevel {
	case ModeTitle:
		g.modeLevel = ModeGame //TODO: make sure to add title screen before switching to game mode
	case ModeGame:
		//check which key was selected
		gameKeys()

		// check get all the fired munitions
		var munitions []graphics.SpriteAmmoObject
		munitions = append(munitions, WolfpackApp.destroyer.GetFiredMunitions()...)

		if graphics.BringTheWolf {
			munitions = append(munitions, WolfpackApp.u103.GetFiredMunitions()...)
		} else {
			for _, uboat := range WolfpackApp.wolfpack {
				munitions = append(munitions, uboat.GetFiredMunitions()...)
			}
		}

		// only check for collisions if munitions have been fired
		if len(munitions) > 0 {
			fmt.Println("-------------- MUNITIONS ------------ FIRED ---------------------")
			// get all the characters
			var characters []graphics.SpriteCharacterObject
			characters = append(characters, WolfpackApp.destroyer)

			if graphics.BringTheWolf {
				characters = append(characters, WolfpackApp.u103)
			} else {
				characters = append(characters, WolfpackApp.wolfpack...)
			}

			//check for collisions
			graphics.CheckCollision(characters, munitions)
		}
	case ModeOver:
		// do cleaning
	}
	return nil
}
