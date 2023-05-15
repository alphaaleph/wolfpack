package ui

import (
	"github.com/alphaaleph/wolfpack/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

var (
	BringTheWolf = false
)

type u103 struct {
	stamp
	sunk bool
}

// NewU103 creates an instance of a new U103 boss
func NewU103() Sprite {
	// create a new u103
	u103Sprite := &u103{}
	u103Sprite.image = New[*u103]().loadSprite(u103Sprite)

	// u103
	u103Sprite.X = util.ScreenWidth / 2.0
	u103Sprite.Y = float64(u103Sprite.image.Bounds().Dy()) * 1.5
	u103Sprite.speed = 15.0 // TODO: create different speed for how hard the game is
	return u103Sprite
}

// getRect returns the u103's image dimensions
func (s *u103) getRect() image.Rectangle {
	var (
		topCorner    = image.Point{64, 0}
		bottomCorner = image.Point{127, 63}
	)
	return image.Rectangle{topCorner, bottomCorner}
}

// Render draws the u103
func (s *u103) Render(screen *ebiten.Image) {
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
