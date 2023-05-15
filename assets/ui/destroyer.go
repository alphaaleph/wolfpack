package ui

import (
	"github.com/alphaaleph/wolfpack/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

// destroyer represents the destroyer's sprite
type destroyer struct {
	stamp
	sunk bool
}

// NewDestroyer creates an instance of a new Destroyer
func NewDestroyer() Sprite {
	// create a new destroyer
	destroyerSprite := &destroyer{}
	destroyerSprite.image = New[*destroyer]().loadSprite(destroyerSprite)

	// destroyer
	destroyerSprite.X = util.ScreenWidth / 2.0
	destroyerSprite.Y = float64(util.DestroyerSection.Bounds().Dy() - destroyerSprite.image.Bounds().Size().Y)
	destroyerSprite.speed = 15 // TODO: create different speed for how hard the gae is
	return destroyerSprite
}

// getRect returns the destroyer's image dimensions
func (p *destroyer) getRect() image.Rectangle {
	var (
		topCorner    = image.Point{0, 0}
		bottomCorner = image.Point{63, 63}
	)
	return image.Rectangle{topCorner, bottomCorner}
}

// Render draws the destroyer
func (p *destroyer) Render(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(1.5, 1.5)
	w, h := p.image.Bounds().Dx(), p.image.Bounds().Dy()
	options.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)

	// readjust to use the middle of the destroyer
	x := p.X - float64(p.image.Bounds().Size().X/2)
	options.GeoM.Translate(x, p.Y)
	options.Filter = ebiten.FilterLinear
	screen.DrawImage(p.image, options)
}
