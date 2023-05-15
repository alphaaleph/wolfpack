package ui

import (
	"github.com/alphaaleph/wolfpack/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type torpedo struct {
	stamp
}

// NewTorpedo creates an instance of a new depth torpedo
func NewTorpedo() Sprite {
	// create a new torpedo
	torpedoSprite := &torpedo{}
	torpedoSprite.image = New[*torpedo]().loadSprite(torpedoSprite)

	// torpedo
	torpedoSprite.X = util.ScreenWidth / 2.0
	torpedoSprite.Y = float64(util.DestroyerSection.Bounds().Dy() - torpedoSprite.image.Bounds().Size().Y)
	torpedoSprite.speed = 15 // TODO: create different speed for how hard the gae is
	return torpedoSprite
}

// getRect returns the torpedo's image dimensions
func (c *torpedo) getRect() image.Rectangle {
	var (
		topCorner    = image.Point{0, 64}
		bottomCorner = image.Point{15, 15}
	)
	return image.Rectangle{topCorner, bottomCorner}
}

// Render draws the torpedo
func (c *torpedo) Render(screen *ebiten.Image) {
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
