package main

import (
	"fmt"
	"log"

	"github.com/andreasatle/go-design-patterns/01-creational-design-patterns/01-abstract-factory/sport"
)

func main() {
	adidasFactory, err := sport.GetSportsFactory("adidas")
	if err != nil {
		log.Fatal(err)
	}
	nikeFactory, err := sport.GetSportsFactory("nike")
	if err != nil {
		log.Fatal(err)
	}
	nikeShoe := nikeFactory.MakeShoe(12)
	nikeShort := nikeFactory.MakeShorts(13)
	adidasShoe := adidasFactory.MakeShoe(14)
	adidasShort := adidasFactory.MakeShorts(15)
	printShoeDetails(nikeShoe)
	printShortsDetails(nikeShort)
	printShoeDetails(adidasShoe)
	printShortsDetails(adidasShort)
}

func printShoeDetails(s sport.ShoeInterface) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

func printShortsDetails(s sport.ShortsInterface) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
