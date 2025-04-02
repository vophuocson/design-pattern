package builder

import "design-pattern-practice/builder/product"

type CarBuilder interface {
	SetPowertrain(string) CarBuilder
	SetChassis(string) CarBuilder
	SetBody(string) CarBuilder
	SetWheels(string) CarBuilder
	GetCar() product.Car
}

type HondaBuilder struct {
	C product.Honda
}

func (c *HondaBuilder) GetCar() product.Car {
	return &c.C
}

func (c *HondaBuilder) SetPowertrain(powertrain string) CarBuilder {
	c.C.Powertrain = powertrain
	return c
}

func (c *HondaBuilder) SetChassis(chassis string) CarBuilder {
	c.C.Chassis = chassis
	return c
}

func (c *HondaBuilder) SetBody(body string) CarBuilder {
	c.C.Body = body
	return c
}

func (c *HondaBuilder) SetWheels(wheels string) CarBuilder {
	c.C.Wheels = wheels
	return c
}

type BMWBuilder struct {
	C product.BMW
}

func (c *BMWBuilder) GetCar() product.Car {
	return &c.C
}

func (c *BMWBuilder) SetPowertrain(powertrain string) CarBuilder {
	c.C.Powertrain = powertrain
	return c
}

func (c *BMWBuilder) SetChassis(chassis string) CarBuilder {
	c.C.Chassis = chassis
	return c
}

func (c *BMWBuilder) SetBody(body string) CarBuilder {
	c.C.Body = body
	return c
}

func (c *BMWBuilder) SetWheels(wheels string) CarBuilder {
	c.C.Wheels = wheels
	return c
}

type ToyotaBuilder struct {
	C product.Toyota
}

func (c *ToyotaBuilder) GetCar() product.Car {
	return &c.C
}

func (c *ToyotaBuilder) SetPowertrain(powertrain string) CarBuilder {
	c.C.Powertrain = powertrain
	return c
}

func (c *ToyotaBuilder) SetChassis(chassis string) CarBuilder {
	c.C.Chassis = chassis
	return c
}

func (c *ToyotaBuilder) SetBody(body string) CarBuilder {
	c.C.Body = body
	return c
}

func (c *ToyotaBuilder) SetWheels(wheels string) CarBuilder {
	c.C.Wheels = wheels
	return c
}
