package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})

	SpidersApp.Player.Render(screen)
	SpidersApp.Spider.Render(screen)
}
