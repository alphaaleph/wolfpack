package util

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

const (
	ScreenWidth  = 600
	ScreenHeight = 950
)

var (
	// TODO: change the code to manage window resizing
	// sections height
	scoreSectionHeight     = 50
	destroyerSectionHeight = 150
	wolfpackSectionHeight  = 250 // 3 sections for a total of 750 pixels height

	// size each section rectangle
	scoreSectionRect     = ebiten.NewImage(ScreenWidth, scoreSectionHeight)
	destroyerSectionRect = ebiten.NewImage(ScreenWidth, destroyerSectionHeight)
	wolfpackSectionRectA = ebiten.NewImage(ScreenWidth, wolfpackSectionHeight)
	wolfpackSectionRectB = ebiten.NewImage(ScreenWidth, wolfpackSectionHeight)
	wolfpackSectionRectC = ebiten.NewImage(ScreenWidth, wolfpackSectionHeight)

	// sections top vertical pixel
	ScoreSectionTopY        = 0
	DestroyerSectionTopY    = scoreSectionHeight
	DestroyerSectionBottomY = DestroyerSectionTopY + destroyerSectionHeight
	WolfpackSectionTopYA    = scoreSectionHeight + destroyerSectionHeight
	WolfpackSectionTopYB    = scoreSectionHeight + destroyerSectionHeight + wolfpackSectionHeight
	WolfpackSectionTopYC    = scoreSectionHeight + destroyerSectionHeight + wolfpackSectionHeight*2
)

func DrawScoresBackground(screen *ebiten.Image) {
	scoreSectionRect.Fill(color.RGBA{0xff, 0xf8, 0xdc, 0x00}) // white
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(0.0, float64(ScoreSectionTopY))
	options.Filter = ebiten.FilterLinear
	screen.DrawImage(scoreSectionRect, nil)
}

func DrawDestroyerBackground(screen *ebiten.Image) {
	destroyerSectionRect.Fill(color.RGBA{0xe0, 0xff, 0xff, 0xff}) // light cyan
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(0.0, float64(DestroyerSectionTopY))
	options.Filter = ebiten.FilterLinear
	screen.DrawImage(destroyerSectionRect, options)
}

func DrawWolfpackBackgroundA(screen *ebiten.Image) {
	wolfpackSectionRectA.Fill(color.RGBA{0x1e, 0x90, 0xff, 0xff}) // dodger blue
	options := &ebiten.DrawImageOptions{}
	options.Filter = ebiten.FilterLinear
	options.GeoM.Translate(0.0, float64(WolfpackSectionTopYA))
	screen.DrawImage(wolfpackSectionRectA, options)
}

func DrawWolfpackBackgroundB(screen *ebiten.Image) {
	wolfpackSectionRectB.Fill(color.RGBA{0x41, 0x69, 0xe1, 0xff}) // royal blue
	options := &ebiten.DrawImageOptions{}
	options.Filter = ebiten.FilterLinear
	options.GeoM.Translate(0.0, float64(WolfpackSectionTopYB))
	screen.DrawImage(wolfpackSectionRectB, options)
}

func DrawWolfpackBackgroundC(screen *ebiten.Image) {
	wolfpackSectionRectC.Fill(color.RGBA{0x00, 0x00, 0x8b, 0xff}) // blue
	options := &ebiten.DrawImageOptions{}
	options.Filter = ebiten.FilterLinear
	options.GeoM.Translate(0.0, float64(WolfpackSectionTopYC))
	screen.DrawImage(wolfpackSectionRectC, options)
}
