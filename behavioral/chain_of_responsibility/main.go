package main

import (
	"fmt"
	"log"
)

type Patient struct {
	Name string
}

type Department interface {
	execute(*Patient) error
	setNext(Department)
}

// ======================== Receptionist

type Receptionist struct {
	next Department
}

func (r *Receptionist) execute(p *Patient) error {
	fmt.Println("Visited Receptionist")

	err := r.next.execute(p)
	if err != nil {
		return err
	}
	return nil
}

func (r *Receptionist) setNext(department Department) {
	r.next = department
}

// ======================== Doctor

type Doctor struct {
	next Department
}

func (r *Doctor) execute(p *Patient) error {
	fmt.Println("Visited Doctor")

	err := r.next.execute(p)
	if err != nil {
		return err
	}
	return nil
}

func (r *Doctor) setNext(department Department) {
	r.next = department
}

// ======================== Medical

type Medical struct {
	next Department
}

func (r *Medical) execute(p *Patient) error {
	fmt.Println("Visited Medical")

	err := r.next.execute(p)
	if err != nil {
		return err
	}
	return nil
}

func (r *Medical) setNext(department Department) {
	r.next = department
}

// ======================== Medical

type Cashier struct {
	next Department
}

func (r *Cashier) execute(p *Patient) error {
	fmt.Println("Visited Cashier")
	return nil
}

func (r *Cashier) setNext(department Department) {
	r.next = department
}

func main() {
	patient := &Patient{
		Name: "Mag",
	}

	cashier := &Cashier{}

	medical := &Medical{}
	medical.setNext(cashier)

	doctor := &Doctor{}
	doctor.setNext(medical)

	receptionist := &Receptionist{}
	receptionist.setNext(doctor)

	err := receptionist.execute(patient)
	if err != nil {
		log.Fatal(err)
	}
}
