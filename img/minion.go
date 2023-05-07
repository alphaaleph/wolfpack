package img

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Minion struct {
	Image *ebiten.Image
	X, Y  float64
	Speed float64
}

// getRect returns the minion's image dimensions
func (m *Minion) getRect() image.Rectangle {
	return image.Rect(128, 0, 181, 63)
}

// load sets the minion's image
func (m *Minion) load(image *ebiten.Image) {
	m.Image = image
}

// Render draws the player
func (m *Minion) Render(screen *ebiten.Image) {
	screen.DrawImage(m.Image, nil)
}
