package graphics

// ammoType tracks the ammo's selection in the sprite sheet
type ammoType int

const (
	deepCharge   ammoType = iota // 0
	uboatTorpedo                 // 1
	u103Torpedo                  // 2
)

var (
	ammoTypes = []ammoType{deepCharge, uboatTorpedo, u103Torpedo}
)

// Int returns the integer value of the ammoType
func (at ammoType) Int() int {
	return [...]int{0, 1, 2}[at]
}
