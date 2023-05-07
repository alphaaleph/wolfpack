package img

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Spider struct {
	Image *ebiten.Image
	X, Y  float64
	Speed float64
}

// getRect returns the spider's image dimensions
func (s *Spider) getRect() image.Rectangle {
	return image.Rect(64, 0, 127, 63)
}

// load sets the spider's image
func (s *Spider) load(image *ebiten.Image) {
	s.Image = image
}

// Render draws the spider
func (s *Spider) Render(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(1.5, 1.5)
	w, h := s.Image.Bounds().Dx(), s.Image.Bounds().Dy()
	options.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)

	// readjust to use the middle of the player
	x := s.X - float64(s.Image.Bounds().Size().X/2)
	options.GeoM.Translate(x, s.Y)
	options.Filter = ebiten.FilterLinear
	screen.DrawImage(s.Image, options)
}
