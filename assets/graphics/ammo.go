package graphics

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

var (
	ammoPixelHalfSide = 32
	ammoPixelSide     = 64
)

type ammo struct {
	atype    ammoType
	seqNum   int
	exploded bool
	fired    bool
	X, Y     float64
	speed    float64
	image    *ebiten.Image
}

// newAmmo returns and instance of a new ammo
func newAmmo(seqNum int, at ammoType, speed float64) *ammo {
	a := &ammo{seqNum: seqNum, atype: at, speed: speed}
	a.image = newSpriteImpl[*ammo]().load(at.Int(), a)
	return a
}

// DecX decreases the X coordinate by the provide amount
func (a *ammo) DecX(x float64) {
	a.X -= x
}

// GetSpeed returns the ammo's speed
func (a *ammo) GetSpeed() float64 {
	return a.speed
}

// HasExploded determines if the ammo has exploded
func (a *ammo) HasExploded() bool {
	return a.exploded
}

// IncX increases the X coordinate by the provide amount
func (a *ammo) IncX(x float64) {
	a.X += x
}

// Render draws the ammo
func (a *ammo) Render(screen *ebiten.Image) {
	// render the ammo
	options := &ebiten.DrawImageOptions{}

	if !a.HasExploded() {
		switch a.atype {
		case deepCharge:
			a.Y = a.Y + destroyerAmmoSpeed
			options.GeoM.Scale(0.30, 0.30)
		case uboatTorpedo:
			a.Y = a.Y - uboatAmmoSpeed
		case u103Torpedo:
			a.Y = a.Y - u103AmmoSpeed
			options.GeoM.Scale(0.50, 0.50)
		}

		// render
		options.GeoM.Translate(a.X, a.Y)
		options.Filter = ebiten.FilterLinear
	}
	screen.DrawImage(a.image, options)
}

// Sprite changes the character's image
func (a *ammo) Sprite(at ammoType) {
	a.atype = at
}

// getRect returns an ammo rectangle coordinates
func (a *ammo) getRect() image.Rectangle {
	rect := image.Rectangle{
		Min: image.Point{int(a.X), int(a.Y)},
		Max: image.Point{int(a.X) + ammoPixelSide, int(a.Y) + ammoPixelSide},
	}
	return rect
}

// TODO: move to configuration
// getSpriteRect returns the ammo's image location in the sprite sheet
func (a *ammo) getSpriteRect(position int) image.Rectangle {
	switch ammoTypes[position] {
	case deepCharge:
		fmt.Println("getting deep charge rectangle")
		return deepChargeSprite
	case uboatTorpedo:
		fmt.Println("getting uboat torpedo rectangle")
		return uboatTorpedoSprite
	case u103Torpedo:
		fmt.Println("getting u103 torpedo rectangle")
		return u103TorpedoSprite
	default:
		return image.Rectangle{}
	}
}

// getSeqNum returns the ammo sequence number
func (a *ammo) getSeqNum() int {
	fmt.Println("get sequence number: ", a.seqNum)
	return a.seqNum
}

// getType returns the ammo type
func (a *ammo) getType() ammoType {
	return a.atype
}

// setExploded sets the ammo explosion flag so it is deleted from the pool
func (a *ammo) setExploded(e bool) {
	a.exploded = e
}
