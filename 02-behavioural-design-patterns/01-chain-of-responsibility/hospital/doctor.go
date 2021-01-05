package hospital

import "fmt"

type Doctor struct {
	next Department
}

func (d *Doctor) Execute(p *Patient) {

	if p.doctorCheckupDone {
		fmt.Println("Doctor Checkup already done.")
		d.next.Execute(p)
		return
	}
	fmt.Printf("Doctor checking patient: %s\n", p.name)
	p.doctorCheckupDone = true
	d.next.Execute(p)
}

func (r *Doctor) SetNext(next Department) {
	r.next = next
}
