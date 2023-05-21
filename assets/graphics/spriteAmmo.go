package graphics

// spriteAmmoConstrains defines the image manipulation constrains for ammunition's
/*type spriteAmmoConstrains interface {
	*ammo
	SpriteAmmoObject
}

// spriteAmmoImpl is used to call functions on the ammo sprite structs
type spriteAmmoImpl[T spriteAmmoConstrains] struct {
}

// New returns an instance of an ammo sprite
func NewAmmoImpl[T spriteAmmoConstrains]() (t *spriteAmmoImpl[T]) {
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

	// retrieve the struct ammo and load it into
	ammo := gameSprites.SubImage(value.GetRect(at)).(*ebiten.Image)
	return ammo
}*/
