package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

const (
	mainTileSide = 128
)

var (
	gameSprites *ebiten.Image
)

// spriteObject defines the image functionality for both
type spriteObject interface {
	DecX(float64)
	IncX(float64)
	GetSpeed() float64
	Render(screen *ebiten.Image)
}

type SpriteCharacterObject interface {
	spriteObject
	Fire(bool)
	GetRect(characterType) image.Rectangle
	IsFiring() bool
	Character(characterType)
}

type SpriteAmmoObject interface {
	spriteObject
	GetRect(ammoType) image.Rectangle
}
