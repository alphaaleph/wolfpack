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
	msg := fmt.Sprintf("SCORE: %0.2f", ebiten.ActualTPS())
	text.Draw(screen, msg, mplusNormalFont, 20, 40, color.Black)

	// render the destroyer
	WolfpackApp.destroyer.Render(screen)

	// render the uboats, and when all are dead bring in the boss U103
	if graphics.BringTheWolf {
		WolfpackApp.u103.Render(screen)
	} else {
		// render uboats as long as we have a wolfpack
		if len(WolfpackApp.wolfpack) > 0 {
			for i, uboat := range WolfpackApp.wolfpack {
				uboat.Render(screen)

				// if the uboat was exploded and
				if uboat.HasExploded() {
					if uboat.StillHasLives() {
						uboat.Reset()
					} else {
						WolfpackApp.wolfpack = append(WolfpackApp.wolfpack[:i], WolfpackApp.wolfpack[i+1:]...)
					}
				}
			}
		} else {
			// when the wolfpack is destroyed bring the boss
			graphics.BringTheWolf = true
		}
	}
}
