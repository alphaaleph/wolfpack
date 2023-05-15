package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// stamp is a generalization of a cut sprite
type stamp struct {
	image *ebiten.Image
	X, Y  float64
	speed float64
}

// DecX decreases the X coordinate by the provide amount
func (s *stamp) DecX(x float64) {
	s.X -= x
}

// IncX increases the X coordinate by the provide amount
func (s *stamp) IncX(x float64) {
	s.X += x
}

// GetSpeed returns the stamp's speed
func (s *stamp) GetSpeed() float64 {
	return s.speed
}

// Render draws the stamp
func (s *stamp) Render(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(1.5, 1.5)
	w, h := s.image.Bounds().Dx(), s.image.Bounds().Dy()
	options.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)

	// readjust to use the middle of the player
	x := s.X - float64(s.image.Bounds().Size().X/2)
	options.GeoM.Translate(x, s.Y)
	options.Filter = ebiten.FilterLinear
	screen.DrawImage(s.image, options)
}
