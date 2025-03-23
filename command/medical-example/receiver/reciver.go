package receiver

import "fmt"

type Device interface {
	On()
	Off()
}

type TV struct{}

func (d *TV) On() {
	fmt.Println("this is TV ON")
}

func (d *TV) Off() {
	fmt.Println("this is TV OFF")
}

type Radio struct{}

func (d *Radio) On() {
	fmt.Println("this is Radio ON")
}

func (d *Radio) Off() {
	fmt.Println("this is Radio OFF")
}
