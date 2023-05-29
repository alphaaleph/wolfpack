package graphics

// depthType tracks the depth selection in the display from top to bottom
type depthType uint8

const (
	depthA depthType = iota // 0x00 top depth section a
	depthB                  // 0x01 middle depth section b
	depthC                  // 0x02 bottom depth section c
)

var (
	depthTypes = []depthType{depthA, depthB, depthC}
)
