package img

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Player struct {
	Image *ebiten.Image
	X, Y  float64
	Speed float64
}

// getRect returns the player's image dimensions
func (p *Player) getRect() image.Rectangle {
	return image.Rect(0, 0, 63, 63)
}

// load sets the player's image
func (p *Player) load(image *ebiten.Image) {
	p.Image = image
}

// Render draws the player
func (p *Player) Render(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(1.5, 1.5)
	w, h := p.Image.Bounds().Dx(), p.Image.Bounds().Dy()
	options.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)

	// readjust to use the middle of the player
	x := p.X - float64(p.Image.Bounds().Size().X/2)
	options.GeoM.Translate(x, p.Y)
	options.Filter = ebiten.FilterLinear
	screen.DrawImage(p.Image, options)
}
