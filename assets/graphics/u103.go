package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

var (
	BringTheWolf  = false
	u103AmmoPack  = 8 // 8 torpedos
	u103AmmoSpeed = 5.0
)

type u103 struct {
	character
	sunk bool
}

// NewU103 creates an instance of a new U103 boss
func NewU103() SpriteCharacterObject {
	// create a new destroyer
	u := &u103{
		character: character{
			ctype: CharacterLeft,
			speed: subSpeed[depthC], //TODO: pick the speed based on the depth that the u103 appears at
			X:     0,                //TODO: calculate the correct random location
			Y:     0,                //TODO: calculate the correct random location
		},
		sunk: false,
	}

	// load the ammo into the u103
	u.munition = u.getMunitionPool(u103Torpedo, u103AmmoSpeed, u103AmmoPack)

	return u
}

// GetRect returns the u103's image location in the sprite sheet
func (u *u103) GetRect(ct characterType) image.Rectangle {
	switch ct {
	case CharacterLeft:
		return u103LeftSprite
	case CharacterRight:
		return u103RightSprite
	case CharacterExplosion:
		return u103ExplosionSprite
	default:
		return image.Rectangle{}
	}
}

// Render draws the u103
func (u *u103) Render(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(1.5, 1.5)
	w, h := u.image.Bounds().Dx(), u.image.Bounds().Dy()
	options.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)

	// readjust to use the middle of the player
	x := u.X - float64(u.image.Bounds().Size().X/2)
	options.GeoM.Translate(x, u.Y)
	options.Filter = ebiten.FilterLinear
	screen.DrawImage(u.image, options)
}
