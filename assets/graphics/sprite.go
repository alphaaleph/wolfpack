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
	getRect() image.Rectangle
	getSpriteRect(int) image.Rectangle
	setExploded(bool)
	GetSpeed() float64
	HasExploded() bool
	DecX(float64)
	IncX(float64)
	Render(screen *ebiten.Image)
}

type SpriteCharacterObject interface {
	spriteObject
	isFiring() bool
	Fire(bool)
	Character(characterImageType)
	GetFiredMunitions() []SpriteAmmoObject
}

type SpriteAmmoObject interface {
	spriteObject
	getType() ammoType
}
