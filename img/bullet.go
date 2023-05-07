package img

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Bullet struct {
	image *ebiten.Image
	x, y  float64
}

// getRect returns the bullet's image dimensions
func (b *Bullet) getRect() image.Rectangle {
	return image.Rect(0, 64, 15, 15)
}

// load sets the bullet's image
func (b *Bullet) load(img *ebiten.Image) {
	b.image = img
}

// Render draws the bullet
func (b *Bullet) Render(screen *ebiten.Image) {
	screen.DrawImage(b.image, nil)
}
