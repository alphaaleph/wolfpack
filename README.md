# wolfpack
Learning [Ebiten](https://ebitengine.org/), and generic constrains with pointers, as well as trying a pool pattern, by 
making a submarine hunt game.  The game consists of a destroyer used by the player, to sink the passing submarines 
below; just like the old [Atari](https://atari.com/) or [Commodore](https://en.wikipedia.org/wiki/Commodore_64games).  
There is a lot to fix and refactor in this game, but it is just a prototype to learn how to write a game and move 
forward with an original one.  Feel free to copy it and improve it, or update this repo.

## Characters

There are three characters displayed in the game:

### Destroyer
![plot](./assets/graphics/sprites/ship_left.png)

Used by the player to sink the enemy submarines. The destroyer fires deep charges, which have to be reloaded at 
certain intervals.  The left and right arrows are used to move the destroyer, and the space bar is used to fire the
deep charges.

### Uboat
![plot](./assets/graphics/sprites/uboat_left.png)

This is the regular enemy submarine.  It will show at different speeds and depths and will try to sink the destroyer
shooting torpedoes at it.  There is a slice of uboats (wolfpack), that has to be completely sunk before the boss (U103) 
submarine shows up.


### U103
![plot](./assets/graphics/sprites/wolf_left.png)

This is the boss submarine.  Once all of the uboats have been sunk, the boss U103 will show up. The U103 will show up
at different depths, using different speeds.




## Ammo

There are three sets of ammo used by the characters:

### Deep Charges 
![plot](./assets/graphics/sprites/deep_charge.png)

### Uboat Torpedoes
![plot](./assets/graphics/sprites/uboat_torpedo.png)

### U103 Torpedoes
![plot](./assets/graphics/sprites/wolf_torpedo.png)


## Ammo Pool
A pool pattern is being used to manage the character's ammo. Currently, each character can only fire as much ammo as
is available in the pool at one time.  Once the ammo leaves the screen, or explodes by hitting another character, the 
ammo is returned to the pool. These are the current amounts for ach character:

* destroyer has 3 deep charges allocated
* uboat has 2 torpedos allocated
* u103 has 5 torpedos allocated



## Generics
Was playing around with generics. Used them to load the sprites into the corresponding character and ammo images. 
**NOTE:** while using generics with structs, and methods with pointer receivers, you have to make sure that your 
interface constrains also use pointers, or it will fail to compile.  Ex."
```
    // spriteConstrains defines the image manipulation constrains
    type spriteConstrains interface {
	    *ammo | *character | *destroyer | *u103 | *uboat
	    spriteObject
    }
```

The generics implementation use the interface for the generic parameters as follows:
```
    // spriteImpl is used to call functions on the sprite structs
    type spriteImpl[T spriteConstrains] struct {
    }

    // new returns an instance of an ammo sprite
    func newSpriteImpl[T spriteConstrains]() (t *spriteImpl[T]) {
	    t = &spriteImpl[T]{}
	    return t
    }
```

But because the interface constrains and methods receivers both have pointers, the code has to call the implementation
using pointers as follows:
```
	d := &destroyer{}
	d.leftImage = newSpriteImpl[*destroyer]().load(0, d)
```


### Running the Program

## Running from Binary
Included under the bin directory there are a couple of already built executables. Run using one of the following the 
following commands:

* for linux :        ./bin/wolfpackx
* for windows 32 bit :      ./bin/wolfpackx86.exe
* for windows 64 bit :      ./bin/wolfpack64.exe


## Building and Running the Program

* make sure that you have Golang installed
* from the project directory build the code for your OS:
  * linux:     go build -o ./bin/wolfpack 
  * windows 32 bit:   env GOOS=windows GOARCH=386 go build -o ./bin/wolfpackx86.exe
  * windows 64 bit:   env GOOS=windows GOARCH=amd64 go build -o ./bin/wolfpack64.exe

    
## Keys

* left arrow - moves the destroyer to the left
* right arrow - moves the destroyer to the right
* space bar - used to shoot deep charges against the subs


### Images

![wolfpack - destroyer against regular subs](https://github.com/alphaaleph/wolfpack/blob/main/extras/images/wolfpack.png)

![wolfpack - destroyer against boss u103](https://github.com/alphaaleph/wolfpack/blob/main/extras/images/u103.png)
