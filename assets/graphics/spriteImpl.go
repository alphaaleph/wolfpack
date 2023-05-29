package graphics

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

// spriteConstrains defines the image manipulation constrains
type spriteConstrains interface {
	*ammo | *character | *destroyer | *u103 | *uboat
	spriteObject
}

// spriteImpl is used to call functions on the sprite structs
type spriteImpl[T spriteConstrains] struct {
}

// new returns an instance of an ammo sprite
func newSpriteImpl[T spriteConstrains]() (t *spriteImpl[T]) {
	t = &spriteImpl[T]{}
	return t
}

// load loads a sprite for a particular spriteObject
func (s *spriteImpl[T]) load(position int, value T) *ebiten.Image {
	// decode the sprite string
	if gameSprites == nil {
		sprites, _, err := image.Decode(bytes.NewReader(byteSprites))
		if err != nil {
			fmt.Println("failed to read sprite string:", err)
			panic(err)
		}
		gameSprites = ebiten.NewImageFromImage(sprites)
	}

	// retrieve the struct ammo and load it into
	rect := value.getSpriteRect(position)
	ammo := gameSprites.SubImage(rect).(*ebiten.Image)
	return ammo
}
