package graphics

import "github.com/hajimehoshi/ebiten/v2"

const (
	mainTileSide = 128
)

var (
	gameSprites *ebiten.Image
)

// spriteObject defines the image functionality for both
type SpriteObject interface {
	DecX(float64)
	IncX(float64)
	GetSpeed() float64
	Render(screen *ebiten.Image)
	Stamp(characterType)
}
