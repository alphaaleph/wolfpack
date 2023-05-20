package graphics

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
func NewUboat() SpriteObject {
	// create a new uboat
	uboatSprite := &uboat{}
	uboatSprite.image = NewCharacterImpl[*uboat]().loadCharacterSprite(CharacterLeft, uboatSprite)

	// uboat
	uboatSprite.speed = 15.34 // TODO: create different speed for how hard the game is
	return uboatSprite
}

// GetRect returns the uboat's image location in the sprite sheet
func (u *uboat) GetRect(ct characterType) image.Rectangle {
	switch ct {
	case CharacterLeft:
		return uboatLeftStamp
	case CharacterRight:
		return uboatRightStamp
	case CharacterExplosion:
		return uboatExplosionStamp
	default:
		return image.Rectangle{}
	}
}

// Render draws a uboat
func (u *uboat) Render(screen *ebiten.Image) {
	if !u.sunk {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Scale(1.5, 1.5)
		w, h := u.image.Bounds().Dx(), u.image.Bounds().Dy()
		options.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)

		// get the location of uboat
		var x, y float64
		//timer = time.AfterFunc(time.Second, func() {
		x, y = u.getRandomLocation()
		//})
		options.GeoM.Translate(x, y)
		options.Filter = ebiten.FilterLinear
		screen.DrawImage(u.image, options)
	}
}

// getRandomLocation picks random coordinates to display a uboat
func (b *uboat) getRandomLocation() (x float64, y float64) {
	x = float64(rand.Intn(util.ScreenWidth))
	y = float64(rand.Intn(util.ScreenHeight))
	return
}
