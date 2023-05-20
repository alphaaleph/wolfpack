package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// stamp is a generalization of a cut sprite
type stamp struct {
	ctype characterType
	X, Y  float64
	speed float64
	image *ebiten.Image
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

// Stamp changes the character's image
func (s *stamp) Stamp(ct characterType) {
	s.ctype = ct
}
