package pooltoys

import "errors"

type PoolToys struct {
	connections []Toy
	size        int
	active      int
	size_limit  int
}

type Toy struct {
	connection interface{}
}

func New(init_toys []interface{}, size_limit int) *PoolToys {
	var toys []Toy
	for toy_interface := range init_toys {
		toys = append(toys, Toy{toy_interface})
	}
	return &PoolToys{
		connections: toys,
		size:        len(init_toys),
		active:      0,
		size_limit:  size_limit,
	}
}

func (toys PoolToys) Add(conn interface{}) (PoolToys, error) {
	if toys.size == toys.size_limit {
		return toys,
			errors.New("Your Connection pool has reached it's size limit, cannot add anymore")
	}
	toys.connections = append(toys.connections, Toy{conn})
	return toys, nil
}

func (toys PoolToys) Take() interface{} {
	toy := toys.connections[toys.active]
	toys.active++

	if toys.active == toys.size+1 {
		toys.active = 0
	}

	return toy.connection
}
