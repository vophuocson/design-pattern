package chainofresponsibility

import (
	"design-pattern-practice/chain-of-responsibility/patient"
	processfactory "design-pattern-practice/chain-of-responsibility/process-factory"
)

func main() {
	var process processfactory.Process
	p := patient.Patient{}
	process = processfactory.FactoryProcess("normal")
	process.ReceivePatient(&p)
}
