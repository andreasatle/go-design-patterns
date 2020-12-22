package ak47

import "github.com/andreasatle/go-design-patterns/01-creational-design-patterns/03-factory/gun"

// Ak47 is the type containing basic info about an AK47 gun.
type Ak47 struct {
	gun.Gun
}

// New creates a new AK47 gun.
func New() gun.Gunner {
	return &Ak47{
		Gun: gun.Gun{
			Name:  "AK47 gun",
			Power: 4,
		},
	}
}
