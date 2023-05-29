package graphics

// characterImageType tracks the character's image selection in the sprite sheet
type characterImageType int

const (
	CharacterLeft      characterImageType = iota // 0
	CharacterRight                               // 1
	CharacterExplosion                           // 2
)

var (
	characterTypes = []characterImageType{CharacterLeft, CharacterRight, CharacterExplosion}
)

// Int returns the integer value of the characterType
func (ct characterImageType) Int() int {
	return [...]int{0, 1, 2}[ct]
}
