package main

import "github.com/andreasatle/go-design-patterns/02-behavioural-design-patterns/01-chain-of-responsibility/hospital"

func main() {
	reception := hospital.NewReception()
	patient := hospital.NewPatient("Sven")
	reception.Execute(patient)
}
