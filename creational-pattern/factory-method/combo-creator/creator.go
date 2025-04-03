package combocreator

import (
	"design-pattern-practice/creational-pattern/factory-method/combo"
	"design-pattern-practice/creational-pattern/factory-method/item"
)

type Creator interface {
	CreateCombo()
}

type ProteinPowerMealCreator struct {
	combo combo.ProteinPowerMeal
}

func (c *ProteinPowerMealCreator) CreateCombo() {
	c.combo.Add(&item.GrilledChickenBreast{})
	c.combo.Add(&item.SteamedBroccoli{})
	c.combo.Add(&item.ScrambledEggs{})
	c.combo.Add(&item.BrownRice{})
	c.combo.Add(&item.GreekYogurt{})
}

type HeartyComfortMeal struct {
	combo combo.HeartyComfortMeal
}

func (c *HeartyComfortMeal) CreateCombo() {
	c.combo.Add(&item.BeefSteak{})
	c.combo.Add(&item.CreamyMushroomSoup{})
	c.combo.Add(&item.CaesarSalad{})
	c.combo.Add(&item.GarlicButterBread{})
	c.combo.Add(&item.ChocolateLavaCake{})
}

func CreateCombo(nameCombo string) combo.Combo {
	var comboResult Creator
	if nameCombo == "Protein Power Meal" {
		combo := ProteinPowerMealCreator{}
		comboResult = &combo
	} else if nameCombo == "Hearty Comfort Meal" {
		combo := ProteinPowerMealCreator{}
		comboResult = &combo
	}
	if comboResult != nil {
		comboResult.CreateCombo()
	}
	return nil
}
