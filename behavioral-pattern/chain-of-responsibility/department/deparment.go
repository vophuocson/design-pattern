package department

import (
	"design-pattern-practice/behavioral-pattern/chain-of-responsibility/patient"
	"fmt"
)

type Department interface {
	Execute(*patient.Patient)
	SetNext(Department)
}

type Reception struct {
	next Department
}

func (d *Reception) Execute(p *patient.Patient) {
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
		d.next.Execute(p)
		return
	}
	p.RegistrationDone = true
	d.next.Execute(p)
}

func (d *Reception) SetNext(department Department) {
	d.next = department
}

type Doctor struct {
	next Department
}

func (d *Doctor) Execute(p *patient.Patient) {
	if p.DoctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.Execute(p)
		return
	}
	p.DoctorCheckUpDone = true
	d.next.Execute(p)
}

func (d *Doctor) SetNext(department Department) {
	d.next = department
}

type Medical struct {
	next Department
}

func (d *Medical) Execute(p *patient.Patient) {
	if p.MedicineDone {
		fmt.Println("Medical checkup already done")
		d.next.Execute(p)
		return
	}
	p.MedicineDone = true
	d.next.Execute(p)
}

func (d *Medical) SetNext(department Department) {
	d.next = department
}

type Cashier struct {
	next Department
}

func (d *Cashier) Execute(p *patient.Patient) {
	if p.PaymentDone {
		fmt.Println("Cashier checkup already done")
		d.next.Execute(p)
		return
	}
	p.PaymentDone = true
	d.next.Execute(p)
}

func (d *Cashier) SetNext(department Department) {
	d.next = department
}
