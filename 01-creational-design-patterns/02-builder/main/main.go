package main

import (
	"fmt"

	"github.com/andreasatle/go-design-patterns/01-creational-design-patterns/02-builder/house"
)

func main() {
	normalBuilder := house.GetBuilder("normal")
	iglooBuilder := house.GetBuilder("igloo")

	director := house.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()
	fmt.Print(normalHouse)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()
	fmt.Print(iglooHouse)
}
