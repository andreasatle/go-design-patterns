package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/andreasatle/go-design-patterns/01-creational-design-patterns/03-factory/gun"
	"github.com/andreasatle/go-design-patterns/01-creational-design-patterns/03-factory/gunfactory"
)

func main() {
	ak47, err := gunfactory.GetGun("ak47")
	if err != nil {
		log.Fatal(err)
	}

	maverick, err := gunfactory.GetGun("maverick")
	if err != nil {
		log.Fatal(err)
	}

	foo, err := gunfactory.GetGun("foo")
	if err != nil {
		log.Printf("This error should occur: %v\n", err)
	}

	guns := []gun.Gunner{ak47, maverick, foo}

	fmt.Printf("Output for Ak47 gun:\n%v\n\n", ak47)
	fmt.Printf("Output for Maverick gun:\n%v\n\n", maverick)
	fmt.Printf("Output for Foo gun:\n%v\n\n", foo)

	fmt.Println(toString(guns))
	increasePowerOfAllGuns(guns)
	fmt.Println(toString(guns))
}

// Increase the power of all initialized guns.
// The point being that we have gun.Gunner as argument.
func increasePowerOfAllGuns(guns []gun.Gunner) {
	for _, gun := range guns {
		if gun == nil {
			continue
		}
		gun.SetPower(gun.GetPower() + 1)
	}
}

// Convert the Gunner-slice to a string for output.
func toString(guns []gun.Gunner) string {
	strs := []string{}
	for _, gun := range guns {
		if gun == nil {
			strs = append(strs, "<nil>")
			continue
		}
		// Use the Stringer interface for gun.Gunner
		strs = append(strs, fmt.Sprint(gun))
	}
	return strings.Join(strs, "\n")
}
