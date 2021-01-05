package hospital

import "fmt"

type Reception struct {
	next Department
}

func (r *Reception) Execute(p *Patient) {

	if p.registrationDone {
		fmt.Println("Patient already registered.")
		r.next.Execute(p)
		return
	}

	fmt.Printf("Reception registering patient: %s\n", p.name)
	p.registrationDone = true
	r.next.Execute(p)
}

func (r *Reception) SetNext(next Department) {
	r.next = next
}

func NewReception() *Reception {
	cashier := &Cashier{}
	medical := &Medical{}
	medical.SetNext(cashier)
	doctor := &Doctor{}
	doctor.SetNext(medical)
	reception := &Reception{}
	reception.SetNext(doctor)
	return reception
}
