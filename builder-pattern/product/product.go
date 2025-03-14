package product

import "fmt"

type Car interface {
	Show()
}

type Honda struct {
	Powertrain string
	Chassis    string
	Body       string
	Wheels     string
}

func (h *Honda) Show() {
	fmt.Print("this is Honda")
}

type BMW struct {
	Powertrain string
	Chassis    string
	Body       string
	Wheels     string
}

func (b *BMW) Show() {
	fmt.Print("this is BMW")
}

type Toyota struct {
	Powertrain string
	Chassis    string
	Body       string
	Wheels     string
}

func (b *Toyota) Show() {
	fmt.Print("this is Toyota")
}
