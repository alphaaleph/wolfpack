package graphics

// ammoType tracks the ammo's selection in the sprite sheet
type ammoType uint8

const (
	deepCharge   ammoType = 1 << iota // 0x01
	uboatTorpedo                      // 0x02
	u103Torpedo                       // 0x04
)

var (
	ammoTypes = []ammoType{deepCharge, uboatTorpedo, u103Torpedo}
)
