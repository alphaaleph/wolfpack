# spiders
Learning Ebiten, and generic constrains with pointers, by making a shootout game






## Generics
Was playing around with generics. Used them to load the tiled images into the corresponding images.  **NOTE:** while 
using generics with structs, and methods with pointer receivers, you have to make sure that your interface constrains
also use pointers, or it will fail to compile.  Ex."
```
// tile defines the images manipulation constrains
type tile interface {
	*Player | *Bullet | *Spider | *Minion | *Poison
	load(*ebiten.Image)
	getRect() image.Rectangle
	GetImage() *ebiten.Image
}
```
The generics implementation do not use pointers for the generic parameters as follows:
```
// tileImpl is used to call functions on the game tile structs
type tileImpl[T tile] struct {
}

// New returns an instance of a tile
func New[T tile]() (t *tileImpl[T]) {
	t = &tileImpl[T]{}
	return t
}
```
But because the interface constrains and methods receivers both have pointers, the code has to call the implementation
using pointers as follows:
```
// init the images
SpidersApp.Player = img.Player{}
SpidersApp.Bullet = img.Bullet{}
SpidersApp.Spider = img.Spider{}
SpidersApp.Minion = img.Minion{}
SpidersApp.Poison = img.Poison{}

// load the tiles
img.New[*tile.Player]().SetTile(eimage, &SpidersApp.Player)
img.New[*tile.Bullet]().SetTile(eimage, &SpidersApp.Bullet)
img.New[*tile.Spider]().SetTile(eimage, &SpidersApp.Spider)
img.New[*tile.Minion]().SetTile(eimage, &SpidersApp.Minion)
img.New[*tile.Poison]().SetTile(eimage, &SpidersApp.Poison)
```



