package hospital

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(p *Patient) {

	if p.paymentDone {
		fmt.Println("Payment alread done.")
		c.next.Execute(p)
		return
	}
	fmt.Printf("Cashier getting payment from patient: %s\n", p.name)
	p.paymentDone = true
}

func (c *Cashier) SetNext(next Department) {
	c.next = next
}
