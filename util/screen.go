package util

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

const (
	ScreenWidth  = 600
	ScreenHeight = 950
)

var (
	// score section 50 pixels
	scoreSectionHeight = 50
	scoreSectionRect   = ebiten.NewImage(ScreenWidth, scoreSectionHeight)
	//scoreSection       = scoreSectionRect.SubImage(image.Rect(0, 0, ScreenWidth-1, scoreSectionHeight-1)).(*ebiten.Image)
	//scoreSection = scoreSectionRect.SubImage(image.Rect(100, 100, 200, 200)).(*ebiten.Image)

	// destroyer section 150 pixels
	destroyerSectionHeight = 150
	destroyerSectionRect   = ebiten.NewImage(ScreenWidth, destroyerSectionHeight)
	//DestroyerSection       = destroyerSectionRect.SubImage(image.Rect(0, scoreSectionHeight, ScreenWidth-1,
	//	scoreSectionHeight+destroyerSectionHeight-1)).(*ebiten.Image)
	DestroyerSection = destroyerSectionRect.SubImage(image.Rect(0, 200, 600, 700)).(*ebiten.Image)

	// wolfpack section 750 pixels, 250 per sub-section
	wolfpackSectionHeight    = 600
	wolfpackSubSectionHeight = 250
	wolfpackSubSections      = 3

	wolfpackSubSectionRect_A = ebiten.NewImage(ScreenWidth, wolfpackSubSectionHeight)
	wolfpackSubSectionRect_B = ebiten.NewImage(ScreenWidth, wolfpackSubSectionHeight)
	wolfpackSubSectionRect_C = ebiten.NewImage(ScreenWidth, wolfpackSubSectionHeight)

	wolfpackSubSection_A = wolfpackSubSectionRect_A.SubImage(image.Rect(0, scoreSectionHeight+destroyerSectionHeight,
		ScreenWidth-1, scoreSectionHeight+destroyerSectionHeight+wolfpackSubSectionHeight)).(*ebiten.Image)
	wolfpackSubSection_B = wolfpackSubSectionRect_B.SubImage(image.Rect(0, wolfpackSubSection_A.Bounds().Dy(),
		ScreenWidth-1, wolfpackSubSection_A.Bounds().Dy()+wolfpackSubSectionHeight)).(*ebiten.Image)
	wolfpackSubSection_C = wolfpackSubSectionRect_C.SubImage(image.Rect(0, wolfpackSubSection_B.Bounds().Dy(),
		ScreenWidth-1, wolfpackSubSection_B.Bounds().Dy()+wolfpackSubSectionHeight)).(*ebiten.Image)
)

func DrawScoresBackground(screen *ebiten.Image) {
	scoreSectionRect.Fill(color.RGBA{0xff, 0xf8, 0xdc, 0x00}) // white
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(0.0, 0.0)
	options.Filter = ebiten.FilterLinear
	screen.DrawImage(scoreSectionRect, nil)
}

func DrawDestroyerBackground(screen *ebiten.Image) {
	destroyerSectionRect.Fill(color.RGBA{0xe0, 0xff, 0xff, 0xff}) // light cyan
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(0.0, 50.0)
	options.Filter = ebiten.FilterLinear
	screen.DrawImage(destroyerSectionRect, options)
}

func DrawWolfpackBackground(screen *ebiten.Image) {

	options := &ebiten.DrawImageOptions{}
	options.Filter = ebiten.FilterLinear

	// draw section A
	wolfpackSubSectionRect_A.Fill(color.RGBA{0x1e, 0x90, 0xff, 0xff}) // dodger blue
	options.GeoM.Translate(0.0, 200.0)
	screen.DrawImage(wolfpackSubSectionRect_A, options)

	// draw section B
	wolfpackSubSectionRect_B.Fill(color.RGBA{0x41, 0x69, 0xe1, 0xff}) // royal blue
	options.GeoM.Translate(0.0, 250.0)
	screen.DrawImage(wolfpackSubSectionRect_B, options)

	// draw section C
	wolfpackSubSectionRect_C.Fill(color.RGBA{0x00, 0x00, 0x8b, 0xff}) // blue
	options.GeoM.Translate(0.0, 250.0)
	screen.DrawImage(wolfpackSubSectionRect_C, options)
}
