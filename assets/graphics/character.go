package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	depthA = 0
	depthB = 1
	depthC = 2
	// subSpeed is based on which depth section the submarine is in sections A, B, and C (top to bottom)
	subSpeed = [3]float64{10.0, 6.0, 4.0}
)

// character is a generalization of a sprite that is not ammo
type character struct {
	ctype    characterType
	fire     bool
	X, Y     float64
	speed    float64
	munition *pool
	image    *ebiten.Image
}

// DecX decreases the X coordinate by the provide amount
func (c *character) DecX(x float64) {
	c.X -= x
}

// IncX increases the X coordinate by the provide amount
func (c *character) IncX(x float64) {
	c.X += x
}

// GetSpeed returns the character's speed
func (c *character) GetSpeed() float64 {
	return c.speed
}

// Character changes the character's image
func (c *character) Character(ct characterType) {
	c.ctype = ct
}

// Fire sets the character's fire flag to shoot ammo
func (c *character) Fire(f bool) {
	c.fire = f
}

// IsFiring checks if the character is firing ammo
func (c *character) IsFiring() bool {
	return c.fire
}

// getMunitionPool returns the pool of ammo
func (c *character) getMunitionPool(at ammoType, speed float64, count int) (p *pool) {
	//create the pool
	munitions := make([]*ammo, 0)
	for i := 0; i < count; i++ {
		a := &ammo{seqNum: i + 1, atype: at, speed: speed}
		a.image = a.loadAmmoSprite(at)
		munitions = append(munitions, a)
	}

	// ready the ammo
	p, _ = weaponReady(munitions)
	return
}
