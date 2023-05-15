package ui

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

const (
	mainTileSide = 128
)

var (
	gameSprites *ebiten.Image
)

// Sprite defines the image functionality
type Sprite interface {
	getRect() image.Rectangle
	DecX(float64)
	IncX(float64)
	GetSpeed() float64
	Render(screen *ebiten.Image)
}

// Sprites defines the image manipulation constrains
type Sprites interface {
	*destroyer | *charge | *u103 | *uboat | *torpedo | *stamp
	Sprite
}

// spritesImpl is used to call functions on the game sprite structs
type spritesImpl[T Sprites] struct {
}

// New returns an instance of a sprite
func New[T Sprites]() (t *spritesImpl[T]) {
	t = &spritesImpl[T]{}
	return t
}

// Load loads a sprite for a particular game piece
func (s *spritesImpl[T]) loadSprite(value T) *ebiten.Image {
	// decode the sprite string
	if gameSprites == nil {
		sprites, _, err := image.Decode(bytes.NewReader(byteSprites))
		if err != nil {
			fmt.Println("failed to read sprite string:", err)
			panic(err)
		}
		gameSprites = ebiten.NewImageFromImage(sprites)
	}

	// retrieve the struct stamp and load it into
	stamp := gameSprites.SubImage(value.getRect()).(*ebiten.Image)
	return stamp
}
