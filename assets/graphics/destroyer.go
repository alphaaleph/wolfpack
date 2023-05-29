package graphics

import (
	"github.com/alphaaleph/wolfpack/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

// TODO: add to configuration
var (
	destroyerSpeed     = 8.0
	destroyerAmmoPack  = 5 // 5 deep charges
	destroyerAmmoSpeed = 1.0
)

// destroyer represents the destroyer's sprite
type destroyer struct {
	character
}

// NewDestroyer creates an instance of a new Destroyer
func NewDestroyer() SpriteCharacterObject {
	// create a new destroyer
	d := &destroyer{
		character: character{
			cImageType: CharacterLeft,
			exploded:   false,
			speed:      destroyerSpeed,
			X:          util.ScreenWidth / 2.0,
			Y:          float64(util.DestroyerSectionTopY + characterPixelHalfSide - 10),
			//Y:        float64(util.DestroyerSectionRect.Bounds().Dy() - util.ScoreSectionRect.Bounds().Dy()),
		},
	}

	// load the destroyer sprites ahead of time
	d.leftImage = newSpriteImpl[*destroyer]().load(0, d)
	d.rightImage = newSpriteImpl[*destroyer]().load(1, d)
	d.explodeImage = newSpriteImpl[*destroyer]().load(2, d)

	// load the ammo into the destroyer
	d.munition = d.getMunitionPool(deepCharge, destroyerAmmoSpeed, destroyerAmmoPack)

	return d
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

// TODO: move to configuration
// getSpriteRect returns the destroyer's image location in the sprite sheet
func (d *destroyer) getSpriteRect(position int) image.Rectangle {
	switch characterTypes[position] {
	case CharacterLeft:
		return destroyerLeftSprite
	case CharacterRight:
		return destroyerRightSprite
	case CharacterExplosion:
		return destroyerExplosionSprite
	default:
		return image.Rectangle{}
	}
}

// Render draws the destroyer
func (d *destroyer) Render(screen *ebiten.Image) {

	// render the destroyer
	d.image = d.getSprite(d.cImageType)
	options := &ebiten.DrawImageOptions{}
	x := d.X - float64(d.image.Bounds().Size().X/2)
	options.GeoM.Translate(x, d.Y)
	options.Filter = ebiten.FilterLinear
	screen.DrawImage(d.image, options)

	// if the destroyer fires a deep charge show it
	if d.fire {
		defer d.Fire(false)
		if deepCharge, _ := d.munition.fire(); deepCharge != nil {
			if d.cImageType == CharacterLeft {
				x = x + float64(characterPixelSide) - 20.0
			}
			deepCharge.X = x
			deepCharge.Y = d.Y + float64(characterPixelHalfSide)
			deepCharge.Render(screen)
		}
	}

	// check if there is an active ammo and render
	for _, munition := range d.munition.active {
		// when the deep charge goes pass the bottom of the screen reload it into the pool
		if munition.Y >= util.ScreenHeight {
			d.munition.reload(munition)
		} else {
			munition.Render(screen)
		}
	}
}
