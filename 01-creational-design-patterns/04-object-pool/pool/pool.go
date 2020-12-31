package pool

import (
	"errors"
	"log"
	"sync"
)

// Object is an interface with method GetId() string.
type Object interface {
	GetID() string
}

// Pool contains a pool of Objects.
type Pool struct {
	idle     []Object
	active   []Object
	capacity int
	mutex    *sync.Mutex
}

// InitPool initializes a Pool containing Objects.
func InitPool(objects []Object) (*Pool, error) {
	if len(objects) == 0 {
		return nil, errors.New("can't create an empty pool")
	}

	pool := &Pool{
		idle:     objects,
		active:   []Object{},
		capacity: len(objects),
		mutex:    new(sync.Mutex),
	}
	return pool, nil
}

// Get retrieves an Object from the pool.
func (p *Pool) Get() (Object, error) {

	p.mutex.Lock()
	defer p.mutex.Unlock()

	if len(p.idle) == 0 {
		return nil, errors.New("pool is empty")
	}

	object := p.idle[0]
	p.idle = p.idle[1:]

	p.active = append(p.active, object)
	log.Printf("Get Pool Object with ID: %s\n", object.GetID())
	return object, nil
}

func (p *Pool) Return(object Object) error {

	p.mutex.Lock()
	defer p.mutex.Unlock()

	err := p.remove(object)
	if err != nil {
		return err
	}

	p.idle = append(p.idle, object)
	log.Printf("Return Pool Object with ID: %s\n", object.GetID())
	return nil
}

func (p *Pool) remove(object Object) error {
	last := len(p.active) - 1
	for i, active := range p.active {
		if active == object {
			// swap current object with last object
			p.active[last], p.active[i] = p.active[i], p.active[last]

			// shrink size of active pool
			p.active = p.active[:last]

			return nil
		}
	}
	return errors.New("object not in active pool")
}
