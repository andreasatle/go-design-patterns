package main

import "github.com/andreasatle/go-design-patterns/02-behavioural-design-patterns/01-chain-of-responsibility/hospital"

func main() {
	cashier := &hospital.Cashier{}
	medical := &hospital.Medical{}
	medical.SetNext(cashier)
	doctor := &hospital.Doctor{}
	doctor.SetNext(medical)
	reception := &hospital.Reception{}
	reception.SetNext(doctor)
	patient := hospital.NewPatient("Sven")
	reception.Execute(patient)
}
