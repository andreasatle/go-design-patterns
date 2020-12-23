package gun

// Maverick is the type containing basic info about a Maverick gun.
type Maverick struct {
	Gun
}

// New creates a new Maverick
func NewMaverick() Gunner {
	return &Maverick{
		Gun: Gun{
			Name:  "Maverick gun",
			Power: 5,
		},
	}
}
