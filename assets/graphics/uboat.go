package graphics

import (
	"github.com/alphaaleph/wolfpack/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"math/rand"
)

var (
	uboatAmmoPack  = 3 // 3 torpedos
	uboatAmmoSpeed = 2.0
)

type uboat struct {
	character
	sunk bool
}

// NewUboat creates an instance of a new Uboat
func NewUboat() SpriteCharacterObject {
	// create a new destroyer
	u := &u103{
		character: character{
			ctype: CharacterLeft,
			speed: subSpeed[depthC], //TODO: pick the correct speed based on the depth the sub is at
			X:     0,                //TODO: calculate the correct random location
			Y:     0,                //TODO: calculate the correct random location
		},
		sunk: false,
	}

	// load the ammo into the u103
	u.munition = u.getMunitionPool(uboatTorpedo, uboatAmmoSpeed, uboatAmmoPack)

	return u
}

// GetRect returns the uboat's image location in the sprite sheet
func (u *uboat) GetRect(ct characterType) image.Rectangle {
	switch ct {
	case CharacterLeft:
		return uboatLeftSprite
	case CharacterRight:
		return uboatRightSprite
	case CharacterExplosion:
		return uboatExplosionSprite
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
