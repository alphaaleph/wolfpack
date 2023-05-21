package graphics

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

// spriteCharacterConstrains defines the image manipulation constrains for characters
type spriteCharacterConstrains interface {
	*destroyer | *u103 | *uboat | *character
	SpriteCharacterObject
}

// spriteCharacterImpl is used to call functions on the character's sprite structs
type spriteCharacterImpl[T spriteCharacterConstrains] struct {
}

// New returns an instance of a sprite
func NewCharacterImpl[T spriteCharacterConstrains]() (t *spriteCharacterImpl[T]) {
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

	// retrieve the struct character sprite and load it into
	character := gameSprites.SubImage(value.GetRect(ct)).(*ebiten.Image)
	return character
}
