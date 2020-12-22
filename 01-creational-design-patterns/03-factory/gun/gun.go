package gun

import "fmt"

// Gunner is the interface that creates a gun.
type Gunner interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}

// Gun is the type containing basic info about a gun.
type Gun struct {
	Name  string
	Power int
}

// SetName sets the name of the gun (part of Gunner interface).
func (g *Gun) SetName(name string) {
	g.Name = name
}

// GetName gets the name of the gun (part of Gunner interface).
func (g *Gun) GetName() string {
	return g.Name
}

// SetPower sets the power of the gun (part of Gunner interface).
func (g *Gun) SetPower(power int) {
	g.Power = power
}

// GetPower gets the power of the gun (part of Gunner interface).
func (g *Gun) GetPower() int {
	return g.Power
}

// String implements the Stringer interface.
func (g *Gun) String() string {
	return fmt.Sprintf("Gun: %s\nPower: %d", g.Name, g.Power)
}
