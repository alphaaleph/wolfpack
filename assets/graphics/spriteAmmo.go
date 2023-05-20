package graphics

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

// spriteAmmoObject defines the image manipulation constrains for ammunition's
type spriteAmmoObject interface {
	*ammo | *stamp
	SpriteObject
	GetRect(ammoType) image.Rectangle
}

// spriteAmmoImpl is used to call functions on the ammo sprite structs
type spriteAmmoImpl[T spriteAmmoObject] struct {
}

// New returns an instance of an ammo sprite
func NewAmmoImpl[T spriteAmmoObject]() (t *spriteAmmoImpl[T]) {
	t = &spriteAmmoImpl[T]{}
	return t
}

// loadAmmoSprite loads a sprite for a particular weapon piece
func (s *spriteAmmoImpl[T]) loadAmmoSprite(at ammoType, value T) *ebiten.Image {
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
	stamp := gameSprites.SubImage(value.GetRect(at)).(*ebiten.Image)
	return stamp
}
