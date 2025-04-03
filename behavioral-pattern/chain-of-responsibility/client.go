package chainofresponsibility

import (
	"design-pattern-practice/behavioral-pattern/chain-of-responsibility/patient"
	processfactory "design-pattern-practice/behavioral-pattern/chain-of-responsibility/process-factory"
)

func main() {
	var process processfactory.Process
	p := patient.Patient{}
	process = processfactory.FactoryProcess("normal")
	process.ReceivePatient(&p)
}
