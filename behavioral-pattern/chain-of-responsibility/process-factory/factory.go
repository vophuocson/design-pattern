package processfactory

import (
	"design-pattern-practice/behavioral-pattern/chain-of-responsibility/department"
	"design-pattern-practice/behavioral-pattern/chain-of-responsibility/patient"
)

type Process interface {
	ReceivePatient(p *patient.Patient)
}

type NormalCare struct {
	D department.Department
}

func (f *NormalCare) ReceivePatient(p *patient.Patient) {
	f.D.Execute(p)
}

func FactoryProcess(processName string) Process {
	if processName == "normal" {
		var d department.Department
		cashier := department.Cashier{}
		medical := department.Medical{}
		medical.SetNext(&cashier)
		doctor := department.Doctor{}
		doctor.SetNext(&medical)
		reception := department.Reception{}
		reception.SetNext(&doctor)
		d = &reception
		return &NormalCare{D: d}
	}
	return nil
}
