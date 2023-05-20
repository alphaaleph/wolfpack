package graphics

// characterType tracks the character's selection in the sprite sheet
type characterType uint8

const (
	CharacterLeft      characterType = 1 << iota // 0x01
	CharacterRight                               // 0x02
	CharacterExplosion                           // 0x04
)

var (
	characterTypes = []characterType{CharacterLeft, CharacterRight, CharacterExplosion}
)
