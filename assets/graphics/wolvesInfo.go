package graphics

import (
	"github.com/alphaaleph/wolfpack/util"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

var (
	// subSpeed is based on which depth section the submarine is in sections A, B, and C (top to bottom)
	subSpeed = [3]float64{6.0, 4.0, 2.0}
)

// wolvesInfo is a structure that handles functions used by the u103 and uboats
type wolvesInfo struct {
	dtype     depthType
	fireCoord float64
}

// checkLocation checks if the uboat is out of the screen, and if so it is reset to come back in at the same depth
func (wi *wolvesInfo) checkLocation(ct *characterImageType, x *float64) {
	if (*ct == CharacterLeft && *x < -128) || (*ct == CharacterRight && *x > util.ScreenWidth) {
		wi.reset(ct, x)
	}
}

// getEntryLocation will return the x,y coordinates where the sub is coming into the game
func (wi *wolvesInfo) getEntryLocation(cit characterImageType, dt depthType) (x float64, y float64) {
	// select which side the sub is coming into the display
	switch cit {
	case CharacterLeft:
		x = util.ScreenWidth
	case CharacterRight:
		x = 0 - 128
	default:
		x = 0 - 128
	}

	// get the depth coordinate
	switch dt {
	case depthA:
		y = float64(util.WolfpackSectionTopYA + 128)
	case depthB:
		y = float64(util.WolfpackSectionTopYB + 128)
	case depthC:
		y = float64(util.WolfpackSectionTopYC + 128)
	default:
		y = float64(util.WolfpackSectionTopYA + 128)
	}

	return
}

// getFileLocation sets a non-techy way when the sub will fire
func (wi *wolvesInfo) getFireLocation() {
	//it can fire in between 10 pixels of each side
	max := util.ScreenWidth - 10
	min := 10
	wi.fireCoord = float64(rand.Intn(max-min+1) + min)
}

// getMovementMultiplier moves the uboat in the correct direction
func (wi *wolvesInfo) getMovementMultiplier(cit characterImageType) int {
	switch cit {
	case CharacterLeft:
		return -1
	case CharacterRight:
		return 1
	default:
		return 0
	}
}

// getRandomDepth will return a sub's random depth location
func (wi *wolvesInfo) getRandomDepth() depthType {
	// pick one of the depths
	return depthTypes[rand.Intn(3)]
}

// getRandomDirectionImage will return a sub image left or right profile
func (wi *wolvesInfo) getRandomDirection() characterImageType {
	// pick one of the first two profiles
	return characterTypes[rand.Intn(2)]
}

// renderTorpedo shows the torpedo in the game screen
func (wi *wolvesInfo) renderTorpedo(screen *ebiten.Image) {
}

// reset redirects the uboat to come back into the playing area at the same depth
func (wi *wolvesInfo) reset(ct *characterImageType, x *float64) {
	*ct = wi.getRandomDirection()
	*x, _ = wi.getEntryLocation(*ct, wi.dtype)

	// if it is the u103 uboat also pick a random depth
	if BringTheWolf {
		wi.dtype = wi.getRandomDepth()
	}
}
