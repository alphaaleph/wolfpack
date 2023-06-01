package graphics

import (
	"fmt"
	"github.com/alphaaleph/wolfpack/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

var (
	BringTheWolf  = false
	u103AmmoPack  = 5 // 5 torpedos
	u103AmmoSpeed = 5.0
)

type u103 struct {
	character
	wolvesInfo
}

// NewU103 creates an instance of a new U103 boss
func NewU103() SpriteCharacterObject {
	// create a new destroyer
	u := &u103{
		character: character{
			exploded: false,
		},
	}

	// get sub info
	u.cImageType = u.getRandomDirection()
	u.dtype = u.getRandomDepth()
	u.X, u.Y = u.getEntryLocation(u.cImageType, u.dtype)
	u.speed = subSpeed[u.dtype]

	// load the u103 sprites ahead of time
	u.leftImage = newSpriteImpl[*u103]().load(0, u)
	u.rightImage = newSpriteImpl[*u103]().load(1, u)
	u.explodeImage = newSpriteImpl[*u103]().load(2, u)

	// load the ammo into the u103
	u.munition = u.getMunitionPool(u103Torpedo, u103AmmoSpeed, u103AmmoPack)

	return u
}

// GetPoints returns the u103 points
func (u *u103) GetPoints() int {
	return u103Points
}

// Render draws the u103
func (u *u103) Render(screen *ebiten.Image) {

	// render the u103
	options := &ebiten.DrawImageOptions{}

	if !u.exploded {
		// move the uboat
		u.X = u.X + (u.speed * float64(u.getMovementMultiplier(u.cImageType)))
		fmt.Printf("rendering uboat with X: %v and Y: %v\n", u.X, u.Y)

		// check the location in case it has to be reset
		u.checkLocation(&u.cImageType, &u.X)
		_, u.Y = u.getEntryLocation(u.cImageType, u.dtype)

		// check when to fire
		if len(u.munition.idle) > 0 {
			if !u.fire {
				u.Fire(true)
				u.getFireLocation()
			}

			if (u.cImageType == CharacterLeft && u.X <= u.fireCoord) || (u.cImageType == CharacterRight && u.X >= u.fireCoord) {
				if u.fire {
					u.Fire(false)
					if torpedo, _ := u.munition.fire(); torpedo != nil {
						torpedo.X = u.X
						torpedo.Y = u.Y
					}
				}
			}
		}

		// check if there is an active ammo and render
		for _, munition := range u.munition.active {
			// when the deep charge goes pass the bottom of the screen reload it into the pool
			if munition.Y <= float64(util.DestroyerSectionBottomY) {
				u.munition.reload(munition)
			} else {
				munition.Render(screen)
			}
		}
	} else {
		u.cImageType = CharacterExplosion
	}

	// render the u103
	u.image = u.getSprite(u.cImageType)
	options.GeoM.Scale(0.75, 0.75)
	options.GeoM.Translate(u.X, u.Y)
	options.Filter = ebiten.FilterLinear
	screen.DrawImage(u.image, options)
}

// TODO:  move to configuration
// getSpriteRect returns the u103's image location in the sprite sheet
func (u *u103) getSpriteRect(positoin int) image.Rectangle {
	switch characterTypes[positoin] {
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
