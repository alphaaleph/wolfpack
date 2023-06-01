package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

var (
	characterPixelHalfSide = 64
	characterPixelSide     = 128

	destroyerPoints = -1000000
	uboatAPoints    = 5
	uboatBPOints    = 10
	uboatCPoints    = 15
	u103Points      = 100
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

// Character changes the character's image
func (c *character) Character(cit characterImageType) {
	c.cImageType = cit
}

// Fire sets the character's fire flag to shoot ammo
func (c *character) Fire(f bool) {
	c.fire = f
}

// GetFiredMunitions returns a list of all munitions fired by the character
func (c *character) GetFiredMunitions() (munitions []SpriteAmmoObject) {
	for _, munition := range c.munition.active {
		munitions = append(munitions, munition)
	}
	return
}

// GetSpeed returns the character's speed
func (c *character) GetSpeed() float64 {
	return c.speed
}

// HasExploded determines if a character has exploded
func (c *character) HasExploded() bool {
	return c.exploded
}

// IncX increases the X coordinate by the provide amount
func (c *character) IncX(x float64) {
	c.X += x
}

// Reset sets the uboat back to a playing state
func (c *character) Reset() {
}

// Reload returns used ammo to the ammo pool
func (c *character) Reload() {
	// check which ammo has exploeded and reloaded into the pool
	for _, usedAmmo := range c.munition.active {
		if usedAmmo.exploded {
			usedAmmo.exploded = false
			c.munition.reload(usedAmmo)
		}
	}
}

// StillHasLives is true of the chracter still has lives
func (c *character) StillHasLives() bool {
	return false
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

// getRect returns a character rectangle coordinates
func (c *character) getRect() image.Rectangle {
	rect := image.Rectangle{
		Min: image.Point{int(c.X), int(c.Y)},
		Max: image.Point{int(c.X) + characterPixelSide, int(c.Y) + characterPixelSide},
	}
	return rect
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

// isFiring checks if the character is firing ammo
func (c *character) isFiring() bool {
	return c.fire
}

// setExploded sets the character explosion flag so it renders the correct exploded image
func (c *character) setExploded(e bool) {
	c.exploded = e
}
