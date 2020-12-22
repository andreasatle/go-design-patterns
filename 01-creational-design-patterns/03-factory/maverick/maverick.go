package maverick

import "github.com/andreasatle/go-design-patterns/01-creational-design-patterns/03-factory/gun"

// Maverick is the type containing basic info about a Maverick gun.
type Maverick struct {
	gun.Gun
}

// New creates a new Maverick gun.
func New() gun.Gunner {
	return &Maverick{
		Gun: gun.Gun{
			Name:  "Maverick gun",
			Power: 5,
		},
	}
}
