package graphics

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

var (
	wolfpack       = 3 // 3 uboats will show at once TODO: move to configuration
	uboatsLives    = 3 // each uboat has 3 lives
	uboatAmmoPack  = 2 // 3 torpedos
	uboatAmmoSpeed = 2.0
)

type uboat struct {
	character
	wolvesInfo
	lives int
}

// NewWolfpack returns a slice of new uboats
func NewWolfpack() []SpriteCharacterObject {
	//create a new map
	var uboats []SpriteCharacterObject

	for i := 0; i < wolfpack; i++ {
		uboats = append(uboats, NewUboat(depthTypes[i]))
	}
	return uboats
}

// NewUboat creates an instance of a new Uboat
func NewUboat(dt depthType) SpriteCharacterObject {
	// create a new destroyer
	u := &uboat{
		character: character{
			exploded: false,
		},
		lives: uboatsLives,
	}

	// get sub info
	u.cImageType = u.getRandomDirection()
	u.dtype = dt
	u.X, u.Y = u.getEntryLocation(u.cImageType, u.dtype)
	u.speed = subSpeed[u.dtype]

	// load the uboat sprites ahead of time
	u.leftImage = newSpriteImpl[*uboat]().load(0, u)
	u.rightImage = newSpriteImpl[*uboat]().load(1, u)
	u.explodeImage = newSpriteImpl[*uboat]().load(2, u)

	// load the ammo into the u103
	u.munition = u.getMunitionPool(uboatTorpedo, uboatAmmoSpeed, uboatAmmoPack)

	fmt.Printf("uboat:  ---------------------\n ctype: %v\n dtype: %v\n fire: %v\n exploded: %v\n"+
		"X: %f\n Y: %f\n speed: %v\n image: %v\n leftImage: %v\n rightImage: %v\n explodedImage: %v\n"+
		"munition: %v\n", u.cImageType, u.dtype, u.fire, u.exploded, u.X, u.Y, u.speed, u.image, u.leftImage,
		u.rightImage, u.explodeImage, u.munition)

	return u
}

// TODO: move to configuration
// getSpriteRect returns the uboat's image location in the sprite sheet
func (u *uboat) getSpriteRect(position int) image.Rectangle {
	switch characterTypes[position] {
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

	options := &ebiten.DrawImageOptions{}

	if !u.exploded {
		// move the uboat
		u.X = u.X + (u.speed * float64(u.getMovementMultiplier(u.cImageType)))
		fmt.Printf("rendering uboat with X: %v and Y: %v\n", u.X, u.Y)

		// check the location in case it has to be reset
		u.checkLocation(&u.cImageType, &u.X)
	} else {
		u.cImageType = CharacterExplosion
		u.lives--
	}

	// render the uboat
	u.image = u.getSprite(u.cImageType)
	options.GeoM.Scale(0.75, 0.75)
	options.GeoM.Translate(u.X, u.Y)
	options.Filter = ebiten.FilterLinear
	screen.DrawImage(u.image, options)
}

// StillHasLives returns true if the uboat has additional lives
func (u *uboat) StillHasLives() bool {
	return u.lives > 0
}

// Reset sets the uboat back to a playing state
func (u *uboat) Reset() {
	u.exploded = false
	u.image = u.leftImage
}
