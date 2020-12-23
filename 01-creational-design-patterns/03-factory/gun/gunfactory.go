package gun

import (
	"fmt"
)

func GetGun(gunType string) (Gunner, error) {
	switch gunType {
	case "ak47":
		return NewAk47(), nil
	case "maverick":
		return NewMaverick(), nil
	default:
		return nil, fmt.Errorf("Unknown gun-type: %s", gunType)
	}
}
