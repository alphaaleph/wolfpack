package graphics

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

// spriteCharacterObject defines the image manipulation constrains for characters
type spriteCharacterObject interface {
	*destroyer | *u103 | *uboat | *stamp
	SpriteObject
	GetRect(characterType) image.Rectangle
}

// spriteCharacterImpl is used to call functions on the character's sprite structs
type spriteCharacterImpl[T spriteCharacterObject] struct {
}

// New returns an instance of a sprite
func NewCharacterImpl[T spriteCharacterObject]() (t *spriteCharacterImpl[T]) {
	t = &spriteCharacterImpl[T]{}
	return t
}

// loadCharacterSprite loads a sprite for a particular character
func (s *spriteCharacterImpl[T]) loadCharacterSprite(ct characterType, value T) *ebiten.Image {
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
	stamp := gameSprites.SubImage(value.GetRect(ct)).(*ebiten.Image)
	return stamp
}
