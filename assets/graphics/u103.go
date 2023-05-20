package graphics

import (
	"github.com/alphaaleph/wolfpack/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

var (
	BringTheWolf = false
)

type u103 struct {
	stamp
	sunk bool
}

// NewU103 creates an instance of a new U103 boss
func NewU103() SpriteObject {
	// create a new u103
	u103Sprite := &u103{}
	u103Sprite.image = NewCharacterImpl[*u103]().loadCharacterSprite(CharacterLeft, u103Sprite)

	// u103
	u103Sprite.X = util.ScreenWidth / 2.0
	u103Sprite.Y = float64(u103Sprite.image.Bounds().Dy()) * 1.5
	u103Sprite.speed = 15.0 // TODO: create different speed for how hard the game is
	return u103Sprite
}

// GetRect returns the u103's image location in the sprite sheet
func (u *u103) GetRect(ct characterType) image.Rectangle {
	switch ct {
	case CharacterLeft:
		return u103LeftStamp
	case CharacterRight:
		return u103RightStamp
	case CharacterExplosion:
		return u103ExplosionStamp
	default:
		return image.Rectangle{}
	}
}

// Render draws the u103
func (u *u103) Render(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(1.5, 1.5)
	w, h := u.image.Bounds().Dx(), u.image.Bounds().Dy()
	options.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)

	// readjust to use the middle of the player
	x := u.X - float64(u.image.Bounds().Size().X/2)
	options.GeoM.Translate(x, u.Y)
	options.Filter = ebiten.FilterLinear
	screen.DrawImage(u.image, options)
}
