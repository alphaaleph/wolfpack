package img

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Poison struct {
	Image *ebiten.Image
	X, Y  float64
	Speed float64
}

// getRect returns the poison's image dimensions
func (p *Poison) getRect() image.Rectangle {
	return image.Rect(16, 64, 31, 79)
}

// load sets the poison's image
func (p *Poison) load(image *ebiten.Image) {
	p.Image = image
}

// Render draws the poison
func (p *Poison) Render(screen *ebiten.Image) {
	screen.DrawImage(p.Image, nil)
}
