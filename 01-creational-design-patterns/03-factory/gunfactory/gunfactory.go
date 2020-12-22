package gunfactory

import (
	"fmt"

	"github.com/andreasatle/go-design-patterns/01-creational-design-patterns/03-factory/ak47"
	"github.com/andreasatle/go-design-patterns/01-creational-design-patterns/03-factory/gun"
	"github.com/andreasatle/go-design-patterns/01-creational-design-patterns/03-factory/maverick"
)

func GetGun(gunType string) (gun.Gunner, error) {
	switch gunType {
	case "ak47":
		return ak47.New(), nil
	case "maverick":
		return maverick.New(), nil
	default:
		return nil, fmt.Errorf("Unknown gun-type: %s", gunType)
	}
}
