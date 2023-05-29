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
	// ammo
	//ammoSprite.X = util.ScreenWidth / 2.0
	//ammoSprite.Y = float64(util.DestroyerSectionTopY + ammoSprite.image.Bounds().Size().Y)
	//ammoSprite.speed = 15 // TODO: create different speed for how hard the gae is
	//return ammoSprite
}

// DecX decreases the X coordinate by the provide amount
func (a *ammo) DecX(x float64) {
	a.X -= x
}

// IncX increases the X coordinate by the provide amount
func (a *ammo) IncX(x float64) {
	a.X += x
}

// getRect returns an ammo rectangle coordinates
func (a *ammo) getRect() image.Rectangle {
	rect := image.Rectangle{
		Min: image.Point{int(a.X), int(a.Y)},
		Max: image.Point{int(a.X) + ammoPixelSide, int(a.Y) + ammoPixelSide},
	}
	return rect
}

// GetSpeed returns the ammo's speed
func (a *ammo) GetSpeed() float64 {
	return a.speed
}

// Sprite changes the character's image
func (a *ammo) Sprite(at ammoType) {
	a.atype = at
}

// getSeqNum returns the ammo sequence number
func (a *ammo) getSeqNum() int {
	fmt.Println("get sequence number: %d", a.seqNum)
	return a.seqNum
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

// Render draws the ammo
func (a *ammo) Render(screen *ebiten.Image) {
	// render the ammo
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(0.35, 0.35)

	if a.fired {
		a.Y = a.Y + destroyerAmmoSpeed
	}

	// render
	options.GeoM.Translate(a.X, a.Y)
	options.Filter = ebiten.FilterLinear
	screen.DrawImage(a.image, options)
}

// setExploded sets the ammo explosion flag so it is deleted from the pool
func (a *ammo) setExploded(e bool) {
	a.exploded = e
}

// HasExploded determines if the ammo has exploded
func (a *ammo) HasExploded() bool {
	return a.exploded
}

// getType returns the ammo type
func (a *ammo) getType() ammoType {
	return a.atype
}
