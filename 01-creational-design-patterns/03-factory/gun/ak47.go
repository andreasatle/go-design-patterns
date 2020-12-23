package gun

import (
	"fmt"
)

// Ak47 is the type containing basic info about an AK47 gun.
type Ak47 struct {
	Gun
	Origin string
}

// New creates a new AK47 gun.
func NewAk47() Gunner {
	return &Ak47{
		Gun: Gun{
			Name:  "AK47 gun",
			Power: 4,
		},
		Origin: "Russia",
	}
}

// String implements the Stringer interface
func (ak *Ak47) String() string {
	return fmt.Sprintf("%s\nOrigin: %s", ak.Gun.String(), ak.Origin)
}
