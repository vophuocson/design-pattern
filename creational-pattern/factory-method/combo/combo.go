package combo

import "design-pattern-practice/creational-pattern/factory-method/item"

type Combo interface {
	ShowCombo()
	Add(it item.Item)
}

type ProteinPowerMeal struct {
	items []item.Item
}

func (i *ProteinPowerMeal) ShowCombo() {
	for _, item := range i.items {
		item.GetName()
	}
}
func (i *ProteinPowerMeal) Add(it item.Item) {
	i.items = append(i.items, it)
}

type HeartyComfortMeal struct {
	items []item.Item
}

func (i *HeartyComfortMeal) ShowCombo() {
	for _, item := range i.items {
		item.GetName()
	}
}

func (i *HeartyComfortMeal) Add(it item.Item) {
	i.items = append(i.items, it)
}
