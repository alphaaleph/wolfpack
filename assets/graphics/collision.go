package graphics

import "fmt"

// will use a simple box collision method as this is not a production project

func CheckCollision(spriteCharacters []SpriteCharacterObject, spriteAmmos []SpriteAmmoObject) bool {

	// check each character against any possible ammo collision
	for _, spriteCharacter := range spriteCharacters {
		// check if the ammo intersect the character
		for _, spriteAmmo := range spriteAmmos {

			// check if it is enemy ammo before checking collision
			var isDestroyer bool
			aType := spriteAmmo.getType()

			switch spriteCharacter.(type) {
			case *destroyer:
				isDestroyer = true
			default:
				isDestroyer = false
			}

			if (isDestroyer && aType != deepCharge) || (!isDestroyer && aType == deepCharge) {
				aRect := spriteAmmo.getRect()
				cRect := spriteCharacter.getRect()

				if aRect.Overlaps(cRect) {
					fmt.Sprintf("HIT HIT HIT ---- character: %v,  with ammo: %v\n\n", spriteCharacter, spriteAmmo)
					spriteCharacter.setExploded(true)
					spriteAmmo.setExploded(true)
					return true
				}
			}
		}
	}

	// there are no collisions
	return false
}
