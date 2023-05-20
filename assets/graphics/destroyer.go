package graphics

import (
	"github.com/alphaaleph/wolfpack/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

// destroyer represents the destroyer's sprite
type destroyer struct {
	stamp
	sunk bool
}

// NewDestroyer creates an instance of a new Destroyer
func NewDestroyer() SpriteObject {
	// create a new destroyer
	destroyerSprite := &destroyer{}
	destroyerSprite.ctype = CharacterLeft

	// destroyer
	destroyerSprite.X = util.ScreenWidth / 2.0
	destroyerSprite.Y = float64(util.DestroyerSectionRect.Bounds().Dy() - util.ScoreSectionRect.Bounds().Dy())
	destroyerSprite.speed = 8 // TODO: create different speed for how hard the gae is
	return destroyerSprite
}

// DecX decreases the X coordinate by the provide amount
func (d *destroyer) DecX(x float64) {
	leftLimit := float64(d.image.Bounds().Dx() / 2)
	if d.X <= leftLimit {
		d.X = leftLimit
		return
	}
	d.X -= x
}

// IncX increases the X coordinate by the provide amount
func (d *destroyer) IncX(x float64) {
	rightLimit := float64(util.ScreenWidth - (d.image.Bounds().Dx() / 2))
	if d.X >= rightLimit {
		d.X = rightLimit
		return
	}
	d.X += x
}

// GetRect returns the destroyer's image location in the sprite sheet
func (d *destroyer) GetRect(ct characterType) image.Rectangle {
	switch ct {
	case CharacterLeft:
		return destroyerLeftStamp
	case CharacterRight:
		return destroyerRightStamp
	case CharacterExplosion:
		return destroyerExplosionStamp
	default:
		return image.Rectangle{}
	}
}

// Render draws the destroyer
func (d *destroyer) Render(screen *ebiten.Image) {
	d.image = NewCharacterImpl[*destroyer]().loadCharacterSprite(d.ctype, d)
	options := &ebiten.DrawImageOptions{}
	x := d.X - float64(d.image.Bounds().Size().X/2)
	options.GeoM.Translate(x, d.Y)
	options.Filter = ebiten.FilterLinear
	screen.DrawImage(d.image, options)
}
