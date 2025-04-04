package client

import "design-pattern-practice/creational-pattern/builder/builder"

func main() {
	var carBuilder builder.CarBuilder
	toyotaBuilder := builder.ToyotaBuilder{}
	carBuilder = &toyotaBuilder
	car := carBuilder.SetBody("this body").
		SetWheels("this wheels").
		SetPowertrain("this powertrain").
		SetChassis("this chassis").GetCar()
	car.Show()
}
