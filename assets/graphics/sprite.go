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
	GetSpeed() float64
	HasExploded() bool
	IncX(float64)
	Render(screen *ebiten.Image)
	getRect() image.Rectangle
	getSpriteRect(int) image.Rectangle
	setExploded(bool)
}

type SpriteCharacterObject interface {
	spriteObject
	Character(characterImageType)
	Fire(bool)
	GetFiredMunitions() []SpriteAmmoObject
	GetPoints() int
	Reload()
	Reset()
	StillHasLives() bool
	isFiring() bool
}

type SpriteAmmoObject interface {
	spriteObject
	getType() ammoType
}
