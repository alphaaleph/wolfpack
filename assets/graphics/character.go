package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

var (
	characterPixelHalfSide = 64
	characterPixelSide     = 128
)

// character is a generalization of a sprite that is not ammo
type character struct {
	cImageType                          characterImageType
	fire                                bool
	exploded                            bool
	X, Y                                float64
	speed                               float64
	munition                            *pool
	image                               *ebiten.Image // current displayed image
	leftImage, rightImage, explodeImage *ebiten.Image // character images
}

// DecX decreases the X coordinate by the provide amount
func (c *character) DecX(x float64) {
	c.X -= x
}

// IncX increases the X coordinate by the provide amount
func (c *character) IncX(x float64) {
	c.X += x
}

// getRect returns a character rectangle coordinates
func (c *character) getRect() image.Rectangle {
	rect := image.Rectangle{
		Min: image.Point{int(c.X), int(c.Y)},
		Max: image.Point{int(c.X) + characterPixelSide, int(c.Y) + characterPixelSide},
	}
	return rect
}

// GetSpeed returns the character's speed
func (c *character) GetSpeed() float64 {
	return c.speed
}

// Character changes the character's image
func (c *character) Character(cit characterImageType) {
	c.cImageType = cit
}

// Fire sets the character's fire flag to shoot ammo
func (c *character) Fire(f bool) {
	c.fire = f
}

// isFiring checks if the character is firing ammo
func (c *character) isFiring() bool {
	return c.fire
}

// getMunitionPool returns the pool of ammo
func (c *character) getMunitionPool(at ammoType, speed float64, count int) (p *pool) {
	//create the pool
	munitions := make([]*ammo, 0)
	for i := 0; i < count; i++ {
		munitions = append(munitions, newAmmo(i+1, at, speed))
	}

	// ready the ammo
	p, _ = weaponReady(munitions)
	return
}

// getSprite return the selected sprite
func (c *character) getSprite(cit characterImageType) *ebiten.Image {
	switch cit {
	case CharacterLeft:
		return c.leftImage
	case CharacterRight:
		return c.rightImage
	case CharacterExplosion:
		return c.explodeImage
	default:
		return c.leftImage
	}
}

// setExploded sets the character explosion flag so it renders the correct exploded image
func (c *character) setExploded(e bool) {
	c.exploded = e
}

// HasExploded determines if a character has exploded
func (c *character) HasExploded() bool {
	return c.exploded
}

// GetFiredMunitions returns a list of all munitions fired by the character
func (c *character) GetFiredMunitions() (munitions []SpriteAmmoObject) {
	for _, munition := range c.munition.active {
		munitions = append(munitions, munition)
	}
	return
}

// Reset sets the uboat back to a playing state
func (c *character) Reset() {
}

// StillHasLives is true of the chracter still has lives
func (c *character) StillHasLives() bool {
	return false
}
