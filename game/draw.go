package game

import (
	"fmt"
	"github.com/alphaaleph/wolfpack/assets/graphics"
	"github.com/alphaaleph/wolfpack/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
)

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {

	// paint the full screen
	screen.Fill(color.RGBA{0xe0, 0xe0, 0xe0, 0xff})

	// render the background
	util.DrawScoresBackground(screen)
	util.DrawDestroyerBackground(screen)
	util.DrawWolfpackBackgroundA(screen)
	util.DrawWolfpackBackgroundB(screen)
	util.DrawWolfpackBackgroundC(screen)

	// render the score
	const dpi = 72
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	mplusNormalFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
	msg := fmt.Sprintf("SCORE: %d", Score)
	text.Draw(screen, msg, mplusNormalFont, 20, 40, color.Black)

	// render the game while it is not over
	if g.modeLevel != ModeOver {
		// render the destroyer
		g.destroyer.Render(screen)

		if g.destroyer.HasExploded() {
			Score = Score + g.destroyer.GetPoints()
			g.modeLevel = ModeOver
		}

		// render the uboats, and when all are dead bring in the boss U103
		if graphics.BringTheWolf {
			if !g.u103.HasExploded() {
				g.u103.Render(screen)
			} else {
				// you kill the boss the games is over
				Score = Score + g.u103.GetPoints()
				g.modeLevel = ModeOver
			}
		} else {
			// render uboats as long as we have a wolfpack
			if len(g.wolfpack) > 0 {
				for i, uboat := range g.wolfpack {
					uboat.Render(screen)

					// if the uboat was exploded and
					if uboat.HasExploded() {

						// score
						Score = Score + uboat.GetPoints()

						// clear all exploded bobms
						g.destroyer.Reload()

						// deal with the uboats
						if uboat.StillHasLives() {
							uboat.Reset()
						} else {
							// remove dead uboats grouping by depth
							g.wolfpack = append(g.wolfpack[:i], g.wolfpack[i+1:]...)
						}
					}
				}
			} else {
				// when the wolfpack is destroyed bring the boss
				graphics.BringTheWolf = true
			}
		}
	} else {
		//show only the destroyer
		g.destroyer.Render(screen)
		g.u103.Render(screen)

		// show game over
		const dpi = 140
		tt, err := opentype.Parse(fonts.PressStart2P_ttf)
		if err != nil {
			log.Fatal(err)
		}
		gameOverFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
			Size:    36,
			DPI:     dpi,
			Hinting: font.HintingVertical,
		})
		if err != nil {
			log.Fatal(err)
		}
		msgOver := fmt.Sprintf("G A M E\n\n\nO V E R")
		text.Draw(screen, msgOver, gameOverFont, 50, util.ScreenHeight/3, color.RGBA{255, 127, 127, 1})
	}
}
