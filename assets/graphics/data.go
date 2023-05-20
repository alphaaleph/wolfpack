package graphics

import (
	_ "embed"
	"image"
)

//go:embed "sprites/spritesheet.png"
var byteSprites []byte

var (
	// destroyer sprites
	destroyerLeftStamp      = image.Rectangle{image.Point{0, 0}, image.Point{127, 127}}
	destroyerRightStamp     = image.Rectangle{image.Point{128, 0}, image.Point{255, 127}}
	destroyerExplosionStamp = image.Rectangle{image.Point{256, 0}, image.Point{383, 127}}

	// uboat sprites
	uboatLeftStamp      = image.Rectangle{image.Point{0, 128}, image.Point{127, 255}}
	uboatRightStamp     = image.Rectangle{image.Point{128, 128}, image.Point{255, 255}}
	uboatExplosionStamp = image.Rectangle{image.Point{256, 128}, image.Point{383, 255}}

	// u103 sprites
	u103LeftStamp      = image.Rectangle{image.Point{0, 256}, image.Point{127, 383}}
	u103RightStamp     = image.Rectangle{image.Point{128, 256}, image.Point{255, 383}}
	u103ExplosionStamp = image.Rectangle{image.Point{256, 256}, image.Point{383, 383}}

	// weapons
	deepChargeStamp   = image.Rectangle{image.Point{384, 63}, image.Point{447, 127}}
	uboatTorpedoStamp = image.Rectangle{image.Point{384, 191}, image.Point{447, 255}}
	u103TorpedoStamp  = image.Rectangle{image.Point{384, 319}, image.Point{447, 383}}
)
