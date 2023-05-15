package ui

import (
	"github.com/alphaaleph/wolfpack/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"math/rand"
)

type uboat struct {
	stamp
	sunk bool
}

// NewUboat creates an instance of a new Uboat
func NewUboat() Sprite {
	// create a new uboat
	uboatSprite := &uboat{}
	uboatSprite.image = New[*uboat]().loadSprite(uboatSprite)

	// uboat
	uboatSprite.speed = 15.34 // TODO: create different speed for how hard the game is
	return uboatSprite
}

// getRect returns the uboat's image dimensions
func (b *uboat) getRect() image.Rectangle {
	var (
		topCorner    = image.Point{128, 0}
		bottomCorner = image.Point{181, 63}
	)
	return image.Rectangle{topCorner, bottomCorner}
}

// Render draws a uboat
func (b *uboat) Render(screen *ebiten.Image) {
	if !b.sunk {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Scale(1.5, 1.5)
		w, h := b.image.Bounds().Dx(), b.image.Bounds().Dy()
		options.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)

		// get the location of uboat
		var x, y float64
		//timer = time.AfterFunc(time.Second, func() {
		x, y = b.getRandomLocation()
		//})
		options.GeoM.Translate(x, y)
		options.Filter = ebiten.FilterLinear
		screen.DrawImage(b.image, options)
	}
}

// getRandomLocation picks random coordinates to display a uboat
func (b *uboat) getRandomLocation() (x float64, y float64) {
	x = float64(rand.Intn(util.ScreenWidth))
	y = float64(rand.Intn(util.ScreenHeight))
	return
}
