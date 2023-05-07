package img

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

// tile defines the tile manipulation constrains
type tile interface {
	*Player | *Bullet | *Spider | *Minion | *Poison
	load(image2 *ebiten.Image)
	getRect() image.Rectangle
	Render(screen *ebiten.Image)
}

// tileImpl is used to call functions on the game tile structs
type tileImpl[T tile] struct {
}

// New returns an instance of a tile
func New[T tile]() (t *tileImpl[T]) {
	t = &tileImpl[T]{}
	return t
}

// LoadTile reads the tile tiles
func (s *tileImpl[T]) LoadTile() (tiles image.Image, err error) {
	// decode the tile string
	tiles, _, err = image.Decode(bytes.NewReader(gameTiles))
	if err != nil {
		fmt.Println("failed to load tile:", err)
		return nil, err
	}
	return tiles, nil
}

// SetTile an image for a particular game piece
func (s *tileImpl[T]) SetTile(image *ebiten.Image, value T) {
	// retrieve the struct image
	img := image.SubImage(value.getRect()).(*ebiten.Image)

	// set the struct tile
	value.load(img)
}
