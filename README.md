# wolfpack
Learning Ebiten, and generic constrains with pointers, by making a submarine hunt game






## Generics
Was playing around with generics. Used them to load the sprites into the corresponding images.  **NOTE:** while 
using generics with structs, and methods with pointer receivers, you have to make sure that your interface constrains
also use pointers, or it will fail to compile.  Ex."
```
    // Sprites defines the image manipulation constrains
    type Sprites interface {
	    *destroyer | *charge | *u103 | *uboat | *torpedo | *stamp
	    Sprite
    }
```

The generics implementation use the interface for the generic parameters as follows:
```
    // spritesImpl is used to call functions on the game sprite structs
    type spritesImpl[T Sprites] struct {
    }

    // New returns an instance of a sprite
    func New[T Sprites]() (t *spritesImpl[T]) {
	    t = &spritesImpl[T]{}
	    return t
    }
```

But because the interface constrains and methods receivers both have pointers, the code has to call the implementation
using pointers as follows:
```
    destroyerSprite := &destroyer{}
    destroyerSprite.image = New[*destroyer]().loadSprite(destroyerSprite)
```



