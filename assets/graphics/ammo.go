package graphics

import (
	"bytes"
	"fmt"
	"github.com/alphaaleph/wolfpack/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type ammo struct {
	atype  ammoType
	seqNum int
	fired  bool
	X, Y   float64
	speed  float64
	image  *ebiten.Image
}

// NewAmmo creates an instance of a new ammo
func NewAmmo(at ammoType, id int) SpriteAmmoObject {
	// create a new ammo
	ammoSprite := &ammo{atype: at, seqNum: id}
	ammoSprite.image = ammoSprite.loadAmmoSprite(0)

	// ammo
	ammoSprite.X = util.ScreenWidth / 2.0
	ammoSprite.Y = float64(util.DestroyerSectionRect.Bounds().Dy() - ammoSprite.image.Bounds().Size().Y)
	ammoSprite.speed = 15 // TODO: create different speed for how hard the gae is
	return ammoSprite
}

// DecX decreases the X coordinate by the provide amount
func (a *ammo) DecX(x float64) {
	a.X -= x
}

// IncX increases the X coordinate by the provide amount
func (a *ammo) IncX(x float64) {
	a.X += x
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

// GetRect returns the ammo's image location in the sprite sheet
func (a *ammo) GetRect(at ammoType) image.Rectangle {
	switch at {
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

// loadAmmoSprite loads a sprite for a particular weapon piece
func (a *ammo) loadAmmoSprite(at ammoType) *ebiten.Image {
	// decode the sprite string
	if gameSprites == nil {
		sprites, _, err := image.Decode(bytes.NewReader(byteSprites))
		if err != nil {
			fmt.Println("failed to read sprite string:", err)
			panic(err)
		}
		gameSprites = ebiten.NewImageFromImage(sprites)
	}

	// retrieve the struct character and load it into
	character := gameSprites.SubImage(a.GetRect(at)).(*ebiten.Image)
	return character
}
