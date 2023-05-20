package graphics

import (
	"github.com/alphaaleph/wolfpack/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type ammo struct {
	stamp
	loaded ammoType
	num    int
}

// NewAmmo creates an instance of a new ammo
func NewAmmo(at ammoType, num int) SpriteObject {
	// create a new ammo
	ammoSprite := &ammo{loaded: at, num: num}
	ammoSprite.image = NewAmmoImpl[*ammo]().loadAmmoSprite(0, ammoSprite)

	// ammo
	ammoSprite.X = util.ScreenWidth / 2.0
	ammoSprite.Y = float64(util.DestroyerSectionRect.Bounds().Dy() - ammoSprite.image.Bounds().Size().Y)
	ammoSprite.speed = 15 // TODO: create different speed for how hard the gae is
	return ammoSprite
}

// getNum returns the ammo number
func (a *ammo) getNum() int {
	return a.num
}

// GetRect returns the ammo's image location in the sprite sheet
func (a *ammo) GetRect(at ammoType) image.Rectangle {
	switch at {
	case deepCharge:
		return deepChargeStamp
	case uboatTorpedo:
		return uboatTorpedoStamp
	case u103Torpedo:
		return u103TorpedoStamp
	default:
		return image.Rectangle{}
	}
}

// Render draws the ammo
func (a *ammo) Render(screen *ebiten.Image) {
}
