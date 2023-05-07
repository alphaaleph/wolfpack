package game

// modeType tracks the game active screen
type modeType uint8

const (
	ModeTitle modeType = 1 << iota // 0x01
	ModeGame                       // 0x02
	ModeOver  modeType = 255       // 0xff
)

var (
	modes = []modeType{ModeTitle, ModeGame, ModeOver}
)
