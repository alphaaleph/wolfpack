package game

import (
	"github.com/alphaaleph/spiders/img"
	"github.com/hajimehoshi/ebiten/v2"
)

// initialize images
func init() {
	// read the image tiles
	image, err := img.New[*img.Player]().LoadTile()
	if err != nil {
		panic(err)
	}
	eimage := ebiten.NewImageFromImage(image)

	// init the game
	SpidersApp = NewGame()

	// init the images
	SpidersApp.Player = img.Player{}
	SpidersApp.Bullet = img.Bullet{}
	SpidersApp.Spider = img.Spider{}
	SpidersApp.Minion = img.Minion{}
	SpidersApp.Poison = img.Poison{}

	// load the images
	img.New[*img.Player]().SetTile(eimage, &SpidersApp.Player)
	img.New[*img.Bullet]().SetTile(eimage, &SpidersApp.Bullet)
	img.New[*img.Spider]().SetTile(eimage, &SpidersApp.Spider)
	img.New[*img.Minion]().SetTile(eimage, &SpidersApp.Minion)
	img.New[*img.Poison]().SetTile(eimage, &SpidersApp.Poison)
}

// initialize starting values
func init() {
	// mode level
	SpidersApp.modeLevel = ModeTitle

	// player
	SpidersApp.Player.X = ScreenWidth / 2.0
	SpidersApp.Player.Y = ScreenHeight - float64(SpidersApp.Player.Image.Bounds().Dy())*1.5
	SpidersApp.Player.Speed = 15.0 // TODO: create different speed for how hard the game is

	// spider
	SpidersApp.Spider.X = ScreenWidth / 2.0
	SpidersApp.Spider.Y = float64(SpidersApp.Spider.Image.Bounds().Dy()) * 1.5
	SpidersApp.Spider.Speed = 15.0 // TODO: create different speed for how hard the game is
}
