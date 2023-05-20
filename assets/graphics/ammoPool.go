package graphics

import (
	"fmt"
	"sync"
)

// poolObject is the interface of ammunition used in the game
type poolObject interface {
	getNum() int
}

// pool manages the different ammo pools
type pool struct {
	idle     []poolObject
	active   []poolObject
	capacity int
	mulock   *sync.Mutex
}

// destroyer handles the ammunition for the destroyer
type destroyerPool struct {
	pool
}

// wolfpackPool handles the ammunition for the uboats
type wolfpackPool struct {
	pool
}

// u103Pool handles the ammunition for the u103 boss
type u103Pool struct {
	pool
}

// loadAmmo initializes the ammo pool
func LoadAmmo(poolObjects []poolObject) (*pool, error) {
	if len(poolObjects) == 0 {
		return nil, fmt.Errorf("Cannot craete a pool of 0 length")
	}
	active := make([]poolObject, 0)
	pool := &pool{
		idle:     poolObjects,
		active:   active,
		capacity: len(poolObjects),
		mulock:   new(sync.Mutex),
	}
	return pool, nil
}

// fire gets an ammo from the pool to shoot it to an adversary
func (p *pool) fire() (poolObject, error) {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	if len(p.idle) == 0 {
		return nil, fmt.Errorf("No pool object free. Please request after sometime")
	}
	obj := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, obj)
	fmt.Printf("fire the ammo number: %s\n", obj.getNum())
	return obj, nil
}

// reload returns ammo into the pool to be reused
func (p *pool) reload(target poolObject) error {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	err := p.discharge(target)
	if err != nil {
		return err
	}
	p.idle = append(p.idle, target)
	fmt.Printf("reload ammo number: %s\n", target.getNum())
	return nil
}

// discharge gets rid of ammo from the pool
func (p *pool) discharge(target poolObject) error {
	currentActiveLength := len(p.active)
	for i, obj := range p.active {
		if obj.getNum() == target.getNum() {
			p.active[currentActiveLength-1], p.active[i] = p.active[i], p.active[currentActiveLength-1]
			p.active = p.active[:currentActiveLength-1]
			return nil
		}
	}
	return fmt.Errorf("ammo doesn't belong to the pool")
}
