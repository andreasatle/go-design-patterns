package hospital

import "fmt"

type Medical struct {
	next Department
}

func (m *Medical) Execute(p *Patient) {

	if p.medicineDone {
		fmt.Println("Medicine already given to patient.")
		m.next.Execute(p)
		return
	}
	fmt.Printf("Medicine to patient: %s\n", p.name)
	p.medicineDone = true
	m.next.Execute(p)
}

func (m *Medical) SetNext(next Department) {
	m.next = next
}
