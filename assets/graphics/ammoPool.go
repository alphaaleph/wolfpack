package graphics

import (
	"fmt"
	"sync"
)

// pool manages the different ammo pools
type pool struct {
	idle     []*ammo
	active   []*ammo
	capacity int
	mulock   *sync.Mutex
}

// weaponReady initializes the ammo pool
func weaponReady(munitions []*ammo) (*pool, error) {
	if len(munitions) == 0 {
		return nil, fmt.Errorf("Cannot craete a pool of 0 length")
	}
	active := make([]*ammo, 0)
	pool := &pool{
		idle:     munitions,
		active:   active,
		capacity: len(munitions),
		mulock:   new(sync.Mutex),
	}
	return pool, nil
}

// fire gets an ammo from the pool to shoot it to an adversary
func (p *pool) fire() (*ammo, error) {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	if len(p.idle) == 0 {
		return nil, fmt.Errorf("No pool object free. Please request after sometime")
	}
	obj := p.idle[0]
	obj.fired = true
	p.idle = p.idle[1:]
	p.active = append(p.active, obj)
	fmt.Printf("fire the ammo sequence number: %s\n", obj.getSeqNum())
	return obj, nil
}

// reload returns ammo into the pool to be reused
func (p *pool) reload(target *ammo) error {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	target.fired = false
	err := p.discharge(target)
	if err != nil {
		return err
	}
	p.idle = append(p.idle, target)
	fmt.Printf("reload ammo sequence number: \n", target.getSeqNum())
	return nil
}

// discharge gets rid of ammo from the pool
func (p *pool) discharge(target *ammo) error {
	currentActiveLength := len(p.active)
	for i, obj := range p.active {
		if obj.getSeqNum() == target.getSeqNum() {
			p.active[currentActiveLength-1], p.active[i] = p.active[i], p.active[currentActiveLength-1]
			p.active = p.active[:currentActiveLength-1]
			return nil
		}
	}
	return fmt.Errorf("ammo doesn't belong to the pool")
}
