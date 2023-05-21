package graphics

import (
	_ "embed"
	"image"
)

//go:embed "sprites/spritesheet.png"
var byteSprites []byte

var (
	// destroyer sprites
	destroyerLeftSprite      = image.Rectangle{image.Point{0, 0}, image.Point{127, 127}}
	destroyerRightSprite     = image.Rectangle{image.Point{128, 0}, image.Point{255, 127}}
	destroyerExplosionSprite = image.Rectangle{image.Point{256, 0}, image.Point{383, 127}}

	// uboat sprites
	uboatLeftSprite      = image.Rectangle{image.Point{0, 128}, image.Point{127, 255}}
	uboatRightSprite     = image.Rectangle{image.Point{128, 128}, image.Point{255, 255}}
	uboatExplosionSprite = image.Rectangle{image.Point{256, 128}, image.Point{383, 255}}

	// u103 sprites
	u103LeftSprite      = image.Rectangle{image.Point{0, 256}, image.Point{127, 383}}
	u103RightSprite     = image.Rectangle{image.Point{128, 256}, image.Point{255, 383}}
	u103ExplosionSprite = image.Rectangle{image.Point{256, 256}, image.Point{383, 383}}

	// weapons
	deepChargeSprite   = image.Rectangle{image.Point{384, 63}, image.Point{447, 127}}
	uboatTorpedoSprite = image.Rectangle{image.Point{384, 191}, image.Point{447, 255}}
	u103TorpedoSprite  = image.Rectangle{image.Point{384, 319}, image.Point{447, 383}}
)
