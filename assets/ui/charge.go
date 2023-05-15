package ui

import (
	"github.com/alphaaleph/wolfpack/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type charge struct {
	stamp
}

// NewCharge creates an instance of a new depth charge
func NewCharge() Sprite {
	// create a new charge
	chargeSprite := &charge{}
	chargeSprite.image = New[*charge]().loadSprite(chargeSprite)

	// charge
	chargeSprite.X = util.ScreenWidth / 2.0
	chargeSprite.Y = float64(util.DestroyerSection.Bounds().Dy() - chargeSprite.image.Bounds().Size().Y)
	chargeSprite.speed = 15 // TODO: create different speed for how hard the gae is
	return chargeSprite
}

// getRect returns the charge's image dimensions
func (c *charge) getRect() image.Rectangle {
	var (
		topCorner    = image.Point{0, 64}
		bottomCorner = image.Point{15, 15}
	)
	return image.Rectangle{topCorner, bottomCorner}
}

// Render draws the charge
func (c *charge) Render(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(1.5, 1.5)
	w, h := c.image.Bounds().Dx(), c.image.Bounds().Dy()
	options.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)

	// readjust to use the middle of the destroyer
	x := c.X - float64(c.image.Bounds().Size().X/2)
	options.GeoM.Translate(x, c.Y)
	options.Filter = ebiten.FilterLinear
	screen.DrawImage(c.image, options)
}
