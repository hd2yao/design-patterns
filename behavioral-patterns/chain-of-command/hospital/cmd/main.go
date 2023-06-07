package main

import "design-patterns/behavioral-patterns/chain-of-command/hospital/handler"

func main() {
	cashier := &handler.Cashier{}

	//Set next for medical department
	medical := &handler.Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &handler.Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &handler.Reception{}
	reception.SetNext(doctor)

	patient := &handler.Patient{Name: "abc"}
	//Patient visiting
	reception.Execute(patient)
}
